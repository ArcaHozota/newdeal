// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"newdeal/ent/hymn"
	"newdeal/ent/student"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// StudentCreate is the builder for creating a Student entity.
type StudentCreate struct {
	config
	mutation *StudentMutation
	hooks    []Hook
}

// SetLoginAccount sets the "login_account" field.
func (sc *StudentCreate) SetLoginAccount(s string) *StudentCreate {
	sc.mutation.SetLoginAccount(s)
	return sc
}

// SetPassword sets the "password" field.
func (sc *StudentCreate) SetPassword(s string) *StudentCreate {
	sc.mutation.SetPassword(s)
	return sc
}

// SetUsername sets the "username" field.
func (sc *StudentCreate) SetUsername(s string) *StudentCreate {
	sc.mutation.SetUsername(s)
	return sc
}

// SetDateOfBirth sets the "date_of_birth" field.
func (sc *StudentCreate) SetDateOfBirth(t time.Time) *StudentCreate {
	sc.mutation.SetDateOfBirth(t)
	return sc
}

// SetEmail sets the "email" field.
func (sc *StudentCreate) SetEmail(s string) *StudentCreate {
	sc.mutation.SetEmail(s)
	return sc
}

// SetUpdatedTime sets the "updated_time" field.
func (sc *StudentCreate) SetUpdatedTime(t time.Time) *StudentCreate {
	sc.mutation.SetUpdatedTime(t)
	return sc
}

// SetVisibleFlg sets the "visible_flg" field.
func (sc *StudentCreate) SetVisibleFlg(b bool) *StudentCreate {
	sc.mutation.SetVisibleFlg(b)
	return sc
}

// SetID sets the "id" field.
func (sc *StudentCreate) SetID(u uuid.UUID) *StudentCreate {
	sc.mutation.SetID(u)
	return sc
}

// AddHymnIDs adds the "hymns" edge to the Hymn entity by IDs.
func (sc *StudentCreate) AddHymnIDs(ids ...uuid.UUID) *StudentCreate {
	sc.mutation.AddHymnIDs(ids...)
	return sc
}

// AddHymns adds the "hymns" edges to the Hymn entity.
func (sc *StudentCreate) AddHymns(h ...*Hymn) *StudentCreate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return sc.AddHymnIDs(ids...)
}

// Mutation returns the StudentMutation object of the builder.
func (sc *StudentCreate) Mutation() *StudentMutation {
	return sc.mutation
}

// Save creates the Student in the database.
func (sc *StudentCreate) Save(ctx context.Context) (*Student, error) {
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StudentCreate) SaveX(ctx context.Context) *Student {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StudentCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StudentCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StudentCreate) check() error {
	if _, ok := sc.mutation.LoginAccount(); !ok {
		return &ValidationError{Name: "login_account", err: errors.New(`ent: missing required field "Student.login_account"`)}
	}
	if _, ok := sc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Student.password"`)}
	}
	if _, ok := sc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "Student.username"`)}
	}
	if _, ok := sc.mutation.DateOfBirth(); !ok {
		return &ValidationError{Name: "date_of_birth", err: errors.New(`ent: missing required field "Student.date_of_birth"`)}
	}
	if _, ok := sc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Student.email"`)}
	}
	if _, ok := sc.mutation.UpdatedTime(); !ok {
		return &ValidationError{Name: "updated_time", err: errors.New(`ent: missing required field "Student.updated_time"`)}
	}
	if _, ok := sc.mutation.VisibleFlg(); !ok {
		return &ValidationError{Name: "visible_flg", err: errors.New(`ent: missing required field "Student.visible_flg"`)}
	}
	return nil
}

func (sc *StudentCreate) sqlSave(ctx context.Context) (*Student, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *StudentCreate) createSpec() (*Student, *sqlgraph.CreateSpec) {
	var (
		_node = &Student{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(student.Table, sqlgraph.NewFieldSpec(student.FieldID, field.TypeUUID))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.LoginAccount(); ok {
		_spec.SetField(student.FieldLoginAccount, field.TypeString, value)
		_node.LoginAccount = value
	}
	if value, ok := sc.mutation.Password(); ok {
		_spec.SetField(student.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := sc.mutation.Username(); ok {
		_spec.SetField(student.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := sc.mutation.DateOfBirth(); ok {
		_spec.SetField(student.FieldDateOfBirth, field.TypeTime, value)
		_node.DateOfBirth = value
	}
	if value, ok := sc.mutation.Email(); ok {
		_spec.SetField(student.FieldEmail, field.TypeString, value)
		_node.Email = &value
	}
	if value, ok := sc.mutation.UpdatedTime(); ok {
		_spec.SetField(student.FieldUpdatedTime, field.TypeTime, value)
		_node.UpdatedTime = &value
	}
	if value, ok := sc.mutation.VisibleFlg(); ok {
		_spec.SetField(student.FieldVisibleFlg, field.TypeBool, value)
		_node.VisibleFlg = value
	}
	if nodes := sc.mutation.HymnsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.HymnsTable,
			Columns: []string{student.HymnsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hymn.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StudentCreateBulk is the builder for creating many Student entities in bulk.
type StudentCreateBulk struct {
	config
	err      error
	builders []*StudentCreate
}

// Save creates the Student entities in the database.
func (scb *StudentCreateBulk) Save(ctx context.Context) ([]*Student, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Student, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StudentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StudentCreateBulk) SaveX(ctx context.Context) []*Student {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StudentCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StudentCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
