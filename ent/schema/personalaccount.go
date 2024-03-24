package schema

import (
	"Savings/ent/schema/base"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// PersonalAccount holds the schema definition for the PersonalAccount entity.
type PersonalAccount struct {
	base.Entity
}

// Fields of the PersonalAccount.
func (PersonalAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("account_id").
			Positive(),
		field.String("type").
			MaxLen(20),
		field.Float32("balance").
			SchemaType(map[string]string{
				dialect.MySQL:    "decimal(15,4)", // Override MySQL.
				dialect.Postgres: "numeric",       // Override Postgres.
			}),
		field.Float32("interest").
			SchemaType(map[string]string{
				dialect.MySQL:    "decimal(10,2)", // Override MySQL.
				dialect.Postgres: "numeric",       // Override Postgres.
			}),
	}
}

func (PersonalAccount) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("account_id", "type").
			Unique(),
	}
}

// Edges of the PersonalAccount.
func (PersonalAccount) Edges() []ent.Edge {
	return nil
}
