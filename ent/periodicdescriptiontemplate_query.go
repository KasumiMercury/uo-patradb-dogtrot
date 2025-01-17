// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/description"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/periodicdescriptiontemplate"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/predicate"
)

// PeriodicDescriptionTemplateQuery is the builder for querying PeriodicDescriptionTemplate entities.
type PeriodicDescriptionTemplateQuery struct {
	config
	ctx              *QueryContext
	order            []periodicdescriptiontemplate.OrderOption
	inters           []Interceptor
	predicates       []predicate.PeriodicDescriptionTemplate
	withDescriptions *DescriptionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PeriodicDescriptionTemplateQuery builder.
func (pdtq *PeriodicDescriptionTemplateQuery) Where(ps ...predicate.PeriodicDescriptionTemplate) *PeriodicDescriptionTemplateQuery {
	pdtq.predicates = append(pdtq.predicates, ps...)
	return pdtq
}

// Limit the number of records to be returned by this query.
func (pdtq *PeriodicDescriptionTemplateQuery) Limit(limit int) *PeriodicDescriptionTemplateQuery {
	pdtq.ctx.Limit = &limit
	return pdtq
}

// Offset to start from.
func (pdtq *PeriodicDescriptionTemplateQuery) Offset(offset int) *PeriodicDescriptionTemplateQuery {
	pdtq.ctx.Offset = &offset
	return pdtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pdtq *PeriodicDescriptionTemplateQuery) Unique(unique bool) *PeriodicDescriptionTemplateQuery {
	pdtq.ctx.Unique = &unique
	return pdtq
}

// Order specifies how the records should be ordered.
func (pdtq *PeriodicDescriptionTemplateQuery) Order(o ...periodicdescriptiontemplate.OrderOption) *PeriodicDescriptionTemplateQuery {
	pdtq.order = append(pdtq.order, o...)
	return pdtq
}

// QueryDescriptions chains the current query on the "descriptions" edge.
func (pdtq *PeriodicDescriptionTemplateQuery) QueryDescriptions() *DescriptionQuery {
	query := (&DescriptionClient{config: pdtq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pdtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pdtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(periodicdescriptiontemplate.Table, periodicdescriptiontemplate.FieldID, selector),
			sqlgraph.To(description.Table, description.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, periodicdescriptiontemplate.DescriptionsTable, periodicdescriptiontemplate.DescriptionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pdtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PeriodicDescriptionTemplate entity from the query.
// Returns a *NotFoundError when no PeriodicDescriptionTemplate was found.
func (pdtq *PeriodicDescriptionTemplateQuery) First(ctx context.Context) (*PeriodicDescriptionTemplate, error) {
	nodes, err := pdtq.Limit(1).All(setContextOp(ctx, pdtq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{periodicdescriptiontemplate.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pdtq *PeriodicDescriptionTemplateQuery) FirstX(ctx context.Context) *PeriodicDescriptionTemplate {
	node, err := pdtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PeriodicDescriptionTemplate ID from the query.
// Returns a *NotFoundError when no PeriodicDescriptionTemplate ID was found.
func (pdtq *PeriodicDescriptionTemplateQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pdtq.Limit(1).IDs(setContextOp(ctx, pdtq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{periodicdescriptiontemplate.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pdtq *PeriodicDescriptionTemplateQuery) FirstIDX(ctx context.Context) string {
	id, err := pdtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PeriodicDescriptionTemplate entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PeriodicDescriptionTemplate entity is found.
// Returns a *NotFoundError when no PeriodicDescriptionTemplate entities are found.
func (pdtq *PeriodicDescriptionTemplateQuery) Only(ctx context.Context) (*PeriodicDescriptionTemplate, error) {
	nodes, err := pdtq.Limit(2).All(setContextOp(ctx, pdtq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{periodicdescriptiontemplate.Label}
	default:
		return nil, &NotSingularError{periodicdescriptiontemplate.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pdtq *PeriodicDescriptionTemplateQuery) OnlyX(ctx context.Context) *PeriodicDescriptionTemplate {
	node, err := pdtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PeriodicDescriptionTemplate ID in the query.
// Returns a *NotSingularError when more than one PeriodicDescriptionTemplate ID is found.
// Returns a *NotFoundError when no entities are found.
func (pdtq *PeriodicDescriptionTemplateQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pdtq.Limit(2).IDs(setContextOp(ctx, pdtq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{periodicdescriptiontemplate.Label}
	default:
		err = &NotSingularError{periodicdescriptiontemplate.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pdtq *PeriodicDescriptionTemplateQuery) OnlyIDX(ctx context.Context) string {
	id, err := pdtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PeriodicDescriptionTemplates.
func (pdtq *PeriodicDescriptionTemplateQuery) All(ctx context.Context) ([]*PeriodicDescriptionTemplate, error) {
	ctx = setContextOp(ctx, pdtq.ctx, "All")
	if err := pdtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PeriodicDescriptionTemplate, *PeriodicDescriptionTemplateQuery]()
	return withInterceptors[[]*PeriodicDescriptionTemplate](ctx, pdtq, qr, pdtq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pdtq *PeriodicDescriptionTemplateQuery) AllX(ctx context.Context) []*PeriodicDescriptionTemplate {
	nodes, err := pdtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PeriodicDescriptionTemplate IDs.
func (pdtq *PeriodicDescriptionTemplateQuery) IDs(ctx context.Context) (ids []string, err error) {
	if pdtq.ctx.Unique == nil && pdtq.path != nil {
		pdtq.Unique(true)
	}
	ctx = setContextOp(ctx, pdtq.ctx, "IDs")
	if err = pdtq.Select(periodicdescriptiontemplate.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pdtq *PeriodicDescriptionTemplateQuery) IDsX(ctx context.Context) []string {
	ids, err := pdtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pdtq *PeriodicDescriptionTemplateQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pdtq.ctx, "Count")
	if err := pdtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pdtq, querierCount[*PeriodicDescriptionTemplateQuery](), pdtq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pdtq *PeriodicDescriptionTemplateQuery) CountX(ctx context.Context) int {
	count, err := pdtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pdtq *PeriodicDescriptionTemplateQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pdtq.ctx, "Exist")
	switch _, err := pdtq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pdtq *PeriodicDescriptionTemplateQuery) ExistX(ctx context.Context) bool {
	exist, err := pdtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PeriodicDescriptionTemplateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pdtq *PeriodicDescriptionTemplateQuery) Clone() *PeriodicDescriptionTemplateQuery {
	if pdtq == nil {
		return nil
	}
	return &PeriodicDescriptionTemplateQuery{
		config:           pdtq.config,
		ctx:              pdtq.ctx.Clone(),
		order:            append([]periodicdescriptiontemplate.OrderOption{}, pdtq.order...),
		inters:           append([]Interceptor{}, pdtq.inters...),
		predicates:       append([]predicate.PeriodicDescriptionTemplate{}, pdtq.predicates...),
		withDescriptions: pdtq.withDescriptions.Clone(),
		// clone intermediate query.
		sql:  pdtq.sql.Clone(),
		path: pdtq.path,
	}
}

// WithDescriptions tells the query-builder to eager-load the nodes that are connected to
// the "descriptions" edge. The optional arguments are used to configure the query builder of the edge.
func (pdtq *PeriodicDescriptionTemplateQuery) WithDescriptions(opts ...func(*DescriptionQuery)) *PeriodicDescriptionTemplateQuery {
	query := (&DescriptionClient{config: pdtq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pdtq.withDescriptions = query
	return pdtq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PeriodicDescriptionTemplate.Query().
//		GroupBy(periodicdescriptiontemplate.FieldText).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pdtq *PeriodicDescriptionTemplateQuery) GroupBy(field string, fields ...string) *PeriodicDescriptionTemplateGroupBy {
	pdtq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PeriodicDescriptionTemplateGroupBy{build: pdtq}
	grbuild.flds = &pdtq.ctx.Fields
	grbuild.label = periodicdescriptiontemplate.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//	}
//
//	client.PeriodicDescriptionTemplate.Query().
//		Select(periodicdescriptiontemplate.FieldText).
//		Scan(ctx, &v)
func (pdtq *PeriodicDescriptionTemplateQuery) Select(fields ...string) *PeriodicDescriptionTemplateSelect {
	pdtq.ctx.Fields = append(pdtq.ctx.Fields, fields...)
	sbuild := &PeriodicDescriptionTemplateSelect{PeriodicDescriptionTemplateQuery: pdtq}
	sbuild.label = periodicdescriptiontemplate.Label
	sbuild.flds, sbuild.scan = &pdtq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PeriodicDescriptionTemplateSelect configured with the given aggregations.
func (pdtq *PeriodicDescriptionTemplateQuery) Aggregate(fns ...AggregateFunc) *PeriodicDescriptionTemplateSelect {
	return pdtq.Select().Aggregate(fns...)
}

func (pdtq *PeriodicDescriptionTemplateQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pdtq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pdtq); err != nil {
				return err
			}
		}
	}
	for _, f := range pdtq.ctx.Fields {
		if !periodicdescriptiontemplate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pdtq.path != nil {
		prev, err := pdtq.path(ctx)
		if err != nil {
			return err
		}
		pdtq.sql = prev
	}
	return nil
}

func (pdtq *PeriodicDescriptionTemplateQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PeriodicDescriptionTemplate, error) {
	var (
		nodes       = []*PeriodicDescriptionTemplate{}
		_spec       = pdtq.querySpec()
		loadedTypes = [1]bool{
			pdtq.withDescriptions != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PeriodicDescriptionTemplate).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PeriodicDescriptionTemplate{config: pdtq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pdtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pdtq.withDescriptions; query != nil {
		if err := pdtq.loadDescriptions(ctx, query, nodes,
			func(n *PeriodicDescriptionTemplate) { n.Edges.Descriptions = []*Description{} },
			func(n *PeriodicDescriptionTemplate, e *Description) {
				n.Edges.Descriptions = append(n.Edges.Descriptions, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pdtq *PeriodicDescriptionTemplateQuery) loadDescriptions(ctx context.Context, query *DescriptionQuery, nodes []*PeriodicDescriptionTemplate, init func(*PeriodicDescriptionTemplate), assign func(*PeriodicDescriptionTemplate, *Description)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*PeriodicDescriptionTemplate)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Description(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(periodicdescriptiontemplate.DescriptionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.periodic_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "periodic_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "periodic_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (pdtq *PeriodicDescriptionTemplateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pdtq.querySpec()
	_spec.Node.Columns = pdtq.ctx.Fields
	if len(pdtq.ctx.Fields) > 0 {
		_spec.Unique = pdtq.ctx.Unique != nil && *pdtq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pdtq.driver, _spec)
}

func (pdtq *PeriodicDescriptionTemplateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(periodicdescriptiontemplate.Table, periodicdescriptiontemplate.Columns, sqlgraph.NewFieldSpec(periodicdescriptiontemplate.FieldID, field.TypeString))
	_spec.From = pdtq.sql
	if unique := pdtq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pdtq.path != nil {
		_spec.Unique = true
	}
	if fields := pdtq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, periodicdescriptiontemplate.FieldID)
		for i := range fields {
			if fields[i] != periodicdescriptiontemplate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pdtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pdtq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pdtq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pdtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pdtq *PeriodicDescriptionTemplateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pdtq.driver.Dialect())
	t1 := builder.Table(periodicdescriptiontemplate.Table)
	columns := pdtq.ctx.Fields
	if len(columns) == 0 {
		columns = periodicdescriptiontemplate.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pdtq.sql != nil {
		selector = pdtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pdtq.ctx.Unique != nil && *pdtq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pdtq.predicates {
		p(selector)
	}
	for _, p := range pdtq.order {
		p(selector)
	}
	if offset := pdtq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pdtq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PeriodicDescriptionTemplateGroupBy is the group-by builder for PeriodicDescriptionTemplate entities.
type PeriodicDescriptionTemplateGroupBy struct {
	selector
	build *PeriodicDescriptionTemplateQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pdtgb *PeriodicDescriptionTemplateGroupBy) Aggregate(fns ...AggregateFunc) *PeriodicDescriptionTemplateGroupBy {
	pdtgb.fns = append(pdtgb.fns, fns...)
	return pdtgb
}

// Scan applies the selector query and scans the result into the given value.
func (pdtgb *PeriodicDescriptionTemplateGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pdtgb.build.ctx, "GroupBy")
	if err := pdtgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PeriodicDescriptionTemplateQuery, *PeriodicDescriptionTemplateGroupBy](ctx, pdtgb.build, pdtgb, pdtgb.build.inters, v)
}

func (pdtgb *PeriodicDescriptionTemplateGroupBy) sqlScan(ctx context.Context, root *PeriodicDescriptionTemplateQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pdtgb.fns))
	for _, fn := range pdtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pdtgb.flds)+len(pdtgb.fns))
		for _, f := range *pdtgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pdtgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pdtgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PeriodicDescriptionTemplateSelect is the builder for selecting fields of PeriodicDescriptionTemplate entities.
type PeriodicDescriptionTemplateSelect struct {
	*PeriodicDescriptionTemplateQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pdts *PeriodicDescriptionTemplateSelect) Aggregate(fns ...AggregateFunc) *PeriodicDescriptionTemplateSelect {
	pdts.fns = append(pdts.fns, fns...)
	return pdts
}

// Scan applies the selector query and scans the result into the given value.
func (pdts *PeriodicDescriptionTemplateSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pdts.ctx, "Select")
	if err := pdts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PeriodicDescriptionTemplateQuery, *PeriodicDescriptionTemplateSelect](ctx, pdts.PeriodicDescriptionTemplateQuery, pdts, pdts.inters, v)
}

func (pdts *PeriodicDescriptionTemplateSelect) sqlScan(ctx context.Context, root *PeriodicDescriptionTemplateQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pdts.fns))
	for _, fn := range pdts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pdts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pdts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
