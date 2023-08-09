// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/messagewithid"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageWithIDUpdate is the builder for updating MessageWithID entities.
type MessageWithIDUpdate struct {
	config
	hooks    []Hook
	mutation *MessageWithIDMutation
}

// Where appends a list predicates to the MessageWithIDUpdate builder.
func (mwiu *MessageWithIDUpdate) Where(ps ...predicate.MessageWithID) *MessageWithIDUpdate {
	mwiu.mutation.Where(ps...)
	return mwiu
}

// Mutation returns the MessageWithIDMutation object of the builder.
func (mwiu *MessageWithIDUpdate) Mutation() *MessageWithIDMutation {
	return mwiu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mwiu *MessageWithIDUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, MessageWithIDMutation](ctx, mwiu.sqlSave, mwiu.mutation, mwiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mwiu *MessageWithIDUpdate) SaveX(ctx context.Context) int {
	affected, err := mwiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mwiu *MessageWithIDUpdate) Exec(ctx context.Context) error {
	_, err := mwiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwiu *MessageWithIDUpdate) ExecX(ctx context.Context) {
	if err := mwiu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mwiu *MessageWithIDUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(messagewithid.Table, messagewithid.Columns, sqlgraph.NewFieldSpec(messagewithid.FieldID, field.TypeInt32))
	if ps := mwiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mwiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{messagewithid.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mwiu.mutation.done = true
	return n, nil
}

// MessageWithIDUpdateOne is the builder for updating a single MessageWithID entity.
type MessageWithIDUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageWithIDMutation
}

// Mutation returns the MessageWithIDMutation object of the builder.
func (mwiuo *MessageWithIDUpdateOne) Mutation() *MessageWithIDMutation {
	return mwiuo.mutation
}

// Where appends a list predicates to the MessageWithIDUpdate builder.
func (mwiuo *MessageWithIDUpdateOne) Where(ps ...predicate.MessageWithID) *MessageWithIDUpdateOne {
	mwiuo.mutation.Where(ps...)
	return mwiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mwiuo *MessageWithIDUpdateOne) Select(field string, fields ...string) *MessageWithIDUpdateOne {
	mwiuo.fields = append([]string{field}, fields...)
	return mwiuo
}

// Save executes the query and returns the updated MessageWithID entity.
func (mwiuo *MessageWithIDUpdateOne) Save(ctx context.Context) (*MessageWithID, error) {
	return withHooks[*MessageWithID, MessageWithIDMutation](ctx, mwiuo.sqlSave, mwiuo.mutation, mwiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mwiuo *MessageWithIDUpdateOne) SaveX(ctx context.Context) *MessageWithID {
	node, err := mwiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mwiuo *MessageWithIDUpdateOne) Exec(ctx context.Context) error {
	_, err := mwiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwiuo *MessageWithIDUpdateOne) ExecX(ctx context.Context) {
	if err := mwiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mwiuo *MessageWithIDUpdateOne) sqlSave(ctx context.Context) (_node *MessageWithID, err error) {
	_spec := sqlgraph.NewUpdateSpec(messagewithid.Table, messagewithid.Columns, sqlgraph.NewFieldSpec(messagewithid.FieldID, field.TypeInt32))
	id, ok := mwiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MessageWithID.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mwiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, messagewithid.FieldID)
		for _, f := range fields {
			if !messagewithid.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != messagewithid.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mwiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &MessageWithID{config: mwiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mwiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{messagewithid.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mwiuo.mutation.done = true
	return _node, nil
}
