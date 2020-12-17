// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/fly-nick/Go-000/Week04/internal/customer_follow/ent/customerfollow"
	"github.com/fly-nick/Go-000/Week04/internal/customer_follow/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// CustomerFollowUpdate is the builder for updating CustomerFollow entities.
type CustomerFollowUpdate struct {
	config
	hooks    []Hook
	mutation *CustomerFollowMutation
}

// Where adds a new predicate for the builder.
func (cfu *CustomerFollowUpdate) Where(ps ...predicate.CustomerFollow) *CustomerFollowUpdate {
	cfu.mutation.predicates = append(cfu.mutation.predicates, ps...)
	return cfu
}

// SetStaffId sets the staffId field.
func (cfu *CustomerFollowUpdate) SetStaffId(i int64) *CustomerFollowUpdate {
	cfu.mutation.ResetStaffId()
	cfu.mutation.SetStaffId(i)
	return cfu
}

// AddStaffId adds i to staffId.
func (cfu *CustomerFollowUpdate) AddStaffId(i int64) *CustomerFollowUpdate {
	cfu.mutation.AddStaffId(i)
	return cfu
}

// SetCustomerId sets the customerId field.
func (cfu *CustomerFollowUpdate) SetCustomerId(i int64) *CustomerFollowUpdate {
	cfu.mutation.ResetCustomerId()
	cfu.mutation.SetCustomerId(i)
	return cfu
}

// AddCustomerId adds i to customerId.
func (cfu *CustomerFollowUpdate) AddCustomerId(i int64) *CustomerFollowUpdate {
	cfu.mutation.AddCustomerId(i)
	return cfu
}

// SetContent sets the content field.
func (cfu *CustomerFollowUpdate) SetContent(s string) *CustomerFollowUpdate {
	cfu.mutation.SetContent(s)
	return cfu
}

// SetCreateTime sets the createTime field.
func (cfu *CustomerFollowUpdate) SetCreateTime(t time.Time) *CustomerFollowUpdate {
	cfu.mutation.SetCreateTime(t)
	return cfu
}

// SetDeleted sets the deleted field.
func (cfu *CustomerFollowUpdate) SetDeleted(i int64) *CustomerFollowUpdate {
	cfu.mutation.ResetDeleted()
	cfu.mutation.SetDeleted(i)
	return cfu
}

// AddDeleted adds i to deleted.
func (cfu *CustomerFollowUpdate) AddDeleted(i int64) *CustomerFollowUpdate {
	cfu.mutation.AddDeleted(i)
	return cfu
}

// SetDeleteBy sets the deleteBy field.
func (cfu *CustomerFollowUpdate) SetDeleteBy(s string) *CustomerFollowUpdate {
	cfu.mutation.SetDeleteBy(s)
	return cfu
}

// SetNillableDeleteBy sets the deleteBy field if the given value is not nil.
func (cfu *CustomerFollowUpdate) SetNillableDeleteBy(s *string) *CustomerFollowUpdate {
	if s != nil {
		cfu.SetDeleteBy(*s)
	}
	return cfu
}

// ClearDeleteBy clears the value of deleteBy.
func (cfu *CustomerFollowUpdate) ClearDeleteBy() *CustomerFollowUpdate {
	cfu.mutation.ClearDeleteBy()
	return cfu
}

// Mutation returns the CustomerFollowMutation object of the builder.
func (cfu *CustomerFollowUpdate) Mutation() *CustomerFollowMutation {
	return cfu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cfu *CustomerFollowUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cfu.hooks) == 0 {
		affected, err = cfu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerFollowMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cfu.mutation = mutation
			affected, err = cfu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cfu.hooks) - 1; i >= 0; i-- {
			mut = cfu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cfu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cfu *CustomerFollowUpdate) SaveX(ctx context.Context) int {
	affected, err := cfu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cfu *CustomerFollowUpdate) Exec(ctx context.Context) error {
	_, err := cfu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cfu *CustomerFollowUpdate) ExecX(ctx context.Context) {
	if err := cfu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cfu *CustomerFollowUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customerfollow.Table,
			Columns: customerfollow.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: customerfollow.FieldID,
			},
		},
	}
	if ps := cfu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cfu.mutation.StaffId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customerfollow.FieldStaffId,
		})
	}
	if value, ok := cfu.mutation.AddedStaffId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customerfollow.FieldStaffId,
		})
	}
	if value, ok := cfu.mutation.CustomerId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customerfollow.FieldCustomerId,
		})
	}
	if value, ok := cfu.mutation.AddedCustomerId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customerfollow.FieldCustomerId,
		})
	}
	if value, ok := cfu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customerfollow.FieldContent,
		})
	}
	if value, ok := cfu.mutation.CreateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customerfollow.FieldCreateTime,
		})
	}
	if value, ok := cfu.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customerfollow.FieldDeleted,
		})
	}
	if value, ok := cfu.mutation.AddedDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customerfollow.FieldDeleted,
		})
	}
	if value, ok := cfu.mutation.DeleteBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customerfollow.FieldDeleteBy,
		})
	}
	if cfu.mutation.DeleteByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customerfollow.FieldDeleteBy,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cfu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customerfollow.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CustomerFollowUpdateOne is the builder for updating a single CustomerFollow entity.
