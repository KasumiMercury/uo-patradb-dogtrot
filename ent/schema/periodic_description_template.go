package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.String("text").Annotations(entproto.Field(2)),
		field.Time("start_use_at").Optional().Annotations(entproto.Field(3)),
		field.Time("last_use_at").Optional().Annotations(entproto.Field(4)),
	}
}

// Edges of the Periodic_description_template.
func (Periodic_description_template) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("descriptions", Description.Type).StorageKey(edge.Column("periodic_template_id")).Annotations(entproto.Field(5)),
	}
}

// Annotations of the Periodic_description_template.
func (Periodic_description_template) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
