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

type Role struct {
	ent.Schema
}

func (Role) Fields() []ent.Field {
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
			Comment("名称").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(40)",
			}),
		field.Bool("visible_flg").
			Comment("論理削除フラグ"),
	}
}

func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("student", Student.Type),
	}
}

func (Role) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Unique(),
	}
}

func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "roles"},
		entsql.WithComments(true),
	}
}
