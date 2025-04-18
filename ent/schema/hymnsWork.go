package schema

import (
	"entgo.io/ent"
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
		field.Int64("work_id").
			Unique().
			Comment("ワークID").
			Annotations(
				entsql.Annotation{Default: "0"}, // 明确设置无 default/identity
			),
		field.Bytes("score").
			Comment("楽譜").
			Optional(),
		field.String("name_jp_rational").
			MaxLen(120).
			Comment("日本語名称").
			Optional(),
		field.Time("updated_time").
			Comment("更新時間"),
		field.String("biko").
			MaxLen(10).
			Comment("備考").
			Optional(),
	}
}

func (HymnsWork) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("linked_hymn", Hymn.Type).
			Ref("work").
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
	}
}
