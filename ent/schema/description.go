package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
	"time"
)

// Description holds the schema definition for the Description entity.
type Description struct {
	ent.Schema
}

// Mixin of the Description.
func (Description) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.MixinWithPrefix("VD"),
	}
}

// Fields of the Description.
func (Description) Fields() []ent.Field {
	return []ent.Field{
		field.String("raw"),
		field.String("variable").Optional(),
		field.String("normalized_variable").Optional(),
		field.Time("created_at").Default(func() time.Time { return time.Now() }),
		field.Time("updated_at").Default(func() time.Time { return time.Now() }).UpdateDefault(func() time.Time { return time.Now() }),
	}
}

// Edges of the Description.
func (Description) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("video", Video.Type).Ref("descriptions").Unique().Required(),
		edge.From("periodic_description_template", Periodic_description_template.Type).Ref("descriptions").Unique(),
		edge.From("category_description_template", Category_description_template.Type).Ref("descriptions").Unique(),
		edge.To("description_changes", Description_change.Type).StorageKey(edge.Column("description_id")),
	}
}
