package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
)

// CategoryDescriptionTemplate holds the schema definition for the CategoryDescriptionTemplate entity.
type CategoryDescriptionTemplate struct {
	ent.Schema
}

// Mixin of the Category_description_template.
func (CategoryDescriptionTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.MixinWithPrefix("CDT"),
	}
}

// Fields of the CategoryDescriptionTemplate.
func (CategoryDescriptionTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").Annotations(entproto.Field(2)),
		field.Time("start_use_at").Optional().Annotations(entproto.Field(3)),
		field.Time("last_use_at").Optional().Annotations(entproto.Field(4)),
	}
}

// Annotations of the CategoryDescriptionTemplate.
func (CategoryDescriptionTemplate) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}

// Edges of the CategoryDescriptionTemplate.
func (CategoryDescriptionTemplate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("descriptions", Description.Type).Ref("category_description_template").Annotations(entproto.Skip()),
	}
}