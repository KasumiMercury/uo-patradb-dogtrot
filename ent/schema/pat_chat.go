package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
	"time"
)

// Pat_chat holds the schema definition for the Pat_chat entity.
type Pat_chat struct {
	ent.Schema
}

// Mixin of the Pat_chat.
func (Pat_chat) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.MixinWithPrefix("PC"),
	}
}

// Fields of the Pat_chat.
func (Pat_chat) Fields() []ent.Field {
	return []ent.Field{
		field.String("message"),
		field.Float("magnitude"),
		field.Float("score"),
		field.Bool("is_negative").Default(false),
		field.Time("published_at"),
		field.Time("created_at").Default(func() time.Time { return time.Now() }),
	}
}

// Edges of the Pat_chat.
func (Pat_chat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("video", Video.Type).Ref("pat_chats").Unique().Required(),
	}
}
