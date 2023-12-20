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
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/categorydescriptiontemplate"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/description"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/descriptionchange"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/periodicdescriptiontemplate"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/predicate"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/video"
)

// DescriptionUpdate is the builder for updating Description entities.
type DescriptionUpdate struct {
	config
	hooks    []Hook
	mutation *DescriptionMutation
}

// Where appends a list predicates to the DescriptionUpdate builder.
func (du *DescriptionUpdate) Where(ps ...predicate.Description) *DescriptionUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetRaw sets the "raw" field.
func (du *DescriptionUpdate) SetRaw(s string) *DescriptionUpdate {
	du.mutation.SetRaw(s)
	return du
}

// SetNillableRaw sets the "raw" field if the given value is not nil.
func (du *DescriptionUpdate) SetNillableRaw(s *string) *DescriptionUpdate {
	if s != nil {
		du.SetRaw(*s)
	}
	return du
}

// SetVariable sets the "variable" field.
func (du *DescriptionUpdate) SetVariable(s string) *DescriptionUpdate {
	du.mutation.SetVariable(s)
	return du
}

// SetNillableVariable sets the "variable" field if the given value is not nil.
func (du *DescriptionUpdate) SetNillableVariable(s *string) *DescriptionUpdate {
	if s != nil {
		du.SetVariable(*s)
	}
	return du
}

// ClearVariable clears the value of the "variable" field.
func (du *DescriptionUpdate) ClearVariable() *DescriptionUpdate {
	du.mutation.ClearVariable()
	return du
}

// SetCreatedAt sets the "created_at" field.
func (du *DescriptionUpdate) SetCreatedAt(t time.Time) *DescriptionUpdate {
	du.mutation.SetCreatedAt(t)
	return du
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (du *DescriptionUpdate) SetNillableCreatedAt(t *time.Time) *DescriptionUpdate {
	if t != nil {
		du.SetCreatedAt(*t)
	}
	return du
}

// SetUpdatedAt sets the "updated_at" field.
func (du *DescriptionUpdate) SetUpdatedAt(t time.Time) *DescriptionUpdate {
	du.mutation.SetUpdatedAt(t)
	return du
}

// SetVideoID sets the "video" edge to the Video entity by ID.
func (du *DescriptionUpdate) SetVideoID(id string) *DescriptionUpdate {
	du.mutation.SetVideoID(id)
	return du
}

// SetVideo sets the "video" edge to the Video entity.
func (du *DescriptionUpdate) SetVideo(v *Video) *DescriptionUpdate {
	return du.SetVideoID(v.ID)
}

// SetPeriodicTemplateID sets the "periodic_template" edge to the PeriodicDescriptionTemplate entity by ID.
func (du *DescriptionUpdate) SetPeriodicTemplateID(id string) *DescriptionUpdate {
	du.mutation.SetPeriodicTemplateID(id)
	return du
}

// SetNillablePeriodicTemplateID sets the "periodic_template" edge to the PeriodicDescriptionTemplate entity by ID if the given value is not nil.
func (du *DescriptionUpdate) SetNillablePeriodicTemplateID(id *string) *DescriptionUpdate {
	if id != nil {
		du = du.SetPeriodicTemplateID(*id)
	}
	return du
}

// SetPeriodicTemplate sets the "periodic_template" edge to the PeriodicDescriptionTemplate entity.
func (du *DescriptionUpdate) SetPeriodicTemplate(p *PeriodicDescriptionTemplate) *DescriptionUpdate {
	return du.SetPeriodicTemplateID(p.ID)
}

// SetCategoryTemplateID sets the "category_template" edge to the CategoryDescriptionTemplate entity by ID.
func (du *DescriptionUpdate) SetCategoryTemplateID(id string) *DescriptionUpdate {
	du.mutation.SetCategoryTemplateID(id)
	return du
}

// SetNillableCategoryTemplateID sets the "category_template" edge to the CategoryDescriptionTemplate entity by ID if the given value is not nil.
func (du *DescriptionUpdate) SetNillableCategoryTemplateID(id *string) *DescriptionUpdate {
	if id != nil {
		du = du.SetCategoryTemplateID(*id)
	}
	return du
}

// SetCategoryTemplate sets the "category_template" edge to the CategoryDescriptionTemplate entity.
func (du *DescriptionUpdate) SetCategoryTemplate(c *CategoryDescriptionTemplate) *DescriptionUpdate {
	return du.SetCategoryTemplateID(c.ID)
}

// AddDescriptionChangeIDs adds the "description_changes" edge to the DescriptionChange entity by IDs.
func (du *DescriptionUpdate) AddDescriptionChangeIDs(ids ...string) *DescriptionUpdate {
	du.mutation.AddDescriptionChangeIDs(ids...)
	return du
}

// AddDescriptionChanges adds the "description_changes" edges to the DescriptionChange entity.
func (du *DescriptionUpdate) AddDescriptionChanges(d ...*DescriptionChange) *DescriptionUpdate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.AddDescriptionChangeIDs(ids...)
}

