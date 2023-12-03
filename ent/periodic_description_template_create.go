// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/description"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/periodic_description_template"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
)

// PeriodicDescriptionTemplateCreate is the builder for creating a Periodic_description_template entity.
type PeriodicDescriptionTemplateCreate struct {
	config
	mutation *PeriodicDescriptionTemplateMutation
	hooks    []Hook
}

// SetText sets the "text" field.
func (pdtc *PeriodicDescriptionTemplateCreate) SetText(s string) *PeriodicDescriptionTemplateCreate {
	pdtc.mutation.SetText(s)
	return pdtc
}

// SetStartUseAt sets the "start_use_at" field.
func (pdtc *PeriodicDescriptionTemplateCreate) SetStartUseAt(t time.Time) *PeriodicDescriptionTemplateCreate {
	pdtc.mutation.SetStartUseAt(t)
	return pdtc
}

// SetNillableStartUseAt sets the "start_use_at" field if the given value is not nil.
func (pdtc *PeriodicDescriptionTemplateCreate) SetNillableStartUseAt(t *time.Time) *PeriodicDescriptionTemplateCreate {
	if t != nil {
		pdtc.SetStartUseAt(*t)
	}
	return pdtc
}

// SetLastUseAt sets the "last_use_at" field.
func (pdtc *PeriodicDescriptionTemplateCreate) SetLastUseAt(t time.Time) *PeriodicDescriptionTemplateCreate {
	pdtc.mutation.SetLastUseAt(t)
	return pdtc
}

// SetNillableLastUseAt sets the "last_use_at" field if the given value is not nil.
func (pdtc *PeriodicDescriptionTemplateCreate) SetNillableLastUseAt(t *time.Time) *PeriodicDescriptionTemplateCreate {
	if t != nil {
		pdtc.SetLastUseAt(*t)
	}
	return pdtc
}

// SetID sets the "id" field.
func (pdtc *PeriodicDescriptionTemplateCreate) SetID(pu pulid.ID) *PeriodicDescriptionTemplateCreate {
	pdtc.mutation.SetID(pu)
	return pdtc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pdtc *PeriodicDescriptionTemplateCreate) SetNillableID(pu *pulid.ID) *PeriodicDescriptionTemplateCreate {
	if pu != nil {
		pdtc.SetID(*pu)
	}
	return pdtc
}

// AddDescriptionIDs adds the "descriptions" edge to the Description entity by IDs.
func (pdtc *PeriodicDescriptionTemplateCreate) AddDescriptionIDs(ids ...pulid.ID) *PeriodicDescriptionTemplateCreate {
	pdtc.mutation.AddDescriptionIDs(ids...)
	return pdtc
}

// AddDescriptions adds the "descriptions" edges to the Description entity.
func (pdtc *PeriodicDescriptionTemplateCreate) AddDescriptions(d ...*Description) *PeriodicDescriptionTemplateCreate {
	ids := make([]pulid.ID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pdtc.AddDescriptionIDs(ids...)
}

// Mutation returns the PeriodicDescriptionTemplateMutation object of the builder.
func (pdtc *PeriodicDescriptionTemplateCreate) Mutation() *PeriodicDescriptionTemplateMutation {
	return pdtc.mutation
}

// Save creates the Periodic_description_template in the database.
func (pdtc *PeriodicDescriptionTemplateCreate) Save(ctx context.Context) (*Periodic_description_template, error) {
	pdtc.defaults()
	return withHooks(ctx, pdtc.sqlSave, pdtc.mutation, pdtc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pdtc *PeriodicDescriptionTemplateCreate) SaveX(ctx context.Context) *Periodic_description_template {
	v, err := pdtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pdtc *PeriodicDescriptionTemplateCreate) Exec(ctx context.Context) error {
	_, err := pdtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pdtc *PeriodicDescriptionTemplateCreate) ExecX(ctx context.Context) {
	if err := pdtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pdtc *PeriodicDescriptionTemplateCreate) defaults() {
	if _, ok := pdtc.mutation.ID(); !ok {
		v := periodic_description_template.DefaultID()
		pdtc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pdtc *PeriodicDescriptionTemplateCreate) check() error {
	if _, ok := pdtc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "Periodic_description_template.text"`)}
	}
	return nil
}

func (pdtc *PeriodicDescriptionTemplateCreate) sqlSave(ctx context.Context) (*Periodic_description_template, error) {
	if err := pdtc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pdtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pdtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(pulid.ID); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Periodic_description_template.ID type: %T", _spec.ID.Value)
		}
	}
	pdtc.mutation.id = &_node.ID
	pdtc.mutation.done = true
	return _node, nil
}

func (pdtc *PeriodicDescriptionTemplateCreate) createSpec() (*Periodic_description_template, *sqlgraph.CreateSpec) {
	var (
		_node = &Periodic_description_template{config: pdtc.config}
		_spec = sqlgraph.NewCreateSpec(periodic_description_template.Table, sqlgraph.NewFieldSpec(periodic_description_template.FieldID, field.TypeString))
	)
	if id, ok := pdtc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pdtc.mutation.Text(); ok {
		_spec.SetField(periodic_description_template.FieldText, field.TypeString, value)
		_node.Text = value
	}
	if value, ok := pdtc.mutation.StartUseAt(); ok {
		_spec.SetField(periodic_description_template.FieldStartUseAt, field.TypeTime, value)
		_node.StartUseAt = value
	}
	if value, ok := pdtc.mutation.LastUseAt(); ok {
		_spec.SetField(periodic_description_template.FieldLastUseAt, field.TypeTime, value)
		_node.LastUseAt = value
	}
	if nodes := pdtc.mutation.DescriptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   periodic_description_template.DescriptionsTable,
			Columns: []string{periodic_description_template.DescriptionsColumn},
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
	return _node, _spec
}

// PeriodicDescriptionTemplateCreateBulk is the builder for creating many Periodic_description_template entities in bulk.
type PeriodicDescriptionTemplateCreateBulk struct {
	config
	err      error
	builders []*PeriodicDescriptionTemplateCreate
}

// Save creates the Periodic_description_template entities in the database.
func (pdtcb *PeriodicDescriptionTemplateCreateBulk) Save(ctx context.Context) ([]*Periodic_description_template, error) {
	if pdtcb.err != nil {
		return nil, pdtcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pdtcb.builders))
	nodes := make([]*Periodic_description_template, len(pdtcb.builders))
	mutators := make([]Mutator, len(pdtcb.builders))
	for i := range pdtcb.builders {
		func(i int, root context.Context) {
			builder := pdtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PeriodicDescriptionTemplateMutation)
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
					_, err = mutators[i+1].Mutate(root, pdtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pdtcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pdtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pdtcb *PeriodicDescriptionTemplateCreateBulk) SaveX(ctx context.Context) []*Periodic_description_template {
	v, err := pdtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pdtcb *PeriodicDescriptionTemplateCreateBulk) Exec(ctx context.Context) error {
	_, err := pdtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pdtcb *PeriodicDescriptionTemplateCreateBulk) ExecX(ctx context.Context) {
	if err := pdtcb.Exec(ctx); err != nil {
		panic(err)
	}
}