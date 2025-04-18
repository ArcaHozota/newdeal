// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"newdeal/ent/hymn"
	"newdeal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HymnDelete is the builder for deleting a Hymn entity.
type HymnDelete struct {
	config
	hooks    []Hook
	mutation *HymnMutation
}

// Where appends a list predicates to the HymnDelete builder.
func (hd *HymnDelete) Where(ps ...predicate.Hymn) *HymnDelete {
	hd.mutation.Where(ps...)
	return hd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hd *HymnDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, hd.sqlExec, hd.mutation, hd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (hd *HymnDelete) ExecX(ctx context.Context) int {
	n, err := hd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hd *HymnDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(hymn.Table, sqlgraph.NewFieldSpec(hymn.FieldID, field.TypeOther))
	if ps := hd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, hd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	hd.mutation.done = true
	return affected, err
}

// HymnDeleteOne is the builder for deleting a single Hymn entity.
type HymnDeleteOne struct {
	hd *HymnDelete
}

// Where appends a list predicates to the HymnDelete builder.
func (hdo *HymnDeleteOne) Where(ps ...predicate.Hymn) *HymnDeleteOne {
	hdo.hd.mutation.Where(ps...)
	return hdo
}

// Exec executes the deletion query.
func (hdo *HymnDeleteOne) Exec(ctx context.Context) error {
	n, err := hdo.hd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{hymn.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hdo *HymnDeleteOne) ExecX(ctx context.Context) {
	if err := hdo.Exec(ctx); err != nil {
		panic(err)
	}
}
