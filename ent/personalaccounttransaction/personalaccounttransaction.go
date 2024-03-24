// Code generated by ent, DO NOT EDIT.

package personalaccounttransaction

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the personalaccounttransaction type in the database.
	Label = "personal_account_transaction"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldPersonalAccountID holds the string denoting the personal_account_id field in the database.
	FieldPersonalAccountID = "personal_account_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeAccount holds the string denoting the account edge name in mutations.
	EdgeAccount = "account"
	// Table holds the table name of the personalaccounttransaction in the database.
	Table = "personal_account_transactions"
	// AccountTable is the table that holds the account relation/edge.
	AccountTable = "personal_account_transactions"
	// AccountInverseTable is the table name for the PersonalAccount entity.
	// It exists in this package in order to avoid circular dependency with the "personalaccount" package.
	AccountInverseTable = "personal_accounts"
	// AccountColumn is the table column denoting the account relation/edge.
	AccountColumn = "personal_account_transactions"
)

// Columns holds all SQL columns for personalaccounttransaction fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldPersonalAccountID,
	FieldType,
	FieldAmount,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "personal_account_transactions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"personal_account_transactions",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// PersonalAccountIDValidator is a validator for the "personal_account_id" field. It is called by the builders before save.
	PersonalAccountIDValidator func(uint64) error
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount float32
	// AmountValidator is a validator for the "amount" field. It is called by the builders before save.
	AmountValidator func(float32) error
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus string
)

// OrderOption defines the ordering options for the PersonalAccountTransaction queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByPersonalAccountID orders the results by the personal_account_id field.
func ByPersonalAccountID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPersonalAccountID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByAccountField orders the results by account field.
func ByAccountField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAccountStep(), sql.OrderByField(field, opts...))
	}
}
func newAccountStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AccountInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
	)
}
