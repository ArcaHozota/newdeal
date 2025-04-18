// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"newdeal/ent/hymn"
	"newdeal/ent/hymnswork"
	"newdeal/ent/student"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// HymnCreate is the builder for creating a Hymn entity.
type HymnCreate struct {
	config
	mutation *HymnMutation
	hooks    []Hook
}

// SetNameJp sets the "name_jp" field.
func (hc *HymnCreate) SetNameJp(s string) *HymnCreate {
	hc.mutation.SetNameJp(s)
	return hc
}

// SetNameKr sets the "name_kr" field.
func (hc *HymnCreate) SetNameKr(s string) *HymnCreate {
	hc.mutation.SetNameKr(s)
	return hc
}

// SetLink sets the "link" field.
func (hc *HymnCreate) SetLink(s string) *HymnCreate {
	hc.mutation.SetLink(s)
	return hc
}

// SetUpdatedTime sets the "updated_time" field.
func (hc *HymnCreate) SetUpdatedTime(t time.Time) *HymnCreate {
	hc.mutation.SetUpdatedTime(t)
	return hc
}

// SetUpdatedUser sets the "updated_user" field.
func (hc *HymnCreate) SetUpdatedUser(i int64) *HymnCreate {
	hc.mutation.SetUpdatedUser(i)
	return hc
}

// SetSerif sets the "serif" field.
func (hc *HymnCreate) SetSerif(s string) *HymnCreate {
	hc.mutation.SetSerif(s)
	return hc
}

// SetVisibleFlg sets the "visible_flg" field.
func (hc *HymnCreate) SetVisibleFlg(b bool) *HymnCreate {
	hc.mutation.SetVisibleFlg(b)
	return hc
}

// SetID sets the "id" field.
func (hc *HymnCreate) SetID(u uuid.UUID) *HymnCreate {
	hc.mutation.SetID(u)
	return hc
}

// SetStudentsID sets the "students" edge to the Student entity by ID.
func (hc *HymnCreate) SetStudentsID(id uuid.UUID) *HymnCreate {
	hc.mutation.SetStudentsID(id)
	return hc
}

// SetNillableStudentsID sets the "students" edge to the Student entity by ID if the given value is not nil.
func (hc *HymnCreate) SetNillableStudentsID(id *uuid.UUID) *HymnCreate {
	if id != nil {
		hc = hc.SetStudentsID(*id)
	}
	return hc
}

// SetStudents sets the "students" edge to the Student entity.
func (hc *HymnCreate) SetStudents(s *Student) *HymnCreate {
	return hc.SetStudentsID(s.ID)
}

// SetHymnsWorkID sets the "hymns_work" edge to the HymnsWork entity by ID.
func (hc *HymnCreate) SetHymnsWorkID(id int) *HymnCreate {
	hc.mutation.SetHymnsWorkID(id)
	return hc
}

// SetNillableHymnsWorkID sets the "hymns_work" edge to the HymnsWork entity by ID if the given value is not nil.
func (hc *HymnCreate) SetNillableHymnsWorkID(id *int) *HymnCreate {
	if id != nil {
		hc = hc.SetHymnsWorkID(*id)
	}
	return hc
}

// SetHymnsWork sets the "hymns_work" edge to the HymnsWork entity.
func (hc *HymnCreate) SetHymnsWork(h *HymnsWork) *HymnCreate {
	return hc.SetHymnsWorkID(h.ID)
}

// Mutation returns the HymnMutation object of the builder.
func (hc *HymnCreate) Mutation() *HymnMutation {
	return hc.mutation
}

