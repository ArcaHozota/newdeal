// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"newdeal/ent/chapter"
	"newdeal/ent/phrase"
	"newdeal/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PhraseQuery is the builder for querying Phrase entities.
type PhraseQuery struct {
	config
	ctx               *QueryContext
	order             []phrase.OrderOption
	inters            []Interceptor
	predicates        []predicate.Phrase
	withPhraseChapter *ChapterQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PhraseQuery builder.
func (pq *PhraseQuery) Where(ps ...predicate.Phrase) *PhraseQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PhraseQuery) Limit(limit int) *PhraseQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PhraseQuery) Offset(offset int) *PhraseQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PhraseQuery) Unique(unique bool) *PhraseQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PhraseQuery) Order(o ...phrase.OrderOption) *PhraseQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryPhraseChapter chains the current query on the "phrase_chapter" edge.
func (pq *PhraseQuery) QueryPhraseChapter() *ChapterQuery {
	query := (&ChapterClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(phrase.Table, phrase.FieldID, selector),
			sqlgraph.To(chapter.Table, chapter.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, phrase.PhraseChapterTable, phrase.PhraseChapterColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Phrase entity from the query.
// Returns a *NotFoundError when no Phrase was found.
func (pq *PhraseQuery) First(ctx context.Context) (*Phrase, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{phrase.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PhraseQuery) FirstX(ctx context.Context) *Phrase {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Phrase ID from the query.
// Returns a *NotFoundError when no Phrase ID was found.
func (pq *PhraseQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{phrase.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PhraseQuery) FirstIDX(ctx context.Context) int64 {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Phrase entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Phrase entity is found.
// Returns a *NotFoundError when no Phrase entities are found.
func (pq *PhraseQuery) Only(ctx context.Context) (*Phrase, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{phrase.Label}
	default:
		return nil, &NotSingularError{phrase.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PhraseQuery) OnlyX(ctx context.Context) *Phrase {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Phrase ID in the query.
// Returns a *NotSingularError when more than one Phrase ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PhraseQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{phrase.Label}
	default:
		err = &NotSingularError{phrase.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PhraseQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Phrases.
func (pq *PhraseQuery) All(ctx context.Context) ([]*Phrase, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryAll)
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Phrase, *PhraseQuery]()
	return withInterceptors[[]*Phrase](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PhraseQuery) AllX(ctx context.Context) []*Phrase {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Phrase IDs.
func (pq *PhraseQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryIDs)
	if err = pq.Select(phrase.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PhraseQuery) IDsX(ctx context.Context) []int64 {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PhraseQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryCount)
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PhraseQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PhraseQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PhraseQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryExist)
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PhraseQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PhraseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PhraseQuery) Clone() *PhraseQuery {
	if pq == nil {
		return nil
	}
	return &PhraseQuery{
		config:            pq.config,
		ctx:               pq.ctx.Clone(),
		order:             append([]phrase.OrderOption{}, pq.order...),
		inters:            append([]Interceptor{}, pq.inters...),
		predicates:        append([]predicate.Phrase{}, pq.predicates...),
		withPhraseChapter: pq.withPhraseChapter.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithPhraseChapter tells the query-builder to eager-load the nodes that are connected to
// the "phrase_chapter" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PhraseQuery) WithPhraseChapter(opts ...func(*ChapterQuery)) *PhraseQuery {
	query := (&ChapterClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withPhraseChapter = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Phrase.Query().
//		GroupBy(phrase.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PhraseQuery) GroupBy(field string, fields ...string) *PhraseGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PhraseGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = phrase.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Phrase.Query().
//		Select(phrase.FieldName).
//		Scan(ctx, &v)
func (pq *PhraseQuery) Select(fields ...string) *PhraseSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PhraseSelect{PhraseQuery: pq}
	sbuild.label = phrase.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PhraseSelect configured with the given aggregations.
func (pq *PhraseQuery) Aggregate(fns ...AggregateFunc) *PhraseSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PhraseQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !phrase.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PhraseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Phrase, error) {
	var (
		nodes       = []*Phrase{}
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withPhraseChapter != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Phrase).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Phrase{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withPhraseChapter; query != nil {
		if err := pq.loadPhraseChapter(ctx, query, nodes, nil,
			func(n *Phrase, e *Chapter) { n.Edges.PhraseChapter = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PhraseQuery) loadPhraseChapter(ctx context.Context, query *ChapterQuery, nodes []*Phrase, init func(*Phrase), assign func(*Phrase, *Chapter)) error {
	ids := make([]int32, 0, len(nodes))
	nodeids := make(map[int32][]*Phrase)
	for i := range nodes {
		fk := nodes[i].ChapterID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(chapter.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "chapter_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pq *PhraseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PhraseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(phrase.Table, phrase.Columns, sqlgraph.NewFieldSpec(phrase.FieldID, field.TypeInt64))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, phrase.FieldID)
		for i := range fields {
			if fields[i] != phrase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if pq.withPhraseChapter != nil {
			_spec.Node.AddColumnOnce(phrase.FieldChapterID)
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PhraseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(phrase.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = phrase.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PhraseGroupBy is the group-by builder for Phrase entities.
type PhraseGroupBy struct {
	selector
	build *PhraseQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PhraseGroupBy) Aggregate(fns ...AggregateFunc) *PhraseGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PhraseGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, ent.OpQueryGroupBy)
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PhraseQuery, *PhraseGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PhraseGroupBy) sqlScan(ctx context.Context, root *PhraseQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PhraseSelect is the builder for selecting fields of Phrase entities.
type PhraseSelect struct {
	*PhraseQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PhraseSelect) Aggregate(fns ...AggregateFunc) *PhraseSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PhraseSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, ent.OpQuerySelect)
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PhraseQuery, *PhraseSelect](ctx, ps.PhraseQuery, ps, ps.inters, v)
}

func (ps *PhraseSelect) sqlScan(ctx context.Context, root *PhraseQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
