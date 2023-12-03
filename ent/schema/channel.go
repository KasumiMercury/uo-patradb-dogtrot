package schema

import (
	"entgo.io/ent"
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
		pulid.MixinWithPrefix("CH"),
	}
}

// Fields of the Channel.
func (Channel) Fields() []ent.Field {
	return []ent.Field{
		field.String("display_name"),
		field.String("channel_id").Unique(),
		field.String("handle"),
		field.String("thumbnail_url"),
	}
}

// Edges of the Channel.
func (Channel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("videos", Video.Type).Ref("channel"),
	}
}
