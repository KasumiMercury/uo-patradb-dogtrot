// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/category_description_template"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/description"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
)

// CategoryDescriptionTemplateCreate is the builder for creating a Category_description_template entity.
type CategoryDescriptionTemplateCreate struct {
	config
	mutation *CategoryDescriptionTemplateMutation
	hooks    []Hook
}

// SetText sets the "text" field.
func (cdtc *CategoryDescriptionTemplateCreate) SetText(s string) *CategoryDescriptionTemplateCreate {
	cdtc.mutation.SetText(s)
	return cdtc
}

// SetStartUseAt sets the "start_use_at" field.
func (cdtc *CategoryDescriptionTemplateCreate) SetStartUseAt(t time.Time) *CategoryDescriptionTemplateCreate {
	cdtc.mutation.SetStartUseAt(t)
	return cdtc
}

// SetNillableStartUseAt sets the "start_use_at" field if the given value is not nil.
func (cdtc *CategoryDescriptionTemplateCreate) SetNillableStartUseAt(t *time.Time) *CategoryDescriptionTemplateCreate {
	if t != nil {
		cdtc.SetStartUseAt(*t)
	}
	return cdtc
}

// SetLastUseAt sets the "last_use_at" field.
func (cdtc *CategoryDescriptionTemplateCreate) SetLastUseAt(t time.Time) *CategoryDescriptionTemplateCreate {
	cdtc.mutation.SetLastUseAt(t)
	return cdtc
}

// SetNillableLastUseAt sets the "last_use_at" field if the given value is not nil.
func (cdtc *CategoryDescriptionTemplateCreate) SetNillableLastUseAt(t *time.Time) *CategoryDescriptionTemplateCreate {
	if t != nil {
		cdtc.SetLastUseAt(*t)
	}
	return cdtc
}

// SetID sets the "id" field.
func (cdtc *CategoryDescriptionTemplateCreate) SetID(pu pulid.ID) *CategoryDescriptionTemplateCreate {
	cdtc.mutation.SetID(pu)
	return cdtc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cdtc *CategoryDescriptionTemplateCreate) SetNillableID(pu *pulid.ID) *CategoryDescriptionTemplateCreate {
	if pu != nil {
		cdtc.SetID(*pu)
	}
	return cdtc
}

// AddDescriptionIDs adds the "descriptions" edge to the Description entity by IDs.
func (cdtc *CategoryDescriptionTemplateCreate) AddDescriptionIDs(ids ...pulid.ID) *CategoryDescriptionTemplateCreate {
	cdtc.mutation.AddDescriptionIDs(ids...)
	return cdtc
}

// AddDescriptions adds the "descriptions" edges to the Description entity.
func (cdtc *CategoryDescriptionTemplateCreate) AddDescriptions(d ...*Description) *CategoryDescriptionTemplateCreate {
	ids := make([]pulid.ID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return cdtc.AddDescriptionIDs(ids...)
}

// Mutation returns the CategoryDescriptionTemplateMutation object of the builder.
func (cdtc *CategoryDescriptionTemplateCreate) Mutation() *CategoryDescriptionTemplateMutation {
	return cdtc.mutation
}

// Save creates the Category_description_template in the database.
func (cdtc *CategoryDescriptionTemplateCreate) Save(ctx context.Context) (*Category_description_template, error) {
	cdtc.defaults()
	return withHooks(ctx, cdtc.sqlSave, cdtc.mutation, cdtc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cdtc *CategoryDescriptionTemplateCreate) SaveX(ctx context.Context) *Category_description_template {
	v, err := cdtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cdtc *CategoryDescriptionTemplateCreate) Exec(ctx context.Context) error {
	_, err := cdtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cdtc *CategoryDescriptionTemplateCreate) ExecX(ctx context.Context) {
	if err := cdtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cdtc *CategoryDescriptionTemplateCreate) defaults() {
	if _, ok := cdtc.mutation.ID(); !ok {
		v := category_description_template.DefaultID()
		cdtc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cdtc *CategoryDescriptionTemplateCreate) check() error {
	if _, ok := cdtc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "Category_description_template.text"`)}
	}
	return nil
}

func (cdtc *CategoryDescriptionTemplateCreate) sqlSave(ctx context.Context) (*Category_description_template, error) {
	if err := cdtc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cdtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cdtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(pulid.ID); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Category_description_template.ID type: %T", _spec.ID.Value)
		}
	}
	cdtc.mutation.id = &_node.ID
	cdtc.mutation.done = true
	return _node, nil
}

func (cdtc *CategoryDescriptionTemplateCreate) createSpec() (*Category_description_template, *sqlgraph.CreateSpec) {
	var (
		_node = &Category_description_template{config: cdtc.config}
		_spec = sqlgraph.NewCreateSpec(category_description_template.Table, sqlgraph.NewFieldSpec(category_description_template.FieldID, field.TypeString))
	)
	if id, ok := cdtc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cdtc.mutation.Text(); ok {
		_spec.SetField(category_description_template.FieldText, field.TypeString, value)
		_node.Text = value
	}
	if value, ok := cdtc.mutation.StartUseAt(); ok {
		_spec.SetField(category_description_template.FieldStartUseAt, field.TypeTime, value)
		_node.StartUseAt = value
	}
	if value, ok := cdtc.mutation.LastUseAt(); ok {
		_spec.SetField(category_description_template.FieldLastUseAt, field.TypeTime, value)
		_node.LastUseAt = value
	}
	if nodes := cdtc.mutation.DescriptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category_description_template.DescriptionsTable,
			Columns: []string{category_description_template.DescriptionsColumn},
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

// CategoryDescriptionTemplateCreateBulk is the builder for creating many Category_description_template entities in bulk.
type CategoryDescriptionTemplateCreateBulk struct {
	config
	err      error
	builders []*CategoryDescriptionTemplateCreate
}

// Save creates the Category_description_template entities in the database.
func (cdtcb *CategoryDescriptionTemplateCreateBulk) Save(ctx context.Context) ([]*Category_description_template, error) {
	if cdtcb.err != nil {
		return nil, cdtcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cdtcb.builders))
	nodes := make([]*Category_description_template, len(cdtcb.builders))
	mutators := make([]Mutator, len(cdtcb.builders))
	for i := range cdtcb.builders {
		func(i int, root context.Context) {
			builder := cdtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CategoryDescriptionTemplateMutation)
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
					_, err = mutators[i+1].Mutate(root, cdtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cdtcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cdtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cdtcb *CategoryDescriptionTemplateCreateBulk) SaveX(ctx context.Context) []*Category_description_template {
	v, err := cdtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cdtcb *CategoryDescriptionTemplateCreateBulk) Exec(ctx context.Context) error {
	_, err := cdtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cdtcb *CategoryDescriptionTemplateCreateBulk) ExecX(ctx context.Context) {
	if err := cdtcb.Exec(ctx); err != nil {
		panic(err)
	}
}