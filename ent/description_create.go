// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/categorydescriptiontemplate"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/description"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/descriptionchange"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/periodicdescriptiontemplate"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/video"
)

// DescriptionCreate is the builder for creating a Description entity.
type DescriptionCreate struct {
	config
	mutation *DescriptionMutation
	hooks    []Hook
}

// SetRaw sets the "raw" field.
func (dc *DescriptionCreate) SetRaw(s string) *DescriptionCreate {
	dc.mutation.SetRaw(s)
	return dc
}

// SetVariable sets the "variable" field.
func (dc *DescriptionCreate) SetVariable(s string) *DescriptionCreate {
	dc.mutation.SetVariable(s)
	return dc
}

// SetNillableVariable sets the "variable" field if the given value is not nil.
func (dc *DescriptionCreate) SetNillableVariable(s *string) *DescriptionCreate {
	if s != nil {
		dc.SetVariable(*s)
	}
	return dc
}

// SetNormalizedVariable sets the "normalized_variable" field.
func (dc *DescriptionCreate) SetNormalizedVariable(s string) *DescriptionCreate {
	dc.mutation.SetNormalizedVariable(s)
	return dc
}

// SetNillableNormalizedVariable sets the "normalized_variable" field if the given value is not nil.
func (dc *DescriptionCreate) SetNillableNormalizedVariable(s *string) *DescriptionCreate {
	if s != nil {
		dc.SetNormalizedVariable(*s)
	}
	return dc
}

// SetCreatedAt sets the "created_at" field.
func (dc *DescriptionCreate) SetCreatedAt(t time.Time) *DescriptionCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DescriptionCreate) SetNillableCreatedAt(t *time.Time) *DescriptionCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// SetUpdatedAt sets the "updated_at" field.
func (dc *DescriptionCreate) SetUpdatedAt(t time.Time) *DescriptionCreate {
	dc.mutation.SetUpdatedAt(t)
	return dc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dc *DescriptionCreate) SetNillableUpdatedAt(t *time.Time) *DescriptionCreate {
	if t != nil {
		dc.SetUpdatedAt(*t)
	}
	return dc
}

// SetID sets the "id" field.
func (dc *DescriptionCreate) SetID(s string) *DescriptionCreate {
	dc.mutation.SetID(s)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DescriptionCreate) SetNillableID(s *string) *DescriptionCreate {
	if s != nil {
		dc.SetID(*s)
	}
	return dc
}

// SetVideoID sets the "video" edge to the Video entity by ID.
func (dc *DescriptionCreate) SetVideoID(id string) *DescriptionCreate {
	dc.mutation.SetVideoID(id)
	return dc
}

// SetVideo sets the "video" edge to the Video entity.
func (dc *DescriptionCreate) SetVideo(v *Video) *DescriptionCreate {
	return dc.SetVideoID(v.ID)
}

// SetPeriodicTemplateID sets the "periodic_template" edge to the PeriodicDescriptionTemplate entity by ID.
func (dc *DescriptionCreate) SetPeriodicTemplateID(id string) *DescriptionCreate {
	dc.mutation.SetPeriodicTemplateID(id)
	return dc
}

// SetNillablePeriodicTemplateID sets the "periodic_template" edge to the PeriodicDescriptionTemplate entity by ID if the given value is not nil.
func (dc *DescriptionCreate) SetNillablePeriodicTemplateID(id *string) *DescriptionCreate {
	if id != nil {
		dc = dc.SetPeriodicTemplateID(*id)
	}
	return dc
}

// SetPeriodicTemplate sets the "periodic_template" edge to the PeriodicDescriptionTemplate entity.
func (dc *DescriptionCreate) SetPeriodicTemplate(p *PeriodicDescriptionTemplate) *DescriptionCreate {
	return dc.SetPeriodicTemplateID(p.ID)
}

// SetCategoryTemplateID sets the "category_template" edge to the CategoryDescriptionTemplate entity by ID.
func (dc *DescriptionCreate) SetCategoryTemplateID(id string) *DescriptionCreate {
	dc.mutation.SetCategoryTemplateID(id)
	return dc
}

