// Code generated by ent, DO NOT EDIT.

package periodicdescriptiontemplate

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldContainsFold(FieldID, id))
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldEQ(FieldText, v))
}

// StartUseAt applies equality check predicate on the "start_use_at" field. It's identical to StartUseAtEQ.
func StartUseAt(v time.Time) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldEQ(FieldStartUseAt, v))
}

// LastUseAt applies equality check predicate on the "last_use_at" field. It's identical to LastUseAtEQ.
func LastUseAt(v time.Time) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldEQ(FieldLastUseAt, v))
}

// Hash applies equality check predicate on the "hash" field. It's identical to HashEQ.
func Hash(v uint64) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldEQ(FieldHash, v))
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldEQ(FieldText, v))
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldNEQ(FieldText, v))
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldIn(FieldText, vs...))
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldNotIn(FieldText, vs...))
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldGT(FieldText, v))
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldGTE(FieldText, v))
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldLT(FieldText, v))
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldLTE(FieldText, v))
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldContains(FieldText, v))
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldHasPrefix(FieldText, v))
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldHasSuffix(FieldText, v))
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldEqualFold(FieldText, v))
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldContainsFold(FieldText, v))
}

// StartUseAtEQ applies the EQ predicate on the "start_use_at" field.
func StartUseAtEQ(v time.Time) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldEQ(FieldStartUseAt, v))
}

// StartUseAtNEQ applies the NEQ predicate on the "start_use_at" field.
func StartUseAtNEQ(v time.Time) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldNEQ(FieldStartUseAt, v))
}

// StartUseAtIn applies the In predicate on the "start_use_at" field.
func StartUseAtIn(vs ...time.Time) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldIn(FieldStartUseAt, vs...))
}

// StartUseAtNotIn applies the NotIn predicate on the "start_use_at" field.
func StartUseAtNotIn(vs ...time.Time) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldNotIn(FieldStartUseAt, vs...))
}

// StartUseAtGT applies the GT predicate on the "start_use_at" field.
func StartUseAtGT(v time.Time) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldGT(FieldStartUseAt, v))
}

// StartUseAtGTE applies the GTE predicate on the "start_use_at" field.
func StartUseAtGTE(v time.Time) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldGTE(FieldStartUseAt, v))
}

// StartUseAtLT applies the LT predicate on the "start_use_at" field.
func StartUseAtLT(v time.Time) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldLT(FieldStartUseAt, v))
}

// StartUseAtLTE applies the LTE predicate on the "start_use_at" field.
func StartUseAtLTE(v time.Time) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldLTE(FieldStartUseAt, v))
}

// StartUseAtIsNil applies the IsNil predicate on the "start_use_at" field.
func StartUseAtIsNil() predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldIsNull(FieldStartUseAt))
}

// StartUseAtNotNil applies the NotNil predicate on the "start_use_at" field.
func StartUseAtNotNil() predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldNotNull(FieldStartUseAt))
}

// LastUseAtEQ applies the EQ predicate on the "last_use_at" field.
func LastUseAtEQ(v time.Time) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldEQ(FieldLastUseAt, v))
}

// LastUseAtNEQ applies the NEQ predicate on the "last_use_at" field.
func LastUseAtNEQ(v time.Time) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldNEQ(FieldLastUseAt, v))
}

// LastUseAtIn applies the In predicate on the "last_use_at" field.
func LastUseAtIn(vs ...time.Time) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldIn(FieldLastUseAt, vs...))
}

// LastUseAtNotIn applies the NotIn predicate on the "last_use_at" field.
func LastUseAtNotIn(vs ...time.Time) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldNotIn(FieldLastUseAt, vs...))
}

// LastUseAtGT applies the GT predicate on the "last_use_at" field.
func LastUseAtGT(v time.Time) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldGT(FieldLastUseAt, v))
}

// LastUseAtGTE applies the GTE predicate on the "last_use_at" field.
func LastUseAtGTE(v time.Time) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldGTE(FieldLastUseAt, v))
}

// LastUseAtLT applies the LT predicate on the "last_use_at" field.
func LastUseAtLT(v time.Time) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldLT(FieldLastUseAt, v))
}

// LastUseAtLTE applies the LTE predicate on the "last_use_at" field.
func LastUseAtLTE(v time.Time) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldLTE(FieldLastUseAt, v))
}

// LastUseAtIsNil applies the IsNil predicate on the "last_use_at" field.
func LastUseAtIsNil() predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldIsNull(FieldLastUseAt))
}

// LastUseAtNotNil applies the NotNil predicate on the "last_use_at" field.
func LastUseAtNotNil() predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldNotNull(FieldLastUseAt))
}

// HashEQ applies the EQ predicate on the "hash" field.
func HashEQ(v uint64) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldEQ(FieldHash, v))
}

// HashNEQ applies the NEQ predicate on the "hash" field.
func HashNEQ(v uint64) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldNEQ(FieldHash, v))
}

// HashIn applies the In predicate on the "hash" field.
func HashIn(vs ...uint64) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldIn(FieldHash, vs...))
}

// HashNotIn applies the NotIn predicate on the "hash" field.
func HashNotIn(vs ...uint64) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldNotIn(FieldHash, vs...))
}

// HashGT applies the GT predicate on the "hash" field.
func HashGT(v uint64) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldGT(FieldHash, v))
}

// HashGTE applies the GTE predicate on the "hash" field.
func HashGTE(v uint64) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldGTE(FieldHash, v))
}

// HashLT applies the LT predicate on the "hash" field.
func HashLT(v uint64) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldLT(FieldHash, v))
}

// HashLTE applies the LTE predicate on the "hash" field.
func HashLTE(v uint64) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.FieldLTE(FieldHash, v))
}

// HasDescriptions applies the HasEdge predicate on the "descriptions" edge.
func HasDescriptions() predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, DescriptionsTable, DescriptionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDescriptionsWith applies the HasEdge predicate on the "descriptions" edge with a given conditions (other predicates).
func HasDescriptionsWith(preds ...predicate.Description) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(func(s *sql.Selector) {
		step := newDescriptionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PeriodicDescriptionTemplate) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PeriodicDescriptionTemplate) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PeriodicDescriptionTemplate) predicate.PeriodicDescriptionTemplate {
	return predicate.PeriodicDescriptionTemplate(sql.NotPredicates(p))
}
