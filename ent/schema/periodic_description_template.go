package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
)

// Periodic_description_template holds the schema definition for the Periodic_description_template entity.
type Periodic_description_template struct {
	ent.Schema
}

// Mixin of the Periodic_description_template.
func (Periodic_description_template) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.MixinWithPrefix("PDT"),
	}
}

// Fields of the Periodic_description_template.
func (Periodic_description_template) Fields() []ent.Field {
	return []ent.Field{
		field.String("text"),
		field.Time("start_use_at").Optional(),
		field.Time("last_use_at").Optional(),
	}
}

// Edges of the Periodic_description_template.
func (Periodic_description_template) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("descriptions", Description.Type).StorageKey(edge.Column("periodic_template_id")),
	}
}
