// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/zibbp/ganymede/ent/livecategory"
	"github.com/zibbp/ganymede/ent/predicate"
)

// LiveCategoryDelete is the builder for deleting a LiveCategory entity.
type LiveCategoryDelete struct {
	config
	hooks    []Hook
	mutation *LiveCategoryMutation
}

// Where appends a list predicates to the LiveCategoryDelete builder.
func (lcd *LiveCategoryDelete) Where(ps ...predicate.LiveCategory) *LiveCategoryDelete {
	lcd.mutation.Where(ps...)
	return lcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lcd *LiveCategoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, LiveCategoryMutation](ctx, lcd.sqlExec, lcd.mutation, lcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (lcd *LiveCategoryDelete) ExecX(ctx context.Context) int {
	n, err := lcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lcd *LiveCategoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(livecategory.Table, sqlgraph.NewFieldSpec(livecategory.FieldID, field.TypeUUID))
	if ps := lcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, lcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	lcd.mutation.done = true
	return affected, err
}

// LiveCategoryDeleteOne is the builder for deleting a single LiveCategory entity.
type LiveCategoryDeleteOne struct {
	lcd *LiveCategoryDelete
}

// Where appends a list predicates to the LiveCategoryDelete builder.
func (lcdo *LiveCategoryDeleteOne) Where(ps ...predicate.LiveCategory) *LiveCategoryDeleteOne {
	lcdo.lcd.mutation.Where(ps...)
	return lcdo
}

// Exec executes the deletion query.
func (lcdo *LiveCategoryDeleteOne) Exec(ctx context.Context) error {
	n, err := lcdo.lcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{livecategory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lcdo *LiveCategoryDeleteOne) ExecX(ctx context.Context) {
	if err := lcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
