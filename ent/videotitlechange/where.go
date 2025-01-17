// Code generated by ent, DO NOT EDIT.

package videotitlechange

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldContainsFold(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldEQ(FieldTitle, v))
}

// ChangedAt applies equality check predicate on the "changed_at" field. It's identical to ChangedAtEQ.
func ChangedAt(v time.Time) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldEQ(FieldChangedAt, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldContainsFold(FieldTitle, v))
}

// ChangedAtEQ applies the EQ predicate on the "changed_at" field.
func ChangedAtEQ(v time.Time) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldEQ(FieldChangedAt, v))
}

// ChangedAtNEQ applies the NEQ predicate on the "changed_at" field.
func ChangedAtNEQ(v time.Time) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldNEQ(FieldChangedAt, v))
}

// ChangedAtIn applies the In predicate on the "changed_at" field.
func ChangedAtIn(vs ...time.Time) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldIn(FieldChangedAt, vs...))
}

// ChangedAtNotIn applies the NotIn predicate on the "changed_at" field.
func ChangedAtNotIn(vs ...time.Time) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldNotIn(FieldChangedAt, vs...))
}

// ChangedAtGT applies the GT predicate on the "changed_at" field.
func ChangedAtGT(v time.Time) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldGT(FieldChangedAt, v))
}

// ChangedAtGTE applies the GTE predicate on the "changed_at" field.
func ChangedAtGTE(v time.Time) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldGTE(FieldChangedAt, v))
}

// ChangedAtLT applies the LT predicate on the "changed_at" field.
func ChangedAtLT(v time.Time) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldLT(FieldChangedAt, v))
}

// ChangedAtLTE applies the LTE predicate on the "changed_at" field.
func ChangedAtLTE(v time.Time) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.FieldLTE(FieldChangedAt, v))
}

// HasVideo applies the HasEdge predicate on the "video" edge.
func HasVideo() predicate.VideoTitleChange {
	return predicate.VideoTitleChange(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VideoTable, VideoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVideoWith applies the HasEdge predicate on the "video" edge with a given conditions (other predicates).
func HasVideoWith(preds ...predicate.Video) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(func(s *sql.Selector) {
		step := newVideoStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VideoTitleChange) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VideoTitleChange) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.VideoTitleChange) predicate.VideoTitleChange {
	return predicate.VideoTitleChange(sql.NotPredicates(p))
}
