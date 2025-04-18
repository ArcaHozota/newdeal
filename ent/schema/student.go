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

type Student struct {
	ent.Schema
}

func (Student) Fields() []ent.Field {
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
		field.String("login_account").
			Comment("アカウント").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(40)",
			}),
		field.String("password").
			Comment("パスワード").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(255)",
			}),
		field.String("username").
			Comment("ユーザ名称").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(40)",
			}),
		field.Time("date_of_birth").
			Comment("生年月日").
			SchemaType(map[string]string{
				dialect.Postgres: "date",
			}),
		field.String("email").
			Comment("メール").
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(60)",
			}),
		field.Time("updated_time").
			Comment("登録時間").
			Optional(),
		field.Bool("visible_flg").
			Comment("論理削除フラグ"),
	}
}

func (Student) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("updated_hymns", Hymn.Type),
	}
}

func (Student) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("login_account").Unique(),
		index.Fields("email").Unique(),
	}
}

func (Student) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "students"},
	}
}