// Mutation returns the DescriptionMutation object of the builder.
func (du *DescriptionUpdate) Mutation() *DescriptionMutation {
	return du.mutation
}

// ClearVideo clears the "video" edge to the Video entity.
func (du *DescriptionUpdate) ClearVideo() *DescriptionUpdate {
	du.mutation.ClearVideo()
	return du
}

// ClearPeriodicTemplate clears the "periodic_template" edge to the PeriodicDescriptionTemplate entity.
func (du *DescriptionUpdate) ClearPeriodicTemplate() *DescriptionUpdate {
	du.mutation.ClearPeriodicTemplate()
	return du
}

// ClearCategoryTemplate clears the "category_template" edge to the CategoryDescriptionTemplate entity.
func (du *DescriptionUpdate) ClearCategoryTemplate() *DescriptionUpdate {
	du.mutation.ClearCategoryTemplate()
	return du
}

// ClearDescriptionChanges clears all "description_changes" edges to the DescriptionChange entity.
func (du *DescriptionUpdate) ClearDescriptionChanges() *DescriptionUpdate {
	du.mutation.ClearDescriptionChanges()
	return du
}

// RemoveDescriptionChangeIDs removes the "description_changes" edge to DescriptionChange entities by IDs.
func (du *DescriptionUpdate) RemoveDescriptionChangeIDs(ids ...string) *DescriptionUpdate {
	du.mutation.RemoveDescriptionChangeIDs(ids...)
	return du
}

