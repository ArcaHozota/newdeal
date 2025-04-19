package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Phrase struct {
	ent.Schema
}

func (Phrase) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Comment("節ID").
			Immutable().
			Unique().
			SchemaType(map[string]string{
				dialect.Postgres: "bigint",
			}).
			Annotations(
				entsql.Annotation{Default: "0"},
			),
		field.String("name").
			Comment("節名称").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(33)",
			}),
		field.Text("text_en").
			Comment("内容"),
		field.Text("text_jp").
			Comment("日本語内容"),
		field.Int32("chapter_id").
			Comment("章節ID").
			SchemaType(map[string]string{
				dialect.Postgres: "integer",
			}),
		field.Bool("change_line").
			Comment("改行フラグ"),
	}
}

func (Phrase) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("phrase_chapter", Chapter.Type).
			Ref("to_phrase").
			Field("chapter_id").
			Required().
			Unique(),
	}
}

func (Phrase) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Phrase) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "phrases"},
		entsql.WithComments(true),
	}
}
