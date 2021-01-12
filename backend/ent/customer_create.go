// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/customer"
	"github.com/tanapon395/playlist-video/ent/fix"
	"github.com/tanapon395/playlist-video/ent/gender"
	"github.com/tanapon395/playlist-video/ent/personal"
	"github.com/tanapon395/playlist-video/ent/title"
)

// CustomerCreate is the builder for creating a Customer entity.
type CustomerCreate struct {
	config
	mutation *CustomerMutation
	hooks    []Hook
}

// SetCustomername sets the Customername field.
func (cc *CustomerCreate) SetCustomername(s string) *CustomerCreate {
	cc.mutation.SetCustomername(s)
	return cc
}

// SetAddress sets the Address field.
func (cc *CustomerCreate) SetAddress(s string) *CustomerCreate {
	cc.mutation.SetAddress(s)
	return cc
}

// SetPhonenumber sets the Phonenumber field.
func (cc *CustomerCreate) SetPhonenumber(s string) *CustomerCreate {
	cc.mutation.SetPhonenumber(s)
	return cc
}

// SetGenderID sets the gender edge to Gender by id.
func (cc *CustomerCreate) SetGenderID(id int) *CustomerCreate {
	cc.mutation.SetGenderID(id)
	return cc
}

// SetNillableGenderID sets the gender edge to Gender by id if the given value is not nil.
func (cc *CustomerCreate) SetNillableGenderID(id *int) *CustomerCreate {
	if id != nil {
		cc = cc.SetGenderID(*id)
	}
	return cc
}

// SetGender sets the gender edge to Gender.
func (cc *CustomerCreate) SetGender(g *Gender) *CustomerCreate {
	return cc.SetGenderID(g.ID)
}

// SetPersonalID sets the personal edge to Personal by id.
func (cc *CustomerCreate) SetPersonalID(id int) *CustomerCreate {
	cc.mutation.SetPersonalID(id)
	return cc
}

// SetNillablePersonalID sets the personal edge to Personal by id if the given value is not nil.
func (cc *CustomerCreate) SetNillablePersonalID(id *int) *CustomerCreate {
	if id != nil {
		cc = cc.SetPersonalID(*id)
	}
	return cc
}

// SetPersonal sets the personal edge to Personal.
func (cc *CustomerCreate) SetPersonal(p *Personal) *CustomerCreate {
	return cc.SetPersonalID(p.ID)
}

// SetTitleID sets the title edge to Title by id.
func (cc *CustomerCreate) SetTitleID(id int) *CustomerCreate {
	cc.mutation.SetTitleID(id)
	return cc
}

// SetNillableTitleID sets the title edge to Title by id if the given value is not nil.
func (cc *CustomerCreate) SetNillableTitleID(id *int) *CustomerCreate {
	if id != nil {
		cc = cc.SetTitleID(*id)
	}
	return cc
}

// SetTitle sets the title edge to Title.
func (cc *CustomerCreate) SetTitle(t *Title) *CustomerCreate {
	return cc.SetTitleID(t.ID)
}

// AddFixIDs adds the fix edge to Fix by ids.
func (cc *CustomerCreate) AddFixIDs(ids ...int) *CustomerCreate {
	cc.mutation.AddFixIDs(ids...)
	return cc
}

// AddFix adds the fix edges to Fix.
func (cc *CustomerCreate) AddFix(f ...*Fix) *CustomerCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return cc.AddFixIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cc *CustomerCreate) Mutation() *CustomerMutation {
	return cc.mutation
}

// Save creates the Customer in the database.
func (cc *CustomerCreate) Save(ctx context.Context) (*Customer, error) {
	if _, ok := cc.mutation.Customername(); !ok {
		return nil, &ValidationError{Name: "Customername", err: errors.New("ent: missing required field \"Customername\"")}
	}
	if _, ok := cc.mutation.Address(); !ok {
		return nil, &ValidationError{Name: "Address", err: errors.New("ent: missing required field \"Address\"")}
	}
	if _, ok := cc.mutation.Phonenumber(); !ok {
		return nil, &ValidationError{Name: "Phonenumber", err: errors.New("ent: missing required field \"Phonenumber\"")}
	}
	var (
		err  error
		node *Customer
	)
	if len(cc.hooks) == 0 {
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CustomerCreate) SaveX(ctx context.Context) *Customer {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *CustomerCreate) sqlSave(ctx context.Context) (*Customer, error) {
	c, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	c.ID = int(id)
	return c, nil
}

func (cc *CustomerCreate) createSpec() (*Customer, *sqlgraph.CreateSpec) {
	var (
		c     = &Customer{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: customer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: customer.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Customername(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldCustomername,
		})
		c.Customername = value
	}
	if value, ok := cc.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldAddress,
		})
		c.Address = value
	}
	if value, ok := cc.mutation.Phonenumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldPhonenumber,
		})
		c.Phonenumber = value
	}
	if nodes := cc.mutation.GenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.GenderTable,
			Columns: []string{customer.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.PersonalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.PersonalTable,
			Columns: []string{customer.PersonalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: personal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.TitleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.TitleTable,
			Columns: []string{customer.TitleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: title.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.FixIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.FixTable,
			Columns: []string{customer.FixColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fix.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return c, _spec
}
