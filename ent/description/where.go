// Code generated by ent, DO NOT EDIT.

package description

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Description {
	return predicate.Description(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Description {
	return predicate.Description(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Description {
	return predicate.Description(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Description {
	return predicate.Description(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Description {
	return predicate.Description(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Description {
	return predicate.Description(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Description {
	return predicate.Description(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Description {
	return predicate.Description(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Description {
	return predicate.Description(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Description {
	return predicate.Description(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Description {
	return predicate.Description(sql.FieldContainsFold(FieldID, id))
}

// Raw applies equality check predicate on the "raw" field. It's identical to RawEQ.
func Raw(v string) predicate.Description {
	return predicate.Description(sql.FieldEQ(FieldRaw, v))
}

// Variable applies equality check predicate on the "variable" field. It's identical to VariableEQ.
func Variable(v string) predicate.Description {
	return predicate.Description(sql.FieldEQ(FieldVariable, v))
}

// NormalizedVariable applies equality check predicate on the "normalized_variable" field. It's identical to NormalizedVariableEQ.
func NormalizedVariable(v string) predicate.Description {
	return predicate.Description(sql.FieldEQ(FieldNormalizedVariable, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Description {
	return predicate.Description(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Description {
	return predicate.Description(sql.FieldEQ(FieldUpdatedAt, v))
}

// RawEQ applies the EQ predicate on the "raw" field.
func RawEQ(v string) predicate.Description {
	return predicate.Description(sql.FieldEQ(FieldRaw, v))
}

// RawNEQ applies the NEQ predicate on the "raw" field.
func RawNEQ(v string) predicate.Description {
	return predicate.Description(sql.FieldNEQ(FieldRaw, v))
}

// RawIn applies the In predicate on the "raw" field.
func RawIn(vs ...string) predicate.Description {
	return predicate.Description(sql.FieldIn(FieldRaw, vs...))
}

// RawNotIn applies the NotIn predicate on the "raw" field.
func RawNotIn(vs ...string) predicate.Description {
	return predicate.Description(sql.FieldNotIn(FieldRaw, vs...))
}

// RawGT applies the GT predicate on the "raw" field.
func RawGT(v string) predicate.Description {
	return predicate.Description(sql.FieldGT(FieldRaw, v))
}

// RawGTE applies the GTE predicate on the "raw" field.
func RawGTE(v string) predicate.Description {
	return predicate.Description(sql.FieldGTE(FieldRaw, v))
}

// RawLT applies the LT predicate on the "raw" field.
func RawLT(v string) predicate.Description {
	return predicate.Description(sql.FieldLT(FieldRaw, v))
}

// RawLTE applies the LTE predicate on the "raw" field.
func RawLTE(v string) predicate.Description {
	return predicate.Description(sql.FieldLTE(FieldRaw, v))
}

// RawContains applies the Contains predicate on the "raw" field.
func RawContains(v string) predicate.Description {
	return predicate.Description(sql.FieldContains(FieldRaw, v))
}

// RawHasPrefix applies the HasPrefix predicate on the "raw" field.
func RawHasPrefix(v string) predicate.Description {
	return predicate.Description(sql.FieldHasPrefix(FieldRaw, v))
}

// RawHasSuffix applies the HasSuffix predicate on the "raw" field.
func RawHasSuffix(v string) predicate.Description {
	return predicate.Description(sql.FieldHasSuffix(FieldRaw, v))
}

// RawEqualFold applies the EqualFold predicate on the "raw" field.
func RawEqualFold(v string) predicate.Description {
	return predicate.Description(sql.FieldEqualFold(FieldRaw, v))
}

// RawContainsFold applies the ContainsFold predicate on the "raw" field.
func RawContainsFold(v string) predicate.Description {
	return predicate.Description(sql.FieldContainsFold(FieldRaw, v))
}

// VariableEQ applies the EQ predicate on the "variable" field.
func VariableEQ(v string) predicate.Description {
	return predicate.Description(sql.FieldEQ(FieldVariable, v))
}

// VariableNEQ applies the NEQ predicate on the "variable" field.
func VariableNEQ(v string) predicate.Description {
	return predicate.Description(sql.FieldNEQ(FieldVariable, v))
}

// VariableIn applies the In predicate on the "variable" field.
func VariableIn(vs ...string) predicate.Description {
	return predicate.Description(sql.FieldIn(FieldVariable, vs...))
}

// VariableNotIn applies the NotIn predicate on the "variable" field.
func VariableNotIn(vs ...string) predicate.Description {
	return predicate.Description(sql.FieldNotIn(FieldVariable, vs...))
}

// VariableGT applies the GT predicate on the "variable" field.
func VariableGT(v string) predicate.Description {
	return predicate.Description(sql.FieldGT(FieldVariable, v))
}

// VariableGTE applies the GTE predicate on the "variable" field.
func VariableGTE(v string) predicate.Description {
	return predicate.Description(sql.FieldGTE(FieldVariable, v))
}

// VariableLT applies the LT predicate on the "variable" field.
func VariableLT(v string) predicate.Description {
	return predicate.Description(sql.FieldLT(FieldVariable, v))
}

// VariableLTE applies the LTE predicate on the "variable" field.
func VariableLTE(v string) predicate.Description {
	return predicate.Description(sql.FieldLTE(FieldVariable, v))
}

// VariableContains applies the Contains predicate on the "variable" field.
func VariableContains(v string) predicate.Description {
	return predicate.Description(sql.FieldContains(FieldVariable, v))
}

// VariableHasPrefix applies the HasPrefix predicate on the "variable" field.
func VariableHasPrefix(v string) predicate.Description {
	return predicate.Description(sql.FieldHasPrefix(FieldVariable, v))
}

// VariableHasSuffix applies the HasSuffix predicate on the "variable" field.
func VariableHasSuffix(v string) predicate.Description {
	return predicate.Description(sql.FieldHasSuffix(FieldVariable, v))
}

// VariableIsNil applies the IsNil predicate on the "variable" field.
func VariableIsNil() predicate.Description {
	return predicate.Description(sql.FieldIsNull(FieldVariable))
}

// VariableNotNil applies the NotNil predicate on the "variable" field.
func VariableNotNil() predicate.Description {
	return predicate.Description(sql.FieldNotNull(FieldVariable))
}

// VariableEqualFold applies the EqualFold predicate on the "variable" field.
func VariableEqualFold(v string) predicate.Description {
	return predicate.Description(sql.FieldEqualFold(FieldVariable, v))
}

// VariableContainsFold applies the ContainsFold predicate on the "variable" field.
func VariableContainsFold(v string) predicate.Description {
	return predicate.Description(sql.FieldContainsFold(FieldVariable, v))
}

// NormalizedVariableEQ applies the EQ predicate on the "normalized_variable" field.
func NormalizedVariableEQ(v string) predicate.Description {
	return predicate.Description(sql.FieldEQ(FieldNormalizedVariable, v))
}

// NormalizedVariableNEQ applies the NEQ predicate on the "normalized_variable" field.
func NormalizedVariableNEQ(v string) predicate.Description {
	return predicate.Description(sql.FieldNEQ(FieldNormalizedVariable, v))
}

// NormalizedVariableIn applies the In predicate on the "normalized_variable" field.
func NormalizedVariableIn(vs ...string) predicate.Description {
	return predicate.Description(sql.FieldIn(FieldNormalizedVariable, vs...))
}

// NormalizedVariableNotIn applies the NotIn predicate on the "normalized_variable" field.
func NormalizedVariableNotIn(vs ...string) predicate.Description {
	return predicate.Description(sql.FieldNotIn(FieldNormalizedVariable, vs...))
}

// NormalizedVariableGT applies the GT predicate on the "normalized_variable" field.
func NormalizedVariableGT(v string) predicate.Description {
	return predicate.Description(sql.FieldGT(FieldNormalizedVariable, v))
}

// NormalizedVariableGTE applies the GTE predicate on the "normalized_variable" field.
func NormalizedVariableGTE(v string) predicate.Description {
	return predicate.Description(sql.FieldGTE(FieldNormalizedVariable, v))
}

// NormalizedVariableLT applies the LT predicate on the "normalized_variable" field.
func NormalizedVariableLT(v string) predicate.Description {
	return predicate.Description(sql.FieldLT(FieldNormalizedVariable, v))
}

// NormalizedVariableLTE applies the LTE predicate on the "normalized_variable" field.
func NormalizedVariableLTE(v string) predicate.Description {
	return predicate.Description(sql.FieldLTE(FieldNormalizedVariable, v))
}

// NormalizedVariableContains applies the Contains predicate on the "normalized_variable" field.
func NormalizedVariableContains(v string) predicate.Description {
	return predicate.Description(sql.FieldContains(FieldNormalizedVariable, v))
}

// NormalizedVariableHasPrefix applies the HasPrefix predicate on the "normalized_variable" field.
func NormalizedVariableHasPrefix(v string) predicate.Description {
	return predicate.Description(sql.FieldHasPrefix(FieldNormalizedVariable, v))
}

// NormalizedVariableHasSuffix applies the HasSuffix predicate on the "normalized_variable" field.
func NormalizedVariableHasSuffix(v string) predicate.Description {
	return predicate.Description(sql.FieldHasSuffix(FieldNormalizedVariable, v))
}

// NormalizedVariableIsNil applies the IsNil predicate on the "normalized_variable" field.
func NormalizedVariableIsNil() predicate.Description {
	return predicate.Description(sql.FieldIsNull(FieldNormalizedVariable))
}

// NormalizedVariableNotNil applies the NotNil predicate on the "normalized_variable" field.
func NormalizedVariableNotNil() predicate.Description {
	return predicate.Description(sql.FieldNotNull(FieldNormalizedVariable))
}

// NormalizedVariableEqualFold applies the EqualFold predicate on the "normalized_variable" field.
func NormalizedVariableEqualFold(v string) predicate.Description {
	return predicate.Description(sql.FieldEqualFold(FieldNormalizedVariable, v))
}

// NormalizedVariableContainsFold applies the ContainsFold predicate on the "normalized_variable" field.
func NormalizedVariableContainsFold(v string) predicate.Description {
	return predicate.Description(sql.FieldContainsFold(FieldNormalizedVariable, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Description {
	return predicate.Description(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Description {
	return predicate.Description(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Description {
	return predicate.Description(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Description {
	return predicate.Description(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Description {
	return predicate.Description(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Description {
	return predicate.Description(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Description {
	return predicate.Description(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Description {
	return predicate.Description(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Description {
	return predicate.Description(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Description {
	return predicate.Description(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Description {
	return predicate.Description(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Description {
	return predicate.Description(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Description {
	return predicate.Description(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Description {
	return predicate.Description(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Description {
	return predicate.Description(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Description {
	return predicate.Description(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasVideo applies the HasEdge predicate on the "video" edge.
func HasVideo() predicate.Description {
	return predicate.Description(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, VideoTable, VideoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVideoWith applies the HasEdge predicate on the "video" edge with a given conditions (other predicates).
func HasVideoWith(preds ...predicate.Video) predicate.Description {
	return predicate.Description(func(s *sql.Selector) {
		step := newVideoStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPeriodicDescriptionTemplate applies the HasEdge predicate on the "periodic_description_template" edge.
func HasPeriodicDescriptionTemplate() predicate.Description {
	return predicate.Description(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PeriodicDescriptionTemplateTable, PeriodicDescriptionTemplateColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPeriodicDescriptionTemplateWith applies the HasEdge predicate on the "periodic_description_template" edge with a given conditions (other predicates).
func HasPeriodicDescriptionTemplateWith(preds ...predicate.PeriodicDescriptionTemplate) predicate.Description {
	return predicate.Description(func(s *sql.Selector) {
		step := newPeriodicDescriptionTemplateStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCategoryDescriptionTemplate applies the HasEdge predicate on the "category_description_template" edge.
func HasCategoryDescriptionTemplate() predicate.Description {
	return predicate.Description(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CategoryDescriptionTemplateTable, CategoryDescriptionTemplateColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoryDescriptionTemplateWith applies the HasEdge predicate on the "category_description_template" edge with a given conditions (other predicates).
func HasCategoryDescriptionTemplateWith(preds ...predicate.CategoryDescriptionTemplate) predicate.Description {
	return predicate.Description(func(s *sql.Selector) {
		step := newCategoryDescriptionTemplateStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDescriptionChanges applies the HasEdge predicate on the "description_changes" edge.
func HasDescriptionChanges() predicate.Description {
	return predicate.Description(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DescriptionChangesTable, DescriptionChangesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDescriptionChangesWith applies the HasEdge predicate on the "description_changes" edge with a given conditions (other predicates).
func HasDescriptionChangesWith(preds ...predicate.DescriptionChange) predicate.Description {
	return predicate.Description(func(s *sql.Selector) {
		step := newDescriptionChangesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Description) predicate.Description {
	return predicate.Description(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Description) predicate.Description {
	return predicate.Description(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Description) predicate.Description {
	return predicate.Description(sql.NotPredicates(p))
}
