// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/tanapon395/playlist-video/ent/department"
	"github.com/tanapon395/playlist-video/ent/gender"
	"github.com/tanapon395/playlist-video/ent/personal"
	"github.com/tanapon395/playlist-video/ent/title"
)

// Personal is the model entity for the Personal schema.
type Personal struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Personalname holds the value of the "Personalname" field.
	Personalname string `json:"Personalname,omitempty"`
	// Email holds the value of the "Email" field.
	Email string `json:"Email,omitempty"`
	// Password holds the value of the "Password" field.
	Password string `json:"Password,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PersonalQuery when eager-loading is set.
	Edges         PersonalEdges `json:"edges"`
	department_id *int
	gender_id     *int
	title_id      *int
}

// PersonalEdges holds the relations/edges for other nodes in the graph.
type PersonalEdges struct {
	// Customer holds the value of the customer edge.
	Customer []*Customer
	// Title holds the value of the title edge.
	Title *Title
	// Department holds the value of the department edge.
	Department *Department
	// Gender holds the value of the gender edge.
	Gender *Gender
	// Product holds the value of the product edge.
	Product []*Product
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading.
func (e PersonalEdges) CustomerOrErr() ([]*Customer, error) {
	if e.loadedTypes[0] {
		return e.Customer, nil
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// TitleOrErr returns the Title value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PersonalEdges) TitleOrErr() (*Title, error) {
	if e.loadedTypes[1] {
		if e.Title == nil {
			// The edge title was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: title.Label}
		}
		return e.Title, nil
	}
	return nil, &NotLoadedError{edge: "title"}
}

// DepartmentOrErr returns the Department value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PersonalEdges) DepartmentOrErr() (*Department, error) {
	if e.loadedTypes[2] {
		if e.Department == nil {
			// The edge department was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: department.Label}
		}
		return e.Department, nil
	}
	return nil, &NotLoadedError{edge: "department"}
}

// GenderOrErr returns the Gender value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PersonalEdges) GenderOrErr() (*Gender, error) {
	if e.loadedTypes[3] {
		if e.Gender == nil {
			// The edge gender was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: gender.Label}
		}
		return e.Gender, nil
	}
	return nil, &NotLoadedError{edge: "gender"}
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading.
func (e PersonalEdges) ProductOrErr() ([]*Product, error) {
	if e.loadedTypes[4] {
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "product"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Personal) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Personalname
		&sql.NullString{}, // Email
		&sql.NullString{}, // Password
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Personal) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // department_id
		&sql.NullInt64{}, // gender_id
		&sql.NullInt64{}, // title_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Personal fields.
func (pe *Personal) assignValues(values ...interface{}) error {
	if m, n := len(values), len(personal.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pe.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Personalname", values[0])
	} else if value.Valid {
		pe.Personalname = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Email", values[1])
	} else if value.Valid {
		pe.Email = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Password", values[2])
	} else if value.Valid {
		pe.Password = value.String
	}
	values = values[3:]
	if len(values) == len(personal.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field department_id", value)
		} else if value.Valid {
			pe.department_id = new(int)
			*pe.department_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field gender_id", value)
		} else if value.Valid {
			pe.gender_id = new(int)
			*pe.gender_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field title_id", value)
		} else if value.Valid {
			pe.title_id = new(int)
			*pe.title_id = int(value.Int64)
		}
	}
	return nil
}

// QueryCustomer queries the customer edge of the Personal.
func (pe *Personal) QueryCustomer() *CustomerQuery {
	return (&PersonalClient{config: pe.config}).QueryCustomer(pe)
}

// QueryTitle queries the title edge of the Personal.
func (pe *Personal) QueryTitle() *TitleQuery {
	return (&PersonalClient{config: pe.config}).QueryTitle(pe)
}

// QueryDepartment queries the department edge of the Personal.
func (pe *Personal) QueryDepartment() *DepartmentQuery {
	return (&PersonalClient{config: pe.config}).QueryDepartment(pe)
}

// QueryGender queries the gender edge of the Personal.
func (pe *Personal) QueryGender() *GenderQuery {
	return (&PersonalClient{config: pe.config}).QueryGender(pe)
}

// QueryProduct queries the product edge of the Personal.
func (pe *Personal) QueryProduct() *ProductQuery {
	return (&PersonalClient{config: pe.config}).QueryProduct(pe)
}

// Update returns a builder for updating this Personal.
// Note that, you need to call Personal.Unwrap() before calling this method, if this Personal
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Personal) Update() *PersonalUpdateOne {
	return (&PersonalClient{config: pe.config}).UpdateOne(pe)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pe *Personal) Unwrap() *Personal {
	tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Personal is not a transactional entity")
	}
	pe.config.driver = tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Personal) String() string {
	var builder strings.Builder
	builder.WriteString("Personal(")
	builder.WriteString(fmt.Sprintf("id=%v", pe.ID))
	builder.WriteString(", Personalname=")
	builder.WriteString(pe.Personalname)
	builder.WriteString(", Email=")
	builder.WriteString(pe.Email)
	builder.WriteString(", Password=")
	builder.WriteString(pe.Password)
	builder.WriteByte(')')
	return builder.String()
}

// Personals is a parsable slice of Personal.
type Personals []*Personal

func (pe Personals) config(cfg config) {
	for _i := range pe {
		pe[_i].config = cfg
	}
}
