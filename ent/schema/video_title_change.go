package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
	"time"
)

// Video_title_change holds the schema definition for the Video_title_change entity.
type Video_title_change struct {
	ent.Schema
}

// Mixin of the Video_title_change.
func (Video_title_change) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.MixinWithPrefix("VTC"),
	}
}

// Fields of the Video_title_change.
func (Video_title_change) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("normalized_title"),
		field.Time("changed_at").Default(func() time.Time { return time.Now() }),
	}
}

// Edges of the Video_title_change.
func (Video_title_change) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("video", Video.Type).Ref("video_title_changes").Required(),
	}
}
