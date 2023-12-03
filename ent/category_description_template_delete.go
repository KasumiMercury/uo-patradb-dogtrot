// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/category_description_template"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/predicate"
)

// CategoryDescriptionTemplateDelete is the builder for deleting a Category_description_template entity.
type CategoryDescriptionTemplateDelete struct {
	config
	hooks    []Hook
	mutation *CategoryDescriptionTemplateMutation
}

// Where appends a list predicates to the CategoryDescriptionTemplateDelete builder.
func (cdtd *CategoryDescriptionTemplateDelete) Where(ps ...predicate.Category_description_template) *CategoryDescriptionTemplateDelete {
	cdtd.mutation.Where(ps...)
	return cdtd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cdtd *CategoryDescriptionTemplateDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cdtd.sqlExec, cdtd.mutation, cdtd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cdtd *CategoryDescriptionTemplateDelete) ExecX(ctx context.Context) int {
	n, err := cdtd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cdtd *CategoryDescriptionTemplateDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(category_description_template.Table, sqlgraph.NewFieldSpec(category_description_template.FieldID, field.TypeString))
	if ps := cdtd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cdtd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cdtd.mutation.done = true
	return affected, err
}

// CategoryDescriptionTemplateDeleteOne is the builder for deleting a single Category_description_template entity.
type CategoryDescriptionTemplateDeleteOne struct {
	cdtd *CategoryDescriptionTemplateDelete
}

// Where appends a list predicates to the CategoryDescriptionTemplateDelete builder.
func (cdtdo *CategoryDescriptionTemplateDeleteOne) Where(ps ...predicate.Category_description_template) *CategoryDescriptionTemplateDeleteOne {
	cdtdo.cdtd.mutation.Where(ps...)
	return cdtdo
}

// Exec executes the deletion query.
func (cdtdo *CategoryDescriptionTemplateDeleteOne) Exec(ctx context.Context) error {
	n, err := cdtdo.cdtd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{category_description_template.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdtdo *CategoryDescriptionTemplateDeleteOne) ExecX(ctx context.Context) {
	if err := cdtdo.Exec(ctx); err != nil {
		panic(err)
	}
}