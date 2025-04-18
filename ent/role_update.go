// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"newdeal/ent/auth"
	"newdeal/ent/predicate"
	"newdeal/ent/role"
	"newdeal/ent/student"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoleUpdate is the builder for updating Role entities.
type RoleUpdate struct {
	config
	hooks    []Hook
	mutation *RoleMutation
}

// Where appends a list predicates to the RoleUpdate builder.
func (ru *RoleUpdate) Where(ps ...predicate.Role) *RoleUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetName sets the "name" field.
func (ru *RoleUpdate) SetName(s string) *RoleUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableName(s *string) *RoleUpdate {
	if s != nil {
		ru.SetName(*s)
	}
	return ru
}

// SetVisibleFlg sets the "visible_flg" field.
func (ru *RoleUpdate) SetVisibleFlg(b bool) *RoleUpdate {
	ru.mutation.SetVisibleFlg(b)
	return ru
}

// SetNillableVisibleFlg sets the "visible_flg" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableVisibleFlg(b *bool) *RoleUpdate {
	if b != nil {
		ru.SetVisibleFlg(*b)
	}
	return ru
}

// AddStudentIDs adds the "student" edge to the Student entity by IDs.
func (ru *RoleUpdate) AddStudentIDs(ids ...int64) *RoleUpdate {
	ru.mutation.AddStudentIDs(ids...)
	return ru
}

// AddStudent adds the "student" edges to the Student entity.
func (ru *RoleUpdate) AddStudent(s ...*Student) *RoleUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ru.AddStudentIDs(ids...)
}

// AddAuthIDs adds the "auths" edge to the Auth entity by IDs.
func (ru *RoleUpdate) AddAuthIDs(ids ...int64) *RoleUpdate {
	ru.mutation.AddAuthIDs(ids...)
	return ru
}

// AddAuths adds the "auths" edges to the Auth entity.
func (ru *RoleUpdate) AddAuths(a ...*Auth) *RoleUpdate {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ru.AddAuthIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (ru *RoleUpdate) Mutation() *RoleMutation {
	return ru.mutation
}

// ClearStudent clears all "student" edges to the Student entity.
func (ru *RoleUpdate) ClearStudent() *RoleUpdate {
	ru.mutation.ClearStudent()
	return ru
}

// RemoveStudentIDs removes the "student" edge to Student entities by IDs.
func (ru *RoleUpdate) RemoveStudentIDs(ids ...int64) *RoleUpdate {
	ru.mutation.RemoveStudentIDs(ids...)
	return ru
}

// RemoveStudent removes "student" edges to Student entities.
func (ru *RoleUpdate) RemoveStudent(s ...*Student) *RoleUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ru.RemoveStudentIDs(ids...)
}

// ClearAuths clears all "auths" edges to the Auth entity.
func (ru *RoleUpdate) ClearAuths() *RoleUpdate {
	ru.mutation.ClearAuths()
	return ru
}

// RemoveAuthIDs removes the "auths" edge to Auth entities by IDs.
func (ru *RoleUpdate) RemoveAuthIDs(ids ...int64) *RoleUpdate {
	ru.mutation.RemoveAuthIDs(ids...)
	return ru
}

