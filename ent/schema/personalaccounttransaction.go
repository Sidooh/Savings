package schema

import (
	"Savings/ent/schema/base"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PersonalAccountTransaction holds the schema definition for the PersonalAccountTransaction entity.
type PersonalAccountTransaction struct {
	base.Entity
}

// Fields of the PersonalAccountTransaction.
func (PersonalAccountTransaction) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("personal_account_id"),
		field.String("type").
			MaxLen(20),
		field.Float32("amount").
			Default(0).
			Min(0).
			SchemaType(map[string]string{
				dialect.MySQL:    "decimal(10,4)", // Override MySQL.
				dialect.Postgres: "numeric",       // Override Postgres.
			}),
		field.Float32("balance").
			Default(0).
			Min(0).
			SchemaType(map[string]string{
				dialect.MySQL:    "decimal(15,4)", // Override MySQL.
				dialect.Postgres: "numeric",       // Override Postgres.
			}),
		field.String("status").
			Default("PENDING"),
	}
}

// Edges of the PersonalAccountTransaction.
func (PersonalAccountTransaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", PersonalAccount.Type).
			Ref("transactions").
			Unique().
			Required().
			Field("personal_account_id"),
	}
}
