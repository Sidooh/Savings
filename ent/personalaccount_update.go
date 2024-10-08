// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Savings/ent/personalaccount"
	"Savings/ent/personalaccounttransaction"
	"Savings/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PersonalAccountUpdate is the builder for updating PersonalAccount entities.
type PersonalAccountUpdate struct {
	config
	hooks    []Hook
	mutation *PersonalAccountMutation
}

// Where appends a list predicates to the PersonalAccountUpdate builder.
func (pau *PersonalAccountUpdate) Where(ps ...predicate.PersonalAccount) *PersonalAccountUpdate {
	pau.mutation.Where(ps...)
	return pau
}

// SetUpdatedAt sets the "updated_at" field.
func (pau *PersonalAccountUpdate) SetUpdatedAt(t time.Time) *PersonalAccountUpdate {
	pau.mutation.SetUpdatedAt(t)
	return pau
}

// SetAccountID sets the "account_id" field.
func (pau *PersonalAccountUpdate) SetAccountID(u uint64) *PersonalAccountUpdate {
	pau.mutation.ResetAccountID()
	pau.mutation.SetAccountID(u)
	return pau
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (pau *PersonalAccountUpdate) SetNillableAccountID(u *uint64) *PersonalAccountUpdate {
	if u != nil {
		pau.SetAccountID(*u)
	}
	return pau
}

// AddAccountID adds u to the "account_id" field.
func (pau *PersonalAccountUpdate) AddAccountID(u int64) *PersonalAccountUpdate {
	pau.mutation.AddAccountID(u)
	return pau
}

// SetType sets the "type" field.
func (pau *PersonalAccountUpdate) SetType(s string) *PersonalAccountUpdate {
	pau.mutation.SetType(s)
	return pau
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pau *PersonalAccountUpdate) SetNillableType(s *string) *PersonalAccountUpdate {
	if s != nil {
		pau.SetType(*s)
	}
	return pau
}

// SetBalance sets the "balance" field.
func (pau *PersonalAccountUpdate) SetBalance(f float32) *PersonalAccountUpdate {
	pau.mutation.ResetBalance()
	pau.mutation.SetBalance(f)
	return pau
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (pau *PersonalAccountUpdate) SetNillableBalance(f *float32) *PersonalAccountUpdate {
	if f != nil {
		pau.SetBalance(*f)
	}
	return pau
}

// AddBalance adds f to the "balance" field.
func (pau *PersonalAccountUpdate) AddBalance(f float32) *PersonalAccountUpdate {
	pau.mutation.AddBalance(f)
	return pau
}

// SetInterest sets the "interest" field.
func (pau *PersonalAccountUpdate) SetInterest(f float32) *PersonalAccountUpdate {
	pau.mutation.ResetInterest()
	pau.mutation.SetInterest(f)
	return pau
}

// SetNillableInterest sets the "interest" field if the given value is not nil.
func (pau *PersonalAccountUpdate) SetNillableInterest(f *float32) *PersonalAccountUpdate {
	if f != nil {
		pau.SetInterest(*f)
	}
	return pau
}

// AddInterest adds f to the "interest" field.
func (pau *PersonalAccountUpdate) AddInterest(f float32) *PersonalAccountUpdate {
	pau.mutation.AddInterest(f)
	return pau
}

// AddTransactionIDs adds the "transactions" edge to the PersonalAccountTransaction entity by IDs.
func (pau *PersonalAccountUpdate) AddTransactionIDs(ids ...uint64) *PersonalAccountUpdate {
	pau.mutation.AddTransactionIDs(ids...)
	return pau
}

// AddTransactions adds the "transactions" edges to the PersonalAccountTransaction entity.
func (pau *PersonalAccountUpdate) AddTransactions(p ...*PersonalAccountTransaction) *PersonalAccountUpdate {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pau.AddTransactionIDs(ids...)
}

// Mutation returns the PersonalAccountMutation object of the builder.
func (pau *PersonalAccountUpdate) Mutation() *PersonalAccountMutation {
	return pau.mutation
}

// ClearTransactions clears all "transactions" edges to the PersonalAccountTransaction entity.
func (pau *PersonalAccountUpdate) ClearTransactions() *PersonalAccountUpdate {
	pau.mutation.ClearTransactions()
	return pau
}

// RemoveTransactionIDs removes the "transactions" edge to PersonalAccountTransaction entities by IDs.
func (pau *PersonalAccountUpdate) RemoveTransactionIDs(ids ...uint64) *PersonalAccountUpdate {
	pau.mutation.RemoveTransactionIDs(ids...)
	return pau
}

// RemoveTransactions removes "transactions" edges to PersonalAccountTransaction entities.
func (pau *PersonalAccountUpdate) RemoveTransactions(p ...*PersonalAccountTransaction) *PersonalAccountUpdate {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pau.RemoveTransactionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pau *PersonalAccountUpdate) Save(ctx context.Context) (int, error) {
	pau.defaults()
	return withHooks(ctx, pau.sqlSave, pau.mutation, pau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pau *PersonalAccountUpdate) SaveX(ctx context.Context) int {
	affected, err := pau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pau *PersonalAccountUpdate) Exec(ctx context.Context) error {
	_, err := pau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pau *PersonalAccountUpdate) ExecX(ctx context.Context) {
	if err := pau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pau *PersonalAccountUpdate) defaults() {
	if _, ok := pau.mutation.UpdatedAt(); !ok {
		v := personalaccount.UpdateDefaultUpdatedAt()
		pau.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pau *PersonalAccountUpdate) check() error {
	if v, ok := pau.mutation.AccountID(); ok {
		if err := personalaccount.AccountIDValidator(v); err != nil {
			return &ValidationError{Name: "account_id", err: fmt.Errorf(`ent: validator failed for field "PersonalAccount.account_id": %w`, err)}
		}
	}
	if v, ok := pau.mutation.GetType(); ok {
		if err := personalaccount.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "PersonalAccount.type": %w`, err)}
		}
	}
	if v, ok := pau.mutation.Balance(); ok {
		if err := personalaccount.BalanceValidator(v); err != nil {
			return &ValidationError{Name: "balance", err: fmt.Errorf(`ent: validator failed for field "PersonalAccount.balance": %w`, err)}
		}
	}
	if v, ok := pau.mutation.Interest(); ok {
		if err := personalaccount.InterestValidator(v); err != nil {
			return &ValidationError{Name: "interest", err: fmt.Errorf(`ent: validator failed for field "PersonalAccount.interest": %w`, err)}
		}
	}
	return nil
}

func (pau *PersonalAccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(personalaccount.Table, personalaccount.Columns, sqlgraph.NewFieldSpec(personalaccount.FieldID, field.TypeUint64))
	if ps := pau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pau.mutation.UpdatedAt(); ok {
		_spec.SetField(personalaccount.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pau.mutation.AccountID(); ok {
		_spec.SetField(personalaccount.FieldAccountID, field.TypeUint64, value)
	}
	if value, ok := pau.mutation.AddedAccountID(); ok {
		_spec.AddField(personalaccount.FieldAccountID, field.TypeUint64, value)
	}
	if value, ok := pau.mutation.GetType(); ok {
		_spec.SetField(personalaccount.FieldType, field.TypeString, value)
	}
	if value, ok := pau.mutation.Balance(); ok {
		_spec.SetField(personalaccount.FieldBalance, field.TypeFloat32, value)
	}
	if value, ok := pau.mutation.AddedBalance(); ok {
		_spec.AddField(personalaccount.FieldBalance, field.TypeFloat32, value)
	}
	if value, ok := pau.mutation.Interest(); ok {
		_spec.SetField(personalaccount.FieldInterest, field.TypeFloat32, value)
	}
	if value, ok := pau.mutation.AddedInterest(); ok {
		_spec.AddField(personalaccount.FieldInterest, field.TypeFloat32, value)
	}
	if pau.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   personalaccount.TransactionsTable,
			Columns: []string{personalaccount.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(personalaccounttransaction.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pau.mutation.RemovedTransactionsIDs(); len(nodes) > 0 && !pau.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   personalaccount.TransactionsTable,
			Columns: []string{personalaccount.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(personalaccounttransaction.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pau.mutation.TransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   personalaccount.TransactionsTable,
			Columns: []string{personalaccount.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(personalaccounttransaction.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{personalaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pau.mutation.done = true
	return n, nil
}

// PersonalAccountUpdateOne is the builder for updating a single PersonalAccount entity.
type PersonalAccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PersonalAccountMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (pauo *PersonalAccountUpdateOne) SetUpdatedAt(t time.Time) *PersonalAccountUpdateOne {
	pauo.mutation.SetUpdatedAt(t)
	return pauo
}

// SetAccountID sets the "account_id" field.
func (pauo *PersonalAccountUpdateOne) SetAccountID(u uint64) *PersonalAccountUpdateOne {
	pauo.mutation.ResetAccountID()
	pauo.mutation.SetAccountID(u)
	return pauo
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (pauo *PersonalAccountUpdateOne) SetNillableAccountID(u *uint64) *PersonalAccountUpdateOne {
	if u != nil {
		pauo.SetAccountID(*u)
	}
	return pauo
}

// AddAccountID adds u to the "account_id" field.
func (pauo *PersonalAccountUpdateOne) AddAccountID(u int64) *PersonalAccountUpdateOne {
	pauo.mutation.AddAccountID(u)
	return pauo
}

// SetType sets the "type" field.
func (pauo *PersonalAccountUpdateOne) SetType(s string) *PersonalAccountUpdateOne {
	pauo.mutation.SetType(s)
	return pauo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pauo *PersonalAccountUpdateOne) SetNillableType(s *string) *PersonalAccountUpdateOne {
	if s != nil {
		pauo.SetType(*s)
	}
	return pauo
}

// SetBalance sets the "balance" field.
func (pauo *PersonalAccountUpdateOne) SetBalance(f float32) *PersonalAccountUpdateOne {
	pauo.mutation.ResetBalance()
	pauo.mutation.SetBalance(f)
	return pauo
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (pauo *PersonalAccountUpdateOne) SetNillableBalance(f *float32) *PersonalAccountUpdateOne {
	if f != nil {
		pauo.SetBalance(*f)
	}
	return pauo
}

// AddBalance adds f to the "balance" field.
func (pauo *PersonalAccountUpdateOne) AddBalance(f float32) *PersonalAccountUpdateOne {
	pauo.mutation.AddBalance(f)
	return pauo
}

// SetInterest sets the "interest" field.
func (pauo *PersonalAccountUpdateOne) SetInterest(f float32) *PersonalAccountUpdateOne {
	pauo.mutation.ResetInterest()
	pauo.mutation.SetInterest(f)
	return pauo
}

// SetNillableInterest sets the "interest" field if the given value is not nil.
func (pauo *PersonalAccountUpdateOne) SetNillableInterest(f *float32) *PersonalAccountUpdateOne {
	if f != nil {
		pauo.SetInterest(*f)
	}
	return pauo
}

// AddInterest adds f to the "interest" field.
func (pauo *PersonalAccountUpdateOne) AddInterest(f float32) *PersonalAccountUpdateOne {
	pauo.mutation.AddInterest(f)
	return pauo
}

// AddTransactionIDs adds the "transactions" edge to the PersonalAccountTransaction entity by IDs.
func (pauo *PersonalAccountUpdateOne) AddTransactionIDs(ids ...uint64) *PersonalAccountUpdateOne {
	pauo.mutation.AddTransactionIDs(ids...)
	return pauo
}

// AddTransactions adds the "transactions" edges to the PersonalAccountTransaction entity.
func (pauo *PersonalAccountUpdateOne) AddTransactions(p ...*PersonalAccountTransaction) *PersonalAccountUpdateOne {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pauo.AddTransactionIDs(ids...)
}

// Mutation returns the PersonalAccountMutation object of the builder.
func (pauo *PersonalAccountUpdateOne) Mutation() *PersonalAccountMutation {
	return pauo.mutation
}

// ClearTransactions clears all "transactions" edges to the PersonalAccountTransaction entity.
func (pauo *PersonalAccountUpdateOne) ClearTransactions() *PersonalAccountUpdateOne {
	pauo.mutation.ClearTransactions()
	return pauo
}

// RemoveTransactionIDs removes the "transactions" edge to PersonalAccountTransaction entities by IDs.
func (pauo *PersonalAccountUpdateOne) RemoveTransactionIDs(ids ...uint64) *PersonalAccountUpdateOne {
	pauo.mutation.RemoveTransactionIDs(ids...)
	return pauo
}

// RemoveTransactions removes "transactions" edges to PersonalAccountTransaction entities.
func (pauo *PersonalAccountUpdateOne) RemoveTransactions(p ...*PersonalAccountTransaction) *PersonalAccountUpdateOne {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pauo.RemoveTransactionIDs(ids...)
}

// Where appends a list predicates to the PersonalAccountUpdate builder.
func (pauo *PersonalAccountUpdateOne) Where(ps ...predicate.PersonalAccount) *PersonalAccountUpdateOne {
	pauo.mutation.Where(ps...)
	return pauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pauo *PersonalAccountUpdateOne) Select(field string, fields ...string) *PersonalAccountUpdateOne {
	pauo.fields = append([]string{field}, fields...)
	return pauo
}

// Save executes the query and returns the updated PersonalAccount entity.
func (pauo *PersonalAccountUpdateOne) Save(ctx context.Context) (*PersonalAccount, error) {
	pauo.defaults()
	return withHooks(ctx, pauo.sqlSave, pauo.mutation, pauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pauo *PersonalAccountUpdateOne) SaveX(ctx context.Context) *PersonalAccount {
	node, err := pauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pauo *PersonalAccountUpdateOne) Exec(ctx context.Context) error {
	_, err := pauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pauo *PersonalAccountUpdateOne) ExecX(ctx context.Context) {
	if err := pauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pauo *PersonalAccountUpdateOne) defaults() {
	if _, ok := pauo.mutation.UpdatedAt(); !ok {
		v := personalaccount.UpdateDefaultUpdatedAt()
		pauo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pauo *PersonalAccountUpdateOne) check() error {
	if v, ok := pauo.mutation.AccountID(); ok {
		if err := personalaccount.AccountIDValidator(v); err != nil {
			return &ValidationError{Name: "account_id", err: fmt.Errorf(`ent: validator failed for field "PersonalAccount.account_id": %w`, err)}
		}
	}
	if v, ok := pauo.mutation.GetType(); ok {
		if err := personalaccount.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "PersonalAccount.type": %w`, err)}
		}
	}
	if v, ok := pauo.mutation.Balance(); ok {
		if err := personalaccount.BalanceValidator(v); err != nil {
			return &ValidationError{Name: "balance", err: fmt.Errorf(`ent: validator failed for field "PersonalAccount.balance": %w`, err)}
		}
	}
	if v, ok := pauo.mutation.Interest(); ok {
		if err := personalaccount.InterestValidator(v); err != nil {
			return &ValidationError{Name: "interest", err: fmt.Errorf(`ent: validator failed for field "PersonalAccount.interest": %w`, err)}
		}
	}
	return nil
}

func (pauo *PersonalAccountUpdateOne) sqlSave(ctx context.Context) (_node *PersonalAccount, err error) {
	if err := pauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(personalaccount.Table, personalaccount.Columns, sqlgraph.NewFieldSpec(personalaccount.FieldID, field.TypeUint64))
	id, ok := pauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PersonalAccount.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, personalaccount.FieldID)
		for _, f := range fields {
			if !personalaccount.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != personalaccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pauo.mutation.UpdatedAt(); ok {
		_spec.SetField(personalaccount.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pauo.mutation.AccountID(); ok {
		_spec.SetField(personalaccount.FieldAccountID, field.TypeUint64, value)
	}
	if value, ok := pauo.mutation.AddedAccountID(); ok {
		_spec.AddField(personalaccount.FieldAccountID, field.TypeUint64, value)
	}
	if value, ok := pauo.mutation.GetType(); ok {
		_spec.SetField(personalaccount.FieldType, field.TypeString, value)
	}
	if value, ok := pauo.mutation.Balance(); ok {
		_spec.SetField(personalaccount.FieldBalance, field.TypeFloat32, value)
	}
	if value, ok := pauo.mutation.AddedBalance(); ok {
		_spec.AddField(personalaccount.FieldBalance, field.TypeFloat32, value)
	}
	if value, ok := pauo.mutation.Interest(); ok {
		_spec.SetField(personalaccount.FieldInterest, field.TypeFloat32, value)
	}
	if value, ok := pauo.mutation.AddedInterest(); ok {
		_spec.AddField(personalaccount.FieldInterest, field.TypeFloat32, value)
	}
	if pauo.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   personalaccount.TransactionsTable,
			Columns: []string{personalaccount.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(personalaccounttransaction.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pauo.mutation.RemovedTransactionsIDs(); len(nodes) > 0 && !pauo.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   personalaccount.TransactionsTable,
			Columns: []string{personalaccount.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(personalaccounttransaction.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pauo.mutation.TransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   personalaccount.TransactionsTable,
			Columns: []string{personalaccount.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(personalaccounttransaction.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PersonalAccount{config: pauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{personalaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pauo.mutation.done = true
	return _node, nil
}
