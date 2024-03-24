// Code generated by ent, DO NOT EDIT.

package personalaccount

import (
	"Savings/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldEQ(FieldUpdatedAt, v))
}

// AccountID applies equality check predicate on the "account_id" field. It's identical to AccountIDEQ.
func AccountID(v uint64) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldEQ(FieldAccountID, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldEQ(FieldType, v))
}

// Balance applies equality check predicate on the "balance" field. It's identical to BalanceEQ.
func Balance(v float32) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldEQ(FieldBalance, v))
}

// Interest applies equality check predicate on the "interest" field. It's identical to InterestEQ.
func Interest(v float32) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldEQ(FieldInterest, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldLTE(FieldUpdatedAt, v))
}

// AccountIDEQ applies the EQ predicate on the "account_id" field.
func AccountIDEQ(v uint64) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldEQ(FieldAccountID, v))
}

// AccountIDNEQ applies the NEQ predicate on the "account_id" field.
func AccountIDNEQ(v uint64) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldNEQ(FieldAccountID, v))
}

// AccountIDIn applies the In predicate on the "account_id" field.
func AccountIDIn(vs ...uint64) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldIn(FieldAccountID, vs...))
}

// AccountIDNotIn applies the NotIn predicate on the "account_id" field.
func AccountIDNotIn(vs ...uint64) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldNotIn(FieldAccountID, vs...))
}

// AccountIDGT applies the GT predicate on the "account_id" field.
func AccountIDGT(v uint64) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldGT(FieldAccountID, v))
}

// AccountIDGTE applies the GTE predicate on the "account_id" field.
func AccountIDGTE(v uint64) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldGTE(FieldAccountID, v))
}

// AccountIDLT applies the LT predicate on the "account_id" field.
func AccountIDLT(v uint64) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldLT(FieldAccountID, v))
}

// AccountIDLTE applies the LTE predicate on the "account_id" field.
func AccountIDLTE(v uint64) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldLTE(FieldAccountID, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldContainsFold(FieldType, v))
}

// BalanceEQ applies the EQ predicate on the "balance" field.
func BalanceEQ(v float32) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldEQ(FieldBalance, v))
}

// BalanceNEQ applies the NEQ predicate on the "balance" field.
func BalanceNEQ(v float32) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldNEQ(FieldBalance, v))
}

// BalanceIn applies the In predicate on the "balance" field.
func BalanceIn(vs ...float32) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldIn(FieldBalance, vs...))
}

// BalanceNotIn applies the NotIn predicate on the "balance" field.
func BalanceNotIn(vs ...float32) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldNotIn(FieldBalance, vs...))
}

// BalanceGT applies the GT predicate on the "balance" field.
func BalanceGT(v float32) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldGT(FieldBalance, v))
}

// BalanceGTE applies the GTE predicate on the "balance" field.
func BalanceGTE(v float32) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldGTE(FieldBalance, v))
}

// BalanceLT applies the LT predicate on the "balance" field.
func BalanceLT(v float32) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldLT(FieldBalance, v))
}

// BalanceLTE applies the LTE predicate on the "balance" field.
func BalanceLTE(v float32) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldLTE(FieldBalance, v))
}

// InterestEQ applies the EQ predicate on the "interest" field.
func InterestEQ(v float32) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldEQ(FieldInterest, v))
}

// InterestNEQ applies the NEQ predicate on the "interest" field.
func InterestNEQ(v float32) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldNEQ(FieldInterest, v))
}

// InterestIn applies the In predicate on the "interest" field.
func InterestIn(vs ...float32) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldIn(FieldInterest, vs...))
}

// InterestNotIn applies the NotIn predicate on the "interest" field.
func InterestNotIn(vs ...float32) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldNotIn(FieldInterest, vs...))
}

// InterestGT applies the GT predicate on the "interest" field.
func InterestGT(v float32) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldGT(FieldInterest, v))
}

// InterestGTE applies the GTE predicate on the "interest" field.
func InterestGTE(v float32) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldGTE(FieldInterest, v))
}

// InterestLT applies the LT predicate on the "interest" field.
func InterestLT(v float32) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldLT(FieldInterest, v))
}

// InterestLTE applies the LTE predicate on the "interest" field.
func InterestLTE(v float32) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.FieldLTE(FieldInterest, v))
}

// HasTransactions applies the HasEdge predicate on the "transactions" edge.
func HasTransactions() predicate.PersonalAccount {
	return predicate.PersonalAccount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TransactionsTable, TransactionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTransactionsWith applies the HasEdge predicate on the "transactions" edge with a given conditions (other predicates).
func HasTransactionsWith(preds ...predicate.PersonalAccountTransaction) predicate.PersonalAccount {
	return predicate.PersonalAccount(func(s *sql.Selector) {
		step := newTransactionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PersonalAccount) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PersonalAccount) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PersonalAccount) predicate.PersonalAccount {
	return predicate.PersonalAccount(sql.NotPredicates(p))
}
