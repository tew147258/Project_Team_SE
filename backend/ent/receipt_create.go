// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/adminrepair"
	"github.com/tanapon395/playlist-video/ent/customer"
	"github.com/tanapon395/playlist-video/ent/paymenttype"
	"github.com/tanapon395/playlist-video/ent/personal"
	"github.com/tanapon395/playlist-video/ent/product"
	"github.com/tanapon395/playlist-video/ent/receipt"
)

// ReceiptCreate is the builder for creating a Receipt entity.
type ReceiptCreate struct {
	config
	mutation *ReceiptMutation
	hooks    []Hook
}

// SetDateTime sets the date_time field.
func (rc *ReceiptCreate) SetDateTime(t time.Time) *ReceiptCreate {
	rc.mutation.SetDateTime(t)
	return rc
}

// SetNillableDateTime sets the date_time field if the given value is not nil.
func (rc *ReceiptCreate) SetNillableDateTime(t *time.Time) *ReceiptCreate {
	if t != nil {
		rc.SetDateTime(*t)
	}
	return rc
}

// SetAddress sets the Address field.
func (rc *ReceiptCreate) SetAddress(s string) *ReceiptCreate {
	rc.mutation.SetAddress(s)
	return rc
}

// SetProductname sets the Productname field.
func (rc *ReceiptCreate) SetProductname(s string) *ReceiptCreate {
	rc.mutation.SetProductname(s)
	return rc
}

// SetReceiptcode sets the Receiptcode field.
func (rc *ReceiptCreate) SetReceiptcode(s string) *ReceiptCreate {
	rc.mutation.SetReceiptcode(s)
	return rc
}

// SetPaymenttypeID sets the paymenttype edge to PaymentType by id.
func (rc *ReceiptCreate) SetPaymenttypeID(id int) *ReceiptCreate {
	rc.mutation.SetPaymenttypeID(id)
	return rc
}

// SetNillablePaymenttypeID sets the paymenttype edge to PaymentType by id if the given value is not nil.
func (rc *ReceiptCreate) SetNillablePaymenttypeID(id *int) *ReceiptCreate {
	if id != nil {
		rc = rc.SetPaymenttypeID(*id)
	}
	return rc
}

// SetPaymenttype sets the paymenttype edge to PaymentType.
func (rc *ReceiptCreate) SetPaymenttype(p *PaymentType) *ReceiptCreate {
	return rc.SetPaymenttypeID(p.ID)
}

// SetAdminrepairID sets the adminrepair edge to Adminrepair by id.
func (rc *ReceiptCreate) SetAdminrepairID(id int) *ReceiptCreate {
	rc.mutation.SetAdminrepairID(id)
	return rc
}

// SetNillableAdminrepairID sets the adminrepair edge to Adminrepair by id if the given value is not nil.
func (rc *ReceiptCreate) SetNillableAdminrepairID(id *int) *ReceiptCreate {
	if id != nil {
		rc = rc.SetAdminrepairID(*id)
	}
	return rc
}

// SetAdminrepair sets the adminrepair edge to Adminrepair.
func (rc *ReceiptCreate) SetAdminrepair(a *Adminrepair) *ReceiptCreate {
	return rc.SetAdminrepairID(a.ID)
}

// SetPersonalID sets the personal edge to Personal by id.
func (rc *ReceiptCreate) SetPersonalID(id int) *ReceiptCreate {
	rc.mutation.SetPersonalID(id)
	return rc
}

// SetNillablePersonalID sets the personal edge to Personal by id if the given value is not nil.
func (rc *ReceiptCreate) SetNillablePersonalID(id *int) *ReceiptCreate {
	if id != nil {
		rc = rc.SetPersonalID(*id)
	}
	return rc
}

// SetPersonal sets the personal edge to Personal.
func (rc *ReceiptCreate) SetPersonal(p *Personal) *ReceiptCreate {
	return rc.SetPersonalID(p.ID)
}

// SetCustomerID sets the customer edge to Customer by id.
func (rc *ReceiptCreate) SetCustomerID(id int) *ReceiptCreate {
	rc.mutation.SetCustomerID(id)
	return rc
}

// SetNillableCustomerID sets the customer edge to Customer by id if the given value is not nil.
func (rc *ReceiptCreate) SetNillableCustomerID(id *int) *ReceiptCreate {
	if id != nil {
		rc = rc.SetCustomerID(*id)
	}
	return rc
}

