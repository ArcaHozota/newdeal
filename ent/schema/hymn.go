package schema

import (
	"newdeal/common/tools"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Hymn struct {
	ent.Schema
}

func (Hymn) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", tools.SnowflakeUUID()).
			SchemaType(map[string]string{
				dialect.Postgres: "bigint",
			}),
		field.String("name_jp"),
		field.String("name_kr"),
		field.String("link").Nillable(),
		field.Time("updated_time"),
		field.Int64("updated_user").
			Positive().
			SchemaType(map[string]string{
				dialect.Postgres: "bigint",
			}),
		field.Text("serif").Nillable(),
		field.Bool("visible_flg"),
	}
}

func (Hymn) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("students", Student.Type).
			Ref("hymns").
			Unique(),
		edge.To("hymns_work", HymnsWork.Type).
			Unique(),
	}
}

func (Hymn) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("link", "name_jp", "name_kr").
			Unique(),
	}
}
