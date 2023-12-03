// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/description"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/description_change"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/predicate"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
)

// DescriptionChangeUpdate is the builder for updating Description_change entities.
type DescriptionChangeUpdate struct {
	config
	hooks    []Hook
	mutation *DescriptionChangeMutation
}

// Where appends a list predicates to the DescriptionChangeUpdate builder.
func (dcu *DescriptionChangeUpdate) Where(ps ...predicate.Description_change) *DescriptionChangeUpdate {
	dcu.mutation.Where(ps...)
	return dcu
}

// SetRaw sets the "raw" field.
func (dcu *DescriptionChangeUpdate) SetRaw(s string) *DescriptionChangeUpdate {
	dcu.mutation.SetRaw(s)
	return dcu
}

// SetNillableRaw sets the "raw" field if the given value is not nil.
func (dcu *DescriptionChangeUpdate) SetNillableRaw(s *string) *DescriptionChangeUpdate {
	if s != nil {
		dcu.SetRaw(*s)
	}
	return dcu
}

// SetVariable sets the "variable" field.
func (dcu *DescriptionChangeUpdate) SetVariable(s string) *DescriptionChangeUpdate {
	dcu.mutation.SetVariable(s)
	return dcu
}

// SetNillableVariable sets the "variable" field if the given value is not nil.
func (dcu *DescriptionChangeUpdate) SetNillableVariable(s *string) *DescriptionChangeUpdate {
	if s != nil {
		dcu.SetVariable(*s)
	}
	return dcu
}

// ClearVariable clears the value of the "variable" field.
func (dcu *DescriptionChangeUpdate) ClearVariable() *DescriptionChangeUpdate {
	dcu.mutation.ClearVariable()
	return dcu
}

// SetNormalizedVariable sets the "normalized_variable" field.
func (dcu *DescriptionChangeUpdate) SetNormalizedVariable(s string) *DescriptionChangeUpdate {
	dcu.mutation.SetNormalizedVariable(s)
	return dcu
}

// SetNillableNormalizedVariable sets the "normalized_variable" field if the given value is not nil.
func (dcu *DescriptionChangeUpdate) SetNillableNormalizedVariable(s *string) *DescriptionChangeUpdate {
	if s != nil {
		dcu.SetNormalizedVariable(*s)
	}
	return dcu
}

// ClearNormalizedVariable clears the value of the "normalized_variable" field.
func (dcu *DescriptionChangeUpdate) ClearNormalizedVariable() *DescriptionChangeUpdate {
	dcu.mutation.ClearNormalizedVariable()
	return dcu
}

// SetChangedAt sets the "changed_at" field.
func (dcu *DescriptionChangeUpdate) SetChangedAt(t time.Time) *DescriptionChangeUpdate {
	dcu.mutation.SetChangedAt(t)
	return dcu
}

// SetNillableChangedAt sets the "changed_at" field if the given value is not nil.
func (dcu *DescriptionChangeUpdate) SetNillableChangedAt(t *time.Time) *DescriptionChangeUpdate {
	if t != nil {
		dcu.SetChangedAt(*t)
	}
	return dcu
}

// SetDescriptionID sets the "description" edge to the Description entity by ID.
func (dcu *DescriptionChangeUpdate) SetDescriptionID(id pulid.ID) *DescriptionChangeUpdate {
	dcu.mutation.SetDescriptionID(id)
	return dcu
}

// SetDescription sets the "description" edge to the Description entity.
func (dcu *DescriptionChangeUpdate) SetDescription(d *Description) *DescriptionChangeUpdate {
	return dcu.SetDescriptionID(d.ID)
}

// Mutation returns the DescriptionChangeMutation object of the builder.
func (dcu *DescriptionChangeUpdate) Mutation() *DescriptionChangeMutation {
	return dcu.mutation
}

// ClearDescription clears the "description" edge to the Description entity.
func (dcu *DescriptionChangeUpdate) ClearDescription() *DescriptionChangeUpdate {
	dcu.mutation.ClearDescription()
	return dcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dcu *DescriptionChangeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, dcu.sqlSave, dcu.mutation, dcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dcu *DescriptionChangeUpdate) SaveX(ctx context.Context) int {
	affected, err := dcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dcu *DescriptionChangeUpdate) Exec(ctx context.Context) error {
	_, err := dcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcu *DescriptionChangeUpdate) ExecX(ctx context.Context) {
	if err := dcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dcu *DescriptionChangeUpdate) check() error {
	if _, ok := dcu.mutation.DescriptionID(); dcu.mutation.DescriptionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Description_change.description"`)
	}
	return nil
}

