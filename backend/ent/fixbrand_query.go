// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/fix"
	"github.com/tanapon395/playlist-video/ent/fixbrand"
	"github.com/tanapon395/playlist-video/ent/predicate"
)

// FixbrandQuery is the builder for querying Fixbrand entities.
type FixbrandQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Fixbrand
	// eager-loading edges.
	withFix *FixQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (fq *FixbrandQuery) Where(ps ...predicate.Fixbrand) *FixbrandQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit adds a limit step to the query.
func (fq *FixbrandQuery) Limit(limit int) *FixbrandQuery {
	fq.limit = &limit
	return fq
}

// Offset adds an offset step to the query.
func (fq *FixbrandQuery) Offset(offset int) *FixbrandQuery {
	fq.offset = &offset
	return fq
}

// Order adds an order step to the query.
func (fq *FixbrandQuery) Order(o ...OrderFunc) *FixbrandQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// QueryFix chains the current query on the fix edge.
func (fq *FixbrandQuery) QueryFix() *FixQuery {
	query := &FixQuery{config: fq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(fixbrand.Table, fixbrand.FieldID, fq.sqlQuery()),
			sqlgraph.To(fix.Table, fix.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, fixbrand.FixTable, fixbrand.FixColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Fixbrand entity in the query. Returns *NotFoundError when no fixbrand was found.
func (fq *FixbrandQuery) First(ctx context.Context) (*Fixbrand, error) {
	fs, err := fq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(fs) == 0 {
		return nil, &NotFoundError{fixbrand.Label}
	}
	return fs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FixbrandQuery) FirstX(ctx context.Context) *Fixbrand {
	f, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return f
}

// FirstID returns the first Fixbrand id in the query. Returns *NotFoundError when no id was found.
func (fq *FixbrandQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{fixbrand.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (fq *FixbrandQuery) FirstXID(ctx context.Context) int {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Fixbrand entity in the query, returns an error if not exactly one entity was returned.
func (fq *FixbrandQuery) Only(ctx context.Context) (*Fixbrand, error) {
	fs, err := fq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(fs) {
	case 1:
		return fs[0], nil
	case 0:
		return nil, &NotFoundError{fixbrand.Label}
	default:
		return nil, &NotSingularError{fixbrand.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FixbrandQuery) OnlyX(ctx context.Context) *Fixbrand {
	f, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return f
}

// OnlyID returns the only Fixbrand id in the query, returns an error if not exactly one id was returned.
func (fq *FixbrandQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{fixbrand.Label}
	default:
		err = &NotSingularError{fixbrand.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FixbrandQuery) OnlyIDX(ctx context.Context) int {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Fixbrands.
func (fq *FixbrandQuery) All(ctx context.Context) ([]*Fixbrand, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fq *FixbrandQuery) AllX(ctx context.Context) []*Fixbrand {
	fs, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return fs
}

// IDs executes the query and returns a list of Fixbrand ids.
func (fq *FixbrandQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := fq.Select(fixbrand.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FixbrandQuery) IDsX(ctx context.Context) []int {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FixbrandQuery) Count(ctx context.Context) (int, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FixbrandQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FixbrandQuery) Exist(ctx context.Context) (bool, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FixbrandQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FixbrandQuery) Clone() *FixbrandQuery {
	return &FixbrandQuery{
		config:     fq.config,
		limit:      fq.limit,
		offset:     fq.offset,
		order:      append([]OrderFunc{}, fq.order...),
		unique:     append([]string{}, fq.unique...),
		predicates: append([]predicate.Fixbrand{}, fq.predicates...),
		// clone intermediate query.
		sql:  fq.sql.Clone(),
		path: fq.path,
	}
}

//  WithFix tells the query-builder to eager-loads the nodes that are connected to
// the "fix" edge. The optional arguments used to configure the query builder of the edge.
func (fq *FixbrandQuery) WithFix(opts ...func(*FixQuery)) *FixbrandQuery {
	query := &FixQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	fq.withFix = query
	return fq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Fixbrandname string `json:"fixbrandname,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Fixbrand.Query().
//		GroupBy(fixbrand.FieldFixbrandname).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fq *FixbrandQuery) GroupBy(field string, fields ...string) *FixbrandGroupBy {
	group := &FixbrandGroupBy{config: fq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Fixbrandname string `json:"fixbrandname,omitempty"`
//	}
//
//	client.Fixbrand.Query().
//		Select(fixbrand.FieldFixbrandname).
//		Scan(ctx, &v)
//
func (fq *FixbrandQuery) Select(field string, fields ...string) *FixbrandSelect {
	selector := &FixbrandSelect{config: fq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fq.sqlQuery(), nil
	}
	return selector
}

func (fq *FixbrandQuery) prepareQuery(ctx context.Context) error {
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FixbrandQuery) sqlAll(ctx context.Context) ([]*Fixbrand, error) {
	var (
		nodes       = []*Fixbrand{}
		_spec       = fq.querySpec()
		loadedTypes = [1]bool{
			fq.withFix != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &Fixbrand{config: fq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := fq.withFix; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Fixbrand)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Fix(func(s *sql.Selector) {
			s.Where(sql.InValues(fixbrand.FixColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.fixbrand_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "fixbrand_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "fixbrand_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Fix = append(node.Edges.Fix, n)
		}
	}

	return nodes, nil
}

func (fq *FixbrandQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FixbrandQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (fq *FixbrandQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fixbrand.Table,
			Columns: fixbrand.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fixbrand.FieldID,
			},
		},
		From:   fq.sql,
		Unique: true,
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FixbrandQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(fixbrand.Table)
	selector := builder.Select(t1.Columns(fixbrand.Columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(fixbrand.Columns...)...)
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FixbrandGroupBy is the builder for group-by Fixbrand entities.
type FixbrandGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FixbrandGroupBy) Aggregate(fns ...AggregateFunc) *FixbrandGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the group-by query and scan the result into the given value.
func (fgb *FixbrandGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fgb.path(ctx)
	if err != nil {
		return err
	}
	fgb.sql = query
	return fgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fgb *FixbrandGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := fgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (fgb *FixbrandGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FixbrandGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fgb *FixbrandGroupBy) StringsX(ctx context.Context) []string {
	v, err := fgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (fgb *FixbrandGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fixbrand.Label}
	default:
		err = fmt.Errorf("ent: FixbrandGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fgb *FixbrandGroupBy) StringX(ctx context.Context) string {
	v, err := fgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (fgb *FixbrandGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FixbrandGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fgb *FixbrandGroupBy) IntsX(ctx context.Context) []int {
	v, err := fgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (fgb *FixbrandGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fixbrand.Label}
	default:
		err = fmt.Errorf("ent: FixbrandGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fgb *FixbrandGroupBy) IntX(ctx context.Context) int {
	v, err := fgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (fgb *FixbrandGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FixbrandGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fgb *FixbrandGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := fgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (fgb *FixbrandGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fixbrand.Label}
	default:
		err = fmt.Errorf("ent: FixbrandGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fgb *FixbrandGroupBy) Float64X(ctx context.Context) float64 {
	v, err := fgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (fgb *FixbrandGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FixbrandGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fgb *FixbrandGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := fgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (fgb *FixbrandGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fixbrand.Label}
	default:
		err = fmt.Errorf("ent: FixbrandGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fgb *FixbrandGroupBy) BoolX(ctx context.Context) bool {
	v, err := fgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fgb *FixbrandGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fgb.sqlQuery().Query()
	if err := fgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fgb *FixbrandGroupBy) sqlQuery() *sql.Selector {
	selector := fgb.sql
	columns := make([]string, 0, len(fgb.fields)+len(fgb.fns))
	columns = append(columns, fgb.fields...)
	for _, fn := range fgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(fgb.fields...)
}

// FixbrandSelect is the builder for select fields of Fixbrand entities.
type FixbrandSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (fs *FixbrandSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := fs.path(ctx)
	if err != nil {
		return err
	}
	fs.sql = query
	return fs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fs *FixbrandSelect) ScanX(ctx context.Context, v interface{}) {
	if err := fs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (fs *FixbrandSelect) Strings(ctx context.Context) ([]string, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FixbrandSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fs *FixbrandSelect) StringsX(ctx context.Context) []string {
	v, err := fs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (fs *FixbrandSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fixbrand.Label}
	default:
		err = fmt.Errorf("ent: FixbrandSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fs *FixbrandSelect) StringX(ctx context.Context) string {
	v, err := fs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (fs *FixbrandSelect) Ints(ctx context.Context) ([]int, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FixbrandSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fs *FixbrandSelect) IntsX(ctx context.Context) []int {
	v, err := fs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (fs *FixbrandSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fixbrand.Label}
	default:
		err = fmt.Errorf("ent: FixbrandSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fs *FixbrandSelect) IntX(ctx context.Context) int {
	v, err := fs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (fs *FixbrandSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FixbrandSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fs *FixbrandSelect) Float64sX(ctx context.Context) []float64 {
	v, err := fs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (fs *FixbrandSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fixbrand.Label}
	default:
		err = fmt.Errorf("ent: FixbrandSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fs *FixbrandSelect) Float64X(ctx context.Context) float64 {
	v, err := fs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (fs *FixbrandSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FixbrandSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fs *FixbrandSelect) BoolsX(ctx context.Context) []bool {
	v, err := fs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (fs *FixbrandSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fixbrand.Label}
	default:
		err = fmt.Errorf("ent: FixbrandSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fs *FixbrandSelect) BoolX(ctx context.Context) bool {
	v, err := fs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fs *FixbrandSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fs.sqlQuery().Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fs *FixbrandSelect) sqlQuery() sql.Querier {
	selector := fs.sql
	selector.Select(selector.Columns(fs.fields...)...)
	return selector
}
