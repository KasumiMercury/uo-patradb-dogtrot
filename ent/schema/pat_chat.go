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
		field.String("message").Annotations(entproto.Field(2)),
		field.Float("magnitude").Annotations(entproto.Skip()),
		field.Float("score").Annotations(entproto.Skip()),
		field.Bool("is_negative").Default(false).Annotations(entproto.Field(3)),
		field.Time("published_at").Annotations(entproto.Field(4)),
		field.Time("created_at").Default(func() time.Time { return time.Now() }).Annotations(entproto.Skip()),
	}
}

// Edges of the Pat_chat.
func (Pat_chat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("video", Video.Type).Ref("Pat_chats").Unique().Required().Annotations(entproto.Field(5)),
	}
}

// Annotations of the Pat_chat.
func (Pat_chat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
