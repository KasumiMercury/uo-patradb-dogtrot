package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
)

// VideoTag holds the schema definition for the VideoTag entity.
type VideoTag struct {
	ent.Schema
}

// Mixin of the VideoTag.
func (VideoTag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.Mixin{},
	}
}

// Fields of the VideoTag.
func (VideoTag) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MaxLen(50).Annotations(entproto.Field(2)),
	}
}

// Annotations of the Video.
func (VideoTag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}

// Edges of the VideoTag.
func (VideoTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("videos", Video.Type).Annotations(entproto.Field(3)),
	}
}

// Indexes of the VideoTag.
func (VideoTag) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("title"),
	}
}
