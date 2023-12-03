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
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/channel"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/description"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/predicate"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/video"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/video_disallow_range"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/video_play_range"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/video_title_change"
)

// VideoQuery is the builder for querying Video entities.
type VideoQuery struct {
	config
	ctx                     *QueryContext
	order                   []video.OrderOption
	inters                  []Interceptor
	predicates              []predicate.Video
	withDescriptions        *DescriptionQuery
	withChannel             *ChannelQuery
	withVideoPlayRanges     *VideoPlayRangeQuery
	withVideoDisallowRanges *VideoDisallowRangeQuery
	withVideoTitleChanges   *VideoTitleChangeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VideoQuery builder.
func (vq *VideoQuery) Where(ps ...predicate.Video) *VideoQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit the number of records to be returned by this query.
func (vq *VideoQuery) Limit(limit int) *VideoQuery {
	vq.ctx.Limit = &limit
	return vq
}

// Offset to start from.
func (vq *VideoQuery) Offset(offset int) *VideoQuery {
	vq.ctx.Offset = &offset
	return vq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vq *VideoQuery) Unique(unique bool) *VideoQuery {
	vq.ctx.Unique = &unique
	return vq
}

// Order specifies how the records should be ordered.
func (vq *VideoQuery) Order(o ...video.OrderOption) *VideoQuery {
	vq.order = append(vq.order, o...)
	return vq
}

// QueryDescriptions chains the current query on the "descriptions" edge.
func (vq *VideoQuery) QueryDescriptions() *DescriptionQuery {
	query := (&DescriptionClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(video.Table, video.FieldID, selector),
			sqlgraph.To(description.Table, description.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, video.DescriptionsTable, video.DescriptionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChannel chains the current query on the "channel" edge.
func (vq *VideoQuery) QueryChannel() *ChannelQuery {
	query := (&ChannelClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(video.Table, video.FieldID, selector),
			sqlgraph.To(channel.Table, channel.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, video.ChannelTable, video.ChannelPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVideoPlayRanges chains the current query on the "video_play_ranges" edge.
func (vq *VideoQuery) QueryVideoPlayRanges() *VideoPlayRangeQuery {
	query := (&VideoPlayRangeClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(video.Table, video.FieldID, selector),
			sqlgraph.To(video_play_range.Table, video_play_range.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, video.VideoPlayRangesTable, video.VideoPlayRangesColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVideoDisallowRanges chains the current query on the "video_disallow_ranges" edge.
func (vq *VideoQuery) QueryVideoDisallowRanges() *VideoDisallowRangeQuery {
	query := (&VideoDisallowRangeClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(video.Table, video.FieldID, selector),
			sqlgraph.To(video_disallow_range.Table, video_disallow_range.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, video.VideoDisallowRangesTable, video.VideoDisallowRangesColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVideoTitleChanges chains the current query on the "video_title_changes" edge.
func (vq *VideoQuery) QueryVideoTitleChanges() *VideoTitleChangeQuery {
	query := (&VideoTitleChangeClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(video.Table, video.FieldID, selector),
			sqlgraph.To(video_title_change.Table, video_title_change.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, video.VideoTitleChangesTable, video.VideoTitleChangesColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Video entity from the query.
// Returns a *NotFoundError when no Video was found.
func (vq *VideoQuery) First(ctx context.Context) (*Video, error) {
	nodes, err := vq.Limit(1).All(setContextOp(ctx, vq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{video.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *VideoQuery) FirstX(ctx context.Context) *Video {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Video ID from the query.
// Returns a *NotFoundError when no Video ID was found.
func (vq *VideoQuery) FirstID(ctx context.Context) (id pulid.ID, err error) {
	var ids []pulid.ID
	if ids, err = vq.Limit(1).IDs(setContextOp(ctx, vq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{video.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vq *VideoQuery) FirstIDX(ctx context.Context) pulid.ID {
	id, err := vq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Video entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Video entity is found.
// Returns a *NotFoundError when no Video entities are found.
func (vq *VideoQuery) Only(ctx context.Context) (*Video, error) {
	nodes, err := vq.Limit(2).All(setContextOp(ctx, vq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{video.Label}
	default:
		return nil, &NotSingularError{video.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *VideoQuery) OnlyX(ctx context.Context) *Video {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Video ID in the query.
// Returns a *NotSingularError when more than one Video ID is found.
// Returns a *NotFoundError when no entities are found.
func (vq *VideoQuery) OnlyID(ctx context.Context) (id pulid.ID, err error) {
	var ids []pulid.ID
	if ids, err = vq.Limit(2).IDs(setContextOp(ctx, vq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{video.Label}
	default:
		err = &NotSingularError{video.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vq *VideoQuery) OnlyIDX(ctx context.Context) pulid.ID {
	id, err := vq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Videos.
func (vq *VideoQuery) All(ctx context.Context) ([]*Video, error) {
	ctx = setContextOp(ctx, vq.ctx, "All")
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Video, *VideoQuery]()
	return withInterceptors[[]*Video](ctx, vq, qr, vq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vq *VideoQuery) AllX(ctx context.Context) []*Video {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Video IDs.
func (vq *VideoQuery) IDs(ctx context.Context) (ids []pulid.ID, err error) {
	if vq.ctx.Unique == nil && vq.path != nil {
		vq.Unique(true)
	}
	ctx = setContextOp(ctx, vq.ctx, "IDs")
	if err = vq.Select(video.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vq *VideoQuery) IDsX(ctx context.Context) []pulid.ID {
	ids, err := vq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vq *VideoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vq.ctx, "Count")
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vq, querierCount[*VideoQuery](), vq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vq *VideoQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *VideoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vq.ctx, "Exist")
	switch _, err := vq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *VideoQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VideoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *VideoQuery) Clone() *VideoQuery {
	if vq == nil {
		return nil
	}
	return &VideoQuery{
		config:                  vq.config,
		ctx:                     vq.ctx.Clone(),
		order:                   append([]video.OrderOption{}, vq.order...),
		inters:                  append([]Interceptor{}, vq.inters...),
		predicates:              append([]predicate.Video{}, vq.predicates...),
		withDescriptions:        vq.withDescriptions.Clone(),
		withChannel:             vq.withChannel.Clone(),
		withVideoPlayRanges:     vq.withVideoPlayRanges.Clone(),
		withVideoDisallowRanges: vq.withVideoDisallowRanges.Clone(),
		withVideoTitleChanges:   vq.withVideoTitleChanges.Clone(),
		// clone intermediate query.
		sql:  vq.sql.Clone(),
		path: vq.path,
	}
}

// WithDescriptions tells the query-builder to eager-load the nodes that are connected to
// the "descriptions" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VideoQuery) WithDescriptions(opts ...func(*DescriptionQuery)) *VideoQuery {
	query := (&DescriptionClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withDescriptions = query
	return vq
}

// WithChannel tells the query-builder to eager-load the nodes that are connected to
// the "channel" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VideoQuery) WithChannel(opts ...func(*ChannelQuery)) *VideoQuery {
	query := (&ChannelClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withChannel = query
	return vq
}

// WithVideoPlayRanges tells the query-builder to eager-load the nodes that are connected to
// the "video_play_ranges" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VideoQuery) WithVideoPlayRanges(opts ...func(*VideoPlayRangeQuery)) *VideoQuery {
	query := (&VideoPlayRangeClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withVideoPlayRanges = query
	return vq
}

// WithVideoDisallowRanges tells the query-builder to eager-load the nodes that are connected to
// the "video_disallow_ranges" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VideoQuery) WithVideoDisallowRanges(opts ...func(*VideoDisallowRangeQuery)) *VideoQuery {
	query := (&VideoDisallowRangeClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withVideoDisallowRanges = query
	return vq
}

// WithVideoTitleChanges tells the query-builder to eager-load the nodes that are connected to
// the "video_title_changes" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VideoQuery) WithVideoTitleChanges(opts ...func(*VideoTitleChangeQuery)) *VideoQuery {
	query := (&VideoTitleChangeClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withVideoTitleChanges = query
	return vq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		VideoID string `json:"video_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Video.Query().
//		GroupBy(video.FieldVideoID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vq *VideoQuery) GroupBy(field string, fields ...string) *VideoGroupBy {
	vq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VideoGroupBy{build: vq}
	grbuild.flds = &vq.ctx.Fields
	grbuild.label = video.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		VideoID string `json:"video_id,omitempty"`
//	}
//
//	client.Video.Query().
//		Select(video.FieldVideoID).
//		Scan(ctx, &v)
func (vq *VideoQuery) Select(fields ...string) *VideoSelect {
	vq.ctx.Fields = append(vq.ctx.Fields, fields...)
	sbuild := &VideoSelect{VideoQuery: vq}
	sbuild.label = video.Label
	sbuild.flds, sbuild.scan = &vq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VideoSelect configured with the given aggregations.
func (vq *VideoQuery) Aggregate(fns ...AggregateFunc) *VideoSelect {
	return vq.Select().Aggregate(fns...)
}

func (vq *VideoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vq); err != nil {
				return err
			}
		}
	}
	for _, f := range vq.ctx.Fields {
		if !video.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	return nil
}

func (vq *VideoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Video, error) {
	var (
		nodes       = []*Video{}
		_spec       = vq.querySpec()
		loadedTypes = [5]bool{
			vq.withDescriptions != nil,
			vq.withChannel != nil,
			vq.withVideoPlayRanges != nil,
			vq.withVideoDisallowRanges != nil,
			vq.withVideoTitleChanges != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Video).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Video{config: vq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := vq.withDescriptions; query != nil {
		if err := vq.loadDescriptions(ctx, query, nodes, nil,
			func(n *Video, e *Description) { n.Edges.Descriptions = e }); err != nil {
			return nil, err
		}
	}
	if query := vq.withChannel; query != nil {
		if err := vq.loadChannel(ctx, query, nodes,
			func(n *Video) { n.Edges.Channel = []*Channel{} },
			func(n *Video, e *Channel) { n.Edges.Channel = append(n.Edges.Channel, e) }); err != nil {
			return nil, err
		}
	}
	if query := vq.withVideoPlayRanges; query != nil {
		if err := vq.loadVideoPlayRanges(ctx, query, nodes,
			func(n *Video) { n.Edges.VideoPlayRanges = []*Video_play_range{} },
			func(n *Video, e *Video_play_range) { n.Edges.VideoPlayRanges = append(n.Edges.VideoPlayRanges, e) }); err != nil {
			return nil, err
		}
	}
	if query := vq.withVideoDisallowRanges; query != nil {
		if err := vq.loadVideoDisallowRanges(ctx, query, nodes,
			func(n *Video) { n.Edges.VideoDisallowRanges = []*Video_disallow_range{} },
			func(n *Video, e *Video_disallow_range) {
				n.Edges.VideoDisallowRanges = append(n.Edges.VideoDisallowRanges, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := vq.withVideoTitleChanges; query != nil {
		if err := vq.loadVideoTitleChanges(ctx, query, nodes,
			func(n *Video) { n.Edges.VideoTitleChanges = []*Video_title_change{} },
			func(n *Video, e *Video_title_change) {
				n.Edges.VideoTitleChanges = append(n.Edges.VideoTitleChanges, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vq *VideoQuery) loadDescriptions(ctx context.Context, query *DescriptionQuery, nodes []*Video, init func(*Video), assign func(*Video, *Description)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[pulid.ID]*Video)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.Description(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(video.DescriptionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.video_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "video_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "video_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (vq *VideoQuery) loadChannel(ctx context.Context, query *ChannelQuery, nodes []*Video, init func(*Video), assign func(*Video, *Channel)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[pulid.ID]*Video)
	nids := make(map[pulid.ID]map[*Video]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(video.ChannelTable)
		s.Join(joinT).On(s.C(channel.FieldID), joinT.C(video.ChannelPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(video.ChannelPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(video.ChannelPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullString)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := pulid.ID(values[0].(*sql.NullString).String)
				inValue := pulid.ID(values[1].(*sql.NullString).String)
				if nids[inValue] == nil {
					nids[inValue] = map[*Video]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Channel](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "channel" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (vq *VideoQuery) loadVideoPlayRanges(ctx context.Context, query *VideoPlayRangeQuery, nodes []*Video, init func(*Video), assign func(*Video, *Video_play_range)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[pulid.ID]*Video)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Video_play_range(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(video.VideoPlayRangesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.video_video_play_ranges
		if fk == nil {
			return fmt.Errorf(`foreign-key "video_video_play_ranges" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "video_video_play_ranges" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (vq *VideoQuery) loadVideoDisallowRanges(ctx context.Context, query *VideoDisallowRangeQuery, nodes []*Video, init func(*Video), assign func(*Video, *Video_disallow_range)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[pulid.ID]*Video)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Video_disallow_range(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(video.VideoDisallowRangesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.video_video_disallow_ranges
		if fk == nil {
			return fmt.Errorf(`foreign-key "video_video_disallow_ranges" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "video_video_disallow_ranges" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (vq *VideoQuery) loadVideoTitleChanges(ctx context.Context, query *VideoTitleChangeQuery, nodes []*Video, init func(*Video), assign func(*Video, *Video_title_change)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[pulid.ID]*Video)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Video_title_change(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(video.VideoTitleChangesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.video_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "video_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "video_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (vq *VideoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
	_spec.Node.Columns = vq.ctx.Fields
	if len(vq.ctx.Fields) > 0 {
		_spec.Unique = vq.ctx.Unique != nil && *vq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *VideoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(video.Table, video.Columns, sqlgraph.NewFieldSpec(video.FieldID, field.TypeString))
	_spec.From = vq.sql
	if unique := vq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vq.path != nil {
		_spec.Unique = true
	}
	if fields := vq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, video.FieldID)
		for i := range fields {
			if fields[i] != video.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vq *VideoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(video.Table)
	columns := vq.ctx.Fields
	if len(columns) == 0 {
		columns = video.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vq.ctx.Unique != nil && *vq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector)
	}
	if offset := vq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VideoGroupBy is the group-by builder for Video entities.
type VideoGroupBy struct {
	selector
	build *VideoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *VideoGroupBy) Aggregate(fns ...AggregateFunc) *VideoGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the selector query and scans the result into the given value.
func (vgb *VideoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vgb.build.ctx, "GroupBy")
	if err := vgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VideoQuery, *VideoGroupBy](ctx, vgb.build, vgb, vgb.build.inters, v)
}

func (vgb *VideoGroupBy) sqlScan(ctx context.Context, root *VideoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vgb.fns))
	for _, fn := range vgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vgb.flds)+len(vgb.fns))
		for _, f := range *vgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VideoSelect is the builder for selecting fields of Video entities.
type VideoSelect struct {
	*VideoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vs *VideoSelect) Aggregate(fns ...AggregateFunc) *VideoSelect {
	vs.fns = append(vs.fns, fns...)
	return vs
}

// Scan applies the selector query and scans the result into the given value.
func (vs *VideoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vs.ctx, "Select")
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VideoQuery, *VideoSelect](ctx, vs.VideoQuery, vs, vs.inters, v)
}

func (vs *VideoSelect) sqlScan(ctx context.Context, root *VideoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vs.fns))
	for _, fn := range vs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}