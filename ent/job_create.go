// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Savings/ent/job"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// JobCreate is the builder for creating a Job entity.
type JobCreate struct {
	config
	mutation *JobMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (jc *JobCreate) SetCreatedAt(t time.Time) *JobCreate {
	jc.mutation.SetCreatedAt(t)
	return jc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (jc *JobCreate) SetNillableCreatedAt(t *time.Time) *JobCreate {
	if t != nil {
		jc.SetCreatedAt(*t)
	}
	return jc
}

// SetUpdatedAt sets the "updated_at" field.
func (jc *JobCreate) SetUpdatedAt(t time.Time) *JobCreate {
	jc.mutation.SetUpdatedAt(t)
	return jc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (jc *JobCreate) SetNillableUpdatedAt(t *time.Time) *JobCreate {
	if t != nil {
		jc.SetUpdatedAt(*t)
	}
	return jc
}

// SetName sets the "name" field.
func (jc *JobCreate) SetName(s string) *JobCreate {
	jc.mutation.SetName(s)
	return jc
}

// SetDate sets the "date" field.
func (jc *JobCreate) SetDate(s string) *JobCreate {
	jc.mutation.SetDate(s)
	return jc
}

// SetStatus sets the "status" field.
func (jc *JobCreate) SetStatus(s string) *JobCreate {
	jc.mutation.SetStatus(s)
	return jc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (jc *JobCreate) SetNillableStatus(s *string) *JobCreate {
	if s != nil {
		jc.SetStatus(*s)
	}
	return jc
}

// SetBatch sets the "batch" field.
func (jc *JobCreate) SetBatch(i int) *JobCreate {
	jc.mutation.SetBatch(i)
	return jc
}

// SetNillableBatch sets the "batch" field if the given value is not nil.
func (jc *JobCreate) SetNillableBatch(i *int) *JobCreate {
	if i != nil {
		jc.SetBatch(*i)
	}
	return jc
}

// SetLastProcessedID sets the "last_processed_id" field.
func (jc *JobCreate) SetLastProcessedID(u uint64) *JobCreate {
	jc.mutation.SetLastProcessedID(u)
	return jc
}

// SetNillableLastProcessedID sets the "last_processed_id" field if the given value is not nil.
func (jc *JobCreate) SetNillableLastProcessedID(u *uint64) *JobCreate {
	if u != nil {
		jc.SetLastProcessedID(*u)
	}
	return jc
}

// SetTotalProcessed sets the "total_processed" field.
func (jc *JobCreate) SetTotalProcessed(u uint) *JobCreate {
	jc.mutation.SetTotalProcessed(u)
	return jc
}

// SetNillableTotalProcessed sets the "total_processed" field if the given value is not nil.
func (jc *JobCreate) SetNillableTotalProcessed(u *uint) *JobCreate {
	if u != nil {
		jc.SetTotalProcessed(*u)
	}
	return jc
}

// SetData sets the "data" field.
func (jc *JobCreate) SetData(m map[string]interface{}) *JobCreate {
	jc.mutation.SetData(m)
	return jc
}

// SetID sets the "id" field.
func (jc *JobCreate) SetID(u uint64) *JobCreate {
	jc.mutation.SetID(u)
	return jc
}

// Mutation returns the JobMutation object of the builder.
func (jc *JobCreate) Mutation() *JobMutation {
	return jc.mutation
}

// Save creates the Job in the database.
func (jc *JobCreate) Save(ctx context.Context) (*Job, error) {
	jc.defaults()
	return withHooks(ctx, jc.sqlSave, jc.mutation, jc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (jc *JobCreate) SaveX(ctx context.Context) *Job {
	v, err := jc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jc *JobCreate) Exec(ctx context.Context) error {
	_, err := jc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jc *JobCreate) ExecX(ctx context.Context) {
	if err := jc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jc *JobCreate) defaults() {
	if _, ok := jc.mutation.CreatedAt(); !ok {
		v := job.DefaultCreatedAt()
		jc.mutation.SetCreatedAt(v)
	}
	if _, ok := jc.mutation.UpdatedAt(); !ok {
		v := job.DefaultUpdatedAt()
		jc.mutation.SetUpdatedAt(v)
	}
	if _, ok := jc.mutation.Status(); !ok {
		v := job.DefaultStatus
		jc.mutation.SetStatus(v)
	}
	if _, ok := jc.mutation.Batch(); !ok {
		v := job.DefaultBatch
		jc.mutation.SetBatch(v)
	}
	if _, ok := jc.mutation.LastProcessedID(); !ok {
		v := job.DefaultLastProcessedID
		jc.mutation.SetLastProcessedID(v)
	}
	if _, ok := jc.mutation.TotalProcessed(); !ok {
		v := job.DefaultTotalProcessed
		jc.mutation.SetTotalProcessed(v)
	}
	if _, ok := jc.mutation.Data(); !ok {
		v := job.DefaultData
		jc.mutation.SetData(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jc *JobCreate) check() error {
	if _, ok := jc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Job.created_at"`)}
	}
	if _, ok := jc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Job.updated_at"`)}
	}
	if _, ok := jc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Job.name"`)}
	}
	if _, ok := jc.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "Job.date"`)}
	}
	if _, ok := jc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Job.status"`)}
	}
	if _, ok := jc.mutation.Batch(); !ok {
		return &ValidationError{Name: "batch", err: errors.New(`ent: missing required field "Job.batch"`)}
	}
	if _, ok := jc.mutation.LastProcessedID(); !ok {
		return &ValidationError{Name: "last_processed_id", err: errors.New(`ent: missing required field "Job.last_processed_id"`)}
	}
	if _, ok := jc.mutation.TotalProcessed(); !ok {
		return &ValidationError{Name: "total_processed", err: errors.New(`ent: missing required field "Job.total_processed"`)}
	}
	if _, ok := jc.mutation.Data(); !ok {
		return &ValidationError{Name: "data", err: errors.New(`ent: missing required field "Job.data"`)}
	}
	return nil
}

func (jc *JobCreate) sqlSave(ctx context.Context) (*Job, error) {
	if err := jc.check(); err != nil {
		return nil, err
	}
	_node, _spec := jc.createSpec()
	if err := sqlgraph.CreateNode(ctx, jc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	jc.mutation.id = &_node.ID
	jc.mutation.done = true
	return _node, nil
}

func (jc *JobCreate) createSpec() (*Job, *sqlgraph.CreateSpec) {
	var (
		_node = &Job{config: jc.config}
		_spec = sqlgraph.NewCreateSpec(job.Table, sqlgraph.NewFieldSpec(job.FieldID, field.TypeUint64))
	)
	if id, ok := jc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := jc.mutation.CreatedAt(); ok {
		_spec.SetField(job.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := jc.mutation.UpdatedAt(); ok {
		_spec.SetField(job.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := jc.mutation.Name(); ok {
		_spec.SetField(job.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := jc.mutation.Date(); ok {
		_spec.SetField(job.FieldDate, field.TypeString, value)
		_node.Date = value
	}
	if value, ok := jc.mutation.Status(); ok {
		_spec.SetField(job.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := jc.mutation.Batch(); ok {
		_spec.SetField(job.FieldBatch, field.TypeInt, value)
		_node.Batch = value
	}
	if value, ok := jc.mutation.LastProcessedID(); ok {
		_spec.SetField(job.FieldLastProcessedID, field.TypeUint64, value)
		_node.LastProcessedID = value
	}
	if value, ok := jc.mutation.TotalProcessed(); ok {
		_spec.SetField(job.FieldTotalProcessed, field.TypeUint, value)
		_node.TotalProcessed = value
	}
	if value, ok := jc.mutation.Data(); ok {
		_spec.SetField(job.FieldData, field.TypeJSON, value)
		_node.Data = value
	}
	return _node, _spec
}

// JobCreateBulk is the builder for creating many Job entities in bulk.
type JobCreateBulk struct {
	config
	err      error
	builders []*JobCreate
}

// Save creates the Job entities in the database.
func (jcb *JobCreateBulk) Save(ctx context.Context) ([]*Job, error) {
	if jcb.err != nil {
		return nil, jcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(jcb.builders))
	nodes := make([]*Job, len(jcb.builders))
	mutators := make([]Mutator, len(jcb.builders))
	for i := range jcb.builders {
		func(i int, root context.Context) {
			builder := jcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*JobMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, jcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, jcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, jcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (jcb *JobCreateBulk) SaveX(ctx context.Context) []*Job {
	v, err := jcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jcb *JobCreateBulk) Exec(ctx context.Context) error {
	_, err := jcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jcb *JobCreateBulk) ExecX(ctx context.Context) {
	if err := jcb.Exec(ctx); err != nil {
		panic(err)
	}
}
