// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/tag"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweet"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweettag"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TweetTagQuery is the builder for querying TweetTag entities.
type TweetTagQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TweetTag
	withTag    *TagQuery
	withTweet  *TweetQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TweetTagQuery builder.
func (ttq *TweetTagQuery) Where(ps ...predicate.TweetTag) *TweetTagQuery {
	ttq.predicates = append(ttq.predicates, ps...)
	return ttq
}

// Limit adds a limit step to the query.
func (ttq *TweetTagQuery) Limit(limit int) *TweetTagQuery {
	ttq.limit = &limit
	return ttq
}

// Offset adds an offset step to the query.
func (ttq *TweetTagQuery) Offset(offset int) *TweetTagQuery {
	ttq.offset = &offset
	return ttq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ttq *TweetTagQuery) Unique(unique bool) *TweetTagQuery {
	ttq.unique = &unique
	return ttq
}

// Order adds an order step to the query.
func (ttq *TweetTagQuery) Order(o ...OrderFunc) *TweetTagQuery {
	ttq.order = append(ttq.order, o...)
	return ttq
}

// QueryTag chains the current query on the "tag" edge.
func (ttq *TweetTagQuery) QueryTag() *TagQuery {
	query := &TagQuery{config: ttq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ttq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tweettag.Table, tweettag.FieldID, selector),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, tweettag.TagTable, tweettag.TagColumn),
		)
		fromU = sqlgraph.SetNeighbors(ttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTweet chains the current query on the "tweet" edge.
func (ttq *TweetTagQuery) QueryTweet() *TweetQuery {
	query := &TweetQuery{config: ttq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ttq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tweettag.Table, tweettag.FieldID, selector),
			sqlgraph.To(tweet.Table, tweet.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, tweettag.TweetTable, tweettag.TweetColumn),
		)
		fromU = sqlgraph.SetNeighbors(ttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TweetTag entity from the query.
// Returns a *NotFoundError when no TweetTag was found.
func (ttq *TweetTagQuery) First(ctx context.Context) (*TweetTag, error) {
	nodes, err := ttq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tweettag.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ttq *TweetTagQuery) FirstX(ctx context.Context) *TweetTag {
	node, err := ttq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TweetTag ID from the query.
// Returns a *NotFoundError when no TweetTag ID was found.
func (ttq *TweetTagQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ttq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tweettag.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ttq *TweetTagQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ttq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TweetTag entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TweetTag entity is found.
// Returns a *NotFoundError when no TweetTag entities are found.
func (ttq *TweetTagQuery) Only(ctx context.Context) (*TweetTag, error) {
	nodes, err := ttq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tweettag.Label}
	default:
		return nil, &NotSingularError{tweettag.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ttq *TweetTagQuery) OnlyX(ctx context.Context) *TweetTag {
	node, err := ttq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TweetTag ID in the query.
// Returns a *NotSingularError when more than one TweetTag ID is found.
// Returns a *NotFoundError when no entities are found.
func (ttq *TweetTagQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ttq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tweettag.Label}
	default:
		err = &NotSingularError{tweettag.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ttq *TweetTagQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ttq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TweetTags.
func (ttq *TweetTagQuery) All(ctx context.Context) ([]*TweetTag, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ttq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ttq *TweetTagQuery) AllX(ctx context.Context) []*TweetTag {
	nodes, err := ttq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TweetTag IDs.
func (ttq *TweetTagQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := ttq.Select(tweettag.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ttq *TweetTagQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ttq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ttq *TweetTagQuery) Count(ctx context.Context) (int, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ttq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ttq *TweetTagQuery) CountX(ctx context.Context) int {
	count, err := ttq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ttq *TweetTagQuery) Exist(ctx context.Context) (bool, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ttq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ttq *TweetTagQuery) ExistX(ctx context.Context) bool {
	exist, err := ttq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TweetTagQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ttq *TweetTagQuery) Clone() *TweetTagQuery {
	if ttq == nil {
		return nil
	}
	return &TweetTagQuery{
		config:     ttq.config,
		limit:      ttq.limit,
		offset:     ttq.offset,
		order:      append([]OrderFunc{}, ttq.order...),
		predicates: append([]predicate.TweetTag{}, ttq.predicates...),
		withTag:    ttq.withTag.Clone(),
		withTweet:  ttq.withTweet.Clone(),
		// clone intermediate query.
		sql:    ttq.sql.Clone(),
		path:   ttq.path,
		unique: ttq.unique,
	}
}

// WithTag tells the query-builder to eager-load the nodes that are connected to
// the "tag" edge. The optional arguments are used to configure the query builder of the edge.
func (ttq *TweetTagQuery) WithTag(opts ...func(*TagQuery)) *TweetTagQuery {
	query := &TagQuery{config: ttq.config}
	for _, opt := range opts {
		opt(query)
	}
	ttq.withTag = query
	return ttq
}

// WithTweet tells the query-builder to eager-load the nodes that are connected to
// the "tweet" edge. The optional arguments are used to configure the query builder of the edge.
func (ttq *TweetTagQuery) WithTweet(opts ...func(*TweetQuery)) *TweetTagQuery {
	query := &TweetQuery{config: ttq.config}
	for _, opt := range opts {
		opt(query)
	}
	ttq.withTweet = query
	return ttq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AddedAt time.Time `json:"added_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TweetTag.Query().
//		GroupBy(tweettag.FieldAddedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ttq *TweetTagQuery) GroupBy(field string, fields ...string) *TweetTagGroupBy {
	grbuild := &TweetTagGroupBy{config: ttq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ttq.sqlQuery(ctx), nil
	}
	grbuild.label = tweettag.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AddedAt time.Time `json:"added_at,omitempty"`
//	}
//
//	client.TweetTag.Query().
//		Select(tweettag.FieldAddedAt).
//		Scan(ctx, &v)
func (ttq *TweetTagQuery) Select(fields ...string) *TweetTagSelect {
	ttq.fields = append(ttq.fields, fields...)
	selbuild := &TweetTagSelect{TweetTagQuery: ttq}
	selbuild.label = tweettag.Label
	selbuild.flds, selbuild.scan = &ttq.fields, selbuild.Scan
	return selbuild
}

func (ttq *TweetTagQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ttq.fields {
		if !tweettag.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ttq.path != nil {
		prev, err := ttq.path(ctx)
		if err != nil {
			return err
		}
		ttq.sql = prev
	}
	return nil
}

func (ttq *TweetTagQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TweetTag, error) {
	var (
		nodes       = []*TweetTag{}
		_spec       = ttq.querySpec()
		loadedTypes = [2]bool{
			ttq.withTag != nil,
			ttq.withTweet != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TweetTag).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TweetTag{config: ttq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ttq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ttq.withTag; query != nil {
		if err := ttq.loadTag(ctx, query, nodes, nil,
			func(n *TweetTag, e *Tag) { n.Edges.Tag = e }); err != nil {
			return nil, err
		}
	}
	if query := ttq.withTweet; query != nil {
		if err := ttq.loadTweet(ctx, query, nodes, nil,
			func(n *TweetTag, e *Tweet) { n.Edges.Tweet = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ttq *TweetTagQuery) loadTag(ctx context.Context, query *TagQuery, nodes []*TweetTag, init func(*TweetTag), assign func(*TweetTag, *Tag)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*TweetTag)
	for i := range nodes {
		fk := nodes[i].TagID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(tag.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "tag_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ttq *TweetTagQuery) loadTweet(ctx context.Context, query *TweetQuery, nodes []*TweetTag, init func(*TweetTag), assign func(*TweetTag, *Tweet)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*TweetTag)
	for i := range nodes {
		fk := nodes[i].TweetID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(tweet.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "tweet_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ttq *TweetTagQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ttq.querySpec()
	_spec.Node.Columns = ttq.fields
	if len(ttq.fields) > 0 {
		_spec.Unique = ttq.unique != nil && *ttq.unique
	}
	return sqlgraph.CountNodes(ctx, ttq.driver, _spec)
}

func (ttq *TweetTagQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := ttq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (ttq *TweetTagQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tweettag.Table,
			Columns: tweettag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: tweettag.FieldID,
			},
		},
		From:   ttq.sql,
		Unique: true,
	}
	if unique := ttq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ttq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tweettag.FieldID)
		for i := range fields {
			if fields[i] != tweettag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ttq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ttq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ttq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ttq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ttq *TweetTagQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ttq.driver.Dialect())
	t1 := builder.Table(tweettag.Table)
	columns := ttq.fields
	if len(columns) == 0 {
		columns = tweettag.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ttq.sql != nil {
		selector = ttq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ttq.unique != nil && *ttq.unique {
		selector.Distinct()
	}
	for _, p := range ttq.predicates {
		p(selector)
	}
	for _, p := range ttq.order {
		p(selector)
	}
	if offset := ttq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ttq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TweetTagGroupBy is the group-by builder for TweetTag entities.
type TweetTagGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ttgb *TweetTagGroupBy) Aggregate(fns ...AggregateFunc) *TweetTagGroupBy {
	ttgb.fns = append(ttgb.fns, fns...)
	return ttgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ttgb *TweetTagGroupBy) Scan(ctx context.Context, v any) error {
	query, err := ttgb.path(ctx)
	if err != nil {
		return err
	}
	ttgb.sql = query
	return ttgb.sqlScan(ctx, v)
}

func (ttgb *TweetTagGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range ttgb.fields {
		if !tweettag.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ttgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ttgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ttgb *TweetTagGroupBy) sqlQuery() *sql.Selector {
	selector := ttgb.sql.Select()
	aggregation := make([]string, 0, len(ttgb.fns))
	for _, fn := range ttgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ttgb.fields)+len(ttgb.fns))
		for _, f := range ttgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ttgb.fields...)...)
}

// TweetTagSelect is the builder for selecting fields of TweetTag entities.
type TweetTagSelect struct {
	*TweetTagQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tts *TweetTagSelect) Scan(ctx context.Context, v any) error {
	if err := tts.prepareQuery(ctx); err != nil {
		return err
	}
	tts.sql = tts.TweetTagQuery.sqlQuery(ctx)
	return tts.sqlScan(ctx, v)
}

func (tts *TweetTagSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := tts.sql.Query()
	if err := tts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