// RemoveDescriptionChanges removes "description_changes" edges to DescriptionChange entities.
func (du *DescriptionUpdate) RemoveDescriptionChanges(d ...*DescriptionChange) *DescriptionUpdate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.RemoveDescriptionChangeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DescriptionUpdate) Save(ctx context.Context) (int, error) {
	du.defaults()
	return withHooks(ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DescriptionUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DescriptionUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DescriptionUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DescriptionUpdate) defaults() {
	if _, ok := du.mutation.UpdatedAt(); !ok {
		v := description.UpdateDefaultUpdatedAt()
		du.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DescriptionUpdate) check() error {
	if v, ok := du.mutation.Raw(); ok {
		if err := description.RawValidator(v); err != nil {
			return &ValidationError{Name: "raw", err: fmt.Errorf(`ent: validator failed for field "Description.raw": %w`, err)}
		}
	}
	if _, ok := du.mutation.VideoID(); du.mutation.VideoCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Description.video"`)
	}
	return nil
}

func (du *DescriptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := du.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(description.Table, description.Columns, sqlgraph.NewFieldSpec(description.FieldID, field.TypeString))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Raw(); ok {
		_spec.SetField(description.FieldRaw, field.TypeString, value)
	}
	if value, ok := du.mutation.Variable(); ok {
		_spec.SetField(description.FieldVariable, field.TypeString, value)
	}
	if du.mutation.VariableCleared() {
		_spec.ClearField(description.FieldVariable, field.TypeString)
	}
	if value, ok := du.mutation.CreatedAt(); ok {
		_spec.SetField(description.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := du.mutation.UpdatedAt(); ok {
		_spec.SetField(description.FieldUpdatedAt, field.TypeTime, value)
	}
	if du.mutation.VideoCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.VideoIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.PeriodicTemplateCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.PeriodicTemplateIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.CategoryTemplateCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.CategoryTemplateIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.DescriptionChangesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedDescriptionChangesIDs(); len(nodes) > 0 && !du.mutation.DescriptionChangesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.DescriptionChangesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{description.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DescriptionUpdateOne is the builder for updating a single Description entity.
type DescriptionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DescriptionMutation
}

// SetRaw sets the "raw" field.
func (duo *DescriptionUpdateOne) SetRaw(s string) *DescriptionUpdateOne {
	duo.mutation.SetRaw(s)
	return duo
}

// SetNillableRaw sets the "raw" field if the given value is not nil.
func (duo *DescriptionUpdateOne) SetNillableRaw(s *string) *DescriptionUpdateOne {
	if s != nil {
		duo.SetRaw(*s)
	}
	return duo
}

// SetVariable sets the "variable" field.
func (duo *DescriptionUpdateOne) SetVariable(s string) *DescriptionUpdateOne {
	duo.mutation.SetVariable(s)
	return duo
}

// SetNillableVariable sets the "variable" field if the given value is not nil.
func (duo *DescriptionUpdateOne) SetNillableVariable(s *string) *DescriptionUpdateOne {
	if s != nil {
		duo.SetVariable(*s)
	}
	return duo
}

// ClearVariable clears the value of the "variable" field.
func (duo *DescriptionUpdateOne) ClearVariable() *DescriptionUpdateOne {
	duo.mutation.ClearVariable()
	return duo
}

// SetCreatedAt sets the "created_at" field.
func (duo *DescriptionUpdateOne) SetCreatedAt(t time.Time) *DescriptionUpdateOne {
	duo.mutation.SetCreatedAt(t)
	return duo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (duo *DescriptionUpdateOne) SetNillableCreatedAt(t *time.Time) *DescriptionUpdateOne {
	if t != nil {
		duo.SetCreatedAt(*t)
	}
	return duo
}

// SetUpdatedAt sets the "updated_at" field.
func (duo *DescriptionUpdateOne) SetUpdatedAt(t time.Time) *DescriptionUpdateOne {
	duo.mutation.SetUpdatedAt(t)
	return duo
}

// SetVideoID sets the "video" edge to the Video entity by ID.
func (duo *DescriptionUpdateOne) SetVideoID(id string) *DescriptionUpdateOne {
	duo.mutation.SetVideoID(id)
	return duo
}

// SetVideo sets the "video" edge to the Video entity.
func (duo *DescriptionUpdateOne) SetVideo(v *Video) *DescriptionUpdateOne {
	return duo.SetVideoID(v.ID)
}

// SetPeriodicTemplateID sets the "periodic_template" edge to the PeriodicDescriptionTemplate entity by ID.
func (duo *DescriptionUpdateOne) SetPeriodicTemplateID(id string) *DescriptionUpdateOne {
	duo.mutation.SetPeriodicTemplateID(id)
	return duo
}

// SetNillablePeriodicTemplateID sets the "periodic_template" edge to the PeriodicDescriptionTemplate entity by ID if the given value is not nil.
func (duo *DescriptionUpdateOne) SetNillablePeriodicTemplateID(id *string) *DescriptionUpdateOne {
	if id != nil {
		duo = duo.SetPeriodicTemplateID(*id)
	}
	return duo
}

// SetPeriodicTemplate sets the "periodic_template" edge to the PeriodicDescriptionTemplate entity.
func (duo *DescriptionUpdateOne) SetPeriodicTemplate(p *PeriodicDescriptionTemplate) *DescriptionUpdateOne {
	return duo.SetPeriodicTemplateID(p.ID)
}

// SetCategoryTemplateID sets the "category_template" edge to the CategoryDescriptionTemplate entity by ID.
func (duo *DescriptionUpdateOne) SetCategoryTemplateID(id string) *DescriptionUpdateOne {
	duo.mutation.SetCategoryTemplateID(id)
	return duo
}

// SetNillableCategoryTemplateID sets the "category_template" edge to the CategoryDescriptionTemplate entity by ID if the given value is not nil.
func (duo *DescriptionUpdateOne) SetNillableCategoryTemplateID(id *string) *DescriptionUpdateOne {
	if id != nil {
		duo = duo.SetCategoryTemplateID(*id)
	}
	return duo
}

// SetCategoryTemplate sets the "category_template" edge to the CategoryDescriptionTemplate entity.
func (duo *DescriptionUpdateOne) SetCategoryTemplate(c *CategoryDescriptionTemplate) *DescriptionUpdateOne {
	return duo.SetCategoryTemplateID(c.ID)
}

// AddDescriptionChangeIDs adds the "description_changes" edge to the DescriptionChange entity by IDs.
func (duo *DescriptionUpdateOne) AddDescriptionChangeIDs(ids ...string) *DescriptionUpdateOne {
	duo.mutation.AddDescriptionChangeIDs(ids...)
	return duo
}

// AddDescriptionChanges adds the "description_changes" edges to the DescriptionChange entity.
func (duo *DescriptionUpdateOne) AddDescriptionChanges(d ...*DescriptionChange) *DescriptionUpdateOne {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.AddDescriptionChangeIDs(ids...)
}

// Mutation returns the DescriptionMutation object of the builder.
func (duo *DescriptionUpdateOne) Mutation() *DescriptionMutation {
	return duo.mutation
}

// ClearVideo clears the "video" edge to the Video entity.
func (duo *DescriptionUpdateOne) ClearVideo() *DescriptionUpdateOne {
	duo.mutation.ClearVideo()
	return duo
}

// ClearPeriodicTemplate clears the "periodic_template" edge to the PeriodicDescriptionTemplate entity.
func (duo *DescriptionUpdateOne) ClearPeriodicTemplate() *DescriptionUpdateOne {
	duo.mutation.ClearPeriodicTemplate()
	return duo
}

// ClearCategoryTemplate clears the "category_template" edge to the CategoryDescriptionTemplate entity.
func (duo *DescriptionUpdateOne) ClearCategoryTemplate() *DescriptionUpdateOne {
	duo.mutation.ClearCategoryTemplate()
	return duo
}

// ClearDescriptionChanges clears all "description_changes" edges to the DescriptionChange entity.
func (duo *DescriptionUpdateOne) ClearDescriptionChanges() *DescriptionUpdateOne {
	duo.mutation.ClearDescriptionChanges()
	return duo
}

// RemoveDescriptionChangeIDs removes the "description_changes" edge to DescriptionChange entities by IDs.
func (duo *DescriptionUpdateOne) RemoveDescriptionChangeIDs(ids ...string) *DescriptionUpdateOne {
	duo.mutation.RemoveDescriptionChangeIDs(ids...)
	return duo
}

// RemoveDescriptionChanges removes "description_changes" edges to DescriptionChange entities.
func (duo *DescriptionUpdateOne) RemoveDescriptionChanges(d ...*DescriptionChange) *DescriptionUpdateOne {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.RemoveDescriptionChangeIDs(ids...)
}

// Where appends a list predicates to the DescriptionUpdate builder.
func (duo *DescriptionUpdateOne) Where(ps ...predicate.Description) *DescriptionUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DescriptionUpdateOne) Select(field string, fields ...string) *DescriptionUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Description entity.
func (duo *DescriptionUpdateOne) Save(ctx context.Context) (*Description, error) {
	duo.defaults()
	return withHooks(ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DescriptionUpdateOne) SaveX(ctx context.Context) *Description {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DescriptionUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DescriptionUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DescriptionUpdateOne) defaults() {
	if _, ok := duo.mutation.UpdatedAt(); !ok {
		v := description.UpdateDefaultUpdatedAt()
		duo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DescriptionUpdateOne) check() error {
	if v, ok := duo.mutation.Raw(); ok {
		if err := description.RawValidator(v); err != nil {
			return &ValidationError{Name: "raw", err: fmt.Errorf(`ent: validator failed for field "Description.raw": %w`, err)}
		}
	}
	if _, ok := duo.mutation.VideoID(); duo.mutation.VideoCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Description.video"`)
	}
	return nil
}

func (duo *DescriptionUpdateOne) sqlSave(ctx context.Context) (_node *Description, err error) {
	if err := duo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(description.Table, description.Columns, sqlgraph.NewFieldSpec(description.FieldID, field.TypeString))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Description.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, description.FieldID)
		for _, f := range fields {
			if !description.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != description.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.Raw(); ok {
		_spec.SetField(description.FieldRaw, field.TypeString, value)
	}
	if value, ok := duo.mutation.Variable(); ok {
		_spec.SetField(description.FieldVariable, field.TypeString, value)
	}
	if duo.mutation.VariableCleared() {
		_spec.ClearField(description.FieldVariable, field.TypeString)
	}
	if value, ok := duo.mutation.CreatedAt(); ok {
		_spec.SetField(description.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := duo.mutation.UpdatedAt(); ok {
		_spec.SetField(description.FieldUpdatedAt, field.TypeTime, value)
	}
	if duo.mutation.VideoCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.VideoIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.PeriodicTemplateCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.PeriodicTemplateIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.CategoryTemplateCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.CategoryTemplateIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.DescriptionChangesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedDescriptionChangesIDs(); len(nodes) > 0 && !duo.mutation.DescriptionChangesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.DescriptionChangesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Description{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{description.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}
