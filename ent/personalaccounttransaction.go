// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Savings/ent/personalaccount"
	"Savings/ent/personalaccounttransaction"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// PersonalAccountTransaction is the model entity for the PersonalAccountTransaction schema.
type PersonalAccountTransaction struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// PersonalAccountID holds the value of the "personal_account_id" field.
	PersonalAccountID uint64 `json:"personal_account_id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount float32 `json:"amount,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PersonalAccountTransactionQuery when eager-loading is set.
	Edges                         PersonalAccountTransactionEdges `json:"edges"`
	personal_account_transactions *uint64
	selectValues                  sql.SelectValues
}

// PersonalAccountTransactionEdges holds the relations/edges for other nodes in the graph.
type PersonalAccountTransactionEdges struct {
	// Account holds the value of the account edge.
	Account *PersonalAccount `json:"account,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AccountOrErr returns the Account value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PersonalAccountTransactionEdges) AccountOrErr() (*PersonalAccount, error) {
	if e.Account != nil {
		return e.Account, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: personalaccount.Label}
	}
	return nil, &NotLoadedError{edge: "account"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PersonalAccountTransaction) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case personalaccounttransaction.FieldAmount:
			values[i] = new(sql.NullFloat64)
		case personalaccounttransaction.FieldID, personalaccounttransaction.FieldPersonalAccountID:
			values[i] = new(sql.NullInt64)
		case personalaccounttransaction.FieldType, personalaccounttransaction.FieldStatus:
			values[i] = new(sql.NullString)
		case personalaccounttransaction.FieldCreatedAt, personalaccounttransaction.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case personalaccounttransaction.ForeignKeys[0]: // personal_account_transactions
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PersonalAccountTransaction fields.
func (pat *PersonalAccountTransaction) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case personalaccounttransaction.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pat.ID = uint64(value.Int64)
		case personalaccounttransaction.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pat.CreatedAt = value.Time
			}
		case personalaccounttransaction.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pat.UpdatedAt = value.Time
			}
		case personalaccounttransaction.FieldPersonalAccountID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field personal_account_id", values[i])
			} else if value.Valid {
				pat.PersonalAccountID = uint64(value.Int64)
			}
		case personalaccounttransaction.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				pat.Type = value.String
			}
		case personalaccounttransaction.FieldAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				pat.Amount = float32(value.Float64)
			}
		case personalaccounttransaction.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pat.Status = value.String
			}
		case personalaccounttransaction.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field personal_account_transactions", value)
			} else if value.Valid {
				pat.personal_account_transactions = new(uint64)
				*pat.personal_account_transactions = uint64(value.Int64)
			}
		default:
			pat.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PersonalAccountTransaction.
// This includes values selected through modifiers, order, etc.
func (pat *PersonalAccountTransaction) Value(name string) (ent.Value, error) {
	return pat.selectValues.Get(name)
}

// QueryAccount queries the "account" edge of the PersonalAccountTransaction entity.
func (pat *PersonalAccountTransaction) QueryAccount() *PersonalAccountQuery {
	return NewPersonalAccountTransactionClient(pat.config).QueryAccount(pat)
}

// Update returns a builder for updating this PersonalAccountTransaction.
// Note that you need to call PersonalAccountTransaction.Unwrap() before calling this method if this PersonalAccountTransaction
// was returned from a transaction, and the transaction was committed or rolled back.
func (pat *PersonalAccountTransaction) Update() *PersonalAccountTransactionUpdateOne {
	return NewPersonalAccountTransactionClient(pat.config).UpdateOne(pat)
}

// Unwrap unwraps the PersonalAccountTransaction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pat *PersonalAccountTransaction) Unwrap() *PersonalAccountTransaction {
	_tx, ok := pat.config.driver.(*txDriver)
	if !ok {
		panic("ent: PersonalAccountTransaction is not a transactional entity")
	}
	pat.config.driver = _tx.drv
	return pat
}

// String implements the fmt.Stringer.
func (pat *PersonalAccountTransaction) String() string {
	var builder strings.Builder
	builder.WriteString("PersonalAccountTransaction(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pat.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pat.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pat.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("personal_account_id=")
	builder.WriteString(fmt.Sprintf("%v", pat.PersonalAccountID))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(pat.Type)
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", pat.Amount))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(pat.Status)
	builder.WriteByte(')')
	return builder.String()
}

// PersonalAccountTransactions is a parsable slice of PersonalAccountTransaction.
type PersonalAccountTransactions []*PersonalAccountTransaction
