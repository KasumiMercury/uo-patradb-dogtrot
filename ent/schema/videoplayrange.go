package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
)

// VideoPlayRange holds the schema definition for the VideoPlayRange entity.
type VideoPlayRange struct {
	ent.Schema
}

// Mixin of the VideoPlayRange.
func (VideoPlayRange) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.Mixin{},
	}
}

// Fields of the VideoPlayRange.
func (VideoPlayRange) Fields() []ent.Field {
	return []ent.Field{
		field.Int("start_seconds").Default(0).Annotations(entproto.Field(2)),
		field.Int("end_seconds").Optional().Annotations(entproto.Field(3)),
	}
}

// Annotations of the VideoPlayRange.
func (VideoPlayRange) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}

// Edges of the VideoPlayRange.
func (VideoPlayRange) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("video", Video.Type).Ref("video_play_ranges").Required().Unique().Annotations(entproto.Skip()),
	}
}
