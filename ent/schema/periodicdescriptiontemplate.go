package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
)

// PeriodicDescriptionTemplate holds the schema definition for the PeriodicDescriptionTemplate entity.
type PeriodicDescriptionTemplate struct {
	ent.Schema
}

// Mixin of the Periodic_description_template.
func (PeriodicDescriptionTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.MixinWithPrefix("PDT"),
	}
}

// Fields of the PeriodicDescriptionTemplate.
func (PeriodicDescriptionTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").Annotations(entproto.Field(2)),
		field.Time("start_use_at").Optional().Annotations(entproto.Field(3)),
		field.Time("last_use_at").Optional().Annotations(entproto.Field(4)),
	}
}

// Annotations of the PeriodicDescriptionTemplate.
func (PeriodicDescriptionTemplate) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}

// Edges of the PeriodicDescriptionTemplate.
func (PeriodicDescriptionTemplate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("descriptions", Description.Type).Ref("periodic_description_template").Annotations(entproto.Skip()),
	}
}
