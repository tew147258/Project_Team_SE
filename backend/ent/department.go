// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/tanapon395/playlist-video/ent/department"
)

// Department is the model entity for the Department schema.
type Department struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Departmentname holds the value of the "Departmentname" field.
	Departmentname string `json:"Departmentname,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DepartmentQuery when eager-loading is set.
	Edges DepartmentEdges `json:"edges"`
}

// DepartmentEdges holds the relations/edges for other nodes in the graph.
type DepartmentEdges struct {
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
func (e DepartmentEdges) PersonalOrErr() ([]*Personal, error) {
	if e.loadedTypes[0] {
		return e.Personal, nil
	}
	return nil, &NotLoadedError{edge: "personal"}
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading.
func (e DepartmentEdges) CustomerOrErr() ([]*Customer, error) {
	if e.loadedTypes[1] {
		return e.Customer, nil
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Department) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Departmentname
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Department fields.
func (d *Department) assignValues(values ...interface{}) error {
	if m, n := len(values), len(department.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	d.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Departmentname", values[0])
	} else if value.Valid {
		d.Departmentname = value.String
	}
	return nil
}

// QueryPersonal queries the personal edge of the Department.
func (d *Department) QueryPersonal() *PersonalQuery {
	return (&DepartmentClient{config: d.config}).QueryPersonal(d)
}

// QueryCustomer queries the customer edge of the Department.
func (d *Department) QueryCustomer() *CustomerQuery {
	return (&DepartmentClient{config: d.config}).QueryCustomer(d)
}

// Update returns a builder for updating this Department.
// Note that, you need to call Department.Unwrap() before calling this method, if this Department
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Department) Update() *DepartmentUpdateOne {
	return (&DepartmentClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (d *Department) Unwrap() *Department {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Department is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Department) String() string {
	var builder strings.Builder
	builder.WriteString("Department(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", Departmentname=")
	builder.WriteString(d.Departmentname)
	builder.WriteByte(')')
	return builder.String()
}

// Departments is a parsable slice of Department.
type Departments []*Department

func (d Departments) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}