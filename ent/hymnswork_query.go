// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"newdeal/ent/hymn"
	"newdeal/ent/hymnswork"
	"newdeal/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// HymnsWorkQuery is the builder for querying HymnsWork entities.
type HymnsWorkQuery struct {
	config
	ctx        *QueryContext
	order      []hymnswork.OrderOption
	inters     []Interceptor
	predicates []predicate.HymnsWork
	withHymns  *HymnQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the HymnsWorkQuery builder.
func (hwq *HymnsWorkQuery) Where(ps ...predicate.HymnsWork) *HymnsWorkQuery {
	hwq.predicates = append(hwq.predicates, ps...)
	return hwq
}

// Limit the number of records to be returned by this query.
func (hwq *HymnsWorkQuery) Limit(limit int) *HymnsWorkQuery {
	hwq.ctx.Limit = &limit
	return hwq
}

// Offset to start from.
func (hwq *HymnsWorkQuery) Offset(offset int) *HymnsWorkQuery {
	hwq.ctx.Offset = &offset
	return hwq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (hwq *HymnsWorkQuery) Unique(unique bool) *HymnsWorkQuery {
	hwq.ctx.Unique = &unique
	return hwq
}

// Order specifies how the records should be ordered.
func (hwq *HymnsWorkQuery) Order(o ...hymnswork.OrderOption) *HymnsWorkQuery {
	hwq.order = append(hwq.order, o...)
	return hwq
}

// QueryHymns chains the current query on the "hymns" edge.
func (hwq *HymnsWorkQuery) QueryHymns() *HymnQuery {
	query := (&HymnClient{config: hwq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hwq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hwq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(hymnswork.Table, hymnswork.FieldID, selector),
			sqlgraph.To(hymn.Table, hymn.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, hymnswork.HymnsTable, hymnswork.HymnsColumn),
		)
		fromU = sqlgraph.SetNeighbors(hwq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first HymnsWork entity from the query.
// Returns a *NotFoundError when no HymnsWork was found.
func (hwq *HymnsWorkQuery) First(ctx context.Context) (*HymnsWork, error) {
	nodes, err := hwq.Limit(1).All(setContextOp(ctx, hwq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{hymnswork.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (hwq *HymnsWorkQuery) FirstX(ctx context.Context) *HymnsWork {
	node, err := hwq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first HymnsWork ID from the query.
// Returns a *NotFoundError when no HymnsWork ID was found.
func (hwq *HymnsWorkQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hwq.Limit(1).IDs(setContextOp(ctx, hwq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{hymnswork.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (hwq *HymnsWorkQuery) FirstIDX(ctx context.Context) int {
	id, err := hwq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single HymnsWork entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one HymnsWork entity is found.
// Returns a *NotFoundError when no HymnsWork entities are found.
func (hwq *HymnsWorkQuery) Only(ctx context.Context) (*HymnsWork, error) {
	nodes, err := hwq.Limit(2).All(setContextOp(ctx, hwq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{hymnswork.Label}
	default:
		return nil, &NotSingularError{hymnswork.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (hwq *HymnsWorkQuery) OnlyX(ctx context.Context) *HymnsWork {
	node, err := hwq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only HymnsWork ID in the query.
// Returns a *NotSingularError when more than one HymnsWork ID is found.
// Returns a *NotFoundError when no entities are found.
func (hwq *HymnsWorkQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hwq.Limit(2).IDs(setContextOp(ctx, hwq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{hymnswork.Label}
	default:
		err = &NotSingularError{hymnswork.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (hwq *HymnsWorkQuery) OnlyIDX(ctx context.Context) int {
	id, err := hwq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of HymnsWorks.
func (hwq *HymnsWorkQuery) All(ctx context.Context) ([]*HymnsWork, error) {
	ctx = setContextOp(ctx, hwq.ctx, ent.OpQueryAll)
	if err := hwq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*HymnsWork, *HymnsWorkQuery]()
	return withInterceptors[[]*HymnsWork](ctx, hwq, qr, hwq.inters)
}

// AllX is like All, but panics if an error occurs.
func (hwq *HymnsWorkQuery) AllX(ctx context.Context) []*HymnsWork {
	nodes, err := hwq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of HymnsWork IDs.
func (hwq *HymnsWorkQuery) IDs(ctx context.Context) (ids []int, err error) {
	if hwq.ctx.Unique == nil && hwq.path != nil {
		hwq.Unique(true)
	}
	ctx = setContextOp(ctx, hwq.ctx, ent.OpQueryIDs)
	if err = hwq.Select(hymnswork.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (hwq *HymnsWorkQuery) IDsX(ctx context.Context) []int {
	ids, err := hwq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (hwq *HymnsWorkQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, hwq.ctx, ent.OpQueryCount)
	if err := hwq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, hwq, querierCount[*HymnsWorkQuery](), hwq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (hwq *HymnsWorkQuery) CountX(ctx context.Context) int {
	count, err := hwq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (hwq *HymnsWorkQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, hwq.ctx, ent.OpQueryExist)
	switch _, err := hwq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (hwq *HymnsWorkQuery) ExistX(ctx context.Context) bool {
	exist, err := hwq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the HymnsWorkQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (hwq *HymnsWorkQuery) Clone() *HymnsWorkQuery {
	if hwq == nil {
		return nil
	}
	return &HymnsWorkQuery{
		config:     hwq.config,
		ctx:        hwq.ctx.Clone(),
		order:      append([]hymnswork.OrderOption{}, hwq.order...),
		inters:     append([]Interceptor{}, hwq.inters...),
		predicates: append([]predicate.HymnsWork{}, hwq.predicates...),
		withHymns:  hwq.withHymns.Clone(),
		// clone intermediate query.
		sql:  hwq.sql.Clone(),
		path: hwq.path,
	}
}

// WithHymns tells the query-builder to eager-load the nodes that are connected to
// the "hymns" edge. The optional arguments are used to configure the query builder of the edge.
func (hwq *HymnsWorkQuery) WithHymns(opts ...func(*HymnQuery)) *HymnsWorkQuery {
	query := (&HymnClient{config: hwq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	hwq.withHymns = query
	return hwq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		WorkID uuid.UUID `json:"work_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.HymnsWork.Query().
//		GroupBy(hymnswork.FieldWorkID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (hwq *HymnsWorkQuery) GroupBy(field string, fields ...string) *HymnsWorkGroupBy {
	hwq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &HymnsWorkGroupBy{build: hwq}
	grbuild.flds = &hwq.ctx.Fields
	grbuild.label = hymnswork.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		WorkID uuid.UUID `json:"work_id,omitempty"`
//	}
//
//	client.HymnsWork.Query().
//		Select(hymnswork.FieldWorkID).
//		Scan(ctx, &v)
func (hwq *HymnsWorkQuery) Select(fields ...string) *HymnsWorkSelect {
	hwq.ctx.Fields = append(hwq.ctx.Fields, fields...)
	sbuild := &HymnsWorkSelect{HymnsWorkQuery: hwq}
	sbuild.label = hymnswork.Label
	sbuild.flds, sbuild.scan = &hwq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a HymnsWorkSelect configured with the given aggregations.
func (hwq *HymnsWorkQuery) Aggregate(fns ...AggregateFunc) *HymnsWorkSelect {
	return hwq.Select().Aggregate(fns...)
}

func (hwq *HymnsWorkQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range hwq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, hwq); err != nil {
				return err
			}
		}
	}
	for _, f := range hwq.ctx.Fields {
		if !hymnswork.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if hwq.path != nil {
		prev, err := hwq.path(ctx)
		if err != nil {
			return err
		}
		hwq.sql = prev
	}
	return nil
}

func (hwq *HymnsWorkQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*HymnsWork, error) {
	var (
		nodes       = []*HymnsWork{}
		withFKs     = hwq.withFKs
		_spec       = hwq.querySpec()
		loadedTypes = [1]bool{
			hwq.withHymns != nil,
		}
	)
	if hwq.withHymns != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, hymnswork.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*HymnsWork).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &HymnsWork{config: hwq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, hwq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := hwq.withHymns; query != nil {
		if err := hwq.loadHymns(ctx, query, nodes, nil,
			func(n *HymnsWork, e *Hymn) { n.Edges.Hymns = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (hwq *HymnsWorkQuery) loadHymns(ctx context.Context, query *HymnQuery, nodes []*HymnsWork, init func(*HymnsWork), assign func(*HymnsWork, *Hymn)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*HymnsWork)
	for i := range nodes {
		if nodes[i].hymn_hymns_work == nil {
			continue
		}
		fk := *nodes[i].hymn_hymns_work
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(hymn.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "hymn_hymns_work" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (hwq *HymnsWorkQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := hwq.querySpec()
	_spec.Node.Columns = hwq.ctx.Fields
	if len(hwq.ctx.Fields) > 0 {
		_spec.Unique = hwq.ctx.Unique != nil && *hwq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, hwq.driver, _spec)
}

func (hwq *HymnsWorkQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(hymnswork.Table, hymnswork.Columns, sqlgraph.NewFieldSpec(hymnswork.FieldID, field.TypeInt))
	_spec.From = hwq.sql
	if unique := hwq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if hwq.path != nil {
		_spec.Unique = true
	}
	if fields := hwq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hymnswork.FieldID)
		for i := range fields {
			if fields[i] != hymnswork.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := hwq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := hwq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := hwq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := hwq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (hwq *HymnsWorkQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(hwq.driver.Dialect())
	t1 := builder.Table(hymnswork.Table)
	columns := hwq.ctx.Fields
	if len(columns) == 0 {
		columns = hymnswork.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if hwq.sql != nil {
		selector = hwq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if hwq.ctx.Unique != nil && *hwq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range hwq.predicates {
		p(selector)
	}
	for _, p := range hwq.order {
		p(selector)
	}
	if offset := hwq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := hwq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// HymnsWorkGroupBy is the group-by builder for HymnsWork entities.
type HymnsWorkGroupBy struct {
	selector
	build *HymnsWorkQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hwgb *HymnsWorkGroupBy) Aggregate(fns ...AggregateFunc) *HymnsWorkGroupBy {
	hwgb.fns = append(hwgb.fns, fns...)
	return hwgb
}

// Scan applies the selector query and scans the result into the given value.
func (hwgb *HymnsWorkGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hwgb.build.ctx, ent.OpQueryGroupBy)
	if err := hwgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HymnsWorkQuery, *HymnsWorkGroupBy](ctx, hwgb.build, hwgb, hwgb.build.inters, v)
}

func (hwgb *HymnsWorkGroupBy) sqlScan(ctx context.Context, root *HymnsWorkQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(hwgb.fns))
	for _, fn := range hwgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*hwgb.flds)+len(hwgb.fns))
		for _, f := range *hwgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*hwgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hwgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// HymnsWorkSelect is the builder for selecting fields of HymnsWork entities.
type HymnsWorkSelect struct {
	*HymnsWorkQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (hws *HymnsWorkSelect) Aggregate(fns ...AggregateFunc) *HymnsWorkSelect {
	hws.fns = append(hws.fns, fns...)
	return hws
}

// Scan applies the selector query and scans the result into the given value.
func (hws *HymnsWorkSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hws.ctx, ent.OpQuerySelect)
	if err := hws.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HymnsWorkQuery, *HymnsWorkSelect](ctx, hws.HymnsWorkQuery, hws, hws.inters, v)
}

func (hws *HymnsWorkSelect) sqlScan(ctx context.Context, root *HymnsWorkQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(hws.fns))
	for _, fn := range hws.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*hws.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
