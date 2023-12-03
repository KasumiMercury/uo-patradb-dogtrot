// Code generated by ent, DO NOT EDIT.

package channel

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/predicate"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
)

// ID filters vertices based on their ID field.
func ID(id pulid.ID) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id pulid.ID) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id pulid.ID) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...pulid.ID) predicate.Channel {
	return predicate.Channel(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...pulid.ID) predicate.Channel {
	return predicate.Channel(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id pulid.ID) predicate.Channel {
	return predicate.Channel(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id pulid.ID) predicate.Channel {
	return predicate.Channel(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id pulid.ID) predicate.Channel {
	return predicate.Channel(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id pulid.ID) predicate.Channel {
	return predicate.Channel(sql.FieldLTE(FieldID, id))
}

// DisplayName applies equality check predicate on the "display_name" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldDisplayName, v))
}

// ChannelID applies equality check predicate on the "channel_id" field. It's identical to ChannelIDEQ.
func ChannelID(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldChannelID, v))
}

// Handle applies equality check predicate on the "handle" field. It's identical to HandleEQ.
func Handle(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldHandle, v))
}

// ThumbnailURL applies equality check predicate on the "thumbnail_url" field. It's identical to ThumbnailURLEQ.
func ThumbnailURL(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldThumbnailURL, v))
}

// DisplayNameEQ applies the EQ predicate on the "display_name" field.
func DisplayNameEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldDisplayName, v))
}

// DisplayNameNEQ applies the NEQ predicate on the "display_name" field.
func DisplayNameNEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldDisplayName, v))
}

// DisplayNameIn applies the In predicate on the "display_name" field.
func DisplayNameIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldIn(FieldDisplayName, vs...))
}

// DisplayNameNotIn applies the NotIn predicate on the "display_name" field.
func DisplayNameNotIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldNotIn(FieldDisplayName, vs...))
}

// DisplayNameGT applies the GT predicate on the "display_name" field.
func DisplayNameGT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGT(FieldDisplayName, v))
}

// DisplayNameGTE applies the GTE predicate on the "display_name" field.
func DisplayNameGTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGTE(FieldDisplayName, v))
}

// DisplayNameLT applies the LT predicate on the "display_name" field.
func DisplayNameLT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLT(FieldDisplayName, v))
}

// DisplayNameLTE applies the LTE predicate on the "display_name" field.
func DisplayNameLTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLTE(FieldDisplayName, v))
}

// DisplayNameContains applies the Contains predicate on the "display_name" field.
func DisplayNameContains(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContains(FieldDisplayName, v))
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "display_name" field.
func DisplayNameHasPrefix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasPrefix(FieldDisplayName, v))
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "display_name" field.
func DisplayNameHasSuffix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasSuffix(FieldDisplayName, v))
}

// DisplayNameEqualFold applies the EqualFold predicate on the "display_name" field.
func DisplayNameEqualFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEqualFold(FieldDisplayName, v))
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "display_name" field.
func DisplayNameContainsFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContainsFold(FieldDisplayName, v))
}

// ChannelIDEQ applies the EQ predicate on the "channel_id" field.
func ChannelIDEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldChannelID, v))
}

// ChannelIDNEQ applies the NEQ predicate on the "channel_id" field.
func ChannelIDNEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldChannelID, v))
}

// ChannelIDIn applies the In predicate on the "channel_id" field.
func ChannelIDIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldIn(FieldChannelID, vs...))
}

// ChannelIDNotIn applies the NotIn predicate on the "channel_id" field.
func ChannelIDNotIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldNotIn(FieldChannelID, vs...))
}

// ChannelIDGT applies the GT predicate on the "channel_id" field.
func ChannelIDGT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGT(FieldChannelID, v))
}

// ChannelIDGTE applies the GTE predicate on the "channel_id" field.
func ChannelIDGTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGTE(FieldChannelID, v))
}

// ChannelIDLT applies the LT predicate on the "channel_id" field.
func ChannelIDLT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLT(FieldChannelID, v))
}

// ChannelIDLTE applies the LTE predicate on the "channel_id" field.
func ChannelIDLTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLTE(FieldChannelID, v))
}

// ChannelIDContains applies the Contains predicate on the "channel_id" field.
func ChannelIDContains(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContains(FieldChannelID, v))
}

// ChannelIDHasPrefix applies the HasPrefix predicate on the "channel_id" field.
func ChannelIDHasPrefix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasPrefix(FieldChannelID, v))
}

// ChannelIDHasSuffix applies the HasSuffix predicate on the "channel_id" field.
func ChannelIDHasSuffix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasSuffix(FieldChannelID, v))
}

