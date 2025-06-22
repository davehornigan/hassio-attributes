package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/davehornigan/hassio-attributes/ent/mixin"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("external_id").NotEmpty().Unique().Annotations(
			entgql.QueryField(),
			entgql.OrderField("ExternalID"),
		),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.From("attributes", Attribute.Type).Ref("user").Field("external_id").Unique(),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoSchemaMixin{
			Single: "user",
			Plural: "users",
		},
	}
}