func (dcu *DescriptionChangeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := dcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(description_change.Table, description_change.Columns, sqlgraph.NewFieldSpec(description_change.FieldID, field.TypeString))
	if ps := dcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dcu.mutation.Raw(); ok {
		_spec.SetField(description_change.FieldRaw, field.TypeString, value)
	}
	if value, ok := dcu.mutation.Variable(); ok {
		_spec.SetField(description_change.FieldVariable, field.TypeString, value)
	}
	if dcu.mutation.VariableCleared() {
		_spec.ClearField(description_change.FieldVariable, field.TypeString)
	}
	if value, ok := dcu.mutation.NormalizedVariable(); ok {
		_spec.SetField(description_change.FieldNormalizedVariable, field.TypeString, value)
	}
	if dcu.mutation.NormalizedVariableCleared() {
		_spec.ClearField(description_change.FieldNormalizedVariable, field.TypeString)
	}
	if value, ok := dcu.mutation.ChangedAt(); ok {
		_spec.SetField(description_change.FieldChangedAt, field.TypeTime, value)
	}
	if dcu.mutation.DescriptionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   description_change.DescriptionTable,
			Columns: []string{description_change.DescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(description.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dcu.mutation.DescriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   description_change.DescriptionTable,
			Columns: []string{description_change.DescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(description.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{description_change.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	dcu.mutation.done = true
	return n, nil
}

// DescriptionChangeUpdateOne is the builder for updating a single Description_change entity.
type DescriptionChangeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DescriptionChangeMutation
}

// SetRaw sets the "raw" field.
func (dcuo *DescriptionChangeUpdateOne) SetRaw(s string) *DescriptionChangeUpdateOne {
	dcuo.mutation.SetRaw(s)
	return dcuo
}

// SetNillableRaw sets the "raw" field if the given value is not nil.
func (dcuo *DescriptionChangeUpdateOne) SetNillableRaw(s *string) *DescriptionChangeUpdateOne {
	if s != nil {
		dcuo.SetRaw(*s)
	}
	return dcuo
}

// SetVariable sets the "variable" field.
func (dcuo *DescriptionChangeUpdateOne) SetVariable(s string) *DescriptionChangeUpdateOne {
	dcuo.mutation.SetVariable(s)
	return dcuo
}

// SetNillableVariable sets the "variable" field if the given value is not nil.
func (dcuo *DescriptionChangeUpdateOne) SetNillableVariable(s *string) *DescriptionChangeUpdateOne {
	if s != nil {
		dcuo.SetVariable(*s)
	}
	return dcuo
}

// ClearVariable clears the value of the "variable" field.
func (dcuo *DescriptionChangeUpdateOne) ClearVariable() *DescriptionChangeUpdateOne {
	dcuo.mutation.ClearVariable()
	return dcuo
}

// SetNormalizedVariable sets the "normalized_variable" field.
func (dcuo *DescriptionChangeUpdateOne) SetNormalizedVariable(s string) *DescriptionChangeUpdateOne {
	dcuo.mutation.SetNormalizedVariable(s)
	return dcuo
}

// SetNillableNormalizedVariable sets the "normalized_variable" field if the given value is not nil.
func (dcuo *DescriptionChangeUpdateOne) SetNillableNormalizedVariable(s *string) *DescriptionChangeUpdateOne {
	if s != nil {
		dcuo.SetNormalizedVariable(*s)
	}
	return dcuo
}

// ClearNormalizedVariable clears the value of the "normalized_variable" field.
func (dcuo *DescriptionChangeUpdateOne) ClearNormalizedVariable() *DescriptionChangeUpdateOne {
	dcuo.mutation.ClearNormalizedVariable()
	return dcuo
}

// SetChangedAt sets the "changed_at" field.
func (dcuo *DescriptionChangeUpdateOne) SetChangedAt(t time.Time) *DescriptionChangeUpdateOne {
	dcuo.mutation.SetChangedAt(t)
	return dcuo
}

// SetNillableChangedAt sets the "changed_at" field if the given value is not nil.
func (dcuo *DescriptionChangeUpdateOne) SetNillableChangedAt(t *time.Time) *DescriptionChangeUpdateOne {
	if t != nil {
		dcuo.SetChangedAt(*t)
	}
	return dcuo
}

// SetDescriptionID sets the "description" edge to the Description entity by ID.
func (dcuo *DescriptionChangeUpdateOne) SetDescriptionID(id pulid.ID) *DescriptionChangeUpdateOne {
	dcuo.mutation.SetDescriptionID(id)
	return dcuo
}

// SetDescription sets the "description" edge to the Description entity.
func (dcuo *DescriptionChangeUpdateOne) SetDescription(d *Description) *DescriptionChangeUpdateOne {
	return dcuo.SetDescriptionID(d.ID)
}

// Mutation returns the DescriptionChangeMutation object of the builder.
func (dcuo *DescriptionChangeUpdateOne) Mutation() *DescriptionChangeMutation {
	return dcuo.mutation
}

// ClearDescription clears the "description" edge to the Description entity.
func (dcuo *DescriptionChangeUpdateOne) ClearDescription() *DescriptionChangeUpdateOne {
	dcuo.mutation.ClearDescription()
	return dcuo
}

// Where appends a list predicates to the DescriptionChangeUpdate builder.
func (dcuo *DescriptionChangeUpdateOne) Where(ps ...predicate.Description_change) *DescriptionChangeUpdateOne {
	dcuo.mutation.Where(ps...)
	return dcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dcuo *DescriptionChangeUpdateOne) Select(field string, fields ...string) *DescriptionChangeUpdateOne {
	dcuo.fields = append([]string{field}, fields...)
	return dcuo
}

// Save executes the query and returns the updated Description_change entity.
func (dcuo *DescriptionChangeUpdateOne) Save(ctx context.Context) (*Description_change, error) {
	return withHooks(ctx, dcuo.sqlSave, dcuo.mutation, dcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dcuo *DescriptionChangeUpdateOne) SaveX(ctx context.Context) *Description_change {
	node, err := dcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dcuo *DescriptionChangeUpdateOne) Exec(ctx context.Context) error {
	_, err := dcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcuo *DescriptionChangeUpdateOne) ExecX(ctx context.Context) {
	if err := dcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dcuo *DescriptionChangeUpdateOne) check() error {
	if _, ok := dcuo.mutation.DescriptionID(); dcuo.mutation.DescriptionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Description_change.description"`)
	}
	return nil
}

func (dcuo *DescriptionChangeUpdateOne) sqlSave(ctx context.Context) (_node *Description_change, err error) {
	if err := dcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(description_change.Table, description_change.Columns, sqlgraph.NewFieldSpec(description_change.FieldID, field.TypeString))
	id, ok := dcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Description_change.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, description_change.FieldID)
		for _, f := range fields {
			if !description_change.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != description_change.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dcuo.mutation.Raw(); ok {
		_spec.SetField(description_change.FieldRaw, field.TypeString, value)
	}
	if value, ok := dcuo.mutation.Variable(); ok {
		_spec.SetField(description_change.FieldVariable, field.TypeString, value)
	}
	if dcuo.mutation.VariableCleared() {
		_spec.ClearField(description_change.FieldVariable, field.TypeString)
	}
	if value, ok := dcuo.mutation.NormalizedVariable(); ok {
		_spec.SetField(description_change.FieldNormalizedVariable, field.TypeString, value)
	}
	if dcuo.mutation.NormalizedVariableCleared() {
		_spec.ClearField(description_change.FieldNormalizedVariable, field.TypeString)
	}
	if value, ok := dcuo.mutation.ChangedAt(); ok {
		_spec.SetField(description_change.FieldChangedAt, field.TypeTime, value)
	}
	if dcuo.mutation.DescriptionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   description_change.DescriptionTable,
			Columns: []string{description_change.DescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(description.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dcuo.mutation.DescriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   description_change.DescriptionTable,
			Columns: []string{description_change.DescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(description.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Description_change{config: dcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{description_change.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	dcuo.mutation.done = true
	return _node, nil
}
