// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/channel"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/description"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/patchat"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/video"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/videodisallowrange"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/videoplayrange"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/videotitlechange"
)

// VideoCreate is the builder for creating a Video entity.
type VideoCreate struct {
	config
	mutation *VideoMutation
	hooks    []Hook
}

// SetVideoID sets the "video_id" field.
func (vc *VideoCreate) SetVideoID(s string) *VideoCreate {
	vc.mutation.SetVideoID(s)
	return vc
}

// SetTitle sets the "title" field.
func (vc *VideoCreate) SetTitle(s string) *VideoCreate {
	vc.mutation.SetTitle(s)
	return vc
}

// SetNormalizedTitle sets the "normalized_title" field.
func (vc *VideoCreate) SetNormalizedTitle(s string) *VideoCreate {
	vc.mutation.SetNormalizedTitle(s)
	return vc
}

// SetDurationSeconds sets the "duration_seconds" field.
func (vc *VideoCreate) SetDurationSeconds(i int) *VideoCreate {
	vc.mutation.SetDurationSeconds(i)
	return vc
}

// SetIsCollaboration sets the "is_collaboration" field.
func (vc *VideoCreate) SetIsCollaboration(b bool) *VideoCreate {
	vc.mutation.SetIsCollaboration(b)
	return vc
}

// SetNillableIsCollaboration sets the "is_collaboration" field if the given value is not nil.
func (vc *VideoCreate) SetNillableIsCollaboration(b *bool) *VideoCreate {
	if b != nil {
		vc.SetIsCollaboration(*b)
	}
	return vc
}

// SetStatus sets the "status" field.
func (vc *VideoCreate) SetStatus(s string) *VideoCreate {
	vc.mutation.SetStatus(s)
	return vc
}

// SetChatID sets the "chat_id" field.
func (vc *VideoCreate) SetChatID(s string) *VideoCreate {
	vc.mutation.SetChatID(s)
	return vc
}

// SetNillableChatID sets the "chat_id" field if the given value is not nil.
func (vc *VideoCreate) SetNillableChatID(s *string) *VideoCreate {
	if s != nil {
		vc.SetChatID(*s)
	}
	return vc
}

// SetHasTimeRange sets the "has_time_range" field.
func (vc *VideoCreate) SetHasTimeRange(b bool) *VideoCreate {
	vc.mutation.SetHasTimeRange(b)
	return vc
}

// SetNillableHasTimeRange sets the "has_time_range" field if the given value is not nil.
func (vc *VideoCreate) SetNillableHasTimeRange(b *bool) *VideoCreate {
	if b != nil {
		vc.SetHasTimeRange(*b)
	}
	return vc
}

// SetScheduledAt sets the "scheduled_at" field.
func (vc *VideoCreate) SetScheduledAt(t time.Time) *VideoCreate {
	vc.mutation.SetScheduledAt(t)
	return vc
}

// SetNillableScheduledAt sets the "scheduled_at" field if the given value is not nil.
func (vc *VideoCreate) SetNillableScheduledAt(t *time.Time) *VideoCreate {
	if t != nil {
		vc.SetScheduledAt(*t)
	}
	return vc
}

// SetActualStartAt sets the "actual_start_at" field.
func (vc *VideoCreate) SetActualStartAt(t time.Time) *VideoCreate {
	vc.mutation.SetActualStartAt(t)
	return vc
}

// SetNillableActualStartAt sets the "actual_start_at" field if the given value is not nil.
func (vc *VideoCreate) SetNillableActualStartAt(t *time.Time) *VideoCreate {
	if t != nil {
		vc.SetActualStartAt(*t)
	}
	return vc
}

// SetPublishedAt sets the "published_at" field.
func (vc *VideoCreate) SetPublishedAt(t time.Time) *VideoCreate {
	vc.mutation.SetPublishedAt(t)
	return vc
}

