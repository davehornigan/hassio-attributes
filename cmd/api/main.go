package main

import (
	"context"
	"github.com/davehornigan/hassio-attributes/pkg/config"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/sirupsen/logrus"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"

	"github.com/davehornigan/hassio-attributes/ent"
	"github.com/davehornigan/hassio-attributes/graph"
)

var dbCfg config.DatabaseConfig

func init() {
	config.LoadConfig(&dbCfg)
}

func main() {
	// Setup logrus
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)

	// Connect to PostgreSQL
	client, err := ent.Open(dialect.Postgres, dbCfg.GetDsnString())
	if err != nil {
		log.WithError(err).Fatal("failed to connect to database")
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.WithError(err).Fatal("failed creating schema")
	}

	// Setup Fiber app
	app := fiber.New()
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${ip} ${status} - ${latency} ${method} ${path} - request_id=${locals:requestid}\n",
		TimeFormat: "2006-01-02T15:04:05Z07:00",
		Output:     os.Stdout,
	}))

	app.Use(healthcheck.New())

	// GraphQL handler
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{Client: client},
	}))

	// Inject request ID into context
	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		if rid := ctx.Value(fiber.HeaderXRequestID); rid != nil {
			ctx = context.WithValue(ctx, "request_id", rid)
			ctx = context.WithValue(ctx, "logger", log.WithField("request_id", rid))
		}
		return next(ctx)
	})

	// Mount GraphQL endpoints via adaptor
	app.All("/query", adaptor.HTTPHandler(srv))
	app.Get("/", adaptor.HTTPHandler(playground.Handler("GraphQL", "/query")))

	log.WithField("port", 8080).Info("🚀 API ready at http://localhost:8080/")
	if err := app.Listen(":8080"); err != nil {
		log.WithError(err).Fatal("server exited with error")
	}
}
