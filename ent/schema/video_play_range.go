package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
)

// Video_play_range holds the schema definition for the Video_play_range entity.
type Video_play_range struct {
	ent.Schema
}

// Mixin of the Video_play_range.
func (Video_play_range) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.MixinWithPrefix("VPR"),
	}
}

// Fields of the Video_play_range.
func (Video_play_range) Fields() []ent.Field {
	return []ent.Field{
		field.Int("start_seconds").Default(0).Annotations(entproto.Field(2)),
		field.Int("end_seconds").Optional().Annotations(entproto.Field(3)),
	}
}

// Edges of the Video_play_range.
func (Video_play_range) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("video", Video.Type).Ref("video_play_ranges").Required().Unique().Annotations(entproto.Skip()),
	}
}

// Annotations of the Video_play_range.
func (Video_play_range) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
