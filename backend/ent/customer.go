// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/tanapon395/playlist-video/ent/customer"
	"github.com/tanapon395/playlist-video/ent/gender"
	"github.com/tanapon395/playlist-video/ent/personal"
	"github.com/tanapon395/playlist-video/ent/title"
)

// Customer is the model entity for the Customer schema.
type Customer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Customername holds the value of the "Customername" field.
	Customername string `json:"Customername,omitempty"`
	// Address holds the value of the "Address" field.
	Address string `json:"Address,omitempty"`
	// Phonenumber holds the value of the "Phonenumber" field.
	Phonenumber string `json:"Phonenumber,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CustomerQuery when eager-loading is set.
	Edges         CustomerEdges `json:"edges"`
	department_id *int
	gender_id     *int
	personal_id   *int
	title_id      *int
}

// CustomerEdges holds the relations/edges for other nodes in the graph.
type CustomerEdges struct {
	// Gender holds the value of the gender edge.
	Gender *Gender
	// Personal holds the value of the personal edge.
	Personal *Personal
	// Title holds the value of the title edge.
	Title *Title
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// GenderOrErr returns the Gender value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CustomerEdges) GenderOrErr() (*Gender, error) {
	if e.loadedTypes[0] {
		if e.Gender == nil {
			// The edge gender was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: gender.Label}
		}
		return e.Gender, nil
	}
	return nil, &NotLoadedError{edge: "gender"}
}

// PersonalOrErr returns the Personal value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CustomerEdges) PersonalOrErr() (*Personal, error) {
	if e.loadedTypes[1] {
		if e.Personal == nil {
			// The edge personal was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: personal.Label}
		}
		return e.Personal, nil
	}
	return nil, &NotLoadedError{edge: "personal"}
}

// TitleOrErr returns the Title value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CustomerEdges) TitleOrErr() (*Title, error) {
	if e.loadedTypes[2] {
		if e.Title == nil {
			// The edge title was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: title.Label}
		}
		return e.Title, nil
	}
	return nil, &NotLoadedError{edge: "title"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Customer) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Customername
		&sql.NullString{}, // Address
		&sql.NullString{}, // Phonenumber
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Customer) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // department_id
		&sql.NullInt64{}, // gender_id
		&sql.NullInt64{}, // personal_id
		&sql.NullInt64{}, // title_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Customer fields.
func (c *Customer) assignValues(values ...interface{}) error {
	if m, n := len(values), len(customer.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Customername", values[0])
	} else if value.Valid {
		c.Customername = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Address", values[1])
	} else if value.Valid {
		c.Address = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Phonenumber", values[2])
	} else if value.Valid {
		c.Phonenumber = value.String
	}
	values = values[3:]
	if len(values) == len(customer.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field department_id", value)
		} else if value.Valid {
			c.department_id = new(int)
			*c.department_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field gender_id", value)
		} else if value.Valid {
			c.gender_id = new(int)
			*c.gender_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field personal_id", value)
		} else if value.Valid {
			c.personal_id = new(int)
			*c.personal_id = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field title_id", value)
		} else if value.Valid {
			c.title_id = new(int)
			*c.title_id = int(value.Int64)
		}
	}
	return nil
}

// QueryGender queries the gender edge of the Customer.
func (c *Customer) QueryGender() *GenderQuery {
	return (&CustomerClient{config: c.config}).QueryGender(c)
}

// QueryPersonal queries the personal edge of the Customer.
func (c *Customer) QueryPersonal() *PersonalQuery {
	return (&CustomerClient{config: c.config}).QueryPersonal(c)
}

// QueryTitle queries the title edge of the Customer.
func (c *Customer) QueryTitle() *TitleQuery {
	return (&CustomerClient{config: c.config}).QueryTitle(c)
}

// Update returns a builder for updating this Customer.
// Note that, you need to call Customer.Unwrap() before calling this method, if this Customer
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Customer) Update() *CustomerUpdateOne {
	return (&CustomerClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Customer) Unwrap() *Customer {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Customer is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Customer) String() string {
	var builder strings.Builder
	builder.WriteString("Customer(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", Customername=")
	builder.WriteString(c.Customername)
	builder.WriteString(", Address=")
	builder.WriteString(c.Address)
	builder.WriteString(", Phonenumber=")
	builder.WriteString(c.Phonenumber)
	builder.WriteByte(')')
	return builder.String()
}

// Customers is a parsable slice of Customer.
type Customers []*Customer

func (c Customers) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
