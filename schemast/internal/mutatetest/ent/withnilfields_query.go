// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/contrib/schemast/internal/mutatetest/ent/predicate"
	"entgo.io/contrib/schemast/internal/mutatetest/ent/withnilfields"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WithNilFieldsQuery is the builder for querying WithNilFields entities.
type WithNilFieldsQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.WithNilFields
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WithNilFieldsQuery builder.
func (wnfq *WithNilFieldsQuery) Where(ps ...predicate.WithNilFields) *WithNilFieldsQuery {
	wnfq.predicates = append(wnfq.predicates, ps...)
	return wnfq
}

// Limit adds a limit step to the query.
func (wnfq *WithNilFieldsQuery) Limit(limit int) *WithNilFieldsQuery {
	wnfq.limit = &limit
	return wnfq
}

// Offset adds an offset step to the query.
func (wnfq *WithNilFieldsQuery) Offset(offset int) *WithNilFieldsQuery {
	wnfq.offset = &offset
	return wnfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wnfq *WithNilFieldsQuery) Unique(unique bool) *WithNilFieldsQuery {
	wnfq.unique = &unique
	return wnfq
}

// Order adds an order step to the query.
func (wnfq *WithNilFieldsQuery) Order(o ...OrderFunc) *WithNilFieldsQuery {
	wnfq.order = append(wnfq.order, o...)
	return wnfq
}

