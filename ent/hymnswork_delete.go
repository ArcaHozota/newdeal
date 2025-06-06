// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"newdeal/ent/hymnswork"
	"newdeal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HymnsWorkDelete is the builder for deleting a HymnsWork entity.
type HymnsWorkDelete struct {
	config
	hooks    []Hook
	mutation *HymnsWorkMutation
}

// Where appends a list predicates to the HymnsWorkDelete builder.
func (hwd *HymnsWorkDelete) Where(ps ...predicate.HymnsWork) *HymnsWorkDelete {
	hwd.mutation.Where(ps...)
	return hwd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hwd *HymnsWorkDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, hwd.sqlExec, hwd.mutation, hwd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (hwd *HymnsWorkDelete) ExecX(ctx context.Context) int {
	n, err := hwd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hwd *HymnsWorkDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(hymnswork.Table, sqlgraph.NewFieldSpec(hymnswork.FieldID, field.TypeInt64))
	if ps := hwd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, hwd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	hwd.mutation.done = true
	return affected, err
}

// HymnsWorkDeleteOne is the builder for deleting a single HymnsWork entity.
type HymnsWorkDeleteOne struct {
	hwd *HymnsWorkDelete
}

// Where appends a list predicates to the HymnsWorkDelete builder.
func (hwdo *HymnsWorkDeleteOne) Where(ps ...predicate.HymnsWork) *HymnsWorkDeleteOne {
	hwdo.hwd.mutation.Where(ps...)
	return hwdo
}

// Exec executes the deletion query.
func (hwdo *HymnsWorkDeleteOne) Exec(ctx context.Context) error {
	n, err := hwdo.hwd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{hymnswork.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hwdo *HymnsWorkDeleteOne) ExecX(ctx context.Context) {
	if err := hwdo.Exec(ctx); err != nil {
		panic(err)
	}
}
