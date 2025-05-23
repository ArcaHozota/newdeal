package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type HymnsWork struct {
	ent.Schema
}

func (HymnsWork) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Comment("ID").
			Immutable().
			Unique().
			SchemaType(map[string]string{
				dialect.Postgres: "bigint",
			}),
		field.Int64("work_id").
			Comment("ワークID"),
		field.Bytes("score").
			Comment("楽譜").
			Optional(),
		field.String("name_jp_rational").
			Comment("日本語名称").
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(120)",
			}),
		field.Time("updated_time").
			Comment("更新時間"),
		field.String("biko").
			Comment("備考").
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(15)",
			}),
	}
}

func (HymnsWork) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("linked_hymn", Hymn.Type).
			Ref("to_work").
			Field("work_id").
			// We add the "Required" method to the builder
			// to make this edge required on entity creation.
			// i.e. Card cannot be created without its owner.
			Required().
			Unique(),
	}
}

func (HymnsWork) Indexes() []ent.Index {
	return []ent.Index{}
}

// Annotations of the HymnsWork.
func (HymnsWork) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "hymns_work"},
		entsql.WithComments(true),
	}
}
