package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
	"time"
)

// VideoTitleChange holds the schema definition for the VideoTitleChange entity.
type VideoTitleChange struct {
	ent.Schema
}

// Mixin of the VideoTitleChange.
func (VideoTitleChange) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.Mixin{},
	}
}

// Fields of the VideoTitleChange.
func (VideoTitleChange) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MaxLen(100),
		field.String("normalized_title").Optional(),
		field.Time("changed_at").Default(func() time.Time { return time.Now() }),
	}
}

// Edges of the VideoTitleChange.
func (VideoTitleChange) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("video", Video.Type).Ref("video_title_changes").Unique().Required(),
	}
}