// SetCustomer sets the customer edge to Customer.
func (rc *ReceiptCreate) SetCustomer(c *Customer) *ReceiptCreate {
	return rc.SetCustomerID(c.ID)
}

// SetProductID sets the product edge to Product by id.
func (rc *ReceiptCreate) SetProductID(id int) *ReceiptCreate {
	rc.mutation.SetProductID(id)
	return rc
}

// SetNillableProductID sets the product edge to Product by id if the given value is not nil.
func (rc *ReceiptCreate) SetNillableProductID(id *int) *ReceiptCreate {
	if id != nil {
		rc = rc.SetProductID(*id)
	}
	return rc
}

// SetProduct sets the product edge to Product.
func (rc *ReceiptCreate) SetProduct(p *Product) *ReceiptCreate {
	return rc.SetProductID(p.ID)
}

// Mutation returns the ReceiptMutation object of the builder.
func (rc *ReceiptCreate) Mutation() *ReceiptMutation {
	return rc.mutation
}

// Save creates the Receipt in the database.
func (rc *ReceiptCreate) Save(ctx context.Context) (*Receipt, error) {
	if _, ok := rc.mutation.DateTime(); !ok {
		v := receipt.DefaultDateTime()
		rc.mutation.SetDateTime(v)
	}
	if _, ok := rc.mutation.Address(); !ok {
		return nil, &ValidationError{Name: "Address", err: errors.New("ent: missing required field \"Address\"")}
	}
	if v, ok := rc.mutation.Address(); ok {
		if err := receipt.AddressValidator(v); err != nil {
			return nil, &ValidationError{Name: "Address", err: fmt.Errorf("ent: validator failed for field \"Address\": %w", err)}
		}
	}
	if _, ok := rc.mutation.Productname(); !ok {
		return nil, &ValidationError{Name: "Productname", err: errors.New("ent: missing required field \"Productname\"")}
	}
	if v, ok := rc.mutation.Productname(); ok {
		if err := receipt.ProductnameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Productname", err: fmt.Errorf("ent: validator failed for field \"Productname\": %w", err)}
		}
	}
	if _, ok := rc.mutation.Receiptcode(); !ok {
		return nil, &ValidationError{Name: "Receiptcode", err: errors.New("ent: missing required field \"Receiptcode\"")}
	}
	if v, ok := rc.mutation.Receiptcode(); ok {
		if err := receipt.ReceiptcodeValidator(v); err != nil {
			return nil, &ValidationError{Name: "Receiptcode", err: fmt.Errorf("ent: validator failed for field \"Receiptcode\": %w", err)}
		}
	}
	var (
		err  error
		node *Receipt
	)
	if len(rc.hooks) == 0 {
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReceiptMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rc.mutation = mutation
			node, err = rc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ReceiptCreate) SaveX(ctx context.Context) *Receipt {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rc *ReceiptCreate) sqlSave(ctx context.Context) (*Receipt, error) {
	r, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	r.ID = int(id)
	return r, nil
}

func (rc *ReceiptCreate) createSpec() (*Receipt, *sqlgraph.CreateSpec) {
	var (
		r     = &Receipt{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: receipt.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: receipt.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.DateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: receipt.FieldDateTime,
		})
		r.DateTime = value
	}
	if value, ok := rc.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: receipt.FieldAddress,
		})
		r.Address = value
	}
	if value, ok := rc.mutation.Productname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: receipt.FieldProductname,
		})
		r.Productname = value
	}
	if value, ok := rc.mutation.Receiptcode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: receipt.FieldReceiptcode,
		})
		r.Receiptcode = value
	}
	if nodes := rc.mutation.PaymenttypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   receipt.PaymenttypeTable,
			Columns: []string{receipt.PaymenttypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: paymenttype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.AdminrepairIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   receipt.AdminrepairTable,
			Columns: []string{receipt.AdminrepairColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adminrepair.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.PersonalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   receipt.PersonalTable,
			Columns: []string{receipt.PersonalColumn},
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
	if nodes := rc.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   receipt.CustomerTable,
			Columns: []string{receipt.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   receipt.ProductTable,
			Columns: []string{receipt.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return r, _spec
}
