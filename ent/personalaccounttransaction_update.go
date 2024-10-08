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

// PersonalAccountTransactionUpdate is the builder for updating PersonalAccountTransaction entities.
type PersonalAccountTransactionUpdate struct {
	config
	hooks    []Hook
	mutation *PersonalAccountTransactionMutation
}

// Where appends a list predicates to the PersonalAccountTransactionUpdate builder.
func (patu *PersonalAccountTransactionUpdate) Where(ps ...predicate.PersonalAccountTransaction) *PersonalAccountTransactionUpdate {
	patu.mutation.Where(ps...)
	return patu
}

// SetUpdatedAt sets the "updated_at" field.
func (patu *PersonalAccountTransactionUpdate) SetUpdatedAt(t time.Time) *PersonalAccountTransactionUpdate {
	patu.mutation.SetUpdatedAt(t)
	return patu
}

// SetPersonalAccountID sets the "personal_account_id" field.
func (patu *PersonalAccountTransactionUpdate) SetPersonalAccountID(u uint64) *PersonalAccountTransactionUpdate {
	patu.mutation.SetPersonalAccountID(u)
	return patu
}

// SetNillablePersonalAccountID sets the "personal_account_id" field if the given value is not nil.
func (patu *PersonalAccountTransactionUpdate) SetNillablePersonalAccountID(u *uint64) *PersonalAccountTransactionUpdate {
	if u != nil {
		patu.SetPersonalAccountID(*u)
	}
	return patu
}

