package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Book struct {
	ent.Schema
}

func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.Int16("id").
			Comment("書別ID").
			Immutable().
			Unique().
			SchemaType(map[string]string{
				dialect.Postgres: "smallint",
			}).
			Annotations(
				entsql.Annotation{Default: "0"},
			),
		field.String("name").
			Comment("書別名").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(33)",
			}),
		field.Text("name_jp").
			Comment("書別日本語名").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(33)",
			}),
	}
}

func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("to_chapter", Chapter.Type),
	}
}

func (Book) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Book) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "books"},
		entsql.WithComments(true),
	}
}
