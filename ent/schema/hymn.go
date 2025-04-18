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
			}),
		field.String("name_jp").
			MaxLen(66).
			Comment("日本語名称"),
		field.String("name_kr").
			MaxLen(66).
			Comment("韓国語名称"),
		field.String("link").
			MaxLen(255).
			Comment("リンク").
			Optional(),
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
		edge.To("work", HymnsWork.Type).
			Unique(),
	}
}

func (Hymn) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("link").
			Unique(),
		index.Fields("name_jp").
			Unique(),
		index.Fields("name_kr").
			Unique(),
	}
}

func (Hymn) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "hymns"},
	}
}
