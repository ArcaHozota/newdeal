package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Auth struct {
	ent.Schema
}

func (Auth) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Comment("ID").
			Immutable().
			Unique().
			SchemaType(map[string]string{
				dialect.Postgres: "bigint",
			}).
			Annotations(
				entsql.Annotation{Default: "0"}, // 明确设置无 default/identity
			),
		field.String("name").
			Comment("権限名称").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(50)",
			}),
		field.String("title").
			Comment("漢字名称").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(40)",
			}),
		field.Int64("category_id").
			Comment("分類ID").
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "bigint",
			}),
	}
}

func (Auth) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("to_roles", Role.Type).
			Ref("auth"),
	}
}

func (Auth) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Unique(),
		index.Fields("title").Unique(),
	}
}

func (Auth) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "authorities"},
		entsql.WithComments(true),
	}
}
