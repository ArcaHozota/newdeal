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
			MaxLen(40).
			Comment("アカウント"),
		field.String("password").
			MaxLen(255).
			Comment("パスワード"),
		field.String("username").
			MaxLen(40).
			Comment("ユーザ名称"),
		field.Time("date_of_birth").
			Comment("生年月日").
			SchemaType(map[string]string{
				dialect.Postgres: "date",
			}),
		field.String("email").
			MaxLen(60).
			Comment("メール").
			Optional(),
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
		index.Fields("login_account").
			Unique(),
		index.Fields("email").
			Unique(),
	}
}

func (Student) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "students"},
	}
}
