// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/pat_chat"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/predicate"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/video"
)

// PatChatQuery is the builder for querying Pat_chat entities.
type PatChatQuery struct {
	config
	ctx        *QueryContext
	order      []pat_chat.OrderOption
	inters     []Interceptor
	predicates []predicate.Pat_chat
	withVideo  *VideoQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PatChatQuery builder.
func (pcq *PatChatQuery) Where(ps ...predicate.Pat_chat) *PatChatQuery {
	pcq.predicates = append(pcq.predicates, ps...)
	return pcq
}

// Limit the number of records to be returned by this query.
func (pcq *PatChatQuery) Limit(limit int) *PatChatQuery {
	pcq.ctx.Limit = &limit
	return pcq
}

// Offset to start from.
func (pcq *PatChatQuery) Offset(offset int) *PatChatQuery {
	pcq.ctx.Offset = &offset
	return pcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pcq *PatChatQuery) Unique(unique bool) *PatChatQuery {
	pcq.ctx.Unique = &unique
	return pcq
}

// Order specifies how the records should be ordered.
func (pcq *PatChatQuery) Order(o ...pat_chat.OrderOption) *PatChatQuery {
	pcq.order = append(pcq.order, o...)
	return pcq
}

// QueryVideo chains the current query on the "video" edge.
func (pcq *PatChatQuery) QueryVideo() *VideoQuery {
	query := (&VideoClient{config: pcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pat_chat.Table, pat_chat.FieldID, selector),
			sqlgraph.To(video.Table, video.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, pat_chat.VideoTable, pat_chat.VideoColumn),
		)
		fromU = sqlgraph.SetNeighbors(pcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Pat_chat entity from the query.
// Returns a *NotFoundError when no Pat_chat was found.
func (pcq *PatChatQuery) First(ctx context.Context) (*Pat_chat, error) {
	nodes, err := pcq.Limit(1).All(setContextOp(ctx, pcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pat_chat.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pcq *PatChatQuery) FirstX(ctx context.Context) *Pat_chat {
	node, err := pcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Pat_chat ID from the query.
// Returns a *NotFoundError when no Pat_chat ID was found.
func (pcq *PatChatQuery) FirstID(ctx context.Context) (id pulid.ID, err error) {
	var ids []pulid.ID
	if ids, err = pcq.Limit(1).IDs(setContextOp(ctx, pcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pat_chat.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pcq *PatChatQuery) FirstIDX(ctx context.Context) pulid.ID {
	id, err := pcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Pat_chat entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Pat_chat entity is found.
// Returns a *NotFoundError when no Pat_chat entities are found.
func (pcq *PatChatQuery) Only(ctx context.Context) (*Pat_chat, error) {
	nodes, err := pcq.Limit(2).All(setContextOp(ctx, pcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pat_chat.Label}
	default:
		return nil, &NotSingularError{pat_chat.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pcq *PatChatQuery) OnlyX(ctx context.Context) *Pat_chat {
	node, err := pcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Pat_chat ID in the query.
// Returns a *NotSingularError when more than one Pat_chat ID is found.
// Returns a *NotFoundError when no entities are found.
func (pcq *PatChatQuery) OnlyID(ctx context.Context) (id pulid.ID, err error) {
	var ids []pulid.ID
	if ids, err = pcq.Limit(2).IDs(setContextOp(ctx, pcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pat_chat.Label}
	default:
		err = &NotSingularError{pat_chat.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pcq *PatChatQuery) OnlyIDX(ctx context.Context) pulid.ID {
	id, err := pcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Pat_chats.
func (pcq *PatChatQuery) All(ctx context.Context) ([]*Pat_chat, error) {
	ctx = setContextOp(ctx, pcq.ctx, "All")
	if err := pcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Pat_chat, *PatChatQuery]()
	return withInterceptors[[]*Pat_chat](ctx, pcq, qr, pcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pcq *PatChatQuery) AllX(ctx context.Context) []*Pat_chat {
	nodes, err := pcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Pat_chat IDs.
func (pcq *PatChatQuery) IDs(ctx context.Context) (ids []pulid.ID, err error) {
	if pcq.ctx.Unique == nil && pcq.path != nil {
		pcq.Unique(true)
	}
	ctx = setContextOp(ctx, pcq.ctx, "IDs")
	if err = pcq.Select(pat_chat.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pcq *PatChatQuery) IDsX(ctx context.Context) []pulid.ID {
	ids, err := pcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pcq *PatChatQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pcq.ctx, "Count")
	if err := pcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pcq, querierCount[*PatChatQuery](), pcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pcq *PatChatQuery) CountX(ctx context.Context) int {
	count, err := pcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pcq *PatChatQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pcq.ctx, "Exist")
	switch _, err := pcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pcq *PatChatQuery) ExistX(ctx context.Context) bool {
	exist, err := pcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PatChatQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pcq *PatChatQuery) Clone() *PatChatQuery {
	if pcq == nil {
		return nil
	}
	return &PatChatQuery{
		config:     pcq.config,
		ctx:        pcq.ctx.Clone(),
		order:      append([]pat_chat.OrderOption{}, pcq.order...),
		inters:     append([]Interceptor{}, pcq.inters...),
		predicates: append([]predicate.Pat_chat{}, pcq.predicates...),
		withVideo:  pcq.withVideo.Clone(),
		// clone intermediate query.
		sql:  pcq.sql.Clone(),
		path: pcq.path,
	}
}

// WithVideo tells the query-builder to eager-load the nodes that are connected to
// the "video" edge. The optional arguments are used to configure the query builder of the edge.
func (pcq *PatChatQuery) WithVideo(opts ...func(*VideoQuery)) *PatChatQuery {
	query := (&VideoClient{config: pcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pcq.withVideo = query
	return pcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Message string `json:"message,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PatChat.Query().
//		GroupBy(pat_chat.FieldMessage).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pcq *PatChatQuery) GroupBy(field string, fields ...string) *PatChatGroupBy {
	pcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PatChatGroupBy{build: pcq}
	grbuild.flds = &pcq.ctx.Fields
	grbuild.label = pat_chat.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Message string `json:"message,omitempty"`
//	}
//
//	client.PatChat.Query().
//		Select(pat_chat.FieldMessage).
//		Scan(ctx, &v)
func (pcq *PatChatQuery) Select(fields ...string) *PatChatSelect {
	pcq.ctx.Fields = append(pcq.ctx.Fields, fields...)
	sbuild := &PatChatSelect{PatChatQuery: pcq}
	sbuild.label = pat_chat.Label
	sbuild.flds, sbuild.scan = &pcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PatChatSelect configured with the given aggregations.
func (pcq *PatChatQuery) Aggregate(fns ...AggregateFunc) *PatChatSelect {
	return pcq.Select().Aggregate(fns...)
}

func (pcq *PatChatQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pcq); err != nil {
				return err
			}
		}
	}
	for _, f := range pcq.ctx.Fields {
		if !pat_chat.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pcq.path != nil {
		prev, err := pcq.path(ctx)
		if err != nil {
			return err
		}
		pcq.sql = prev
	}
	return nil
}

func (pcq *PatChatQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Pat_chat, error) {
	var (
		nodes       = []*Pat_chat{}
		withFKs     = pcq.withFKs
		_spec       = pcq.querySpec()
		loadedTypes = [1]bool{
			pcq.withVideo != nil,
		}
	)
	if pcq.withVideo != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, pat_chat.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Pat_chat).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Pat_chat{config: pcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pcq.withVideo; query != nil {
		if err := pcq.loadVideo(ctx, query, nodes, nil,
			func(n *Pat_chat, e *Video) { n.Edges.Video = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pcq *PatChatQuery) loadVideo(ctx context.Context, query *VideoQuery, nodes []*Pat_chat, init func(*Pat_chat), assign func(*Pat_chat, *Video)) error {
	ids := make([]pulid.ID, 0, len(nodes))
	nodeids := make(map[pulid.ID][]*Pat_chat)
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

func (pcq *PatChatQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pcq.querySpec()
	_spec.Node.Columns = pcq.ctx.Fields
	if len(pcq.ctx.Fields) > 0 {
		_spec.Unique = pcq.ctx.Unique != nil && *pcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pcq.driver, _spec)
}

func (pcq *PatChatQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(pat_chat.Table, pat_chat.Columns, sqlgraph.NewFieldSpec(pat_chat.FieldID, field.TypeString))
	_spec.From = pcq.sql
	if unique := pcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pcq.path != nil {
		_spec.Unique = true
	}
	if fields := pcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pat_chat.FieldID)
		for i := range fields {
			if fields[i] != pat_chat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pcq *PatChatQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pcq.driver.Dialect())
	t1 := builder.Table(pat_chat.Table)
	columns := pcq.ctx.Fields
	if len(columns) == 0 {
		columns = pat_chat.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pcq.sql != nil {
		selector = pcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pcq.ctx.Unique != nil && *pcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pcq.predicates {
		p(selector)
	}
	for _, p := range pcq.order {
		p(selector)
	}
	if offset := pcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PatChatGroupBy is the group-by builder for Pat_chat entities.
type PatChatGroupBy struct {
	selector
	build *PatChatQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pcgb *PatChatGroupBy) Aggregate(fns ...AggregateFunc) *PatChatGroupBy {
	pcgb.fns = append(pcgb.fns, fns...)
	return pcgb
}

// Scan applies the selector query and scans the result into the given value.
func (pcgb *PatChatGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pcgb.build.ctx, "GroupBy")
	if err := pcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PatChatQuery, *PatChatGroupBy](ctx, pcgb.build, pcgb, pcgb.build.inters, v)
}

func (pcgb *PatChatGroupBy) sqlScan(ctx context.Context, root *PatChatQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pcgb.fns))
	for _, fn := range pcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pcgb.flds)+len(pcgb.fns))
		for _, f := range *pcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PatChatSelect is the builder for selecting fields of PatChat entities.
type PatChatSelect struct {
	*PatChatQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pcs *PatChatSelect) Aggregate(fns ...AggregateFunc) *PatChatSelect {
	pcs.fns = append(pcs.fns, fns...)
	return pcs
}

// Scan applies the selector query and scans the result into the given value.
func (pcs *PatChatSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pcs.ctx, "Select")
	if err := pcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PatChatQuery, *PatChatSelect](ctx, pcs.PatChatQuery, pcs, pcs.inters, v)
}

func (pcs *PatChatSelect) sqlScan(ctx context.Context, root *PatChatQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pcs.fns))
	for _, fn := range pcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}