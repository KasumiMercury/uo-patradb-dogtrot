package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
)

// VideoDisallowRange holds the schema definition for the VideoDisallowRange entity.
type VideoDisallowRange struct {
	ent.Schema
}

// Mixin of the VideoDisallowRange.
func (VideoDisallowRange) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.Mixin{},
	}
}

// Fields of the VideoDisallowRange.
func (VideoDisallowRange) Fields() []ent.Field {
	return []ent.Field{
		field.Int("start_seconds").NonNegative().Annotations(entproto.Field(2)),
		field.Int("end_seconds").Positive().Annotations(entproto.Field(3)),
	}
}

// Annotations of the VideoDisallowRange.
func (VideoDisallowRange) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}

// Edges of the VideoDisallowRange.
func (VideoDisallowRange) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("video", Video.Type).Ref("video_disallow_ranges").Required().Unique().Annotations(entproto.Skip()),
	}
}
