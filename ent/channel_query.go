// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/channel"
	"github.com/zibbp/ganymede/ent/live"
	"github.com/zibbp/ganymede/ent/predicate"
	"github.com/zibbp/ganymede/ent/vod"
)

// ChannelQuery is the builder for querying Channel entities.
type ChannelQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.Channel
	withVods   *VodQuery
	withLive   *LiveQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ChannelQuery builder.
func (cq *ChannelQuery) Where(ps ...predicate.Channel) *ChannelQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *ChannelQuery) Limit(limit int) *ChannelQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *ChannelQuery) Offset(offset int) *ChannelQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ChannelQuery) Unique(unique bool) *ChannelQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *ChannelQuery) Order(o ...OrderFunc) *ChannelQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryVods chains the current query on the "vods" edge.
func (cq *ChannelQuery) QueryVods() *VodQuery {
	query := (&VodClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(channel.Table, channel.FieldID, selector),
			sqlgraph.To(vod.Table, vod.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, channel.VodsTable, channel.VodsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLive chains the current query on the "live" edge.
func (cq *ChannelQuery) QueryLive() *LiveQuery {
	query := (&LiveClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(channel.Table, channel.FieldID, selector),
			sqlgraph.To(live.Table, live.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, channel.LiveTable, channel.LiveColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Channel entity from the query.
// Returns a *NotFoundError when no Channel was found.
func (cq *ChannelQuery) First(ctx context.Context) (*Channel, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{channel.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ChannelQuery) FirstX(ctx context.Context) *Channel {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Channel ID from the query.
// Returns a *NotFoundError when no Channel ID was found.
func (cq *ChannelQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{channel.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ChannelQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Channel entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Channel entity is found.
// Returns a *NotFoundError when no Channel entities are found.
func (cq *ChannelQuery) Only(ctx context.Context) (*Channel, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{channel.Label}
	default:
		return nil, &NotSingularError{channel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ChannelQuery) OnlyX(ctx context.Context) *Channel {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Channel ID in the query.
// Returns a *NotSingularError when more than one Channel ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ChannelQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{channel.Label}
	default:
		err = &NotSingularError{channel.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ChannelQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Channels.
func (cq *ChannelQuery) All(ctx context.Context) ([]*Channel, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Channel, *ChannelQuery]()
	return withInterceptors[[]*Channel](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *ChannelQuery) AllX(ctx context.Context) []*Channel {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Channel IDs.
func (cq *ChannelQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(channel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ChannelQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ChannelQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*ChannelQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ChannelQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ChannelQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ChannelQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ChannelQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ChannelQuery) Clone() *ChannelQuery {
	if cq == nil {
		return nil
	}
	return &ChannelQuery{
		config:     cq.config,
		ctx:        cq.ctx.Clone(),
		order:      append([]OrderFunc{}, cq.order...),
		inters:     append([]Interceptor{}, cq.inters...),
		predicates: append([]predicate.Channel{}, cq.predicates...),
		withVods:   cq.withVods.Clone(),
		withLive:   cq.withLive.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithVods tells the query-builder to eager-load the nodes that are connected to
// the "vods" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ChannelQuery) WithVods(opts ...func(*VodQuery)) *ChannelQuery {
	query := (&VodClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withVods = query
	return cq
}

// WithLive tells the query-builder to eager-load the nodes that are connected to
// the "live" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ChannelQuery) WithLive(opts ...func(*LiveQuery)) *ChannelQuery {
	query := (&LiveClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withLive = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ExtID string `json:"ext_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Channel.Query().
//		GroupBy(channel.FieldExtID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *ChannelQuery) GroupBy(field string, fields ...string) *ChannelGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ChannelGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = channel.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ExtID string `json:"ext_id,omitempty"`
//	}
//
//	client.Channel.Query().
//		Select(channel.FieldExtID).
//		Scan(ctx, &v)
func (cq *ChannelQuery) Select(fields ...string) *ChannelSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &ChannelSelect{ChannelQuery: cq}
	sbuild.label = channel.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ChannelSelect configured with the given aggregations.
func (cq *ChannelQuery) Aggregate(fns ...AggregateFunc) *ChannelSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *ChannelQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !channel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *ChannelQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Channel, error) {
	var (
		nodes       = []*Channel{}
		_spec       = cq.querySpec()
		loadedTypes = [2]bool{
			cq.withVods != nil,
			cq.withLive != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Channel).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Channel{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withVods; query != nil {
		if err := cq.loadVods(ctx, query, nodes,
			func(n *Channel) { n.Edges.Vods = []*Vod{} },
			func(n *Channel, e *Vod) { n.Edges.Vods = append(n.Edges.Vods, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withLive; query != nil {
		if err := cq.loadLive(ctx, query, nodes,
			func(n *Channel) { n.Edges.Live = []*Live{} },
			func(n *Channel, e *Live) { n.Edges.Live = append(n.Edges.Live, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *ChannelQuery) loadVods(ctx context.Context, query *VodQuery, nodes []*Channel, init func(*Channel), assign func(*Channel, *Vod)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Channel)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Vod(func(s *sql.Selector) {
		s.Where(sql.InValues(channel.VodsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.channel_vods
		if fk == nil {
			return fmt.Errorf(`foreign-key "channel_vods" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "channel_vods" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *ChannelQuery) loadLive(ctx context.Context, query *LiveQuery, nodes []*Channel, init func(*Channel), assign func(*Channel, *Live)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Channel)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Live(func(s *sql.Selector) {
		s.Where(sql.InValues(channel.LiveColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.channel_live
		if fk == nil {
			return fmt.Errorf(`foreign-key "channel_live" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "channel_live" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *ChannelQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ChannelQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(channel.Table, channel.Columns, sqlgraph.NewFieldSpec(channel.FieldID, field.TypeUUID))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, channel.FieldID)
		for i := range fields {
			if fields[i] != channel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *ChannelQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(channel.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = channel.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ChannelGroupBy is the group-by builder for Channel entities.
type ChannelGroupBy struct {
	selector
	build *ChannelQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ChannelGroupBy) Aggregate(fns ...AggregateFunc) *ChannelGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *ChannelGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChannelQuery, *ChannelGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *ChannelGroupBy) sqlScan(ctx context.Context, root *ChannelQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ChannelSelect is the builder for selecting fields of Channel entities.
type ChannelSelect struct {
	*ChannelQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *ChannelSelect) Aggregate(fns ...AggregateFunc) *ChannelSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ChannelSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChannelQuery, *ChannelSelect](ctx, cs.ChannelQuery, cs, cs.inters, v)
}

func (cs *ChannelSelect) sqlScan(ctx context.Context, root *ChannelQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
