// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"newdeal/ent/hymn"
	"newdeal/ent/predicate"
	"newdeal/ent/student"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StudentUpdate is the builder for updating Student entities.
type StudentUpdate struct {
	config
	hooks    []Hook
	mutation *StudentMutation
}

// Where appends a list predicates to the StudentUpdate builder.
func (su *StudentUpdate) Where(ps ...predicate.Student) *StudentUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetLoginAccount sets the "login_account" field.
func (su *StudentUpdate) SetLoginAccount(s string) *StudentUpdate {
	su.mutation.SetLoginAccount(s)
	return su
}

// SetNillableLoginAccount sets the "login_account" field if the given value is not nil.
func (su *StudentUpdate) SetNillableLoginAccount(s *string) *StudentUpdate {
	if s != nil {
		su.SetLoginAccount(*s)
	}
	return su
}

// SetPassword sets the "password" field.
func (su *StudentUpdate) SetPassword(s string) *StudentUpdate {
	su.mutation.SetPassword(s)
	return su
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (su *StudentUpdate) SetNillablePassword(s *string) *StudentUpdate {
	if s != nil {
		su.SetPassword(*s)
	}
	return su
}

// SetUsername sets the "username" field.
func (su *StudentUpdate) SetUsername(s string) *StudentUpdate {
	su.mutation.SetUsername(s)
	return su
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (su *StudentUpdate) SetNillableUsername(s *string) *StudentUpdate {
	if s != nil {
		su.SetUsername(*s)
	}
	return su
}

// SetDateOfBirth sets the "date_of_birth" field.
func (su *StudentUpdate) SetDateOfBirth(s string) *StudentUpdate {
	su.mutation.SetDateOfBirth(s)
	return su
}

// SetNillableDateOfBirth sets the "date_of_birth" field if the given value is not nil.
func (su *StudentUpdate) SetNillableDateOfBirth(s *string) *StudentUpdate {
	if s != nil {
		su.SetDateOfBirth(*s)
	}
	return su
}

// SetEmail sets the "email" field.
func (su *StudentUpdate) SetEmail(s string) *StudentUpdate {
	su.mutation.SetEmail(s)
	return su
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (su *StudentUpdate) SetNillableEmail(s *string) *StudentUpdate {
	if s != nil {
		su.SetEmail(*s)
	}
	return su
}

// SetUpdatedTime sets the "updated_time" field.
func (su *StudentUpdate) SetUpdatedTime(t time.Time) *StudentUpdate {
	su.mutation.SetUpdatedTime(t)
	return su
}

// SetNillableUpdatedTime sets the "updated_time" field if the given value is not nil.
func (su *StudentUpdate) SetNillableUpdatedTime(t *time.Time) *StudentUpdate {
	if t != nil {
		su.SetUpdatedTime(*t)
	}
	return su
}

// SetVisibleFlg sets the "visible_flg" field.
func (su *StudentUpdate) SetVisibleFlg(b bool) *StudentUpdate {
	su.mutation.SetVisibleFlg(b)
	return su
}

// SetNillableVisibleFlg sets the "visible_flg" field if the given value is not nil.
func (su *StudentUpdate) SetNillableVisibleFlg(b *bool) *StudentUpdate {
	if b != nil {
		su.SetVisibleFlg(*b)
	}
	return su
}

// AddHymnIDs adds the "hymns" edge to the Hymn entity by IDs.
func (su *StudentUpdate) AddHymnIDs(ids ...int64) *StudentUpdate {
	su.mutation.AddHymnIDs(ids...)
	return su
}

// AddHymns adds the "hymns" edges to the Hymn entity.
func (su *StudentUpdate) AddHymns(h ...*Hymn) *StudentUpdate {
	ids := make([]int64, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return su.AddHymnIDs(ids...)
}

// Mutation returns the StudentMutation object of the builder.
func (su *StudentUpdate) Mutation() *StudentMutation {
	return su.mutation
}

// ClearHymns clears all "hymns" edges to the Hymn entity.
func (su *StudentUpdate) ClearHymns() *StudentUpdate {
	su.mutation.ClearHymns()
	return su
}

// RemoveHymnIDs removes the "hymns" edge to Hymn entities by IDs.
func (su *StudentUpdate) RemoveHymnIDs(ids ...int64) *StudentUpdate {
	su.mutation.RemoveHymnIDs(ids...)
	return su
}

// RemoveHymns removes "hymns" edges to Hymn entities.
func (su *StudentUpdate) RemoveHymns(h ...*Hymn) *StudentUpdate {
	ids := make([]int64, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return su.RemoveHymnIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StudentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StudentUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StudentUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StudentUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *StudentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(student.Table, student.Columns, sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt64))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.LoginAccount(); ok {
		_spec.SetField(student.FieldLoginAccount, field.TypeString, value)
	}
	if value, ok := su.mutation.Password(); ok {
		_spec.SetField(student.FieldPassword, field.TypeString, value)
	}
	if value, ok := su.mutation.Username(); ok {
		_spec.SetField(student.FieldUsername, field.TypeString, value)
	}
	if value, ok := su.mutation.DateOfBirth(); ok {
		_spec.SetField(student.FieldDateOfBirth, field.TypeString, value)
	}
	if value, ok := su.mutation.Email(); ok {
		_spec.SetField(student.FieldEmail, field.TypeString, value)
	}
	if value, ok := su.mutation.UpdatedTime(); ok {
		_spec.SetField(student.FieldUpdatedTime, field.TypeTime, value)
	}
	if value, ok := su.mutation.VisibleFlg(); ok {
		_spec.SetField(student.FieldVisibleFlg, field.TypeBool, value)
	}
	if su.mutation.HymnsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.HymnsTable,
			Columns: []string{student.HymnsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hymn.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedHymnsIDs(); len(nodes) > 0 && !su.mutation.HymnsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.HymnsTable,
			Columns: []string{student.HymnsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hymn.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.HymnsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.HymnsTable,
			Columns: []string{student.HymnsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hymn.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{student.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// StudentUpdateOne is the builder for updating a single Student entity.
type StudentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StudentMutation
}

// SetLoginAccount sets the "login_account" field.
func (suo *StudentUpdateOne) SetLoginAccount(s string) *StudentUpdateOne {
	suo.mutation.SetLoginAccount(s)
	return suo
}

// SetNillableLoginAccount sets the "login_account" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableLoginAccount(s *string) *StudentUpdateOne {
	if s != nil {
		suo.SetLoginAccount(*s)
	}
	return suo
}

// SetPassword sets the "password" field.
func (suo *StudentUpdateOne) SetPassword(s string) *StudentUpdateOne {
	suo.mutation.SetPassword(s)
	return suo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillablePassword(s *string) *StudentUpdateOne {
	if s != nil {
		suo.SetPassword(*s)
	}
	return suo
}

// SetUsername sets the "username" field.
func (suo *StudentUpdateOne) SetUsername(s string) *StudentUpdateOne {
	suo.mutation.SetUsername(s)
	return suo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableUsername(s *string) *StudentUpdateOne {
	if s != nil {
		suo.SetUsername(*s)
	}
	return suo
}

// SetDateOfBirth sets the "date_of_birth" field.
func (suo *StudentUpdateOne) SetDateOfBirth(s string) *StudentUpdateOne {
	suo.mutation.SetDateOfBirth(s)
	return suo
}

// SetNillableDateOfBirth sets the "date_of_birth" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableDateOfBirth(s *string) *StudentUpdateOne {
	if s != nil {
		suo.SetDateOfBirth(*s)
	}
	return suo
}

// SetEmail sets the "email" field.
func (suo *StudentUpdateOne) SetEmail(s string) *StudentUpdateOne {
	suo.mutation.SetEmail(s)
	return suo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableEmail(s *string) *StudentUpdateOne {
	if s != nil {
		suo.SetEmail(*s)
	}
	return suo
}

// SetUpdatedTime sets the "updated_time" field.
func (suo *StudentUpdateOne) SetUpdatedTime(t time.Time) *StudentUpdateOne {
	suo.mutation.SetUpdatedTime(t)
	return suo
}

// SetNillableUpdatedTime sets the "updated_time" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableUpdatedTime(t *time.Time) *StudentUpdateOne {
	if t != nil {
		suo.SetUpdatedTime(*t)
	}
	return suo
}

// SetVisibleFlg sets the "visible_flg" field.
func (suo *StudentUpdateOne) SetVisibleFlg(b bool) *StudentUpdateOne {
	suo.mutation.SetVisibleFlg(b)
	return suo
}

// SetNillableVisibleFlg sets the "visible_flg" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableVisibleFlg(b *bool) *StudentUpdateOne {
	if b != nil {
		suo.SetVisibleFlg(*b)
	}
	return suo
}

// AddHymnIDs adds the "hymns" edge to the Hymn entity by IDs.
func (suo *StudentUpdateOne) AddHymnIDs(ids ...int64) *StudentUpdateOne {
	suo.mutation.AddHymnIDs(ids...)
	return suo
}

// AddHymns adds the "hymns" edges to the Hymn entity.
func (suo *StudentUpdateOne) AddHymns(h ...*Hymn) *StudentUpdateOne {
	ids := make([]int64, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return suo.AddHymnIDs(ids...)
}

// Mutation returns the StudentMutation object of the builder.
func (suo *StudentUpdateOne) Mutation() *StudentMutation {
	return suo.mutation
}

// ClearHymns clears all "hymns" edges to the Hymn entity.
func (suo *StudentUpdateOne) ClearHymns() *StudentUpdateOne {
	suo.mutation.ClearHymns()
	return suo
}

// RemoveHymnIDs removes the "hymns" edge to Hymn entities by IDs.
func (suo *StudentUpdateOne) RemoveHymnIDs(ids ...int64) *StudentUpdateOne {
	suo.mutation.RemoveHymnIDs(ids...)
	return suo
}

// RemoveHymns removes "hymns" edges to Hymn entities.
func (suo *StudentUpdateOne) RemoveHymns(h ...*Hymn) *StudentUpdateOne {
	ids := make([]int64, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return suo.RemoveHymnIDs(ids...)
}

// Where appends a list predicates to the StudentUpdate builder.
func (suo *StudentUpdateOne) Where(ps ...predicate.Student) *StudentUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StudentUpdateOne) Select(field string, fields ...string) *StudentUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Student entity.
func (suo *StudentUpdateOne) Save(ctx context.Context) (*Student, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StudentUpdateOne) SaveX(ctx context.Context) *Student {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StudentUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StudentUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *StudentUpdateOne) sqlSave(ctx context.Context) (_node *Student, err error) {
	_spec := sqlgraph.NewUpdateSpec(student.Table, student.Columns, sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt64))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Student.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, student.FieldID)
		for _, f := range fields {
			if !student.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != student.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.LoginAccount(); ok {
		_spec.SetField(student.FieldLoginAccount, field.TypeString, value)
	}
	if value, ok := suo.mutation.Password(); ok {
		_spec.SetField(student.FieldPassword, field.TypeString, value)
	}
	if value, ok := suo.mutation.Username(); ok {
		_spec.SetField(student.FieldUsername, field.TypeString, value)
	}
	if value, ok := suo.mutation.DateOfBirth(); ok {
		_spec.SetField(student.FieldDateOfBirth, field.TypeString, value)
	}
	if value, ok := suo.mutation.Email(); ok {
		_spec.SetField(student.FieldEmail, field.TypeString, value)
	}
	if value, ok := suo.mutation.UpdatedTime(); ok {
		_spec.SetField(student.FieldUpdatedTime, field.TypeTime, value)
	}
	if value, ok := suo.mutation.VisibleFlg(); ok {
		_spec.SetField(student.FieldVisibleFlg, field.TypeBool, value)
	}
	if suo.mutation.HymnsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.HymnsTable,
			Columns: []string{student.HymnsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hymn.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedHymnsIDs(); len(nodes) > 0 && !suo.mutation.HymnsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.HymnsTable,
			Columns: []string{student.HymnsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hymn.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.HymnsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.HymnsTable,
			Columns: []string{student.HymnsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hymn.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Student{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{student.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
