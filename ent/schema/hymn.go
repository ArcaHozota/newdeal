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

type Hymn struct {
	ent.Schema
}

func (Hymn) Fields() []ent.Field {
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
		field.String("name_jp").
			Comment("日本語名称").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(66)",
			}),
		field.String("name_kr").
			Comment("韓国語名称").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(66)",
			}),
		field.String("link").
			Comment("リンク").
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(255)",
			}),
		field.Int64("updated_user").
			Comment("更新者ID"),
		field.Time("updated_time").
			Comment("更新時間"),
		field.Text("serif").
			Comment("歌詞").
			Optional(),
		field.Bool("visible_flg").
			Comment("論理削除フラグ"),
	}
}

func (Hymn) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("updated_by", Student.Type).
			Ref("updated_hymns").  // 反向关系字段
			Field("updated_user"). // 外键字段
			Required().
			Unique(),
		edge.To("to_work", HymnsWork.Type).
			Unique(),
	}
}

func (Hymn) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("link").Unique(),
		index.Fields("name_jp").Unique(),
		index.Fields("name_kr").Unique(),
	}
}

func (Hymn) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "hymns"},
		entsql.WithComments(true),
	}
}
