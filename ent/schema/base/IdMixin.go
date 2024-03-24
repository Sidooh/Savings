package base

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// IdMixin implements the ent.Mixin for sharing
// ID field with package schemas.
type IdMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (IdMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"). // make unsigned
					Unique().
					Immutable(),
	}
}
