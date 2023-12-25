package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
)

// StreamSchedule holds the schema definition for the StreamSchedule entity.
type StreamSchedule struct {
	ent.Schema
}

// Mixin of the StreamSchedule.
func (StreamSchedule) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.Mixin{},
	}
}

// Fields of the StreamSchedule.
func (StreamSchedule) Fields() []ent.Field {
	return []ent.Field{
		field.Time("scheduled_at").Annotations(entproto.Field(2)),
		field.String("Title").MaxLen(100).Annotations(entproto.Field(3)),
		field.Time("created_at").Annotations(entproto.Skip()),
		field.Time("updated_at").Annotations(entproto.Skip()),
	}
}

// Annotations of the StreamSchedule.
func (StreamSchedule) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}

// Edges of the StreamSchedule.
func (StreamSchedule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("videos", Video.Type).StorageKey(edge.Column("video_id")).Unique().Annotations(entproto.Skip()),
	}
}
