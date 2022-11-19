// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"myapp/ent/predicate"
	"myapp/ent/system_parameter"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SystemParameterDelete is the builder for deleting a System_parameter entity.
type SystemParameterDelete struct {
	config
	hooks    []Hook
	mutation *SystemParameterMutation
}

// Where appends a list predicates to the SystemParameterDelete builder.
func (spd *SystemParameterDelete) Where(ps ...predicate.System_parameter) *SystemParameterDelete {
	spd.mutation.Where(ps...)
	return spd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (spd *SystemParameterDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(spd.hooks) == 0 {
		affected, err = spd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SystemParameterMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			spd.mutation = mutation
			affected, err = spd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(spd.hooks) - 1; i >= 0; i-- {
			if spd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = spd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, spd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (spd *SystemParameterDelete) ExecX(ctx context.Context) int {
	n, err := spd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (spd *SystemParameterDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: system_parameter.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: system_parameter.FieldID,
			},
		},
	}
	if ps := spd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, spd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// SystemParameterDeleteOne is the builder for deleting a single System_parameter entity.
type SystemParameterDeleteOne struct {
	spd *SystemParameterDelete
}

// Exec executes the deletion query.
func (spdo *SystemParameterDeleteOne) Exec(ctx context.Context) error {
	n, err := spdo.spd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{system_parameter.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (spdo *SystemParameterDeleteOne) ExecX(ctx context.Context) {
	spdo.spd.ExecX(ctx)
}
