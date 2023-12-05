package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
)

// Video_disallow_range holds the schema definition for the Video_disallow_range entity.
type Video_disallow_range struct {
	ent.Schema
}

// Mixin of the Video_disallow_range.
func (Video_disallow_range) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.MixinWithPrefix("VDR"),
	}
}

// Fields of the Video_disallow_range.
func (Video_disallow_range) Fields() []ent.Field {
	return []ent.Field{
		field.Int("start_seconds").Annotations(entproto.Field(2)),
		field.Int("end_seconds").Annotations(entproto.Field(3)),
	}
}

// Edges of the Video_disallow_range.
func (Video_disallow_range) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("video", Video.Type).Ref("video_disallow_ranges").Required().Unique().Annotations(entproto.Skip()),
	}
}

// Annotations of the Video_disallow_range.
func (Video_disallow_range) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
