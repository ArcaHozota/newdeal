package schema

import (
	"newdeal/common/tools"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Student struct {
	ent.Schema
}

func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(tools.SnowflakeID()).
			SchemaType(map[string]string{
				dialect.Postgres: "bigint",
			}),
		field.String("login_account"),
		field.String("password"),
		field.String("username"),
		field.Time("date_of_birth").
			SchemaType(map[string]string{
				dialect.Postgres: "date",
			}),
		field.String("email").Nillable(),
		field.Time("updated_time").Nillable(),
		field.Bool("visible_flg"),
	}
}

func (Student) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hymns", Hymn.Type),
	}
}

func (Student) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "login_account", "email").
			Unique(),
	}
}
