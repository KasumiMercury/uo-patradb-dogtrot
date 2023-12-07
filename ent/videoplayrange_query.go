// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/predicate"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/video"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/videoplayrange"
)

// VideoPlayRangeQuery is the builder for querying VideoPlayRange entities.
type VideoPlayRangeQuery struct {
	config
	ctx        *QueryContext
	order      []videoplayrange.OrderOption
	inters     []Interceptor
	predicates []predicate.VideoPlayRange
	withVideo  *VideoQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VideoPlayRangeQuery builder.
func (vprq *VideoPlayRangeQuery) Where(ps ...predicate.VideoPlayRange) *VideoPlayRangeQuery {
	vprq.predicates = append(vprq.predicates, ps...)
	return vprq
}

// Limit the number of records to be returned by this query.
func (vprq *VideoPlayRangeQuery) Limit(limit int) *VideoPlayRangeQuery {
	vprq.ctx.Limit = &limit
	return vprq
}

// Offset to start from.
func (vprq *VideoPlayRangeQuery) Offset(offset int) *VideoPlayRangeQuery {
	vprq.ctx.Offset = &offset
	return vprq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vprq *VideoPlayRangeQuery) Unique(unique bool) *VideoPlayRangeQuery {
	vprq.ctx.Unique = &unique
	return vprq
}

// Order specifies how the records should be ordered.
func (vprq *VideoPlayRangeQuery) Order(o ...videoplayrange.OrderOption) *VideoPlayRangeQuery {
	vprq.order = append(vprq.order, o...)
	return vprq
}

// QueryVideo chains the current query on the "video" edge.
func (vprq *VideoPlayRangeQuery) QueryVideo() *VideoQuery {
	query := (&VideoClient{config: vprq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vprq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vprq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(videoplayrange.Table, videoplayrange.FieldID, selector),
			sqlgraph.To(video.Table, video.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, videoplayrange.VideoTable, videoplayrange.VideoColumn),
		)
		fromU = sqlgraph.SetNeighbors(vprq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first VideoPlayRange entity from the query.
// Returns a *NotFoundError when no VideoPlayRange was found.
func (vprq *VideoPlayRangeQuery) First(ctx context.Context) (*VideoPlayRange, error) {
	nodes, err := vprq.Limit(1).All(setContextOp(ctx, vprq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{videoplayrange.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vprq *VideoPlayRangeQuery) FirstX(ctx context.Context) *VideoPlayRange {
	node, err := vprq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VideoPlayRange ID from the query.
// Returns a *NotFoundError when no VideoPlayRange ID was found.
func (vprq *VideoPlayRangeQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = vprq.Limit(1).IDs(setContextOp(ctx, vprq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{videoplayrange.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vprq *VideoPlayRangeQuery) FirstIDX(ctx context.Context) string {
	id, err := vprq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VideoPlayRange entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one VideoPlayRange entity is found.
// Returns a *NotFoundError when no VideoPlayRange entities are found.
func (vprq *VideoPlayRangeQuery) Only(ctx context.Context) (*VideoPlayRange, error) {
	nodes, err := vprq.Limit(2).All(setContextOp(ctx, vprq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{videoplayrange.Label}
	default:
		return nil, &NotSingularError{videoplayrange.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vprq *VideoPlayRangeQuery) OnlyX(ctx context.Context) *VideoPlayRange {
	node, err := vprq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VideoPlayRange ID in the query.
// Returns a *NotSingularError when more than one VideoPlayRange ID is found.
// Returns a *NotFoundError when no entities are found.
func (vprq *VideoPlayRangeQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = vprq.Limit(2).IDs(setContextOp(ctx, vprq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{videoplayrange.Label}
	default:
		err = &NotSingularError{videoplayrange.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vprq *VideoPlayRangeQuery) OnlyIDX(ctx context.Context) string {
	id, err := vprq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VideoPlayRanges.
func (vprq *VideoPlayRangeQuery) All(ctx context.Context) ([]*VideoPlayRange, error) {
	ctx = setContextOp(ctx, vprq.ctx, "All")
	if err := vprq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*VideoPlayRange, *VideoPlayRangeQuery]()
	return withInterceptors[[]*VideoPlayRange](ctx, vprq, qr, vprq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vprq *VideoPlayRangeQuery) AllX(ctx context.Context) []*VideoPlayRange {
	nodes, err := vprq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VideoPlayRange IDs.
func (vprq *VideoPlayRangeQuery) IDs(ctx context.Context) (ids []string, err error) {
	if vprq.ctx.Unique == nil && vprq.path != nil {
		vprq.Unique(true)
	}
	ctx = setContextOp(ctx, vprq.ctx, "IDs")
	if err = vprq.Select(videoplayrange.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vprq *VideoPlayRangeQuery) IDsX(ctx context.Context) []string {
	ids, err := vprq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vprq *VideoPlayRangeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vprq.ctx, "Count")
	if err := vprq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vprq, querierCount[*VideoPlayRangeQuery](), vprq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vprq *VideoPlayRangeQuery) CountX(ctx context.Context) int {
	count, err := vprq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vprq *VideoPlayRangeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vprq.ctx, "Exist")
	switch _, err := vprq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vprq *VideoPlayRangeQuery) ExistX(ctx context.Context) bool {
	exist, err := vprq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VideoPlayRangeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vprq *VideoPlayRangeQuery) Clone() *VideoPlayRangeQuery {
	if vprq == nil {
		return nil
	}
	return &VideoPlayRangeQuery{
		config:     vprq.config,
		ctx:        vprq.ctx.Clone(),
		order:      append([]videoplayrange.OrderOption{}, vprq.order...),
		inters:     append([]Interceptor{}, vprq.inters...),
		predicates: append([]predicate.VideoPlayRange{}, vprq.predicates...),
		withVideo:  vprq.withVideo.Clone(),
		// clone intermediate query.
		sql:  vprq.sql.Clone(),
		path: vprq.path,
	}
}

// WithVideo tells the query-builder to eager-load the nodes that are connected to
// the "video" edge. The optional arguments are used to configure the query builder of the edge.
func (vprq *VideoPlayRangeQuery) WithVideo(opts ...func(*VideoQuery)) *VideoPlayRangeQuery {
	query := (&VideoClient{config: vprq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vprq.withVideo = query
	return vprq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		StartSeconds int `json:"start_seconds,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.VideoPlayRange.Query().
//		GroupBy(videoplayrange.FieldStartSeconds).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vprq *VideoPlayRangeQuery) GroupBy(field string, fields ...string) *VideoPlayRangeGroupBy {
	vprq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VideoPlayRangeGroupBy{build: vprq}
	grbuild.flds = &vprq.ctx.Fields
	grbuild.label = videoplayrange.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		StartSeconds int `json:"start_seconds,omitempty"`
//	}
//
//	client.VideoPlayRange.Query().
//		Select(videoplayrange.FieldStartSeconds).
//		Scan(ctx, &v)
func (vprq *VideoPlayRangeQuery) Select(fields ...string) *VideoPlayRangeSelect {
	vprq.ctx.Fields = append(vprq.ctx.Fields, fields...)
	sbuild := &VideoPlayRangeSelect{VideoPlayRangeQuery: vprq}
	sbuild.label = videoplayrange.Label
	sbuild.flds, sbuild.scan = &vprq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VideoPlayRangeSelect configured with the given aggregations.
func (vprq *VideoPlayRangeQuery) Aggregate(fns ...AggregateFunc) *VideoPlayRangeSelect {
	return vprq.Select().Aggregate(fns...)
}

func (vprq *VideoPlayRangeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vprq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vprq); err != nil {
				return err
			}
		}
	}
	for _, f := range vprq.ctx.Fields {
		if !videoplayrange.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vprq.path != nil {
		prev, err := vprq.path(ctx)
		if err != nil {
			return err
		}
		vprq.sql = prev
	}
	return nil
}

func (vprq *VideoPlayRangeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*VideoPlayRange, error) {
	var (
		nodes       = []*VideoPlayRange{}
		withFKs     = vprq.withFKs
		_spec       = vprq.querySpec()
		loadedTypes = [1]bool{
			vprq.withVideo != nil,
		}
	)
	if vprq.withVideo != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, videoplayrange.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*VideoPlayRange).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &VideoPlayRange{config: vprq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vprq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := vprq.withVideo; query != nil {
		if err := vprq.loadVideo(ctx, query, nodes, nil,
			func(n *VideoPlayRange, e *Video) { n.Edges.Video = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vprq *VideoPlayRangeQuery) loadVideo(ctx context.Context, query *VideoQuery, nodes []*VideoPlayRange, init func(*VideoPlayRange), assign func(*VideoPlayRange, *Video)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*VideoPlayRange)
	for i := range nodes {
		if nodes[i].video_id == nil {
			continue
		}
		fk := *nodes[i].video_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(video.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "video_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (vprq *VideoPlayRangeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vprq.querySpec()
	_spec.Node.Columns = vprq.ctx.Fields
	if len(vprq.ctx.Fields) > 0 {
		_spec.Unique = vprq.ctx.Unique != nil && *vprq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vprq.driver, _spec)
}

func (vprq *VideoPlayRangeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(videoplayrange.Table, videoplayrange.Columns, sqlgraph.NewFieldSpec(videoplayrange.FieldID, field.TypeString))
	_spec.From = vprq.sql
	if unique := vprq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vprq.path != nil {
		_spec.Unique = true
	}
	if fields := vprq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, videoplayrange.FieldID)
		for i := range fields {
			if fields[i] != videoplayrange.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vprq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vprq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vprq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vprq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vprq *VideoPlayRangeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vprq.driver.Dialect())
	t1 := builder.Table(videoplayrange.Table)
	columns := vprq.ctx.Fields
	if len(columns) == 0 {
		columns = videoplayrange.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vprq.sql != nil {
		selector = vprq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vprq.ctx.Unique != nil && *vprq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vprq.predicates {
		p(selector)
	}
	for _, p := range vprq.order {
		p(selector)
	}
	if offset := vprq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vprq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VideoPlayRangeGroupBy is the group-by builder for VideoPlayRange entities.
type VideoPlayRangeGroupBy struct {
	selector
	build *VideoPlayRangeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vprgb *VideoPlayRangeGroupBy) Aggregate(fns ...AggregateFunc) *VideoPlayRangeGroupBy {
	vprgb.fns = append(vprgb.fns, fns...)
	return vprgb
}

// Scan applies the selector query and scans the result into the given value.
func (vprgb *VideoPlayRangeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vprgb.build.ctx, "GroupBy")
	if err := vprgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VideoPlayRangeQuery, *VideoPlayRangeGroupBy](ctx, vprgb.build, vprgb, vprgb.build.inters, v)
}

func (vprgb *VideoPlayRangeGroupBy) sqlScan(ctx context.Context, root *VideoPlayRangeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vprgb.fns))
	for _, fn := range vprgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vprgb.flds)+len(vprgb.fns))
		for _, f := range *vprgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vprgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vprgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VideoPlayRangeSelect is the builder for selecting fields of VideoPlayRange entities.
type VideoPlayRangeSelect struct {
	*VideoPlayRangeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vprs *VideoPlayRangeSelect) Aggregate(fns ...AggregateFunc) *VideoPlayRangeSelect {
	vprs.fns = append(vprs.fns, fns...)
	return vprs
}

// Scan applies the selector query and scans the result into the given value.
func (vprs *VideoPlayRangeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vprs.ctx, "Select")
	if err := vprs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VideoPlayRangeQuery, *VideoPlayRangeSelect](ctx, vprs.VideoPlayRangeQuery, vprs, vprs.inters, v)
}

func (vprs *VideoPlayRangeSelect) sqlScan(ctx context.Context, root *VideoPlayRangeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vprs.fns))
	for _, fn := range vprs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vprs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vprs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}