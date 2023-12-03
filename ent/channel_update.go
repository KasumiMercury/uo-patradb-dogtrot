// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/channel"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/predicate"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/video"
)

// ChannelUpdate is the builder for updating Channel entities.
type ChannelUpdate struct {
	config
	hooks    []Hook
	mutation *ChannelMutation
}

// Where appends a list predicates to the ChannelUpdate builder.
func (cu *ChannelUpdate) Where(ps ...predicate.Channel) *ChannelUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetDisplayName sets the "display_name" field.
func (cu *ChannelUpdate) SetDisplayName(s string) *ChannelUpdate {
	cu.mutation.SetDisplayName(s)
	return cu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (cu *ChannelUpdate) SetNillableDisplayName(s *string) *ChannelUpdate {
	if s != nil {
		cu.SetDisplayName(*s)
	}
	return cu
}

// SetChannelID sets the "channel_id" field.
func (cu *ChannelUpdate) SetChannelID(s string) *ChannelUpdate {
	cu.mutation.SetChannelID(s)
	return cu
}

// SetNillableChannelID sets the "channel_id" field if the given value is not nil.
func (cu *ChannelUpdate) SetNillableChannelID(s *string) *ChannelUpdate {
	if s != nil {
		cu.SetChannelID(*s)
	}
	return cu
}

// SetHandle sets the "handle" field.
func (cu *ChannelUpdate) SetHandle(s string) *ChannelUpdate {
	cu.mutation.SetHandle(s)
	return cu
}

// SetNillableHandle sets the "handle" field if the given value is not nil.
func (cu *ChannelUpdate) SetNillableHandle(s *string) *ChannelUpdate {
	if s != nil {
		cu.SetHandle(*s)
	}
	return cu
}

// SetThumbnailURL sets the "thumbnail_url" field.
func (cu *ChannelUpdate) SetThumbnailURL(s string) *ChannelUpdate {
	cu.mutation.SetThumbnailURL(s)
	return cu
}

// SetNillableThumbnailURL sets the "thumbnail_url" field if the given value is not nil.
func (cu *ChannelUpdate) SetNillableThumbnailURL(s *string) *ChannelUpdate {
	if s != nil {
		cu.SetThumbnailURL(*s)
	}
	return cu
}

// AddVideoIDs adds the "videos" edge to the Video entity by IDs.
func (cu *ChannelUpdate) AddVideoIDs(ids ...pulid.ID) *ChannelUpdate {
	cu.mutation.AddVideoIDs(ids...)
	return cu
}

// AddVideos adds the "videos" edges to the Video entity.
func (cu *ChannelUpdate) AddVideos(v ...*Video) *ChannelUpdate {
	ids := make([]pulid.ID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cu.AddVideoIDs(ids...)
}

// Mutation returns the ChannelMutation object of the builder.
func (cu *ChannelUpdate) Mutation() *ChannelMutation {
	return cu.mutation
}

// ClearVideos clears all "videos" edges to the Video entity.
func (cu *ChannelUpdate) ClearVideos() *ChannelUpdate {
	cu.mutation.ClearVideos()
	return cu
}

// RemoveVideoIDs removes the "videos" edge to Video entities by IDs.
func (cu *ChannelUpdate) RemoveVideoIDs(ids ...pulid.ID) *ChannelUpdate {
	cu.mutation.RemoveVideoIDs(ids...)
	return cu
}

// RemoveVideos removes "videos" edges to Video entities.
func (cu *ChannelUpdate) RemoveVideos(v ...*Video) *ChannelUpdate {
	ids := make([]pulid.ID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cu.RemoveVideoIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ChannelUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChannelUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChannelUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChannelUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *ChannelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(channel.Table, channel.Columns, sqlgraph.NewFieldSpec(channel.FieldID, field.TypeString))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.DisplayName(); ok {
		_spec.SetField(channel.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := cu.mutation.ChannelID(); ok {
		_spec.SetField(channel.FieldChannelID, field.TypeString, value)
	}
	if value, ok := cu.mutation.Handle(); ok {
		_spec.SetField(channel.FieldHandle, field.TypeString, value)
	}
	if value, ok := cu.mutation.ThumbnailURL(); ok {
		_spec.SetField(channel.FieldThumbnailURL, field.TypeString, value)
	}
	if cu.mutation.VideosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   channel.VideosTable,
			Columns: channel.VideosPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(video.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedVideosIDs(); len(nodes) > 0 && !cu.mutation.VideosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   channel.VideosTable,
			Columns: channel.VideosPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(video.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.VideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   channel.VideosTable,
			Columns: channel.VideosPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(video.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{channel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ChannelUpdateOne is the builder for updating a single Channel entity.
type ChannelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChannelMutation
}

// SetDisplayName sets the "display_name" field.
func (cuo *ChannelUpdateOne) SetDisplayName(s string) *ChannelUpdateOne {
	cuo.mutation.SetDisplayName(s)
	return cuo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (cuo *ChannelUpdateOne) SetNillableDisplayName(s *string) *ChannelUpdateOne {
	if s != nil {
		cuo.SetDisplayName(*s)
	}
	return cuo
}

// SetChannelID sets the "channel_id" field.
func (cuo *ChannelUpdateOne) SetChannelID(s string) *ChannelUpdateOne {
	cuo.mutation.SetChannelID(s)
	return cuo
}

// SetNillableChannelID sets the "channel_id" field if the given value is not nil.
func (cuo *ChannelUpdateOne) SetNillableChannelID(s *string) *ChannelUpdateOne {
	if s != nil {
		cuo.SetChannelID(*s)
	}
	return cuo
}

// SetHandle sets the "handle" field.
func (cuo *ChannelUpdateOne) SetHandle(s string) *ChannelUpdateOne {
	cuo.mutation.SetHandle(s)
	return cuo
}

// SetNillableHandle sets the "handle" field if the given value is not nil.
func (cuo *ChannelUpdateOne) SetNillableHandle(s *string) *ChannelUpdateOne {
	if s != nil {
		cuo.SetHandle(*s)
	}
	return cuo
}

// SetThumbnailURL sets the "thumbnail_url" field.
func (cuo *ChannelUpdateOne) SetThumbnailURL(s string) *ChannelUpdateOne {
	cuo.mutation.SetThumbnailURL(s)
	return cuo
}

// SetNillableThumbnailURL sets the "thumbnail_url" field if the given value is not nil.
func (cuo *ChannelUpdateOne) SetNillableThumbnailURL(s *string) *ChannelUpdateOne {
	if s != nil {
		cuo.SetThumbnailURL(*s)
	}
	return cuo
}

// AddVideoIDs adds the "videos" edge to the Video entity by IDs.
func (cuo *ChannelUpdateOne) AddVideoIDs(ids ...pulid.ID) *ChannelUpdateOne {
	cuo.mutation.AddVideoIDs(ids...)
	return cuo
}

// AddVideos adds the "videos" edges to the Video entity.
func (cuo *ChannelUpdateOne) AddVideos(v ...*Video) *ChannelUpdateOne {
	ids := make([]pulid.ID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cuo.AddVideoIDs(ids...)
}

// Mutation returns the ChannelMutation object of the builder.
func (cuo *ChannelUpdateOne) Mutation() *ChannelMutation {
	return cuo.mutation
}

// ClearVideos clears all "videos" edges to the Video entity.
func (cuo *ChannelUpdateOne) ClearVideos() *ChannelUpdateOne {
	cuo.mutation.ClearVideos()
	return cuo
}

// RemoveVideoIDs removes the "videos" edge to Video entities by IDs.
func (cuo *ChannelUpdateOne) RemoveVideoIDs(ids ...pulid.ID) *ChannelUpdateOne {
	cuo.mutation.RemoveVideoIDs(ids...)
	return cuo
}

// RemoveVideos removes "videos" edges to Video entities.
func (cuo *ChannelUpdateOne) RemoveVideos(v ...*Video) *ChannelUpdateOne {
	ids := make([]pulid.ID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cuo.RemoveVideoIDs(ids...)
}

// Where appends a list predicates to the ChannelUpdate builder.
func (cuo *ChannelUpdateOne) Where(ps ...predicate.Channel) *ChannelUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ChannelUpdateOne) Select(field string, fields ...string) *ChannelUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Channel entity.
func (cuo *ChannelUpdateOne) Save(ctx context.Context) (*Channel, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChannelUpdateOne) SaveX(ctx context.Context) *Channel {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ChannelUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChannelUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *ChannelUpdateOne) sqlSave(ctx context.Context) (_node *Channel, err error) {
	_spec := sqlgraph.NewUpdateSpec(channel.Table, channel.Columns, sqlgraph.NewFieldSpec(channel.FieldID, field.TypeString))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Channel.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, channel.FieldID)
		for _, f := range fields {
			if !channel.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != channel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.DisplayName(); ok {
		_spec.SetField(channel.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.ChannelID(); ok {
		_spec.SetField(channel.FieldChannelID, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Handle(); ok {
		_spec.SetField(channel.FieldHandle, field.TypeString, value)
	}
	if value, ok := cuo.mutation.ThumbnailURL(); ok {
		_spec.SetField(channel.FieldThumbnailURL, field.TypeString, value)
	}
	if cuo.mutation.VideosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   channel.VideosTable,
			Columns: channel.VideosPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(video.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedVideosIDs(); len(nodes) > 0 && !cuo.mutation.VideosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   channel.VideosTable,
			Columns: channel.VideosPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(video.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.VideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   channel.VideosTable,
			Columns: channel.VideosPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(video.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Channel{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{channel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}