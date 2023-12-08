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

// PatChat holds the schema definition for the PatChat entity.
type PatChat struct {
	ent.Schema
}

// Mixin of the PatChat.
func (PatChat) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.Mixin{},
	}
}

// Fields of the PatChat.
func (PatChat) Fields() []ent.Field {
	return []ent.Field{
		field.String("message").MaxLen(200).Annotations(entproto.Field(2)),
		field.Float("magnitude").Min(0.0).Annotations(entproto.Skip()),
		field.Float("score").Range(-1.0, 1.0).Annotations(entproto.Skip()),
		field.Bool("is_negative").Default(false).Annotations(entproto.Field(3)),
		field.Time("published_at").Annotations(entproto.Field(4)),
		field.Time("created_at").Default(func() time.Time { return time.Now() }).Annotations(entproto.Skip()),
	}
}

// Annotations of the PatChat.
func (PatChat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}

// Edges of the PatChat.
func (PatChat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("video", Video.Type).Ref("Pat_chats").Unique().Required().Annotations(entproto.Skip()),
	}
}
