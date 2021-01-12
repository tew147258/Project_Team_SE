// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/tanapon395/playlist-video/ent/receipt"
)

// Receipt is the model entity for the Receipt schema.
type Receipt struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Receipt) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Receipt fields.
func (r *Receipt) assignValues(values ...interface{}) error {
	if m, n := len(values), len(receipt.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	r.ID = int(value.Int64)
	values = values[1:]
	return nil
}

// Update returns a builder for updating this Receipt.
// Note that, you need to call Receipt.Unwrap() before calling this method, if this Receipt
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Receipt) Update() *ReceiptUpdateOne {
	return (&ReceiptClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (r *Receipt) Unwrap() *Receipt {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Receipt is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Receipt) String() string {
	var builder strings.Builder
	builder.WriteString("Receipt(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Receipts is a parsable slice of Receipt.
type Receipts []*Receipt

func (r Receipts) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}