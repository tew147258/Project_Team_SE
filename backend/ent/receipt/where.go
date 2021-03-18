// Code generated by entc, DO NOT EDIT.

package receipt

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/tanapon395/playlist-video/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DateTime applies equality check predicate on the "date_time" field. It's identical to DateTimeEQ.
func DateTime(v time.Time) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateTime), v))
	})
}

// Address applies equality check predicate on the "Address" field. It's identical to AddressEQ.
func Address(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// Productname applies equality check predicate on the "Productname" field. It's identical to ProductnameEQ.
func Productname(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductname), v))
	})
}

// Receiptcode applies equality check predicate on the "Receiptcode" field. It's identical to ReceiptcodeEQ.
func Receiptcode(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReceiptcode), v))
	})
}

// DateTimeEQ applies the EQ predicate on the "date_time" field.
func DateTimeEQ(v time.Time) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateTime), v))
	})
}

// DateTimeNEQ applies the NEQ predicate on the "date_time" field.
func DateTimeNEQ(v time.Time) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDateTime), v))
	})
}

// DateTimeIn applies the In predicate on the "date_time" field.
func DateTimeIn(vs ...time.Time) predicate.Receipt {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Receipt(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDateTime), v...))
	})
}

// DateTimeNotIn applies the NotIn predicate on the "date_time" field.
func DateTimeNotIn(vs ...time.Time) predicate.Receipt {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Receipt(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDateTime), v...))
	})
}

// DateTimeGT applies the GT predicate on the "date_time" field.
func DateTimeGT(v time.Time) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDateTime), v))
	})
}

// DateTimeGTE applies the GTE predicate on the "date_time" field.
func DateTimeGTE(v time.Time) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDateTime), v))
	})
}

// DateTimeLT applies the LT predicate on the "date_time" field.
func DateTimeLT(v time.Time) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDateTime), v))
	})
}

// DateTimeLTE applies the LTE predicate on the "date_time" field.
func DateTimeLTE(v time.Time) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDateTime), v))
	})
}

// AddressEQ applies the EQ predicate on the "Address" field.
func AddressEQ(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// AddressNEQ applies the NEQ predicate on the "Address" field.
func AddressNEQ(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddress), v))
	})
}

// AddressIn applies the In predicate on the "Address" field.
func AddressIn(vs ...string) predicate.Receipt {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Receipt(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAddress), v...))
	})
}

// AddressNotIn applies the NotIn predicate on the "Address" field.
func AddressNotIn(vs ...string) predicate.Receipt {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Receipt(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAddress), v...))
	})
}

// AddressGT applies the GT predicate on the "Address" field.
func AddressGT(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddress), v))
	})
}

// AddressGTE applies the GTE predicate on the "Address" field.
func AddressGTE(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddress), v))
	})
}

// AddressLT applies the LT predicate on the "Address" field.
func AddressLT(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddress), v))
	})
}

// AddressLTE applies the LTE predicate on the "Address" field.
func AddressLTE(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddress), v))
	})
}

// AddressContains applies the Contains predicate on the "Address" field.
func AddressContains(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAddress), v))
	})
}

// AddressHasPrefix applies the HasPrefix predicate on the "Address" field.
func AddressHasPrefix(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAddress), v))
	})
}

// AddressHasSuffix applies the HasSuffix predicate on the "Address" field.
func AddressHasSuffix(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAddress), v))
	})
}

// AddressEqualFold applies the EqualFold predicate on the "Address" field.
func AddressEqualFold(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAddress), v))
	})
}

// AddressContainsFold applies the ContainsFold predicate on the "Address" field.
func AddressContainsFold(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAddress), v))
	})
}

// ProductnameEQ applies the EQ predicate on the "Productname" field.
func ProductnameEQ(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductname), v))
	})
}

// ProductnameNEQ applies the NEQ predicate on the "Productname" field.
func ProductnameNEQ(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProductname), v))
	})
}

// ProductnameIn applies the In predicate on the "Productname" field.
func ProductnameIn(vs ...string) predicate.Receipt {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Receipt(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProductname), v...))
	})
}

// ProductnameNotIn applies the NotIn predicate on the "Productname" field.
func ProductnameNotIn(vs ...string) predicate.Receipt {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Receipt(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProductname), v...))
	})
}

// ProductnameGT applies the GT predicate on the "Productname" field.
func ProductnameGT(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProductname), v))
	})
}

// ProductnameGTE applies the GTE predicate on the "Productname" field.
func ProductnameGTE(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProductname), v))
	})
}

// ProductnameLT applies the LT predicate on the "Productname" field.
func ProductnameLT(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProductname), v))
	})
}

// ProductnameLTE applies the LTE predicate on the "Productname" field.
func ProductnameLTE(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProductname), v))
	})
}

