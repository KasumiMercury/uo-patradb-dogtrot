// Code generated by ent, DO NOT EDIT.

package videotag

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldContainsFold(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldEQ(FieldTitle, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.VideoTag {
	return predicate.VideoTag(sql.FieldContainsFold(FieldTitle, v))
}

// HasVideos applies the HasEdge predicate on the "videos" edge.
func HasVideos() predicate.VideoTag {
	return predicate.VideoTag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, VideosTable, VideosPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVideosWith applies the HasEdge predicate on the "videos" edge with a given conditions (other predicates).
func HasVideosWith(preds ...predicate.Video) predicate.VideoTag {
	return predicate.VideoTag(func(s *sql.Selector) {
		step := newVideosStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VideoTag) predicate.VideoTag {
	return predicate.VideoTag(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VideoTag) predicate.VideoTag {
	return predicate.VideoTag(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.VideoTag) predicate.VideoTag {
	return predicate.VideoTag(sql.NotPredicates(p))
}
