// Code generated by ent, DO NOT EDIT.

package descriptionchange

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldContainsFold(FieldID, id))
}

// Raw applies equality check predicate on the "raw" field. It's identical to RawEQ.
func Raw(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldEQ(FieldRaw, v))
}

// Variable applies equality check predicate on the "variable" field. It's identical to VariableEQ.
func Variable(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldEQ(FieldVariable, v))
}

// ChangedAt applies equality check predicate on the "changed_at" field. It's identical to ChangedAtEQ.
func ChangedAt(v time.Time) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldEQ(FieldChangedAt, v))
}

// RawEQ applies the EQ predicate on the "raw" field.
func RawEQ(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldEQ(FieldRaw, v))
}

// RawNEQ applies the NEQ predicate on the "raw" field.
func RawNEQ(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldNEQ(FieldRaw, v))
}

// RawIn applies the In predicate on the "raw" field.
func RawIn(vs ...string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldIn(FieldRaw, vs...))
}

// RawNotIn applies the NotIn predicate on the "raw" field.
func RawNotIn(vs ...string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldNotIn(FieldRaw, vs...))
}

// RawGT applies the GT predicate on the "raw" field.
func RawGT(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldGT(FieldRaw, v))
}

// RawGTE applies the GTE predicate on the "raw" field.
func RawGTE(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldGTE(FieldRaw, v))
}

// RawLT applies the LT predicate on the "raw" field.
func RawLT(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldLT(FieldRaw, v))
}

// RawLTE applies the LTE predicate on the "raw" field.
func RawLTE(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldLTE(FieldRaw, v))
}

// RawContains applies the Contains predicate on the "raw" field.
func RawContains(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldContains(FieldRaw, v))
}

// RawHasPrefix applies the HasPrefix predicate on the "raw" field.
func RawHasPrefix(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldHasPrefix(FieldRaw, v))
}

// RawHasSuffix applies the HasSuffix predicate on the "raw" field.
func RawHasSuffix(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldHasSuffix(FieldRaw, v))
}

// RawEqualFold applies the EqualFold predicate on the "raw" field.
func RawEqualFold(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldEqualFold(FieldRaw, v))
}

// RawContainsFold applies the ContainsFold predicate on the "raw" field.
func RawContainsFold(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldContainsFold(FieldRaw, v))
}

// VariableEQ applies the EQ predicate on the "variable" field.
func VariableEQ(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldEQ(FieldVariable, v))
}

// VariableNEQ applies the NEQ predicate on the "variable" field.
func VariableNEQ(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldNEQ(FieldVariable, v))
}

// VariableIn applies the In predicate on the "variable" field.
func VariableIn(vs ...string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldIn(FieldVariable, vs...))
}

// VariableNotIn applies the NotIn predicate on the "variable" field.
func VariableNotIn(vs ...string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldNotIn(FieldVariable, vs...))
}

// VariableGT applies the GT predicate on the "variable" field.
func VariableGT(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldGT(FieldVariable, v))
}

// VariableGTE applies the GTE predicate on the "variable" field.
func VariableGTE(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldGTE(FieldVariable, v))
}

// VariableLT applies the LT predicate on the "variable" field.
func VariableLT(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldLT(FieldVariable, v))
}

// VariableLTE applies the LTE predicate on the "variable" field.
func VariableLTE(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldLTE(FieldVariable, v))
}

// VariableContains applies the Contains predicate on the "variable" field.
func VariableContains(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldContains(FieldVariable, v))
}

// VariableHasPrefix applies the HasPrefix predicate on the "variable" field.
func VariableHasPrefix(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldHasPrefix(FieldVariable, v))
}

// VariableHasSuffix applies the HasSuffix predicate on the "variable" field.
func VariableHasSuffix(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldHasSuffix(FieldVariable, v))
}

// VariableIsNil applies the IsNil predicate on the "variable" field.
func VariableIsNil() predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldIsNull(FieldVariable))
}

// VariableNotNil applies the NotNil predicate on the "variable" field.
func VariableNotNil() predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldNotNull(FieldVariable))
}

// VariableEqualFold applies the EqualFold predicate on the "variable" field.
func VariableEqualFold(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldEqualFold(FieldVariable, v))
}

// VariableContainsFold applies the ContainsFold predicate on the "variable" field.
func VariableContainsFold(v string) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldContainsFold(FieldVariable, v))
}

// ChangedAtEQ applies the EQ predicate on the "changed_at" field.
func ChangedAtEQ(v time.Time) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldEQ(FieldChangedAt, v))
}

// ChangedAtNEQ applies the NEQ predicate on the "changed_at" field.
func ChangedAtNEQ(v time.Time) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldNEQ(FieldChangedAt, v))
}

// ChangedAtIn applies the In predicate on the "changed_at" field.
func ChangedAtIn(vs ...time.Time) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldIn(FieldChangedAt, vs...))
}

// ChangedAtNotIn applies the NotIn predicate on the "changed_at" field.
func ChangedAtNotIn(vs ...time.Time) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldNotIn(FieldChangedAt, vs...))
}

// ChangedAtGT applies the GT predicate on the "changed_at" field.
func ChangedAtGT(v time.Time) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldGT(FieldChangedAt, v))
}

// ChangedAtGTE applies the GTE predicate on the "changed_at" field.
func ChangedAtGTE(v time.Time) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldGTE(FieldChangedAt, v))
}

// ChangedAtLT applies the LT predicate on the "changed_at" field.
func ChangedAtLT(v time.Time) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldLT(FieldChangedAt, v))
}

// ChangedAtLTE applies the LTE predicate on the "changed_at" field.
func ChangedAtLTE(v time.Time) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.FieldLTE(FieldChangedAt, v))
}

// HasDescription applies the HasEdge predicate on the "description" edge.
func HasDescription() predicate.DescriptionChange {
	return predicate.DescriptionChange(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DescriptionTable, DescriptionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDescriptionWith applies the HasEdge predicate on the "description" edge with a given conditions (other predicates).
func HasDescriptionWith(preds ...predicate.Description) predicate.DescriptionChange {
	return predicate.DescriptionChange(func(s *sql.Selector) {
		step := newDescriptionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DescriptionChange) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DescriptionChange) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DescriptionChange) predicate.DescriptionChange {
	return predicate.DescriptionChange(sql.NotPredicates(p))
}
