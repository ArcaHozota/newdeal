// Code generated by ent, DO NOT EDIT.

package auth

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the auth type in the database.
	Label = "auth"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldCategoryID holds the string denoting the category_id field in the database.
	FieldCategoryID = "category_id"
	// EdgeToRoles holds the string denoting the to_roles edge name in mutations.
	EdgeToRoles = "to_roles"
	// Table holds the table name of the auth in the database.
	Table = "authorities"
	// ToRolesTable is the table that holds the to_roles relation/edge. The primary key declared below.
	ToRolesTable = "role_auth"
	// ToRolesInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	ToRolesInverseTable = "roles"
)

// Columns holds all SQL columns for auth fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldTitle,
	FieldCategoryID,
}

var (
	// ToRolesPrimaryKey and ToRolesColumn2 are the table columns denoting the
	// primary key for the to_roles relation (M2M).
	ToRolesPrimaryKey = []string{"role_id", "auth_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Auth queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByCategoryID orders the results by the category_id field.
func ByCategoryID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategoryID, opts...).ToFunc()
}

// ByToRolesCount orders the results by to_roles count.
func ByToRolesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newToRolesStep(), opts...)
	}
}

// ByToRoles orders the results by to_roles terms.
func ByToRoles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newToRolesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newToRolesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ToRolesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ToRolesTable, ToRolesPrimaryKey...),
	)
}
