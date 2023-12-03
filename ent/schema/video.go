package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
	"time"
)

// Video holds the schema definition for the Video entity.
type Video struct {
	ent.Schema
}

// Mixin of the Video.
func (Video) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.MixinWithPrefix("VM"),
	}
}

// Fields of the Video.
func (Video) Fields() []ent.Field {
	return []ent.Field{
		field.String("video_id").Unique(),
		field.String("title"),
		field.String("normalized_title"),
		field.Int("duration_seconds"),
		field.Bool("is_collaboration"),
		field.String("status"),
		field.String("chat_id").Optional(),
		field.Bool("has_time_range").Default(false),
		field.Time("scheduled_at").Optional(),
		field.Time("actual_start_at").Optional(),
		field.Time("published_at"),
		field.Time("created_at").Default(func() time.Time { return time.Now() }),
		field.Time("updated_at").Default(func() time.Time { return time.Now() }).UpdateDefault(func() time.Time { return time.Now() }),
	}
}

// Edges of the Video.
func (Video) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("descriptions", Description.Type).Unique().StorageKey(edge.Column("video_id")),
		edge.To("channel", Channel.Type),
		edge.To("video_play_range", Video_play_range.Type).StorageKey(edge.Column("video_id")),
		edge.To("video_disallow_range", Video_disallow_range.Type).StorageKey(edge.Column("video_id")),
		edge.To("video_title_change", Video_title_change.Type).StorageKey(edge.Column("video_id")),
	}
}
