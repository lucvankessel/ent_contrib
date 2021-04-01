// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/contrib/entproto/internal/entprototest/ent/validmessage"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ValidMessageCreate is the builder for creating a ValidMessage entity.
type ValidMessageCreate struct {
	config
	mutation *ValidMessageMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (vmc *ValidMessageCreate) SetName(s string) *ValidMessageCreate {
	vmc.mutation.SetName(s)
	return vmc
}

// SetTs sets the "ts" field.
func (vmc *ValidMessageCreate) SetTs(t time.Time) *ValidMessageCreate {
	vmc.mutation.SetTs(t)
	return vmc
}

// SetUUID sets the "uuid" field.
func (vmc *ValidMessageCreate) SetUUID(u uuid.UUID) *ValidMessageCreate {
	vmc.mutation.SetUUID(u)
	return vmc
}

// Mutation returns the ValidMessageMutation object of the builder.
func (vmc *ValidMessageCreate) Mutation() *ValidMessageMutation {
	return vmc.mutation
}

// Save creates the ValidMessage in the database.
func (vmc *ValidMessageCreate) Save(ctx context.Context) (*ValidMessage, error) {
	var (
		err  error
		node *ValidMessage
	)
	if len(vmc.hooks) == 0 {
		if err = vmc.check(); err != nil {
			return nil, err
		}
		node, err = vmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ValidMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vmc.check(); err != nil {
				return nil, err
			}
			vmc.mutation = mutation
			node, err = vmc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vmc.hooks) - 1; i >= 0; i-- {
			mut = vmc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vmc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vmc *ValidMessageCreate) SaveX(ctx context.Context) *ValidMessage {
	v, err := vmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (vmc *ValidMessageCreate) check() error {
	if _, ok := vmc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := vmc.mutation.Ts(); !ok {
		return &ValidationError{Name: "ts", err: errors.New("ent: missing required field \"ts\"")}
	}
	if _, ok := vmc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New("ent: missing required field \"uuid\"")}
	}
	return nil
}

func (vmc *ValidMessageCreate) sqlSave(ctx context.Context) (*ValidMessage, error) {
	_node, _spec := vmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vmc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (vmc *ValidMessageCreate) createSpec() (*ValidMessage, *sqlgraph.CreateSpec) {
	var (
		_node = &ValidMessage{config: vmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: validmessage.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: validmessage.FieldID,
			},
		}
	)
	if value, ok := vmc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: validmessage.FieldName,
		})
		_node.Name = value
	}
	if value, ok := vmc.mutation.Ts(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: validmessage.FieldTs,
		})
		_node.Ts = value
	}
	if value, ok := vmc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: validmessage.FieldUUID,
		})
		_node.UUID = value
	}
	return _node, _spec
}

// ValidMessageCreateBulk is the builder for creating many ValidMessage entities in bulk.
type ValidMessageCreateBulk struct {
	config
	builders []*ValidMessageCreate
}

// Save creates the ValidMessage entities in the database.
func (vmcb *ValidMessageCreateBulk) Save(ctx context.Context) ([]*ValidMessage, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vmcb.builders))
	nodes := make([]*ValidMessage, len(vmcb.builders))
	mutators := make([]Mutator, len(vmcb.builders))
	for i := range vmcb.builders {
		func(i int, root context.Context) {
			builder := vmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ValidMessageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vmcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vmcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vmcb *ValidMessageCreateBulk) SaveX(ctx context.Context) []*ValidMessage {
	v, err := vmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}