// ChannelIDEqualFold applies the EqualFold predicate on the "channel_id" field.
func ChannelIDEqualFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEqualFold(FieldChannelID, v))
}

// ChannelIDContainsFold applies the ContainsFold predicate on the "channel_id" field.
func ChannelIDContainsFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContainsFold(FieldChannelID, v))
}

// HandleEQ applies the EQ predicate on the "handle" field.
func HandleEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldHandle, v))
}

// HandleNEQ applies the NEQ predicate on the "handle" field.
func HandleNEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldHandle, v))
}

// HandleIn applies the In predicate on the "handle" field.
func HandleIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldIn(FieldHandle, vs...))
}

// HandleNotIn applies the NotIn predicate on the "handle" field.
func HandleNotIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldNotIn(FieldHandle, vs...))
}

// HandleGT applies the GT predicate on the "handle" field.
func HandleGT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGT(FieldHandle, v))
}

// HandleGTE applies the GTE predicate on the "handle" field.
func HandleGTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGTE(FieldHandle, v))
}

// HandleLT applies the LT predicate on the "handle" field.
func HandleLT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLT(FieldHandle, v))
}

// HandleLTE applies the LTE predicate on the "handle" field.
func HandleLTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLTE(FieldHandle, v))
}

// HandleContains applies the Contains predicate on the "handle" field.
func HandleContains(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContains(FieldHandle, v))
}

// HandleHasPrefix applies the HasPrefix predicate on the "handle" field.
func HandleHasPrefix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasPrefix(FieldHandle, v))
}

// HandleHasSuffix applies the HasSuffix predicate on the "handle" field.
func HandleHasSuffix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasSuffix(FieldHandle, v))
}

// HandleEqualFold applies the EqualFold predicate on the "handle" field.
func HandleEqualFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEqualFold(FieldHandle, v))
}

// HandleContainsFold applies the ContainsFold predicate on the "handle" field.
func HandleContainsFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContainsFold(FieldHandle, v))
}

// ThumbnailURLEQ applies the EQ predicate on the "thumbnail_url" field.
func ThumbnailURLEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldThumbnailURL, v))
}

// ThumbnailURLNEQ applies the NEQ predicate on the "thumbnail_url" field.
func ThumbnailURLNEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldThumbnailURL, v))
}

// ThumbnailURLIn applies the In predicate on the "thumbnail_url" field.
func ThumbnailURLIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldIn(FieldThumbnailURL, vs...))
}

// ThumbnailURLNotIn applies the NotIn predicate on the "thumbnail_url" field.
func ThumbnailURLNotIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldNotIn(FieldThumbnailURL, vs...))
}

// ThumbnailURLGT applies the GT predicate on the "thumbnail_url" field.
func ThumbnailURLGT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGT(FieldThumbnailURL, v))
}

// ThumbnailURLGTE applies the GTE predicate on the "thumbnail_url" field.
func ThumbnailURLGTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGTE(FieldThumbnailURL, v))
}

// ThumbnailURLLT applies the LT predicate on the "thumbnail_url" field.
func ThumbnailURLLT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLT(FieldThumbnailURL, v))
}

// ThumbnailURLLTE applies the LTE predicate on the "thumbnail_url" field.
func ThumbnailURLLTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLTE(FieldThumbnailURL, v))
}

// ThumbnailURLContains applies the Contains predicate on the "thumbnail_url" field.
func ThumbnailURLContains(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContains(FieldThumbnailURL, v))
}

// ThumbnailURLHasPrefix applies the HasPrefix predicate on the "thumbnail_url" field.
func ThumbnailURLHasPrefix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasPrefix(FieldThumbnailURL, v))
}

// ThumbnailURLHasSuffix applies the HasSuffix predicate on the "thumbnail_url" field.
func ThumbnailURLHasSuffix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasSuffix(FieldThumbnailURL, v))
}

// ThumbnailURLEqualFold applies the EqualFold predicate on the "thumbnail_url" field.
func ThumbnailURLEqualFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEqualFold(FieldThumbnailURL, v))
}

// ThumbnailURLContainsFold applies the ContainsFold predicate on the "thumbnail_url" field.
func ThumbnailURLContainsFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContainsFold(FieldThumbnailURL, v))
}

// HasVideos applies the HasEdge predicate on the "videos" edge.
func HasVideos() predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, VideosTable, VideosPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVideosWith applies the HasEdge predicate on the "videos" edge with a given conditions (other predicates).
func HasVideosWith(preds ...predicate.Video) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		step := newVideosStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Channel) predicate.Channel {
	return predicate.Channel(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Channel) predicate.Channel {
	return predicate.Channel(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Channel) predicate.Channel {
	return predicate.Channel(sql.NotPredicates(p))
}