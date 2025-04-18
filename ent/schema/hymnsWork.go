package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type HymnsWork struct {
	ent.Schema
}

func (HymnsWork) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("work_id").Immutable(),
		field.Bytes("score").Nillable(),
		field.String("name_jp_rational").Nillable(),
		field.Time("updated_time"),
		field.String("biko").Nillable(),
	}
}

func (HymnsWork) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("hymns", Hymn.Type).
			Ref("hymns_work").
			Unique().
			// We add the "Required" method to the builder
			// to make this edge required on entity creation.
			// i.e. Card cannot be created without its owner.
			Required(),
	}
}

func (HymnsWork) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("work_id").Unique(),
	}
}
