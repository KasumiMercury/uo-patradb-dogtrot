package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		pulid.Mixin{},
	}
}

// Fields of the Description.
func (Description) Fields() []ent.Field {
	return []ent.Field{
		field.String("raw").MaxLen(5000).Annotations(entproto.Field(2)),
		field.String("variable").Optional().Annotations(entproto.Field(3)),
		field.Time("created_at").Default(func() time.Time { return time.Now() }).Annotations(entproto.Skip()),
		field.Time("updated_at").Default(func() time.Time { return time.Now() }).UpdateDefault(func() time.Time { return time.Now() }).Annotations(entproto.Field(4)),
	}
}

// Annotations of the Description.
func (Description) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}

// Edges of the Description.
func (Description) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("video", Video.Type).Ref("descriptions").Unique().Required().Annotations(entproto.Field(5)),
		edge.To("periodic_template", PeriodicDescriptionTemplate.Type).Unique().StorageKey(edge.Column("periodic_id")).Annotations(entproto.Field(6)),
		edge.To("category_template", CategoryDescriptionTemplate.Type).Unique().StorageKey(edge.Column("category_id")).Annotations(entproto.Field(7)),
		edge.To("description_changes", DescriptionChange.Type).StorageKey(edge.Column("description_id")).Annotations(entproto.Skip()),
	}
}