// SetNillableCategoryTemplateID sets the "category_template" edge to the CategoryDescriptionTemplate entity by ID if the given value is not nil.
func (dc *DescriptionCreate) SetNillableCategoryTemplateID(id *string) *DescriptionCreate {
	if id != nil {
		dc = dc.SetCategoryTemplateID(*id)
	}
	return dc
}

// SetCategoryTemplate sets the "category_template" edge to the CategoryDescriptionTemplate entity.
func (dc *DescriptionCreate) SetCategoryTemplate(c *CategoryDescriptionTemplate) *DescriptionCreate {
	return dc.SetCategoryTemplateID(c.ID)
}

// AddDescriptionChangeIDs adds the "description_changes" edge to the DescriptionChange entity by IDs.
func (dc *DescriptionCreate) AddDescriptionChangeIDs(ids ...string) *DescriptionCreate {
	dc.mutation.AddDescriptionChangeIDs(ids...)
	return dc
}

// AddDescriptionChanges adds the "description_changes" edges to the DescriptionChange entity.
func (dc *DescriptionCreate) AddDescriptionChanges(d ...*DescriptionChange) *DescriptionCreate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddDescriptionChangeIDs(ids...)
}

// Mutation returns the DescriptionMutation object of the builder.
func (dc *DescriptionCreate) Mutation() *DescriptionMutation {
	return dc.mutation
}

// Save creates the Description in the database.
func (dc *DescriptionCreate) Save(ctx context.Context) (*Description, error) {
	dc.defaults()
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DescriptionCreate) SaveX(ctx context.Context) *Description {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DescriptionCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DescriptionCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DescriptionCreate) defaults() {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		v := description.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		v := description.DefaultUpdatedAt()
		dc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dc.mutation.ID(); !ok {
		v := description.DefaultID()
		dc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DescriptionCreate) check() error {
	if _, ok := dc.mutation.Raw(); !ok {
		return &ValidationError{Name: "raw", err: errors.New(`ent: missing required field "Description.raw"`)}
	}
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Description.created_at"`)}
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Description.updated_at"`)}
	}
	if v, ok := dc.mutation.ID(); ok {
		if err := description.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Description.id": %w`, err)}
		}
	}
	if _, ok := dc.mutation.VideoID(); !ok {
		return &ValidationError{Name: "video", err: errors.New(`ent: missing required edge "Description.video"`)}
	}
	return nil
}

func (dc *DescriptionCreate) sqlSave(ctx context.Context) (*Description, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Description.ID type: %T", _spec.ID.Value)
		}
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DescriptionCreate) createSpec() (*Description, *sqlgraph.CreateSpec) {
	var (
		_node = &Description{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(description.Table, sqlgraph.NewFieldSpec(description.FieldID, field.TypeString))
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.Raw(); ok {
		_spec.SetField(description.FieldRaw, field.TypeString, value)
		_node.Raw = value
	}
	if value, ok := dc.mutation.Variable(); ok {
		_spec.SetField(description.FieldVariable, field.TypeString, value)
		_node.Variable = value
	}
	if value, ok := dc.mutation.NormalizedVariable(); ok {
		_spec.SetField(description.FieldNormalizedVariable, field.TypeString, value)
		_node.NormalizedVariable = value
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.SetField(description.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := dc.mutation.UpdatedAt(); ok {
		_spec.SetField(description.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := dc.mutation.VideoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   description.VideoTable,
			Columns: []string{description.VideoColumn},
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
	if nodes := dc.mutation.PeriodicTemplateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   description.PeriodicTemplateTable,
			Columns: []string{description.PeriodicTemplateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(periodicdescriptiontemplate.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.periodic_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.CategoryTemplateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   description.CategoryTemplateTable,
			Columns: []string{description.CategoryTemplateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(categorydescriptiontemplate.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.category_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.DescriptionChangesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   description.DescriptionChangesTable,
			Columns: []string{description.DescriptionChangesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(descriptionchange.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DescriptionCreateBulk is the builder for creating many Description entities in bulk.
type DescriptionCreateBulk struct {
	config
	err      error
	builders []*DescriptionCreate
}

// Save creates the Description entities in the database.
func (dcb *DescriptionCreateBulk) Save(ctx context.Context) ([]*Description, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Description, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DescriptionMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DescriptionCreateBulk) SaveX(ctx context.Context) []*Description {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DescriptionCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DescriptionCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
