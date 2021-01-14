// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/tanapon395/playlist-video/ent/brand"
)

// Brand is the model entity for the Brand schema.
type Brand struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Brandname holds the value of the "Brandname" field.
	Brandname string `json:"Brandname,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BrandQuery when eager-loading is set.
	Edges BrandEdges `json:"edges"`
}

// BrandEdges holds the relations/edges for other nodes in the graph.
type BrandEdges struct {
	// Product holds the value of the product edge.
	Product []*Product
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading.
func (e BrandEdges) ProductOrErr() ([]*Product, error) {
	if e.loadedTypes[0] {
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "product"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Brand) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Brandname
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Brand fields.
func (b *Brand) assignValues(values ...interface{}) error {
	if m, n := len(values), len(brand.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	b.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Brandname", values[0])
	} else if value.Valid {
		b.Brandname = value.String
	}
	return nil
}

// QueryProduct queries the product edge of the Brand.
func (b *Brand) QueryProduct() *ProductQuery {
	return (&BrandClient{config: b.config}).QueryProduct(b)
}

// Update returns a builder for updating this Brand.
// Note that, you need to call Brand.Unwrap() before calling this method, if this Brand
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Brand) Update() *BrandUpdateOne {
	return (&BrandClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (b *Brand) Unwrap() *Brand {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Brand is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Brand) String() string {
	var builder strings.Builder
	builder.WriteString("Brand(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", Brandname=")
	builder.WriteString(b.Brandname)
	builder.WriteByte(')')
	return builder.String()
}

// Brands is a parsable slice of Brand.
type Brands []*Brand

func (b Brands) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
