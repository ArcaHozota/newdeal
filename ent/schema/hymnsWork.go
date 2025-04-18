package schema

import (
	"newdeal/common/tools"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type HymnsWork struct {
	ent.Schema
}

func (HymnsWork) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Default(int64(tools.SnowflakeID())).
			SchemaType(map[string]string{
				dialect.Postgres: "bigint",
			}),
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
	return []ent.Index{}
}
