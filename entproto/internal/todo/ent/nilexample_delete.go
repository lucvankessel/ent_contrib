// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/contrib/entproto/internal/todo/ent/nilexample"
	"entgo.io/contrib/entproto/internal/todo/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NilExampleDelete is the builder for deleting a NilExample entity.
type NilExampleDelete struct {
	config
	hooks    []Hook
	mutation *NilExampleMutation
}

// Where appends a list predicates to the NilExampleDelete builder.
func (ned *NilExampleDelete) Where(ps ...predicate.NilExample) *NilExampleDelete {
	ned.mutation.Where(ps...)
	return ned
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ned *NilExampleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, NilExampleMutation](ctx, ned.sqlExec, ned.mutation, ned.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ned *NilExampleDelete) ExecX(ctx context.Context) int {
	n, err := ned.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ned *NilExampleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(nilexample.Table, sqlgraph.NewFieldSpec(nilexample.FieldID, field.TypeInt))
	if ps := ned.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ned.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ned.mutation.done = true
	return affected, err
}

// NilExampleDeleteOne is the builder for deleting a single NilExample entity.
type NilExampleDeleteOne struct {
	ned *NilExampleDelete
}

// Where appends a list predicates to the NilExampleDelete builder.
func (nedo *NilExampleDeleteOne) Where(ps ...predicate.NilExample) *NilExampleDeleteOne {
	nedo.ned.mutation.Where(ps...)
	return nedo
}

// Exec executes the deletion query.
func (nedo *NilExampleDeleteOne) Exec(ctx context.Context) error {
	n, err := nedo.ned.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{nilexample.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (nedo *NilExampleDeleteOne) ExecX(ctx context.Context) {
	if err := nedo.Exec(ctx); err != nil {
		panic(err)
	}
}
