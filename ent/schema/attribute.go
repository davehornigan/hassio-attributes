package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/davehornigan/hassio-attributes/ent/mixin"
)

type Attribute struct {
	ent.Schema
}

func (Attribute) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").NotEmpty().Annotations(entgql.OrderField("USER_ID")),
		field.String("attribute_name").NotEmpty().Annotations(entgql.OrderField("ATTRIBUTE_NAME")),
		field.JSON("json_value", map[string]interface{}{}),
	}
}

func (Attribute) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.To("user", User.Type).Unique().Required().Field("user_id"),
	}
}

func (Attribute) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "attribute_name").Unique(),
	}
}

func (Attribute) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoSchemaMixin{
			Single: "attribute",
			Plural: "attributes",
		},
	}
}
