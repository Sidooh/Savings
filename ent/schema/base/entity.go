package base

import (
	"entgo.io/ent"
)

// Entity holds the base schema definition for other entities.
type Entity struct {
	ent.Schema
}

// Mixin of the entity.
func (Entity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IdMixin{},
		TimestampMixin{},
	}
}
