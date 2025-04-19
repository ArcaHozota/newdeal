// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AuthoritiesColumns holds the columns for the "authorities" table.
	AuthoritiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, Comment: "ID", Default: "0", SchemaType: map[string]string{"postgres": "bigint"}},
		{Name: "name", Type: field.TypeString, Comment: "権限名称", SchemaType: map[string]string{"postgres": "varchar(50)"}},
		{Name: "title", Type: field.TypeString, Comment: "漢字名称", SchemaType: map[string]string{"postgres": "varchar(40)"}},
		{Name: "category_id", Type: field.TypeInt64, Nullable: true, Comment: "分類ID"},
	}
	// AuthoritiesTable holds the schema information for the "authorities" table.
	AuthoritiesTable = &schema.Table{
		Name:       "authorities",
		Columns:    AuthoritiesColumns,
		PrimaryKey: []*schema.Column{AuthoritiesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "auth_name",
				Unique:  true,
				Columns: []*schema.Column{AuthoritiesColumns[1]},
			},
			{
				Name:    "auth_title",
				Unique:  true,
				Columns: []*schema.Column{AuthoritiesColumns[2]},
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
		{Name: "biko", Type: field.TypeString, Nullable: true, Comment: "備考", SchemaType: map[string]string{"postgres": "varchar(10)"}},
		{Name: "work_id", Type: field.TypeInt64, Unique: true, Comment: "ワークID", SchemaType: map[string]string{"postgres": "bigint"}},
	}
	// HymnsWorkTable holds the schema information for the "hymns_work" table.
	HymnsWorkTable = &schema.Table{
		Name:       "hymns_work",
		Columns:    HymnsWorkColumns,
		PrimaryKey: []*schema.Column{HymnsWorkColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "hymns_work_hymns_work",
				Columns:    []*schema.Column{HymnsWorkColumns[5]},
				RefColumns: []*schema.Column{HymnsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "hymnswork_work_id",
				Unique:  true,
				Columns: []*schema.Column{HymnsWorkColumns[5]},
			},
		},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, Comment: "ID", Default: "0", SchemaType: map[string]string{"postgres": "bigint"}},
		{Name: "name", Type: field.TypeString, Comment: "名称", SchemaType: map[string]string{"postgres": "varchar(40)"}},
		{Name: "visible_flg", Type: field.TypeBool, Comment: "論理削除フラグ"},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "role_name",
				Unique:  true,
				Columns: []*schema.Column{RolesColumns[1]},
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
		{Name: "role_id", Type: field.TypeInt64, Comment: "役割ID", SchemaType: map[string]string{"postgres": "bigint"}},
	}
	// StudentsTable holds the schema information for the "students" table.
	StudentsTable = &schema.Table{
		Name:       "students",
		Columns:    StudentsColumns,
		PrimaryKey: []*schema.Column{StudentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "students_roles_student",
				Columns:    []*schema.Column{StudentsColumns[8]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
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
	// RoleAuthsColumns holds the columns for the "role_auths" table.
	RoleAuthsColumns = []*schema.Column{
		{Name: "role_id", Type: field.TypeInt64, SchemaType: map[string]string{"postgres": "bigint"}},
		{Name: "auth_id", Type: field.TypeInt64, SchemaType: map[string]string{"postgres": "bigint"}},
	}
	// RoleAuthsTable holds the schema information for the "role_auths" table.
	RoleAuthsTable = &schema.Table{
		Name:       "role_auths",
		Columns:    RoleAuthsColumns,
		PrimaryKey: []*schema.Column{RoleAuthsColumns[0], RoleAuthsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_auths_role_id",
				Columns:    []*schema.Column{RoleAuthsColumns[0]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "role_auths_auth_id",
				Columns:    []*schema.Column{RoleAuthsColumns[1]},
				RefColumns: []*schema.Column{AuthoritiesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AuthoritiesTable,
		HymnsTable,
		HymnsWorkTable,
		RolesTable,
		StudentsTable,
		RoleAuthsTable,
	}
)

func init() {
	AuthoritiesTable.Annotation = &entsql.Annotation{
		Table: "authorities",
	}
	HymnsTable.ForeignKeys[0].RefTable = StudentsTable
	HymnsTable.Annotation = &entsql.Annotation{
		Table: "hymns",
	}
	HymnsWorkTable.ForeignKeys[0].RefTable = HymnsTable
	HymnsWorkTable.Annotation = &entsql.Annotation{
		Table: "hymns_work",
	}
	RolesTable.Annotation = &entsql.Annotation{
		Table: "roles",
	}
	StudentsTable.ForeignKeys[0].RefTable = RolesTable
	StudentsTable.Annotation = &entsql.Annotation{
		Table: "students",
	}
	RoleAuthsTable.ForeignKeys[0].RefTable = RolesTable
	RoleAuthsTable.ForeignKeys[1].RefTable = AuthoritiesTable
}
