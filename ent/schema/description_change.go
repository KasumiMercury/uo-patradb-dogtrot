package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
	"time"
)

// Description_change holds the schema definition for the Description_change entity.
type Description_change struct {
	ent.Schema
}

// Mixin of the Description_change.
func (Description_change) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.MixinWithPrefix("VDC"),
	}
}

// Fields of the Description_change.
func (Description_change) Fields() []ent.Field {
	return []ent.Field{
		field.String("raw"),
		field.String("variable").Optional(),
		field.String("normalized_variable").Optional(),
		field.Time("changed_at").Default(func() time.Time { return time.Now() }),
	}
}

// Edges of the Description_change.
func (Description_change) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("description", Description.Type).Ref("description_changes").Unique().Required(),
	}
}