// Save creates the Hymn in the database.
func (hc *HymnCreate) Save(ctx context.Context) (*Hymn, error) {
	return withHooks(ctx, hc.sqlSave, hc.mutation, hc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HymnCreate) SaveX(ctx context.Context) *Hymn {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hc *HymnCreate) Exec(ctx context.Context) error {
	_, err := hc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hc *HymnCreate) ExecX(ctx context.Context) {
	if err := hc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hc *HymnCreate) check() error {
	if _, ok := hc.mutation.NameJp(); !ok {
		return &ValidationError{Name: "name_jp", err: errors.New(`ent: missing required field "Hymn.name_jp"`)}
	}
	if _, ok := hc.mutation.NameKr(); !ok {
		return &ValidationError{Name: "name_kr", err: errors.New(`ent: missing required field "Hymn.name_kr"`)}
	}
	if _, ok := hc.mutation.Link(); !ok {
		return &ValidationError{Name: "link", err: errors.New(`ent: missing required field "Hymn.link"`)}
	}
	if _, ok := hc.mutation.UpdatedTime(); !ok {
		return &ValidationError{Name: "updated_time", err: errors.New(`ent: missing required field "Hymn.updated_time"`)}
	}
	if _, ok := hc.mutation.UpdatedUser(); !ok {
		return &ValidationError{Name: "updated_user", err: errors.New(`ent: missing required field "Hymn.updated_user"`)}
	}
	if v, ok := hc.mutation.UpdatedUser(); ok {
		if err := hymn.UpdatedUserValidator(v); err != nil {
			return &ValidationError{Name: "updated_user", err: fmt.Errorf(`ent: validator failed for field "Hymn.updated_user": %w`, err)}
		}
	}
	if _, ok := hc.mutation.Serif(); !ok {
		return &ValidationError{Name: "serif", err: errors.New(`ent: missing required field "Hymn.serif"`)}
	}
	if _, ok := hc.mutation.VisibleFlg(); !ok {
		return &ValidationError{Name: "visible_flg", err: errors.New(`ent: missing required field "Hymn.visible_flg"`)}
	}
	return nil
}

func (hc *HymnCreate) sqlSave(ctx context.Context) (*Hymn, error) {
	if err := hc.check(); err != nil {
		return nil, err
	}
	_node, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
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
	hc.mutation.id = &_node.ID
	hc.mutation.done = true
	return _node, nil
}

func (hc *HymnCreate) createSpec() (*Hymn, *sqlgraph.CreateSpec) {
	var (
		_node = &Hymn{config: hc.config}
		_spec = sqlgraph.NewCreateSpec(hymn.Table, sqlgraph.NewFieldSpec(hymn.FieldID, field.TypeUUID))
	)
	if id, ok := hc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := hc.mutation.NameJp(); ok {
		_spec.SetField(hymn.FieldNameJp, field.TypeString, value)
		_node.NameJp = value
	}
	if value, ok := hc.mutation.NameKr(); ok {
		_spec.SetField(hymn.FieldNameKr, field.TypeString, value)
		_node.NameKr = value
	}
	if value, ok := hc.mutation.Link(); ok {
		_spec.SetField(hymn.FieldLink, field.TypeString, value)
		_node.Link = &value
	}
	if value, ok := hc.mutation.UpdatedTime(); ok {
		_spec.SetField(hymn.FieldUpdatedTime, field.TypeTime, value)
		_node.UpdatedTime = value
	}
	if value, ok := hc.mutation.UpdatedUser(); ok {
		_spec.SetField(hymn.FieldUpdatedUser, field.TypeInt64, value)
		_node.UpdatedUser = value
	}
	if value, ok := hc.mutation.Serif(); ok {
		_spec.SetField(hymn.FieldSerif, field.TypeString, value)
		_node.Serif = &value
	}
	if value, ok := hc.mutation.VisibleFlg(); ok {
		_spec.SetField(hymn.FieldVisibleFlg, field.TypeBool, value)
		_node.VisibleFlg = value
	}
	if nodes := hc.mutation.StudentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hymn.StudentsTable,
			Columns: []string{hymn.StudentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.student_hymns = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.HymnsWorkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   hymn.HymnsWorkTable,
			Columns: []string{hymn.HymnsWorkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hymnswork.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HymnCreateBulk is the builder for creating many Hymn entities in bulk.
type HymnCreateBulk struct {
	config
	err      error
	builders []*HymnCreate
}

// Save creates the Hymn entities in the database.
func (hcb *HymnCreateBulk) Save(ctx context.Context) ([]*Hymn, error) {
	if hcb.err != nil {
		return nil, hcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*Hymn, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HymnMutation)
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
					_, err = mutators[i+1].Mutate(root, hcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, hcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hcb *HymnCreateBulk) SaveX(ctx context.Context) []*Hymn {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hcb *HymnCreateBulk) Exec(ctx context.Context) error {
	_, err := hcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcb *HymnCreateBulk) ExecX(ctx context.Context) {
	if err := hcb.Exec(ctx); err != nil {
		panic(err)
	}
}