type CustomerFollowUpdateOne struct {
	config
	hooks    []Hook
	mutation *CustomerFollowMutation
}

// SetStaffId sets the staffId field.
func (cfuo *CustomerFollowUpdateOne) SetStaffId(i int64) *CustomerFollowUpdateOne {
	cfuo.mutation.ResetStaffId()
	cfuo.mutation.SetStaffId(i)
	return cfuo
}

// AddStaffId adds i to staffId.
func (cfuo *CustomerFollowUpdateOne) AddStaffId(i int64) *CustomerFollowUpdateOne {
	cfuo.mutation.AddStaffId(i)
	return cfuo
}

// SetCustomerId sets the customerId field.
func (cfuo *CustomerFollowUpdateOne) SetCustomerId(i int64) *CustomerFollowUpdateOne {
	cfuo.mutation.ResetCustomerId()
	cfuo.mutation.SetCustomerId(i)
	return cfuo
}

// AddCustomerId adds i to customerId.
func (cfuo *CustomerFollowUpdateOne) AddCustomerId(i int64) *CustomerFollowUpdateOne {
	cfuo.mutation.AddCustomerId(i)
	return cfuo
}

// SetContent sets the content field.
func (cfuo *CustomerFollowUpdateOne) SetContent(s string) *CustomerFollowUpdateOne {
	cfuo.mutation.SetContent(s)
	return cfuo
}

// SetCreateTime sets the createTime field.
func (cfuo *CustomerFollowUpdateOne) SetCreateTime(t time.Time) *CustomerFollowUpdateOne {
	cfuo.mutation.SetCreateTime(t)
	return cfuo
}

// SetDeleted sets the deleted field.
func (cfuo *CustomerFollowUpdateOne) SetDeleted(i int64) *CustomerFollowUpdateOne {
	cfuo.mutation.ResetDeleted()
	cfuo.mutation.SetDeleted(i)
	return cfuo
}

// AddDeleted adds i to deleted.
func (cfuo *CustomerFollowUpdateOne) AddDeleted(i int64) *CustomerFollowUpdateOne {
	cfuo.mutation.AddDeleted(i)
	return cfuo
}

// SetDeleteBy sets the deleteBy field.
func (cfuo *CustomerFollowUpdateOne) SetDeleteBy(s string) *CustomerFollowUpdateOne {
	cfuo.mutation.SetDeleteBy(s)
	return cfuo
}

// SetNillableDeleteBy sets the deleteBy field if the given value is not nil.
func (cfuo *CustomerFollowUpdateOne) SetNillableDeleteBy(s *string) *CustomerFollowUpdateOne {
	if s != nil {
		cfuo.SetDeleteBy(*s)
	}
	return cfuo
}

// ClearDeleteBy clears the value of deleteBy.
func (cfuo *CustomerFollowUpdateOne) ClearDeleteBy() *CustomerFollowUpdateOne {
	cfuo.mutation.ClearDeleteBy()
	return cfuo
}

// Mutation returns the CustomerFollowMutation object of the builder.
func (cfuo *CustomerFollowUpdateOne) Mutation() *CustomerFollowMutation {
	return cfuo.mutation
}

// Save executes the query and returns the updated entity.
func (cfuo *CustomerFollowUpdateOne) Save(ctx context.Context) (*CustomerFollow, error) {
	var (
		err  error
		node *CustomerFollow
	)
	if len(cfuo.hooks) == 0 {
		node, err = cfuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerFollowMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cfuo.mutation = mutation
			node, err = cfuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cfuo.hooks) - 1; i >= 0; i-- {
			mut = cfuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cfuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cfuo *CustomerFollowUpdateOne) SaveX(ctx context.Context) *CustomerFollow {
	node, err := cfuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cfuo *CustomerFollowUpdateOne) Exec(ctx context.Context) error {
	_, err := cfuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cfuo *CustomerFollowUpdateOne) ExecX(ctx context.Context) {
	if err := cfuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cfuo *CustomerFollowUpdateOne) sqlSave(ctx context.Context) (_node *CustomerFollow, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customerfollow.Table,
			Columns: customerfollow.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: customerfollow.FieldID,
			},
		},
	}
	id, ok := cfuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing CustomerFollow.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := cfuo.mutation.StaffId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customerfollow.FieldStaffId,
		})
	}
	if value, ok := cfuo.mutation.AddedStaffId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customerfollow.FieldStaffId,
		})
	}
	if value, ok := cfuo.mutation.CustomerId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customerfollow.FieldCustomerId,
		})
	}
	if value, ok := cfuo.mutation.AddedCustomerId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customerfollow.FieldCustomerId,
		})
	}
	if value, ok := cfuo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customerfollow.FieldContent,
		})
	}
	if value, ok := cfuo.mutation.CreateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customerfollow.FieldCreateTime,
		})
	}
	if value, ok := cfuo.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customerfollow.FieldDeleted,
		})
	}
	if value, ok := cfuo.mutation.AddedDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customerfollow.FieldDeleted,
		})
	}
	if value, ok := cfuo.mutation.DeleteBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customerfollow.FieldDeleteBy,
		})
	}
	if cfuo.mutation.DeleteByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customerfollow.FieldDeleteBy,
		})
	}
	_node = &CustomerFollow{config: cfuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cfuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customerfollow.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
