// Code generated by ent, DO NOT EDIT.

package patchat

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.PatChat {
	return predicate.PatChat(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.PatChat {
	return predicate.PatChat(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.PatChat {
	return predicate.PatChat(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.PatChat {
	return predicate.PatChat(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.PatChat {
	return predicate.PatChat(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.PatChat {
	return predicate.PatChat(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.PatChat {
	return predicate.PatChat(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.PatChat {
	return predicate.PatChat(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.PatChat {
	return predicate.PatChat(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.PatChat {
	return predicate.PatChat(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.PatChat {
	return predicate.PatChat(sql.FieldContainsFold(FieldID, id))
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v string) predicate.PatChat {
	return predicate.PatChat(sql.FieldEQ(FieldMessage, v))
}

// Magnitude applies equality check predicate on the "magnitude" field. It's identical to MagnitudeEQ.
func Magnitude(v float64) predicate.PatChat {
	return predicate.PatChat(sql.FieldEQ(FieldMagnitude, v))
}

// Score applies equality check predicate on the "score" field. It's identical to ScoreEQ.
func Score(v float64) predicate.PatChat {
	return predicate.PatChat(sql.FieldEQ(FieldScore, v))
}

// IsNegative applies equality check predicate on the "is_negative" field. It's identical to IsNegativeEQ.
func IsNegative(v bool) predicate.PatChat {
	return predicate.PatChat(sql.FieldEQ(FieldIsNegative, v))
}

// PublishedAt applies equality check predicate on the "published_at" field. It's identical to PublishedAtEQ.
func PublishedAt(v time.Time) predicate.PatChat {
	return predicate.PatChat(sql.FieldEQ(FieldPublishedAt, v))
}

// FromFreechat applies equality check predicate on the "from_freechat" field. It's identical to FromFreechatEQ.
func FromFreechat(v bool) predicate.PatChat {
	return predicate.PatChat(sql.FieldEQ(FieldFromFreechat, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.PatChat {
	return predicate.PatChat(sql.FieldEQ(FieldCreatedAt, v))
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v string) predicate.PatChat {
	return predicate.PatChat(sql.FieldEQ(FieldMessage, v))
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v string) predicate.PatChat {
	return predicate.PatChat(sql.FieldNEQ(FieldMessage, v))
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...string) predicate.PatChat {
	return predicate.PatChat(sql.FieldIn(FieldMessage, vs...))
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...string) predicate.PatChat {
	return predicate.PatChat(sql.FieldNotIn(FieldMessage, vs...))
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v string) predicate.PatChat {
	return predicate.PatChat(sql.FieldGT(FieldMessage, v))
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v string) predicate.PatChat {
	return predicate.PatChat(sql.FieldGTE(FieldMessage, v))
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v string) predicate.PatChat {
	return predicate.PatChat(sql.FieldLT(FieldMessage, v))
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v string) predicate.PatChat {
	return predicate.PatChat(sql.FieldLTE(FieldMessage, v))
}

// MessageContains applies the Contains predicate on the "message" field.
func MessageContains(v string) predicate.PatChat {
	return predicate.PatChat(sql.FieldContains(FieldMessage, v))
}

// MessageHasPrefix applies the HasPrefix predicate on the "message" field.
func MessageHasPrefix(v string) predicate.PatChat {
	return predicate.PatChat(sql.FieldHasPrefix(FieldMessage, v))
}

// MessageHasSuffix applies the HasSuffix predicate on the "message" field.
func MessageHasSuffix(v string) predicate.PatChat {
	return predicate.PatChat(sql.FieldHasSuffix(FieldMessage, v))
}

// MessageEqualFold applies the EqualFold predicate on the "message" field.
func MessageEqualFold(v string) predicate.PatChat {
	return predicate.PatChat(sql.FieldEqualFold(FieldMessage, v))
}

// MessageContainsFold applies the ContainsFold predicate on the "message" field.
func MessageContainsFold(v string) predicate.PatChat {
	return predicate.PatChat(sql.FieldContainsFold(FieldMessage, v))
}

// MagnitudeEQ applies the EQ predicate on the "magnitude" field.
func MagnitudeEQ(v float64) predicate.PatChat {
	return predicate.PatChat(sql.FieldEQ(FieldMagnitude, v))
}

// MagnitudeNEQ applies the NEQ predicate on the "magnitude" field.
func MagnitudeNEQ(v float64) predicate.PatChat {
	return predicate.PatChat(sql.FieldNEQ(FieldMagnitude, v))
}

// MagnitudeIn applies the In predicate on the "magnitude" field.
func MagnitudeIn(vs ...float64) predicate.PatChat {
	return predicate.PatChat(sql.FieldIn(FieldMagnitude, vs...))
}

// MagnitudeNotIn applies the NotIn predicate on the "magnitude" field.
func MagnitudeNotIn(vs ...float64) predicate.PatChat {
	return predicate.PatChat(sql.FieldNotIn(FieldMagnitude, vs...))
}

// MagnitudeGT applies the GT predicate on the "magnitude" field.
func MagnitudeGT(v float64) predicate.PatChat {
	return predicate.PatChat(sql.FieldGT(FieldMagnitude, v))
}

// MagnitudeGTE applies the GTE predicate on the "magnitude" field.
func MagnitudeGTE(v float64) predicate.PatChat {
	return predicate.PatChat(sql.FieldGTE(FieldMagnitude, v))
}

// MagnitudeLT applies the LT predicate on the "magnitude" field.
func MagnitudeLT(v float64) predicate.PatChat {
	return predicate.PatChat(sql.FieldLT(FieldMagnitude, v))
}

// MagnitudeLTE applies the LTE predicate on the "magnitude" field.
func MagnitudeLTE(v float64) predicate.PatChat {
	return predicate.PatChat(sql.FieldLTE(FieldMagnitude, v))
}

// ScoreEQ applies the EQ predicate on the "score" field.
func ScoreEQ(v float64) predicate.PatChat {
	return predicate.PatChat(sql.FieldEQ(FieldScore, v))
}

// ScoreNEQ applies the NEQ predicate on the "score" field.
func ScoreNEQ(v float64) predicate.PatChat {
	return predicate.PatChat(sql.FieldNEQ(FieldScore, v))
}

// ScoreIn applies the In predicate on the "score" field.
func ScoreIn(vs ...float64) predicate.PatChat {
	return predicate.PatChat(sql.FieldIn(FieldScore, vs...))
}

// ScoreNotIn applies the NotIn predicate on the "score" field.
func ScoreNotIn(vs ...float64) predicate.PatChat {
	return predicate.PatChat(sql.FieldNotIn(FieldScore, vs...))
}

// ScoreGT applies the GT predicate on the "score" field.
func ScoreGT(v float64) predicate.PatChat {
	return predicate.PatChat(sql.FieldGT(FieldScore, v))
}

// ScoreGTE applies the GTE predicate on the "score" field.
func ScoreGTE(v float64) predicate.PatChat {
	return predicate.PatChat(sql.FieldGTE(FieldScore, v))
}

// ScoreLT applies the LT predicate on the "score" field.
func ScoreLT(v float64) predicate.PatChat {
	return predicate.PatChat(sql.FieldLT(FieldScore, v))
}

// ScoreLTE applies the LTE predicate on the "score" field.
func ScoreLTE(v float64) predicate.PatChat {
	return predicate.PatChat(sql.FieldLTE(FieldScore, v))
}

// IsNegativeEQ applies the EQ predicate on the "is_negative" field.
func IsNegativeEQ(v bool) predicate.PatChat {
	return predicate.PatChat(sql.FieldEQ(FieldIsNegative, v))
}

// IsNegativeNEQ applies the NEQ predicate on the "is_negative" field.
func IsNegativeNEQ(v bool) predicate.PatChat {
	return predicate.PatChat(sql.FieldNEQ(FieldIsNegative, v))
}

// PublishedAtEQ applies the EQ predicate on the "published_at" field.
func PublishedAtEQ(v time.Time) predicate.PatChat {
	return predicate.PatChat(sql.FieldEQ(FieldPublishedAt, v))
}

// PublishedAtNEQ applies the NEQ predicate on the "published_at" field.
func PublishedAtNEQ(v time.Time) predicate.PatChat {
	return predicate.PatChat(sql.FieldNEQ(FieldPublishedAt, v))
}

// PublishedAtIn applies the In predicate on the "published_at" field.
func PublishedAtIn(vs ...time.Time) predicate.PatChat {
	return predicate.PatChat(sql.FieldIn(FieldPublishedAt, vs...))
}

// PublishedAtNotIn applies the NotIn predicate on the "published_at" field.
func PublishedAtNotIn(vs ...time.Time) predicate.PatChat {
	return predicate.PatChat(sql.FieldNotIn(FieldPublishedAt, vs...))
}

// PublishedAtGT applies the GT predicate on the "published_at" field.
func PublishedAtGT(v time.Time) predicate.PatChat {
	return predicate.PatChat(sql.FieldGT(FieldPublishedAt, v))
}

// PublishedAtGTE applies the GTE predicate on the "published_at" field.
func PublishedAtGTE(v time.Time) predicate.PatChat {
	return predicate.PatChat(sql.FieldGTE(FieldPublishedAt, v))
}

// PublishedAtLT applies the LT predicate on the "published_at" field.
func PublishedAtLT(v time.Time) predicate.PatChat {
	return predicate.PatChat(sql.FieldLT(FieldPublishedAt, v))
}

// PublishedAtLTE applies the LTE predicate on the "published_at" field.
func PublishedAtLTE(v time.Time) predicate.PatChat {
	return predicate.PatChat(sql.FieldLTE(FieldPublishedAt, v))
}

// FromFreechatEQ applies the EQ predicate on the "from_freechat" field.
func FromFreechatEQ(v bool) predicate.PatChat {
	return predicate.PatChat(sql.FieldEQ(FieldFromFreechat, v))
}

// FromFreechatNEQ applies the NEQ predicate on the "from_freechat" field.
func FromFreechatNEQ(v bool) predicate.PatChat {
	return predicate.PatChat(sql.FieldNEQ(FieldFromFreechat, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.PatChat {
	return predicate.PatChat(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.PatChat {
	return predicate.PatChat(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.PatChat {
	return predicate.PatChat(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.PatChat {
	return predicate.PatChat(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.PatChat {
	return predicate.PatChat(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.PatChat {
	return predicate.PatChat(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.PatChat {
	return predicate.PatChat(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.PatChat {
	return predicate.PatChat(sql.FieldLTE(FieldCreatedAt, v))
}

// HasVideo applies the HasEdge predicate on the "video" edge.
func HasVideo() predicate.PatChat {
	return predicate.PatChat(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VideoTable, VideoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVideoWith applies the HasEdge predicate on the "video" edge with a given conditions (other predicates).
func HasVideoWith(preds ...predicate.Video) predicate.PatChat {
	return predicate.PatChat(func(s *sql.Selector) {
		step := newVideoStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PatChat) predicate.PatChat {
	return predicate.PatChat(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PatChat) predicate.PatChat {
	return predicate.PatChat(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PatChat) predicate.PatChat {
	return predicate.PatChat(sql.NotPredicates(p))
}