// SetType sets the "type" field.
func (patu *PersonalAccountTransactionUpdate) SetType(s string) *PersonalAccountTransactionUpdate {
	patu.mutation.SetType(s)
	return patu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (patu *PersonalAccountTransactionUpdate) SetNillableType(s *string) *PersonalAccountTransactionUpdate {
	if s != nil {
		patu.SetType(*s)
	}
	return patu
}

// SetAmount sets the "amount" field.
func (patu *PersonalAccountTransactionUpdate) SetAmount(f float32) *PersonalAccountTransactionUpdate {
	patu.mutation.ResetAmount()
	patu.mutation.SetAmount(f)
	return patu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (patu *PersonalAccountTransactionUpdate) SetNillableAmount(f *float32) *PersonalAccountTransactionUpdate {
	if f != nil {
		patu.SetAmount(*f)
	}
	return patu
}

// AddAmount adds f to the "amount" field.
func (patu *PersonalAccountTransactionUpdate) AddAmount(f float32) *PersonalAccountTransactionUpdate {
	patu.mutation.AddAmount(f)
	return patu
}

// SetBalance sets the "balance" field.
func (patu *PersonalAccountTransactionUpdate) SetBalance(f float32) *PersonalAccountTransactionUpdate {
	patu.mutation.ResetBalance()
	patu.mutation.SetBalance(f)
	return patu
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (patu *PersonalAccountTransactionUpdate) SetNillableBalance(f *float32) *PersonalAccountTransactionUpdate {
	if f != nil {
		patu.SetBalance(*f)
	}
	return patu
}

// AddBalance adds f to the "balance" field.
func (patu *PersonalAccountTransactionUpdate) AddBalance(f float32) *PersonalAccountTransactionUpdate {
	patu.mutation.AddBalance(f)
	return patu
}

// SetDescription sets the "description" field.
func (patu *PersonalAccountTransactionUpdate) SetDescription(s string) *PersonalAccountTransactionUpdate {
	patu.mutation.SetDescription(s)
	return patu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (patu *PersonalAccountTransactionUpdate) SetNillableDescription(s *string) *PersonalAccountTransactionUpdate {
	if s != nil {
		patu.SetDescription(*s)
	}
	return patu
}

// SetStatus sets the "status" field.
func (patu *PersonalAccountTransactionUpdate) SetStatus(s string) *PersonalAccountTransactionUpdate {
	patu.mutation.SetStatus(s)
	return patu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (patu *PersonalAccountTransactionUpdate) SetNillableStatus(s *string) *PersonalAccountTransactionUpdate {
	if s != nil {
		patu.SetStatus(*s)
	}
	return patu
}

// SetAccountID sets the "account" edge to the PersonalAccount entity by ID.
func (patu *PersonalAccountTransactionUpdate) SetAccountID(id uint64) *PersonalAccountTransactionUpdate {
	patu.mutation.SetAccountID(id)
	return patu
}

// SetAccount sets the "account" edge to the PersonalAccount entity.
func (patu *PersonalAccountTransactionUpdate) SetAccount(p *PersonalAccount) *PersonalAccountTransactionUpdate {
	return patu.SetAccountID(p.ID)
}

// Mutation returns the PersonalAccountTransactionMutation object of the builder.
func (patu *PersonalAccountTransactionUpdate) Mutation() *PersonalAccountTransactionMutation {
	return patu.mutation
}

// ClearAccount clears the "account" edge to the PersonalAccount entity.
func (patu *PersonalAccountTransactionUpdate) ClearAccount() *PersonalAccountTransactionUpdate {
	patu.mutation.ClearAccount()
	return patu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (patu *PersonalAccountTransactionUpdate) Save(ctx context.Context) (int, error) {
	patu.defaults()
	return withHooks(ctx, patu.sqlSave, patu.mutation, patu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (patu *PersonalAccountTransactionUpdate) SaveX(ctx context.Context) int {
	affected, err := patu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (patu *PersonalAccountTransactionUpdate) Exec(ctx context.Context) error {
	_, err := patu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (patu *PersonalAccountTransactionUpdate) ExecX(ctx context.Context) {
	if err := patu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (patu *PersonalAccountTransactionUpdate) defaults() {
	if _, ok := patu.mutation.UpdatedAt(); !ok {
		v := personalaccounttransaction.UpdateDefaultUpdatedAt()
		patu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (patu *PersonalAccountTransactionUpdate) check() error {
	if v, ok := patu.mutation.GetType(); ok {
		if err := personalaccounttransaction.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "PersonalAccountTransaction.type": %w`, err)}
		}
	}
	if v, ok := patu.mutation.Amount(); ok {
		if err := personalaccounttransaction.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "PersonalAccountTransaction.amount": %w`, err)}
		}
	}
	if v, ok := patu.mutation.Balance(); ok {
		if err := personalaccounttransaction.BalanceValidator(v); err != nil {
			return &ValidationError{Name: "balance", err: fmt.Errorf(`ent: validator failed for field "PersonalAccountTransaction.balance": %w`, err)}
		}
	}
	if _, ok := patu.mutation.AccountID(); patu.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PersonalAccountTransaction.account"`)
	}
	return nil
}

func (patu *PersonalAccountTransactionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := patu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(personalaccounttransaction.Table, personalaccounttransaction.Columns, sqlgraph.NewFieldSpec(personalaccounttransaction.FieldID, field.TypeUint64))
	if ps := patu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := patu.mutation.UpdatedAt(); ok {
		_spec.SetField(personalaccounttransaction.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := patu.mutation.GetType(); ok {
		_spec.SetField(personalaccounttransaction.FieldType, field.TypeString, value)
	}
	if value, ok := patu.mutation.Amount(); ok {
		_spec.SetField(personalaccounttransaction.FieldAmount, field.TypeFloat32, value)
	}
	if value, ok := patu.mutation.AddedAmount(); ok {
		_spec.AddField(personalaccounttransaction.FieldAmount, field.TypeFloat32, value)
	}
	if value, ok := patu.mutation.Balance(); ok {
		_spec.SetField(personalaccounttransaction.FieldBalance, field.TypeFloat32, value)
	}
	if value, ok := patu.mutation.AddedBalance(); ok {
		_spec.AddField(personalaccounttransaction.FieldBalance, field.TypeFloat32, value)
	}
	if value, ok := patu.mutation.Description(); ok {
		_spec.SetField(personalaccounttransaction.FieldDescription, field.TypeString, value)
	}
	if value, ok := patu.mutation.Status(); ok {
		_spec.SetField(personalaccounttransaction.FieldStatus, field.TypeString, value)
	}
	if patu.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   personalaccounttransaction.AccountTable,
			Columns: []string{personalaccounttransaction.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(personalaccount.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := patu.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   personalaccounttransaction.AccountTable,
			Columns: []string{personalaccounttransaction.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(personalaccount.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, patu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{personalaccounttransaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	patu.mutation.done = true
	return n, nil
}

// PersonalAccountTransactionUpdateOne is the builder for updating a single PersonalAccountTransaction entity.
type PersonalAccountTransactionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PersonalAccountTransactionMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (patuo *PersonalAccountTransactionUpdateOne) SetUpdatedAt(t time.Time) *PersonalAccountTransactionUpdateOne {
	patuo.mutation.SetUpdatedAt(t)
	return patuo
}

// SetPersonalAccountID sets the "personal_account_id" field.
func (patuo *PersonalAccountTransactionUpdateOne) SetPersonalAccountID(u uint64) *PersonalAccountTransactionUpdateOne {
	patuo.mutation.SetPersonalAccountID(u)
	return patuo
}

// SetNillablePersonalAccountID sets the "personal_account_id" field if the given value is not nil.
func (patuo *PersonalAccountTransactionUpdateOne) SetNillablePersonalAccountID(u *uint64) *PersonalAccountTransactionUpdateOne {
	if u != nil {
		patuo.SetPersonalAccountID(*u)
	}
	return patuo
}

// SetType sets the "type" field.
func (patuo *PersonalAccountTransactionUpdateOne) SetType(s string) *PersonalAccountTransactionUpdateOne {
	patuo.mutation.SetType(s)
	return patuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (patuo *PersonalAccountTransactionUpdateOne) SetNillableType(s *string) *PersonalAccountTransactionUpdateOne {
	if s != nil {
		patuo.SetType(*s)
	}
	return patuo
}

// SetAmount sets the "amount" field.
func (patuo *PersonalAccountTransactionUpdateOne) SetAmount(f float32) *PersonalAccountTransactionUpdateOne {
	patuo.mutation.ResetAmount()
	patuo.mutation.SetAmount(f)
	return patuo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (patuo *PersonalAccountTransactionUpdateOne) SetNillableAmount(f *float32) *PersonalAccountTransactionUpdateOne {
	if f != nil {
		patuo.SetAmount(*f)
	}
	return patuo
}

// AddAmount adds f to the "amount" field.
func (patuo *PersonalAccountTransactionUpdateOne) AddAmount(f float32) *PersonalAccountTransactionUpdateOne {
	patuo.mutation.AddAmount(f)
	return patuo
}

// SetBalance sets the "balance" field.
func (patuo *PersonalAccountTransactionUpdateOne) SetBalance(f float32) *PersonalAccountTransactionUpdateOne {
	patuo.mutation.ResetBalance()
	patuo.mutation.SetBalance(f)
	return patuo
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (patuo *PersonalAccountTransactionUpdateOne) SetNillableBalance(f *float32) *PersonalAccountTransactionUpdateOne {
	if f != nil {
		patuo.SetBalance(*f)
	}
	return patuo
}

// AddBalance adds f to the "balance" field.
func (patuo *PersonalAccountTransactionUpdateOne) AddBalance(f float32) *PersonalAccountTransactionUpdateOne {
	patuo.mutation.AddBalance(f)
	return patuo
}

// SetDescription sets the "description" field.
func (patuo *PersonalAccountTransactionUpdateOne) SetDescription(s string) *PersonalAccountTransactionUpdateOne {
	patuo.mutation.SetDescription(s)
	return patuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (patuo *PersonalAccountTransactionUpdateOne) SetNillableDescription(s *string) *PersonalAccountTransactionUpdateOne {
	if s != nil {
		patuo.SetDescription(*s)
	}
	return patuo
}

// SetStatus sets the "status" field.
func (patuo *PersonalAccountTransactionUpdateOne) SetStatus(s string) *PersonalAccountTransactionUpdateOne {
	patuo.mutation.SetStatus(s)
	return patuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (patuo *PersonalAccountTransactionUpdateOne) SetNillableStatus(s *string) *PersonalAccountTransactionUpdateOne {
	if s != nil {
		patuo.SetStatus(*s)
	}
	return patuo
}

// SetAccountID sets the "account" edge to the PersonalAccount entity by ID.
func (patuo *PersonalAccountTransactionUpdateOne) SetAccountID(id uint64) *PersonalAccountTransactionUpdateOne {
	patuo.mutation.SetAccountID(id)
	return patuo
}

// SetAccount sets the "account" edge to the PersonalAccount entity.
func (patuo *PersonalAccountTransactionUpdateOne) SetAccount(p *PersonalAccount) *PersonalAccountTransactionUpdateOne {
	return patuo.SetAccountID(p.ID)
}

// Mutation returns the PersonalAccountTransactionMutation object of the builder.
func (patuo *PersonalAccountTransactionUpdateOne) Mutation() *PersonalAccountTransactionMutation {
	return patuo.mutation
}

// ClearAccount clears the "account" edge to the PersonalAccount entity.
func (patuo *PersonalAccountTransactionUpdateOne) ClearAccount() *PersonalAccountTransactionUpdateOne {
	patuo.mutation.ClearAccount()
	return patuo
}

// Where appends a list predicates to the PersonalAccountTransactionUpdate builder.
func (patuo *PersonalAccountTransactionUpdateOne) Where(ps ...predicate.PersonalAccountTransaction) *PersonalAccountTransactionUpdateOne {
	patuo.mutation.Where(ps...)
	return patuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (patuo *PersonalAccountTransactionUpdateOne) Select(field string, fields ...string) *PersonalAccountTransactionUpdateOne {
	patuo.fields = append([]string{field}, fields...)
	return patuo
}

// Save executes the query and returns the updated PersonalAccountTransaction entity.
func (patuo *PersonalAccountTransactionUpdateOne) Save(ctx context.Context) (*PersonalAccountTransaction, error) {
	patuo.defaults()
	return withHooks(ctx, patuo.sqlSave, patuo.mutation, patuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (patuo *PersonalAccountTransactionUpdateOne) SaveX(ctx context.Context) *PersonalAccountTransaction {
	node, err := patuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (patuo *PersonalAccountTransactionUpdateOne) Exec(ctx context.Context) error {
	_, err := patuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (patuo *PersonalAccountTransactionUpdateOne) ExecX(ctx context.Context) {
	if err := patuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (patuo *PersonalAccountTransactionUpdateOne) defaults() {
	if _, ok := patuo.mutation.UpdatedAt(); !ok {
		v := personalaccounttransaction.UpdateDefaultUpdatedAt()
		patuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (patuo *PersonalAccountTransactionUpdateOne) check() error {
	if v, ok := patuo.mutation.GetType(); ok {
		if err := personalaccounttransaction.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "PersonalAccountTransaction.type": %w`, err)}
		}
	}
	if v, ok := patuo.mutation.Amount(); ok {
		if err := personalaccounttransaction.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "PersonalAccountTransaction.amount": %w`, err)}
		}
	}
	if v, ok := patuo.mutation.Balance(); ok {
		if err := personalaccounttransaction.BalanceValidator(v); err != nil {
			return &ValidationError{Name: "balance", err: fmt.Errorf(`ent: validator failed for field "PersonalAccountTransaction.balance": %w`, err)}
		}
	}
	if _, ok := patuo.mutation.AccountID(); patuo.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PersonalAccountTransaction.account"`)
	}
	return nil
}

func (patuo *PersonalAccountTransactionUpdateOne) sqlSave(ctx context.Context) (_node *PersonalAccountTransaction, err error) {
	if err := patuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(personalaccounttransaction.Table, personalaccounttransaction.Columns, sqlgraph.NewFieldSpec(personalaccounttransaction.FieldID, field.TypeUint64))
	id, ok := patuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PersonalAccountTransaction.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := patuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, personalaccounttransaction.FieldID)
		for _, f := range fields {
			if !personalaccounttransaction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != personalaccounttransaction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := patuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := patuo.mutation.UpdatedAt(); ok {
		_spec.SetField(personalaccounttransaction.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := patuo.mutation.GetType(); ok {
		_spec.SetField(personalaccounttransaction.FieldType, field.TypeString, value)
	}
	if value, ok := patuo.mutation.Amount(); ok {
		_spec.SetField(personalaccounttransaction.FieldAmount, field.TypeFloat32, value)
	}
	if value, ok := patuo.mutation.AddedAmount(); ok {
		_spec.AddField(personalaccounttransaction.FieldAmount, field.TypeFloat32, value)
	}
	if value, ok := patuo.mutation.Balance(); ok {
		_spec.SetField(personalaccounttransaction.FieldBalance, field.TypeFloat32, value)
	}
	if value, ok := patuo.mutation.AddedBalance(); ok {
		_spec.AddField(personalaccounttransaction.FieldBalance, field.TypeFloat32, value)
	}
	if value, ok := patuo.mutation.Description(); ok {
		_spec.SetField(personalaccounttransaction.FieldDescription, field.TypeString, value)
	}
	if value, ok := patuo.mutation.Status(); ok {
		_spec.SetField(personalaccounttransaction.FieldStatus, field.TypeString, value)
	}
	if patuo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   personalaccounttransaction.AccountTable,
			Columns: []string{personalaccounttransaction.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(personalaccount.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := patuo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   personalaccounttransaction.AccountTable,
			Columns: []string{personalaccounttransaction.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(personalaccount.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PersonalAccountTransaction{config: patuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, patuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{personalaccounttransaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	patuo.mutation.done = true
	return _node, nil
}
