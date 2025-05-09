// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BooksColumns holds the columns for the "books" table.
	BooksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt16, Increment: true, Comment: "書別ID", Default: "0", SchemaType: map[string]string{"postgres": "smallint"}},
		{Name: "name", Type: field.TypeString, Comment: "書別名", SchemaType: map[string]string{"postgres": "varchar(33)"}},
		{Name: "name_jp", Type: field.TypeString, Size: 2147483647, Comment: "書別日本語名", SchemaType: map[string]string{"postgres": "varchar(33)"}},
	}
	// BooksTable holds the schema information for the "books" table.
	BooksTable = &schema.Table{
		Name:       "books",
		Columns:    BooksColumns,
		PrimaryKey: []*schema.Column{BooksColumns[0]},
	}
	// ChaptersColumns holds the columns for the "chapters" table.
	ChaptersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true, Comment: "章節ID", Default: "0", SchemaType: map[string]string{"postgres": "integer"}},
		{Name: "name", Type: field.TypeString, Comment: "章節名", SchemaType: map[string]string{"postgres": "varchar(33)"}},
		{Name: "name_jp", Type: field.TypeString, Size: 2147483647, Comment: "章節日本語名", SchemaType: map[string]string{"postgres": "varchar(33)"}},
		{Name: "book_id", Type: field.TypeInt16, Comment: "書別ID", SchemaType: map[string]string{"postgres": "smallint"}},
	}
	// ChaptersTable holds the schema information for the "chapters" table.
	ChaptersTable = &schema.Table{
		Name:       "chapters",
		Columns:    ChaptersColumns,
		PrimaryKey: []*schema.Column{ChaptersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "chapters_books_to_chapter",
				Columns:    []*schema.Column{ChaptersColumns[3]},
				RefColumns: []*schema.Column{BooksColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// HymnsColumns holds the columns for the "hymns" table.
	HymnsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, Comment: "ID", Default: "0", SchemaType: map[string]string{"postgres": "bigint"}},
		{Name: "name_jp", Type: field.TypeString, Comment: "日本語名称", SchemaType: map[string]string{"postgres": "varchar(66)"}},
		{Name: "name_kr", Type: field.TypeString, Comment: "韓国語名称", SchemaType: map[string]string{"postgres": "varchar(66)"}},
		{Name: "link", Type: field.TypeString, Nullable: true, Comment: "リンク", SchemaType: map[string]string{"postgres": "varchar(255)"}},
		{Name: "updated_time", Type: field.TypeTime, Comment: "更新時間"},
		{Name: "serif", Type: field.TypeString, Nullable: true, Size: 2147483647, Comment: "歌詞"},
		{Name: "visible_flg", Type: field.TypeBool, Comment: "論理削除フラグ"},
		{Name: "updated_user", Type: field.TypeInt64, Comment: "更新者ID", SchemaType: map[string]string{"postgres": "bigint"}},
	}
	// HymnsTable holds the schema information for the "hymns" table.
	HymnsTable = &schema.Table{
		Name:       "hymns",
		Columns:    HymnsColumns,
		PrimaryKey: []*schema.Column{HymnsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "hymns_students_updated_hymns",
				Columns:    []*schema.Column{HymnsColumns[7]},
				RefColumns: []*schema.Column{StudentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "hymn_link",
				Unique:  true,
				Columns: []*schema.Column{HymnsColumns[3]},
			},
			{
				Name:    "hymn_name_jp",
				Unique:  true,
				Columns: []*schema.Column{HymnsColumns[1]},
			},
			{
				Name:    "hymn_name_kr",
				Unique:  true,
				Columns: []*schema.Column{HymnsColumns[2]},
			},
		},
	}
	// HymnsWorkColumns holds the columns for the "hymns_work" table.
	HymnsWorkColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, Comment: "ID", SchemaType: map[string]string{"postgres": "bigint"}},
		{Name: "score", Type: field.TypeBytes, Nullable: true, Comment: "楽譜"},
		{Name: "name_jp_rational", Type: field.TypeString, Nullable: true, Comment: "日本語名称", SchemaType: map[string]string{"postgres": "varchar(120)"}},
		{Name: "updated_time", Type: field.TypeTime, Comment: "更新時間"},
		{Name: "biko", Type: field.TypeString, Nullable: true, Comment: "備考", SchemaType: map[string]string{"postgres": "varchar(15)"}},
		{Name: "work_id", Type: field.TypeInt64, Unique: true, Comment: "ワークID", SchemaType: map[string]string{"postgres": "bigint"}},
	}
	// HymnsWorkTable holds the schema information for the "hymns_work" table.
	HymnsWorkTable = &schema.Table{
		Name:       "hymns_work",
		Columns:    HymnsWorkColumns,
		PrimaryKey: []*schema.Column{HymnsWorkColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "hymns_work_hymns_to_work",
				Columns:    []*schema.Column{HymnsWorkColumns[5]},
				RefColumns: []*schema.Column{HymnsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PhrasesColumns holds the columns for the "phrases" table.
	PhrasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, Comment: "節ID", Default: "0", SchemaType: map[string]string{"postgres": "bigint"}},
		{Name: "name", Type: field.TypeString, Comment: "節名称", SchemaType: map[string]string{"postgres": "varchar(33)"}},
		{Name: "text_en", Type: field.TypeString, Size: 2147483647, Comment: "内容"},
		{Name: "text_jp", Type: field.TypeString, Size: 2147483647, Comment: "日本語内容"},
		{Name: "change_line", Type: field.TypeBool, Comment: "改行フラグ"},
		{Name: "chapter_id", Type: field.TypeInt32, Comment: "章節ID", SchemaType: map[string]string{"postgres": "integer"}},
	}
	// PhrasesTable holds the schema information for the "phrases" table.
	PhrasesTable = &schema.Table{
		Name:       "phrases",
		Columns:    PhrasesColumns,
		PrimaryKey: []*schema.Column{PhrasesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "phrases_chapters_to_phrase",
				Columns:    []*schema.Column{PhrasesColumns[5]},
				RefColumns: []*schema.Column{ChaptersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// StudentsColumns holds the columns for the "students" table.
	StudentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, Comment: "ID", Default: "0", SchemaType: map[string]string{"postgres": "bigint"}},
		{Name: "login_account", Type: field.TypeString, Comment: "アカウント", SchemaType: map[string]string{"postgres": "varchar(40)"}},
		{Name: "password", Type: field.TypeString, Comment: "パスワード", SchemaType: map[string]string{"postgres": "varchar(255)"}},
		{Name: "username", Type: field.TypeString, Comment: "ユーザ名称", SchemaType: map[string]string{"postgres": "varchar(40)"}},
		{Name: "date_of_birth", Type: field.TypeTime, Comment: "生年月日", SchemaType: map[string]string{"postgres": "date"}},
		{Name: "email", Type: field.TypeString, Nullable: true, Comment: "メール", SchemaType: map[string]string{"postgres": "varchar(60)"}},
		{Name: "updated_time", Type: field.TypeTime, Nullable: true, Comment: "登録時間"},
		{Name: "visible_flg", Type: field.TypeBool, Comment: "論理削除フラグ"},
	}
	// StudentsTable holds the schema information for the "students" table.
	StudentsTable = &schema.Table{
		Name:       "students",
		Columns:    StudentsColumns,
		PrimaryKey: []*schema.Column{StudentsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "student_login_account",
				Unique:  true,
				Columns: []*schema.Column{StudentsColumns[1]},
			},
			{
				Name:    "student_email",
				Unique:  true,
				Columns: []*schema.Column{StudentsColumns[5]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BooksTable,
		ChaptersTable,
		HymnsTable,
		HymnsWorkTable,
		PhrasesTable,
		StudentsTable,
	}
)

func init() {
	BooksTable.Annotation = &entsql.Annotation{
		Table: "books",
	}
	ChaptersTable.ForeignKeys[0].RefTable = BooksTable
	ChaptersTable.Annotation = &entsql.Annotation{
		Table: "chapters",
	}
	HymnsTable.ForeignKeys[0].RefTable = StudentsTable
	HymnsTable.Annotation = &entsql.Annotation{
		Table: "hymns",
	}
	HymnsWorkTable.ForeignKeys[0].RefTable = HymnsTable
	HymnsWorkTable.Annotation = &entsql.Annotation{
		Table: "hymns_work",
	}
	PhrasesTable.ForeignKeys[0].RefTable = ChaptersTable
	PhrasesTable.Annotation = &entsql.Annotation{
		Table: "phrases",
	}
	StudentsTable.Annotation = &entsql.Annotation{
		Table: "students",
	}
}