// ProductnameContains applies the Contains predicate on the "Productname" field.
func ProductnameContains(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProductname), v))
	})
}

// ProductnameHasPrefix applies the HasPrefix predicate on the "Productname" field.
func ProductnameHasPrefix(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProductname), v))
	})
}

// ProductnameHasSuffix applies the HasSuffix predicate on the "Productname" field.
func ProductnameHasSuffix(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProductname), v))
	})
}

// ProductnameEqualFold applies the EqualFold predicate on the "Productname" field.
func ProductnameEqualFold(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProductname), v))
	})
}

// ProductnameContainsFold applies the ContainsFold predicate on the "Productname" field.
func ProductnameContainsFold(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProductname), v))
	})
}

// ReceiptcodeEQ applies the EQ predicate on the "Receiptcode" field.
func ReceiptcodeEQ(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReceiptcode), v))
	})
}

// ReceiptcodeNEQ applies the NEQ predicate on the "Receiptcode" field.
func ReceiptcodeNEQ(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReceiptcode), v))
	})
}

// ReceiptcodeIn applies the In predicate on the "Receiptcode" field.
func ReceiptcodeIn(vs ...string) predicate.Receipt {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Receipt(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldReceiptcode), v...))
	})
}

// ReceiptcodeNotIn applies the NotIn predicate on the "Receiptcode" field.
func ReceiptcodeNotIn(vs ...string) predicate.Receipt {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Receipt(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldReceiptcode), v...))
	})
}

// ReceiptcodeGT applies the GT predicate on the "Receiptcode" field.
func ReceiptcodeGT(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReceiptcode), v))
	})
}

// ReceiptcodeGTE applies the GTE predicate on the "Receiptcode" field.
func ReceiptcodeGTE(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReceiptcode), v))
	})
}

// ReceiptcodeLT applies the LT predicate on the "Receiptcode" field.
func ReceiptcodeLT(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReceiptcode), v))
	})
}

// ReceiptcodeLTE applies the LTE predicate on the "Receiptcode" field.
func ReceiptcodeLTE(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReceiptcode), v))
	})
}

// ReceiptcodeContains applies the Contains predicate on the "Receiptcode" field.
func ReceiptcodeContains(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldReceiptcode), v))
	})
}

// ReceiptcodeHasPrefix applies the HasPrefix predicate on the "Receiptcode" field.
func ReceiptcodeHasPrefix(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldReceiptcode), v))
	})
}

// ReceiptcodeHasSuffix applies the HasSuffix predicate on the "Receiptcode" field.
func ReceiptcodeHasSuffix(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldReceiptcode), v))
	})
}

// ReceiptcodeEqualFold applies the EqualFold predicate on the "Receiptcode" field.
func ReceiptcodeEqualFold(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldReceiptcode), v))
	})
}

// ReceiptcodeContainsFold applies the ContainsFold predicate on the "Receiptcode" field.
func ReceiptcodeContainsFold(v string) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldReceiptcode), v))
	})
}

// HasPaymenttype applies the HasEdge predicate on the "paymenttype" edge.
func HasPaymenttype() predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PaymenttypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PaymenttypeTable, PaymenttypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPaymenttypeWith applies the HasEdge predicate on the "paymenttype" edge with a given conditions (other predicates).
func HasPaymenttypeWith(preds ...predicate.PaymentType) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PaymenttypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PaymenttypeTable, PaymenttypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAdminrepair applies the HasEdge predicate on the "adminrepair" edge.
func HasAdminrepair() predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AdminrepairTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AdminrepairTable, AdminrepairColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAdminrepairWith applies the HasEdge predicate on the "adminrepair" edge with a given conditions (other predicates).
func HasAdminrepairWith(preds ...predicate.Adminrepair) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AdminrepairInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AdminrepairTable, AdminrepairColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPersonal applies the HasEdge predicate on the "personal" edge.
func HasPersonal() predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PersonalTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PersonalTable, PersonalColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPersonalWith applies the HasEdge predicate on the "personal" edge with a given conditions (other predicates).
func HasPersonalWith(preds ...predicate.Personal) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PersonalInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PersonalTable, PersonalColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCustomer applies the HasEdge predicate on the "customer" edge.
func HasCustomer() predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CustomerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CustomerTable, CustomerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCustomerWith applies the HasEdge predicate on the "customer" edge with a given conditions (other predicates).
func HasCustomerWith(preds ...predicate.Customer) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CustomerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CustomerTable, CustomerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProduct applies the HasEdge predicate on the "product" edge.
func HasProduct() predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProductTable, ProductColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductWith applies the HasEdge predicate on the "product" edge with a given conditions (other predicates).
func HasProductWith(preds ...predicate.Product) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProductTable, ProductColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Receipt) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Receipt) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Receipt) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		p(s.Not())
	})
}
