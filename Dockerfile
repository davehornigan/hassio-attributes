ARG GOLANG_VERSION=1.23

FROM golang:${GOLANG_VERSION}-alpine AS build
LABEL authors="Dave Hornigan<dave@hornigan.com>"

WORKDIR /app

COPY . /app/
RUN go mod download

RUN go build -o /app/build/api /app/cmd/api/main.go

FROM alpine AS api

RUN apk add ca-certificates

WORKDIR /app
COPY --from=build /app/build/ /app/

CMD ["/app/api"]