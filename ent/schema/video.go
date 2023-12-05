package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.String("video_id").Unique().Annotations(entproto.Field(2)),
		field.String("title").Annotations(entproto.Field(3)),
		field.String("normalized_title").Annotations(entproto.Skip()),
		field.Int("duration_seconds").Annotations(entproto.Field(4)),
		field.Bool("is_collaboration").Annotations(entproto.Field(5)),
		field.String("status").Annotations(entproto.Field(6)),
		field.String("chat_id").Optional().Annotations(entproto.Skip()),
		field.Bool("has_time_range").Default(false).Annotations(entproto.Field(7)),
		field.Time("scheduled_at").Optional().Annotations(entproto.Field(8)),
		field.Time("actual_start_at").Optional().Annotations(entproto.Field(9)),
		field.Time("published_at").Annotations(entproto.Field(10)),
		field.Time("created_at").Default(func() time.Time { return time.Now() }).Annotations(entproto.Skip()),
		field.Time("updated_at").Default(func() time.Time { return time.Now() }).UpdateDefault(func() time.Time { return time.Now() }).Annotations(entproto.Field(11)),
	}
}

// Edges of the Video.
func (Video) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("descriptions", Description.Type).Unique().StorageKey(edge.Column("video_id")).Annotations(entproto.Skip()),
		edge.To("channel", Channel.Type).Annotations(entproto.Field(13)),
		edge.To("video_play_ranges", Video_play_range.Type).StorageKey(edge.Column("video_id")).Annotations(entproto.Skip()),
		edge.To("video_disallow_ranges", Video_disallow_range.Type).StorageKey(edge.Column("video_id")).Annotations(entproto.Skip()),
		edge.To("video_title_changes", Video_title_change.Type).StorageKey(edge.Column("video_id")).Annotations(entproto.Skip()),
		edge.To("Pat_chats", Pat_chat.Type).StorageKey(edge.Column("video_id")).Annotations(entproto.Skip()),
	}
}

// Annotations of the Video.
func (Video) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
