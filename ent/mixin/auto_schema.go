package mixin

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type AutoSchemaMixin struct {
	Single string
	Plural string
}

func (m AutoSchemaMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New).Unique().Annotations(
			entgql.OrderField("ID"),
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		),
		field.Time("created_at").Default(time.Now).Immutable().Annotations(
			entgql.OrderField("CREATED_AT"),
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(
			entgql.OrderField("UPDATED_AT"),
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		),
	}
}

func (m AutoSchemaMixin) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (m AutoSchemaMixin) Indexes() []ent.Index {
	return []ent.Index{}
}

func (m AutoSchemaMixin) Hooks() []ent.Hook {
	return []ent.Hook{}
}

func (m AutoSchemaMixin) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{}
}

func (m AutoSchemaMixin) Policy() ent.Policy {
	return nil
}

func (m AutoSchemaMixin) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

func (m AutoSchemaMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(m.Single),
		entgql.QueryField(m.Plural),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}
