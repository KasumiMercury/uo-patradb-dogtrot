package schema

import (
	"entgo.io/ent"
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
	return nil
}