// RemoveAuths removes "auths" edges to Auth entities.
func (ru *RoleUpdate) RemoveAuths(a ...*Auth) *RoleUpdate {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ru.RemoveAuthIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RoleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoleUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoleUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoleUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(role.Table, role.Columns, sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt64))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(role.FieldName, field.TypeString, value)
	}
	if value, ok := ru.mutation.VisibleFlg(); ok {
		_spec.SetField(role.FieldVisibleFlg, field.TypeBool, value)
	}
	if ru.mutation.StudentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.StudentTable,
			Columns: []string{role.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedStudentIDs(); len(nodes) > 0 && !ru.mutation.StudentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.StudentTable,
			Columns: []string{role.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.StudentTable,
			Columns: []string{role.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.AuthsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.AuthsTable,
			Columns: role.AuthsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(auth.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedAuthsIDs(); len(nodes) > 0 && !ru.mutation.AuthsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.AuthsTable,
			Columns: role.AuthsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(auth.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.AuthsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.AuthsTable,
			Columns: role.AuthsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(auth.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RoleUpdateOne is the builder for updating a single Role entity.
type RoleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoleMutation
}

// SetName sets the "name" field.
func (ruo *RoleUpdateOne) SetName(s string) *RoleUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableName(s *string) *RoleUpdateOne {
	if s != nil {
		ruo.SetName(*s)
	}
	return ruo
}

// SetVisibleFlg sets the "visible_flg" field.
func (ruo *RoleUpdateOne) SetVisibleFlg(b bool) *RoleUpdateOne {
	ruo.mutation.SetVisibleFlg(b)
	return ruo
}

// SetNillableVisibleFlg sets the "visible_flg" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableVisibleFlg(b *bool) *RoleUpdateOne {
	if b != nil {
		ruo.SetVisibleFlg(*b)
	}
	return ruo
}

// AddStudentIDs adds the "student" edge to the Student entity by IDs.
func (ruo *RoleUpdateOne) AddStudentIDs(ids ...int64) *RoleUpdateOne {
	ruo.mutation.AddStudentIDs(ids...)
	return ruo
}

// AddStudent adds the "student" edges to the Student entity.
func (ruo *RoleUpdateOne) AddStudent(s ...*Student) *RoleUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ruo.AddStudentIDs(ids...)
}

// AddAuthIDs adds the "auths" edge to the Auth entity by IDs.
func (ruo *RoleUpdateOne) AddAuthIDs(ids ...int64) *RoleUpdateOne {
	ruo.mutation.AddAuthIDs(ids...)
	return ruo
}

// AddAuths adds the "auths" edges to the Auth entity.
func (ruo *RoleUpdateOne) AddAuths(a ...*Auth) *RoleUpdateOne {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ruo.AddAuthIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (ruo *RoleUpdateOne) Mutation() *RoleMutation {
	return ruo.mutation
}

// ClearStudent clears all "student" edges to the Student entity.
func (ruo *RoleUpdateOne) ClearStudent() *RoleUpdateOne {
	ruo.mutation.ClearStudent()
	return ruo
}

// RemoveStudentIDs removes the "student" edge to Student entities by IDs.
func (ruo *RoleUpdateOne) RemoveStudentIDs(ids ...int64) *RoleUpdateOne {
	ruo.mutation.RemoveStudentIDs(ids...)
	return ruo
}

// RemoveStudent removes "student" edges to Student entities.
func (ruo *RoleUpdateOne) RemoveStudent(s ...*Student) *RoleUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ruo.RemoveStudentIDs(ids...)
}

// ClearAuths clears all "auths" edges to the Auth entity.
func (ruo *RoleUpdateOne) ClearAuths() *RoleUpdateOne {
	ruo.mutation.ClearAuths()
	return ruo
}

// RemoveAuthIDs removes the "auths" edge to Auth entities by IDs.
func (ruo *RoleUpdateOne) RemoveAuthIDs(ids ...int64) *RoleUpdateOne {
	ruo.mutation.RemoveAuthIDs(ids...)
	return ruo
}

// RemoveAuths removes "auths" edges to Auth entities.
func (ruo *RoleUpdateOne) RemoveAuths(a ...*Auth) *RoleUpdateOne {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ruo.RemoveAuthIDs(ids...)
}

// Where appends a list predicates to the RoleUpdate builder.
func (ruo *RoleUpdateOne) Where(ps ...predicate.Role) *RoleUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RoleUpdateOne) Select(field string, fields ...string) *RoleUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Role entity.
func (ruo *RoleUpdateOne) Save(ctx context.Context) (*Role, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoleUpdateOne) SaveX(ctx context.Context) *Role {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RoleUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoleUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RoleUpdateOne) sqlSave(ctx context.Context) (_node *Role, err error) {
	_spec := sqlgraph.NewUpdateSpec(role.Table, role.Columns, sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt64))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Role.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, role.FieldID)
		for _, f := range fields {
			if !role.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != role.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(role.FieldName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.VisibleFlg(); ok {
		_spec.SetField(role.FieldVisibleFlg, field.TypeBool, value)
	}
	if ruo.mutation.StudentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.StudentTable,
			Columns: []string{role.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedStudentIDs(); len(nodes) > 0 && !ruo.mutation.StudentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.StudentTable,
			Columns: []string{role.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.StudentTable,
			Columns: []string{role.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.AuthsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.AuthsTable,
			Columns: role.AuthsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(auth.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedAuthsIDs(); len(nodes) > 0 && !ruo.mutation.AuthsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.AuthsTable,
			Columns: role.AuthsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(auth.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.AuthsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.AuthsTable,
			Columns: role.AuthsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(auth.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Role{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
