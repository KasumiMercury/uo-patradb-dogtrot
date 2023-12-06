package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
)

// Channel holds the schema definition for the Channel entity.
type Channel struct {
	ent.Schema
}

// Mixin of the Channel.
func (Channel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.Mixin{},
	}
}

// Fields of the Channel.
func (Channel) Fields() []ent.Field {
	return []ent.Field{
		field.String("display_name").Annotations(entproto.Field(2)),
		field.String("channel_id").Unique().Annotations(entproto.Field(3)),
		field.String("handle").Annotations(entproto.Field(4)),
		field.String("thumbnail_url").Annotations(entproto.Field(5)),
	}
}

// Annotations of the Channel.
func (Channel) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}

// Edges of the Channel.
func (Channel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("videos", Video.Type).Ref("channel").Annotations(entproto.Field(6)),
	}
}
