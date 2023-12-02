package pulid

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type Mixin struct {
	mixin.Schema
	prefix string
}

func MixinWithPrefix(prefix string) Mixin {
	return Mixin{prefix: prefix}
}

func (m Mixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(ID("")).
			DefaultFunc(func() ID { return MustNew(m.prefix) }),
	}
}