// SetCreatedAt sets the "created_at" field.
func (vc *VideoCreate) SetCreatedAt(t time.Time) *VideoCreate {
	vc.mutation.SetCreatedAt(t)
	return vc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vc *VideoCreate) SetNillableCreatedAt(t *time.Time) *VideoCreate {
	if t != nil {
		vc.SetCreatedAt(*t)
	}
	return vc
}

// SetUpdatedAt sets the "updated_at" field.
func (vc *VideoCreate) SetUpdatedAt(t time.Time) *VideoCreate {
	vc.mutation.SetUpdatedAt(t)
	return vc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (vc *VideoCreate) SetNillableUpdatedAt(t *time.Time) *VideoCreate {
	if t != nil {
		vc.SetUpdatedAt(*t)
	}
	return vc
}

// SetID sets the "id" field.
func (vc *VideoCreate) SetID(s string) *VideoCreate {
	vc.mutation.SetID(s)
	return vc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (vc *VideoCreate) SetNillableID(s *string) *VideoCreate {
	if s != nil {
		vc.SetID(*s)
	}
	return vc
}

// SetDescriptionsID sets the "descriptions" edge to the Description entity by ID.
func (vc *VideoCreate) SetDescriptionsID(id string) *VideoCreate {
	vc.mutation.SetDescriptionsID(id)
	return vc
}

// SetNillableDescriptionsID sets the "descriptions" edge to the Description entity by ID if the given value is not nil.
func (vc *VideoCreate) SetNillableDescriptionsID(id *string) *VideoCreate {
	if id != nil {
		vc = vc.SetDescriptionsID(*id)
	}
	return vc
}

// SetDescriptions sets the "descriptions" edge to the Description entity.
func (vc *VideoCreate) SetDescriptions(d *Description) *VideoCreate {
	return vc.SetDescriptionsID(d.ID)
}

// AddChannelIDs adds the "channel" edge to the Channel entity by IDs.
func (vc *VideoCreate) AddChannelIDs(ids ...string) *VideoCreate {
	vc.mutation.AddChannelIDs(ids...)
	return vc
}

// AddChannel adds the "channel" edges to the Channel entity.
func (vc *VideoCreate) AddChannel(c ...*Channel) *VideoCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return vc.AddChannelIDs(ids...)
}

// AddVideoPlayRangeIDs adds the "video_play_ranges" edge to the VideoPlayRange entity by IDs.
func (vc *VideoCreate) AddVideoPlayRangeIDs(ids ...string) *VideoCreate {
	vc.mutation.AddVideoPlayRangeIDs(ids...)
	return vc
}

// AddVideoPlayRanges adds the "video_play_ranges" edges to the VideoPlayRange entity.
func (vc *VideoCreate) AddVideoPlayRanges(v ...*VideoPlayRange) *VideoCreate {
	ids := make([]string, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vc.AddVideoPlayRangeIDs(ids...)
}

// AddVideoDisallowRangeIDs adds the "video_disallow_ranges" edge to the VideoDisallowRange entity by IDs.
func (vc *VideoCreate) AddVideoDisallowRangeIDs(ids ...string) *VideoCreate {
	vc.mutation.AddVideoDisallowRangeIDs(ids...)
	return vc
}

// AddVideoDisallowRanges adds the "video_disallow_ranges" edges to the VideoDisallowRange entity.
func (vc *VideoCreate) AddVideoDisallowRanges(v ...*VideoDisallowRange) *VideoCreate {
	ids := make([]string, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vc.AddVideoDisallowRangeIDs(ids...)
}

// AddVideoTitleChangeIDs adds the "video_title_changes" edge to the VideoTitleChange entity by IDs.
func (vc *VideoCreate) AddVideoTitleChangeIDs(ids ...string) *VideoCreate {
	vc.mutation.AddVideoTitleChangeIDs(ids...)
	return vc
}

// AddVideoTitleChanges adds the "video_title_changes" edges to the VideoTitleChange entity.
func (vc *VideoCreate) AddVideoTitleChanges(v ...*VideoTitleChange) *VideoCreate {
	ids := make([]string, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vc.AddVideoTitleChangeIDs(ids...)
}

// AddPatChatIDs adds the "Pat_chats" edge to the PatChat entity by IDs.
func (vc *VideoCreate) AddPatChatIDs(ids ...string) *VideoCreate {
	vc.mutation.AddPatChatIDs(ids...)
	return vc
}

// AddPatChats adds the "Pat_chats" edges to the PatChat entity.
func (vc *VideoCreate) AddPatChats(p ...*PatChat) *VideoCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return vc.AddPatChatIDs(ids...)
}

// Mutation returns the VideoMutation object of the builder.
func (vc *VideoCreate) Mutation() *VideoMutation {
	return vc.mutation
}

// Save creates the Video in the database.
func (vc *VideoCreate) Save(ctx context.Context) (*Video, error) {
	vc.defaults()
	return withHooks(ctx, vc.sqlSave, vc.mutation, vc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VideoCreate) SaveX(ctx context.Context) *Video {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VideoCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VideoCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vc *VideoCreate) defaults() {
	if _, ok := vc.mutation.IsCollaboration(); !ok {
		v := video.DefaultIsCollaboration
		vc.mutation.SetIsCollaboration(v)
	}
	if _, ok := vc.mutation.HasTimeRange(); !ok {
		v := video.DefaultHasTimeRange
		vc.mutation.SetHasTimeRange(v)
	}
	if _, ok := vc.mutation.CreatedAt(); !ok {
		v := video.DefaultCreatedAt()
		vc.mutation.SetCreatedAt(v)
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		v := video.DefaultUpdatedAt()
		vc.mutation.SetUpdatedAt(v)
	}
	if _, ok := vc.mutation.ID(); !ok {
		v := video.DefaultID()
		vc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VideoCreate) check() error {
	if _, ok := vc.mutation.VideoID(); !ok {
		return &ValidationError{Name: "video_id", err: errors.New(`ent: missing required field "Video.video_id"`)}
	}
	if v, ok := vc.mutation.VideoID(); ok {
		if err := video.VideoIDValidator(v); err != nil {
			return &ValidationError{Name: "video_id", err: fmt.Errorf(`ent: validator failed for field "Video.video_id": %w`, err)}
		}
	}
	if _, ok := vc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Video.title"`)}
	}
	if v, ok := vc.mutation.Title(); ok {
		if err := video.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Video.title": %w`, err)}
		}
	}
	if _, ok := vc.mutation.NormalizedTitle(); !ok {
		return &ValidationError{Name: "normalized_title", err: errors.New(`ent: missing required field "Video.normalized_title"`)}
	}
	if _, ok := vc.mutation.DurationSeconds(); !ok {
		return &ValidationError{Name: "duration_seconds", err: errors.New(`ent: missing required field "Video.duration_seconds"`)}
	}
	if v, ok := vc.mutation.DurationSeconds(); ok {
		if err := video.DurationSecondsValidator(v); err != nil {
			return &ValidationError{Name: "duration_seconds", err: fmt.Errorf(`ent: validator failed for field "Video.duration_seconds": %w`, err)}
		}
	}
	if _, ok := vc.mutation.IsCollaboration(); !ok {
		return &ValidationError{Name: "is_collaboration", err: errors.New(`ent: missing required field "Video.is_collaboration"`)}
	}
	if _, ok := vc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Video.status"`)}
	}
	if _, ok := vc.mutation.HasTimeRange(); !ok {
		return &ValidationError{Name: "has_time_range", err: errors.New(`ent: missing required field "Video.has_time_range"`)}
	}
	if _, ok := vc.mutation.PublishedAt(); !ok {
		return &ValidationError{Name: "published_at", err: errors.New(`ent: missing required field "Video.published_at"`)}
	}
	if _, ok := vc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Video.created_at"`)}
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Video.updated_at"`)}
	}
	if v, ok := vc.mutation.ID(); ok {
		if err := video.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Video.id": %w`, err)}
		}
	}
	return nil
}

