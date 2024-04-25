// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Savings/ent/job"
	"Savings/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// JobDelete is the builder for deleting a Job entity.
type JobDelete struct {
	config
	hooks    []Hook
	mutation *JobMutation
}

// Where appends a list predicates to the JobDelete builder.
func (jd *JobDelete) Where(ps ...predicate.Job) *JobDelete {
	jd.mutation.Where(ps...)
	return jd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (jd *JobDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, jd.sqlExec, jd.mutation, jd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (jd *JobDelete) ExecX(ctx context.Context) int {
	n, err := jd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (jd *JobDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(job.Table, sqlgraph.NewFieldSpec(job.FieldID, field.TypeUint64))
	if ps := jd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, jd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	jd.mutation.done = true
	return affected, err
}

// JobDeleteOne is the builder for deleting a single Job entity.
type JobDeleteOne struct {
	jd *JobDelete
}

// Where appends a list predicates to the JobDelete builder.
func (jdo *JobDeleteOne) Where(ps ...predicate.Job) *JobDeleteOne {
	jdo.jd.mutation.Where(ps...)
	return jdo
}

// Exec executes the deletion query.
func (jdo *JobDeleteOne) Exec(ctx context.Context) error {
	n, err := jdo.jd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{job.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (jdo *JobDeleteOne) ExecX(ctx context.Context) {
	if err := jdo.Exec(ctx); err != nil {
		panic(err)
	}
}
