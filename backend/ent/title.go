// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/tanapon395/playlist-video/ent/title"
)

// Title is the model entity for the Title schema.
type Title struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Titlename holds the value of the "titlename" field.
	Titlename string `json:"titlename,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TitleQuery when eager-loading is set.
	Edges TitleEdges `json:"edges"`
}

// TitleEdges holds the relations/edges for other nodes in the graph.
type TitleEdges struct {
	// Personal holds the value of the personal edge.
	Personal []*Personal
	// Customer holds the value of the customer edge.
	Customer []*Customer
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PersonalOrErr returns the Personal value or an error if the edge
// was not loaded in eager-loading.
func (e TitleEdges) PersonalOrErr() ([]*Personal, error) {
	if e.loadedTypes[0] {
		return e.Personal, nil
	}
	return nil, &NotLoadedError{edge: "personal"}
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading.
func (e TitleEdges) CustomerOrErr() ([]*Customer, error) {
	if e.loadedTypes[1] {
		return e.Customer, nil
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Title) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // titlename
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Title fields.
func (t *Title) assignValues(values ...interface{}) error {
	if m, n := len(values), len(title.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	t.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field titlename", values[0])
	} else if value.Valid {
		t.Titlename = value.String
	}
	return nil
}

// QueryPersonal queries the personal edge of the Title.
func (t *Title) QueryPersonal() *PersonalQuery {
	return (&TitleClient{config: t.config}).QueryPersonal(t)
}

// QueryCustomer queries the customer edge of the Title.
func (t *Title) QueryCustomer() *CustomerQuery {
	return (&TitleClient{config: t.config}).QueryCustomer(t)
}

// Update returns a builder for updating this Title.
// Note that, you need to call Title.Unwrap() before calling this method, if this Title
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Title) Update() *TitleUpdateOne {
	return (&TitleClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (t *Title) Unwrap() *Title {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Title is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Title) String() string {
	var builder strings.Builder
	builder.WriteString("Title(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", titlename=")
	builder.WriteString(t.Titlename)
	builder.WriteByte(')')
	return builder.String()
}

// Titles is a parsable slice of Title.
type Titles []*Title

func (t Titles) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
