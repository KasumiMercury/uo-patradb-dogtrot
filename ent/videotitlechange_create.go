// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/video"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/videotitlechange"
)

// VideoTitleChangeCreate is the builder for creating a VideoTitleChange entity.
type VideoTitleChangeCreate struct {
	config
	mutation *VideoTitleChangeMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (vtcc *VideoTitleChangeCreate) SetTitle(s string) *VideoTitleChangeCreate {
	vtcc.mutation.SetTitle(s)
	return vtcc
}

// SetNormalizedTitle sets the "normalized_title" field.
func (vtcc *VideoTitleChangeCreate) SetNormalizedTitle(s string) *VideoTitleChangeCreate {
	vtcc.mutation.SetNormalizedTitle(s)
	return vtcc
}

// SetChangedAt sets the "changed_at" field.
func (vtcc *VideoTitleChangeCreate) SetChangedAt(t time.Time) *VideoTitleChangeCreate {
	vtcc.mutation.SetChangedAt(t)
	return vtcc
}

// SetNillableChangedAt sets the "changed_at" field if the given value is not nil.
func (vtcc *VideoTitleChangeCreate) SetNillableChangedAt(t *time.Time) *VideoTitleChangeCreate {
	if t != nil {
		vtcc.SetChangedAt(*t)
	}
	return vtcc
}

// SetID sets the "id" field.
func (vtcc *VideoTitleChangeCreate) SetID(s string) *VideoTitleChangeCreate {
	vtcc.mutation.SetID(s)
	return vtcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (vtcc *VideoTitleChangeCreate) SetNillableID(s *string) *VideoTitleChangeCreate {
	if s != nil {
		vtcc.SetID(*s)
	}
	return vtcc
}

// SetVideoID sets the "video" edge to the Video entity by ID.
func (vtcc *VideoTitleChangeCreate) SetVideoID(id string) *VideoTitleChangeCreate {
	vtcc.mutation.SetVideoID(id)
	return vtcc
}

// SetVideo sets the "video" edge to the Video entity.
func (vtcc *VideoTitleChangeCreate) SetVideo(v *Video) *VideoTitleChangeCreate {
	return vtcc.SetVideoID(v.ID)
}

// Mutation returns the VideoTitleChangeMutation object of the builder.
func (vtcc *VideoTitleChangeCreate) Mutation() *VideoTitleChangeMutation {
	return vtcc.mutation
}

// Save creates the VideoTitleChange in the database.
func (vtcc *VideoTitleChangeCreate) Save(ctx context.Context) (*VideoTitleChange, error) {
	vtcc.defaults()
	return withHooks(ctx, vtcc.sqlSave, vtcc.mutation, vtcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vtcc *VideoTitleChangeCreate) SaveX(ctx context.Context) *VideoTitleChange {
	v, err := vtcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vtcc *VideoTitleChangeCreate) Exec(ctx context.Context) error {
	_, err := vtcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vtcc *VideoTitleChangeCreate) ExecX(ctx context.Context) {
	if err := vtcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vtcc *VideoTitleChangeCreate) defaults() {
	if _, ok := vtcc.mutation.ChangedAt(); !ok {
		v := videotitlechange.DefaultChangedAt()
		vtcc.mutation.SetChangedAt(v)
	}
	if _, ok := vtcc.mutation.ID(); !ok {
		v := videotitlechange.DefaultID()
		vtcc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vtcc *VideoTitleChangeCreate) check() error {
	if _, ok := vtcc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "VideoTitleChange.title"`)}
	}
	if _, ok := vtcc.mutation.NormalizedTitle(); !ok {
		return &ValidationError{Name: "normalized_title", err: errors.New(`ent: missing required field "VideoTitleChange.normalized_title"`)}
	}
	if _, ok := vtcc.mutation.ChangedAt(); !ok {
		return &ValidationError{Name: "changed_at", err: errors.New(`ent: missing required field "VideoTitleChange.changed_at"`)}
	}
	if v, ok := vtcc.mutation.ID(); ok {
		if err := videotitlechange.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "VideoTitleChange.id": %w`, err)}
		}
	}
	if _, ok := vtcc.mutation.VideoID(); !ok {
		return &ValidationError{Name: "video", err: errors.New(`ent: missing required edge "VideoTitleChange.video"`)}
	}
	return nil
}

func (vtcc *VideoTitleChangeCreate) sqlSave(ctx context.Context) (*VideoTitleChange, error) {
	if err := vtcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vtcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vtcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected VideoTitleChange.ID type: %T", _spec.ID.Value)
		}
	}
	vtcc.mutation.id = &_node.ID
	vtcc.mutation.done = true
	return _node, nil
}

func (vtcc *VideoTitleChangeCreate) createSpec() (*VideoTitleChange, *sqlgraph.CreateSpec) {
	var (
		_node = &VideoTitleChange{config: vtcc.config}
		_spec = sqlgraph.NewCreateSpec(videotitlechange.Table, sqlgraph.NewFieldSpec(videotitlechange.FieldID, field.TypeString))
	)
	if id, ok := vtcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := vtcc.mutation.Title(); ok {
		_spec.SetField(videotitlechange.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := vtcc.mutation.NormalizedTitle(); ok {
		_spec.SetField(videotitlechange.FieldNormalizedTitle, field.TypeString, value)
		_node.NormalizedTitle = value
	}
	if value, ok := vtcc.mutation.ChangedAt(); ok {
		_spec.SetField(videotitlechange.FieldChangedAt, field.TypeTime, value)
		_node.ChangedAt = value
	}
	if nodes := vtcc.mutation.VideoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   videotitlechange.VideoTable,
			Columns: []string{videotitlechange.VideoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(video.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.video_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VideoTitleChangeCreateBulk is the builder for creating many VideoTitleChange entities in bulk.
type VideoTitleChangeCreateBulk struct {
	config
	err      error
	builders []*VideoTitleChangeCreate
}

// Save creates the VideoTitleChange entities in the database.
func (vtccb *VideoTitleChangeCreateBulk) Save(ctx context.Context) ([]*VideoTitleChange, error) {
	if vtccb.err != nil {
		return nil, vtccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(vtccb.builders))
	nodes := make([]*VideoTitleChange, len(vtccb.builders))
	mutators := make([]Mutator, len(vtccb.builders))
	for i := range vtccb.builders {
		func(i int, root context.Context) {
			builder := vtccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VideoTitleChangeMutation)
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
					_, err = mutators[i+1].Mutate(root, vtccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vtccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vtccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vtccb *VideoTitleChangeCreateBulk) SaveX(ctx context.Context) []*VideoTitleChange {
	v, err := vtccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vtccb *VideoTitleChangeCreateBulk) Exec(ctx context.Context) error {
	_, err := vtccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vtccb *VideoTitleChangeCreateBulk) ExecX(ctx context.Context) {
	if err := vtccb.Exec(ctx); err != nil {
		panic(err)
	}
}