func (vc *VideoCreate) sqlSave(ctx context.Context) (*Video, error) {
	if err := vc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Video.ID type: %T", _spec.ID.Value)
		}
	}
	vc.mutation.id = &_node.ID
	vc.mutation.done = true
	return _node, nil
}

func (vc *VideoCreate) createSpec() (*Video, *sqlgraph.CreateSpec) {
	var (
		_node = &Video{config: vc.config}
		_spec = sqlgraph.NewCreateSpec(video.Table, sqlgraph.NewFieldSpec(video.FieldID, field.TypeString))
	)
	if id, ok := vc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := vc.mutation.VideoID(); ok {
		_spec.SetField(video.FieldVideoID, field.TypeString, value)
		_node.VideoID = value
	}
	if value, ok := vc.mutation.Title(); ok {
		_spec.SetField(video.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := vc.mutation.NormalizedTitle(); ok {
		_spec.SetField(video.FieldNormalizedTitle, field.TypeString, value)
		_node.NormalizedTitle = value
	}
	if value, ok := vc.mutation.DurationSeconds(); ok {
		_spec.SetField(video.FieldDurationSeconds, field.TypeInt, value)
		_node.DurationSeconds = value
	}
	if value, ok := vc.mutation.IsCollaboration(); ok {
		_spec.SetField(video.FieldIsCollaboration, field.TypeBool, value)
		_node.IsCollaboration = value
	}
	if value, ok := vc.mutation.Status(); ok {
		_spec.SetField(video.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := vc.mutation.ChatID(); ok {
		_spec.SetField(video.FieldChatID, field.TypeString, value)
		_node.ChatID = value
	}
	if value, ok := vc.mutation.HasTimeRange(); ok {
		_spec.SetField(video.FieldHasTimeRange, field.TypeBool, value)
		_node.HasTimeRange = value
	}
	if value, ok := vc.mutation.ScheduledAt(); ok {
		_spec.SetField(video.FieldScheduledAt, field.TypeTime, value)
		_node.ScheduledAt = value
	}
	if value, ok := vc.mutation.ActualStartAt(); ok {
		_spec.SetField(video.FieldActualStartAt, field.TypeTime, value)
		_node.ActualStartAt = value
	}
	if value, ok := vc.mutation.PublishedAt(); ok {
		_spec.SetField(video.FieldPublishedAt, field.TypeTime, value)
		_node.PublishedAt = value
	}
	if value, ok := vc.mutation.CreatedAt(); ok {
		_spec.SetField(video.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := vc.mutation.UpdatedAt(); ok {
		_spec.SetField(video.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := vc.mutation.DescriptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   video.DescriptionsTable,
			Columns: []string{video.DescriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(description.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.ChannelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   video.ChannelTable,
			Columns: video.ChannelPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(channel.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.VideoPlayRangesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.VideoPlayRangesTable,
			Columns: []string{video.VideoPlayRangesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(videoplayrange.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.VideoDisallowRangesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.VideoDisallowRangesTable,
			Columns: []string{video.VideoDisallowRangesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(videodisallowrange.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.VideoTitleChangesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.VideoTitleChangesTable,
			Columns: []string{video.VideoTitleChangesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(videotitlechange.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.PatChatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.PatChatsTable,
			Columns: []string{video.PatChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(patchat.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VideoCreateBulk is the builder for creating many Video entities in bulk.
type VideoCreateBulk struct {
	config
	err      error
	builders []*VideoCreate
}

// Save creates the Video entities in the database.
func (vcb *VideoCreateBulk) Save(ctx context.Context) ([]*Video, error) {
	if vcb.err != nil {
		return nil, vcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Video, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VideoMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VideoCreateBulk) SaveX(ctx context.Context) []*Video {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VideoCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VideoCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
