// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"newdeal/ent/book"
	"newdeal/ent/chapter"
	"newdeal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BookUpdate is the builder for updating Book entities.
type BookUpdate struct {
	config
	hooks    []Hook
	mutation *BookMutation
}

// Where appends a list predicates to the BookUpdate builder.
func (bu *BookUpdate) Where(ps ...predicate.Book) *BookUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetName sets the "name" field.
func (bu *BookUpdate) SetName(s string) *BookUpdate {
	bu.mutation.SetName(s)
	return bu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (bu *BookUpdate) SetNillableName(s *string) *BookUpdate {
	if s != nil {
		bu.SetName(*s)
	}
	return bu
}

// SetNameJp sets the "name_jp" field.
func (bu *BookUpdate) SetNameJp(s string) *BookUpdate {
	bu.mutation.SetNameJp(s)
	return bu
}

// SetNillableNameJp sets the "name_jp" field if the given value is not nil.
func (bu *BookUpdate) SetNillableNameJp(s *string) *BookUpdate {
	if s != nil {
		bu.SetNameJp(*s)
	}
	return bu
}

// AddToChapterIDs adds the "to_chapter" edge to the Chapter entity by IDs.
func (bu *BookUpdate) AddToChapterIDs(ids ...int32) *BookUpdate {
	bu.mutation.AddToChapterIDs(ids...)
	return bu
}

// AddToChapter adds the "to_chapter" edges to the Chapter entity.
func (bu *BookUpdate) AddToChapter(c ...*Chapter) *BookUpdate {
	ids := make([]int32, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bu.AddToChapterIDs(ids...)
}

// Mutation returns the BookMutation object of the builder.
func (bu *BookUpdate) Mutation() *BookMutation {
	return bu.mutation
}

// ClearToChapter clears all "to_chapter" edges to the Chapter entity.
func (bu *BookUpdate) ClearToChapter() *BookUpdate {
	bu.mutation.ClearToChapter()
	return bu
}

// RemoveToChapterIDs removes the "to_chapter" edge to Chapter entities by IDs.
func (bu *BookUpdate) RemoveToChapterIDs(ids ...int32) *BookUpdate {
	bu.mutation.RemoveToChapterIDs(ids...)
	return bu
}

// RemoveToChapter removes "to_chapter" edges to Chapter entities.
func (bu *BookUpdate) RemoveToChapter(c ...*Chapter) *BookUpdate {
	ids := make([]int32, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bu.RemoveToChapterIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BookUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BookUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BookUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BookUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BookUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(book.Table, book.Columns, sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt16))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Name(); ok {
		_spec.SetField(book.FieldName, field.TypeString, value)
	}
	if value, ok := bu.mutation.NameJp(); ok {
		_spec.SetField(book.FieldNameJp, field.TypeString, value)
	}
	if bu.mutation.ToChapterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   book.ToChapterTable,
			Columns: []string{book.ToChapterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chapter.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedToChapterIDs(); len(nodes) > 0 && !bu.mutation.ToChapterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   book.ToChapterTable,
			Columns: []string{book.ToChapterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chapter.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.ToChapterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   book.ToChapterTable,
			Columns: []string{book.ToChapterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chapter.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{book.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BookUpdateOne is the builder for updating a single Book entity.
type BookUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookMutation
}

// SetName sets the "name" field.
func (buo *BookUpdateOne) SetName(s string) *BookUpdateOne {
	buo.mutation.SetName(s)
	return buo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableName(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetName(*s)
	}
	return buo
}

// SetNameJp sets the "name_jp" field.
func (buo *BookUpdateOne) SetNameJp(s string) *BookUpdateOne {
	buo.mutation.SetNameJp(s)
	return buo
}

// SetNillableNameJp sets the "name_jp" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableNameJp(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetNameJp(*s)
	}
	return buo
}

// AddToChapterIDs adds the "to_chapter" edge to the Chapter entity by IDs.
func (buo *BookUpdateOne) AddToChapterIDs(ids ...int32) *BookUpdateOne {
	buo.mutation.AddToChapterIDs(ids...)
	return buo
}

// AddToChapter adds the "to_chapter" edges to the Chapter entity.
func (buo *BookUpdateOne) AddToChapter(c ...*Chapter) *BookUpdateOne {
	ids := make([]int32, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return buo.AddToChapterIDs(ids...)
}

// Mutation returns the BookMutation object of the builder.
func (buo *BookUpdateOne) Mutation() *BookMutation {
	return buo.mutation
}

// ClearToChapter clears all "to_chapter" edges to the Chapter entity.
func (buo *BookUpdateOne) ClearToChapter() *BookUpdateOne {
	buo.mutation.ClearToChapter()
	return buo
}

// RemoveToChapterIDs removes the "to_chapter" edge to Chapter entities by IDs.
func (buo *BookUpdateOne) RemoveToChapterIDs(ids ...int32) *BookUpdateOne {
	buo.mutation.RemoveToChapterIDs(ids...)
	return buo
}

// RemoveToChapter removes "to_chapter" edges to Chapter entities.
func (buo *BookUpdateOne) RemoveToChapter(c ...*Chapter) *BookUpdateOne {
	ids := make([]int32, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return buo.RemoveToChapterIDs(ids...)
}

// Where appends a list predicates to the BookUpdate builder.
func (buo *BookUpdateOne) Where(ps ...predicate.Book) *BookUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BookUpdateOne) Select(field string, fields ...string) *BookUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Book entity.
func (buo *BookUpdateOne) Save(ctx context.Context) (*Book, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BookUpdateOne) SaveX(ctx context.Context) *Book {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BookUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BookUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BookUpdateOne) sqlSave(ctx context.Context) (_node *Book, err error) {
	_spec := sqlgraph.NewUpdateSpec(book.Table, book.Columns, sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt16))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Book.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, book.FieldID)
		for _, f := range fields {
			if !book.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != book.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Name(); ok {
		_spec.SetField(book.FieldName, field.TypeString, value)
	}
	if value, ok := buo.mutation.NameJp(); ok {
		_spec.SetField(book.FieldNameJp, field.TypeString, value)
	}
	if buo.mutation.ToChapterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   book.ToChapterTable,
			Columns: []string{book.ToChapterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chapter.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedToChapterIDs(); len(nodes) > 0 && !buo.mutation.ToChapterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   book.ToChapterTable,
			Columns: []string{book.ToChapterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chapter.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.ToChapterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   book.ToChapterTable,
			Columns: []string{book.ToChapterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chapter.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Book{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{book.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
