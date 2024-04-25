package schema

import (
	"Savings/ent/schema/base"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Job holds the schema definition for the Job entity.
type Job struct {
	base.Entity
}

// Fields of the Job.
func (Job) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("date"),
		field.String("status").
			Default("PENDING"),
		field.Int("batch").
			Default(1000),
		field.Uint64("last_processed_id").
			Default(0),
		field.Uint("total_processed").
			Default(0),
		field.JSON("data", map[string]interface{}{}).Default(map[string]interface{}{}),
	}
}
