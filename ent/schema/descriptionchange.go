package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
	"time"
)

// DescriptionChange holds the schema definition for the DescriptionChange entity.
type DescriptionChange struct {
	ent.Schema
}

// Mixin of the DescriptionChange.
func (DescriptionChange) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.Mixin{},
	}
}

// Fields of the DescriptionChange.
func (DescriptionChange) Fields() []ent.Field {
	return []ent.Field{
		field.String("raw"),
		field.String("variable").Optional(),
		field.String("normalized_variable").Optional(),
		field.Time("changed_at").Default(func() time.Time { return time.Now() }),
	}
}

// Edges of the DescriptionChange.
func (DescriptionChange) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("description", Description.Type).Ref("description_changes").Unique().Required(),
	}
}
