package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Chapter struct {
	ent.Schema
}

func (Chapter) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").
			Comment("章節ID").
			Immutable().
			Unique().
			SchemaType(map[string]string{
				dialect.Postgres: "integer",
			}).
			Annotations(
				entsql.Annotation{Default: "0"},
			),
		field.String("name").
			Comment("章節名").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(33)",
			}),
		field.Text("name_jp").
			Comment("章節日本語名").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(33)",
			}),
		field.Int16("book_id").
			Comment("書別ID").
			SchemaType(map[string]string{
				dialect.Postgres: "smallint",
			}),
	}
}

func (Chapter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("to_phrase", Phrase.Type),
		edge.From("chapter_book", Book.Type).
			Ref("to_chapter").
			Field("book_id").
			Required().
			Unique(),
	}
}

func (Chapter) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Chapter) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "chapters"},
		entsql.WithComments(true),
	}
}
