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
		field.String("source_id").MaxLen(12).Unique().Annotations(entproto.Field(3)),
		field.String("raw").MaxLen(5000).Annotations(entproto.Field(4)),
		field.String("variable").MaxLen(5000).Optional().Annotations(entproto.Field(5)),
		field.Bool("template_confidence").Default(false).Annotations(entproto.Field(6)),
		field.Time("created_at").Default(func() time.Time { return time.Now() }).Annotations(entproto.Skip()),
		field.Time("updated_at").Default(func() time.Time { return time.Now() }).UpdateDefault(func() time.Time { return time.Now() }).Annotations(entproto.Field(8)),
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
		edge.From("video", Video.Type).Ref("descriptions").Unique().Required().Annotations(entproto.Field(9)),
		edge.To("periodic_template", PeriodicDescriptionTemplate.Type).Unique().StorageKey(edge.Column("periodic_id")).Annotations(entproto.Field(7)),
		edge.To("description_changes", DescriptionChange.Type).StorageKey(edge.Column("description_id")).Annotations(entproto.Skip()),
	}
}