// First returns the first WithNilFields entity from the query.
// Returns a *NotFoundError when no WithNilFields was found.
func (wnfq *WithNilFieldsQuery) First(ctx context.Context) (*WithNilFields, error) {
	nodes, err := wnfq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{withnilfields.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) FirstX(ctx context.Context) *WithNilFields {
	node, err := wnfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WithNilFields ID from the query.
// Returns a *NotFoundError when no WithNilFields ID was found.
func (wnfq *WithNilFieldsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wnfq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{withnilfields.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) FirstIDX(ctx context.Context) int {
	id, err := wnfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WithNilFields entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WithNilFields entity is found.
// Returns a *NotFoundError when no WithNilFields entities are found.
func (wnfq *WithNilFieldsQuery) Only(ctx context.Context) (*WithNilFields, error) {
	nodes, err := wnfq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{withnilfields.Label}
	default:
		return nil, &NotSingularError{withnilfields.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) OnlyX(ctx context.Context) *WithNilFields {
	node, err := wnfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WithNilFields ID in the query.
// Returns a *NotSingularError when more than one WithNilFields ID is found.
// Returns a *NotFoundError when no entities are found.
func (wnfq *WithNilFieldsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wnfq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{withnilfields.Label}
	default:
		err = &NotSingularError{withnilfields.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) OnlyIDX(ctx context.Context) int {
	id, err := wnfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WithNilFieldsSlice.
func (wnfq *WithNilFieldsQuery) All(ctx context.Context) ([]*WithNilFields, error) {
	if err := wnfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wnfq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) AllX(ctx context.Context) []*WithNilFields {
	nodes, err := wnfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WithNilFields IDs.
func (wnfq *WithNilFieldsQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := wnfq.Select(withnilfields.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) IDsX(ctx context.Context) []int {
	ids, err := wnfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wnfq *WithNilFieldsQuery) Count(ctx context.Context) (int, error) {
	if err := wnfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wnfq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) CountX(ctx context.Context) int {
	count, err := wnfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wnfq *WithNilFieldsQuery) Exist(ctx context.Context) (bool, error) {
	if err := wnfq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wnfq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) ExistX(ctx context.Context) bool {
	exist, err := wnfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WithNilFieldsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wnfq *WithNilFieldsQuery) Clone() *WithNilFieldsQuery {
	if wnfq == nil {
		return nil
	}
	return &WithNilFieldsQuery{
		config:     wnfq.config,
		limit:      wnfq.limit,
		offset:     wnfq.offset,
		order:      append([]OrderFunc{}, wnfq.order...),
		predicates: append([]predicate.WithNilFields{}, wnfq.predicates...),
		// clone intermediate query.
		sql:    wnfq.sql.Clone(),
		path:   wnfq.path,
		unique: wnfq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (wnfq *WithNilFieldsQuery) GroupBy(field string, fields ...string) *WithNilFieldsGroupBy {
	group := &WithNilFieldsGroupBy{config: wnfq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wnfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wnfq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (wnfq *WithNilFieldsQuery) Select(fields ...string) *WithNilFieldsSelect {
	wnfq.fields = append(wnfq.fields, fields...)
	return &WithNilFieldsSelect{WithNilFieldsQuery: wnfq}
}

func (wnfq *WithNilFieldsQuery) prepareQuery(ctx context.Context) error {
	for _, f := range wnfq.fields {
		if !withnilfields.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wnfq.path != nil {
		prev, err := wnfq.path(ctx)
		if err != nil {
			return err
		}
		wnfq.sql = prev
	}
	return nil
}

func (wnfq *WithNilFieldsQuery) sqlAll(ctx context.Context) ([]*WithNilFields, error) {
	var (
		nodes = []*WithNilFields{}
		_spec = wnfq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &WithNilFields{config: wnfq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, wnfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (wnfq *WithNilFieldsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wnfq.querySpec()
	_spec.Node.Columns = wnfq.fields
	if len(wnfq.fields) > 0 {
		_spec.Unique = wnfq.unique != nil && *wnfq.unique
	}
	return sqlgraph.CountNodes(ctx, wnfq.driver, _spec)
}

func (wnfq *WithNilFieldsQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wnfq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (wnfq *WithNilFieldsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   withnilfields.Table,
			Columns: withnilfields.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: withnilfields.FieldID,
			},
		},
		From:   wnfq.sql,
		Unique: true,
	}
	if unique := wnfq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := wnfq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, withnilfields.FieldID)
		for i := range fields {
			if fields[i] != withnilfields.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wnfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wnfq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wnfq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wnfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wnfq *WithNilFieldsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wnfq.driver.Dialect())
	t1 := builder.Table(withnilfields.Table)
	columns := wnfq.fields
	if len(columns) == 0 {
		columns = withnilfields.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wnfq.sql != nil {
		selector = wnfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wnfq.unique != nil && *wnfq.unique {
		selector.Distinct()
	}
	for _, p := range wnfq.predicates {
		p(selector)
	}
	for _, p := range wnfq.order {
		p(selector)
	}
	if offset := wnfq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wnfq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNilFieldsGroupBy is the group-by builder for WithNilFields entities.
type WithNilFieldsGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wnfgb *WithNilFieldsGroupBy) Aggregate(fns ...AggregateFunc) *WithNilFieldsGroupBy {
	wnfgb.fns = append(wnfgb.fns, fns...)
	return wnfgb
}

// Scan applies the group-by query and scans the result into the given value.
func (wnfgb *WithNilFieldsGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := wnfgb.path(ctx)
	if err != nil {
		return err
	}
	wnfgb.sql = query
	return wnfgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wnfgb *WithNilFieldsGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := wnfgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (wnfgb *WithNilFieldsGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(wnfgb.fields) > 1 {
		return nil, errors.New("ent: WithNilFieldsGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := wnfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wnfgb *WithNilFieldsGroupBy) StringsX(ctx context.Context) []string {
	v, err := wnfgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wnfgb *WithNilFieldsGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wnfgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withnilfields.Label}
	default:
		err = fmt.Errorf("ent: WithNilFieldsGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wnfgb *WithNilFieldsGroupBy) StringX(ctx context.Context) string {
	v, err := wnfgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (wnfgb *WithNilFieldsGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(wnfgb.fields) > 1 {
		return nil, errors.New("ent: WithNilFieldsGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := wnfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wnfgb *WithNilFieldsGroupBy) IntsX(ctx context.Context) []int {
	v, err := wnfgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wnfgb *WithNilFieldsGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wnfgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withnilfields.Label}
	default:
		err = fmt.Errorf("ent: WithNilFieldsGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wnfgb *WithNilFieldsGroupBy) IntX(ctx context.Context) int {
	v, err := wnfgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (wnfgb *WithNilFieldsGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(wnfgb.fields) > 1 {
		return nil, errors.New("ent: WithNilFieldsGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := wnfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wnfgb *WithNilFieldsGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := wnfgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wnfgb *WithNilFieldsGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wnfgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withnilfields.Label}
	default:
		err = fmt.Errorf("ent: WithNilFieldsGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wnfgb *WithNilFieldsGroupBy) Float64X(ctx context.Context) float64 {
	v, err := wnfgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (wnfgb *WithNilFieldsGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(wnfgb.fields) > 1 {
		return nil, errors.New("ent: WithNilFieldsGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := wnfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wnfgb *WithNilFieldsGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := wnfgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wnfgb *WithNilFieldsGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wnfgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withnilfields.Label}
	default:
		err = fmt.Errorf("ent: WithNilFieldsGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wnfgb *WithNilFieldsGroupBy) BoolX(ctx context.Context) bool {
	v, err := wnfgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wnfgb *WithNilFieldsGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range wnfgb.fields {
		if !withnilfields.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wnfgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wnfgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wnfgb *WithNilFieldsGroupBy) sqlQuery() *sql.Selector {
	selector := wnfgb.sql.Select()
	aggregation := make([]string, 0, len(wnfgb.fns))
	for _, fn := range wnfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(wnfgb.fields)+len(wnfgb.fns))
		for _, f := range wnfgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(wnfgb.fields...)...)
}

// WithNilFieldsSelect is the builder for selecting fields of WithNilFields entities.
type WithNilFieldsSelect struct {
	*WithNilFieldsQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (wnfs *WithNilFieldsSelect) Scan(ctx context.Context, v interface{}) error {
	if err := wnfs.prepareQuery(ctx); err != nil {
		return err
	}
	wnfs.sql = wnfs.WithNilFieldsQuery.sqlQuery(ctx)
	return wnfs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wnfs *WithNilFieldsSelect) ScanX(ctx context.Context, v interface{}) {
	if err := wnfs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (wnfs *WithNilFieldsSelect) Strings(ctx context.Context) ([]string, error) {
	if len(wnfs.fields) > 1 {
		return nil, errors.New("ent: WithNilFieldsSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := wnfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wnfs *WithNilFieldsSelect) StringsX(ctx context.Context) []string {
	v, err := wnfs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (wnfs *WithNilFieldsSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wnfs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withnilfields.Label}
	default:
		err = fmt.Errorf("ent: WithNilFieldsSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wnfs *WithNilFieldsSelect) StringX(ctx context.Context) string {
	v, err := wnfs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (wnfs *WithNilFieldsSelect) Ints(ctx context.Context) ([]int, error) {
	if len(wnfs.fields) > 1 {
		return nil, errors.New("ent: WithNilFieldsSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := wnfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wnfs *WithNilFieldsSelect) IntsX(ctx context.Context) []int {
	v, err := wnfs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (wnfs *WithNilFieldsSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wnfs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withnilfields.Label}
	default:
		err = fmt.Errorf("ent: WithNilFieldsSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wnfs *WithNilFieldsSelect) IntX(ctx context.Context) int {
	v, err := wnfs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (wnfs *WithNilFieldsSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(wnfs.fields) > 1 {
		return nil, errors.New("ent: WithNilFieldsSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := wnfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wnfs *WithNilFieldsSelect) Float64sX(ctx context.Context) []float64 {
	v, err := wnfs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (wnfs *WithNilFieldsSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wnfs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withnilfields.Label}
	default:
		err = fmt.Errorf("ent: WithNilFieldsSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wnfs *WithNilFieldsSelect) Float64X(ctx context.Context) float64 {
	v, err := wnfs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (wnfs *WithNilFieldsSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(wnfs.fields) > 1 {
		return nil, errors.New("ent: WithNilFieldsSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := wnfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wnfs *WithNilFieldsSelect) BoolsX(ctx context.Context) []bool {
	v, err := wnfs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (wnfs *WithNilFieldsSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wnfs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withnilfields.Label}
	default:
		err = fmt.Errorf("ent: WithNilFieldsSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wnfs *WithNilFieldsSelect) BoolX(ctx context.Context) bool {
	v, err := wnfs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wnfs *WithNilFieldsSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := wnfs.sql.Query()
	if err := wnfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
