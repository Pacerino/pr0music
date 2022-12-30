// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/Pacerino/pr0music/ent/comments"
	"github.com/Pacerino/pr0music/ent/items"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vmihailenco/msgpack/v5"
)

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

// MarshalGQL implements graphql.Marshaler interface.
func (o OrderDirection) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(o.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (o *OrderDirection) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("order direction %T must be a string", val)
	}
	*o = OrderDirection(str)
	return o.Validate()
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

func (o OrderDirection) orderFunc(field string) OrderFunc {
	if o == OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

func cursorsToPredicates(direction OrderDirection, after, before *Cursor, field, idField string) []func(s *sql.Selector) {
	var predicates []func(s *sql.Selector)
	if after != nil {
		if after.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeGT
			} else {
				predicate = sql.CompositeLT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					after.Value, after.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.GT
			} else {
				predicate = sql.LT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					after.ID,
				))
			})
		}
	}
	if before != nil {
		if before.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeLT
			} else {
				predicate = sql.CompositeGT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					before.Value, before.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.LT
			} else {
				predicate = sql.GT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					before.ID,
				))
			})
		}
	}
	return predicates
}

// PageInfo of a connection type.
type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *Cursor `json:"startCursor"`
	EndCursor       *Cursor `json:"endCursor"`
}

// Cursor of an edge type.
type Cursor struct {
	ID    int   `msgpack:"i"`
	Value Value `msgpack:"v,omitempty"`
}

// MarshalGQL implements graphql.Marshaler interface.
func (c Cursor) MarshalGQL(w io.Writer) {
	quote := []byte{'"'}
	w.Write(quote)
	defer w.Write(quote)
	wc := base64.NewEncoder(base64.RawStdEncoding, w)
	defer wc.Close()
	_ = msgpack.NewEncoder(wc).Encode(c)
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (c *Cursor) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a string", v)
	}
	if err := msgpack.NewDecoder(
		base64.NewDecoder(
			base64.RawStdEncoding,
			strings.NewReader(s),
		),
	).Decode(c); err != nil {
		return fmt.Errorf("cannot decode cursor: %w", err)
	}
	return nil
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// CommentsEdge is the edge representation of Comments.
type CommentsEdge struct {
	Node   *Comments `json:"node"`
	Cursor Cursor    `json:"cursor"`
}

// CommentsConnection is the connection containing edges to Comments.
type CommentsConnection struct {
	Edges      []*CommentsEdge `json:"edges"`
	PageInfo   PageInfo        `json:"pageInfo"`
	TotalCount int             `json:"totalCount"`
}

func (c *CommentsConnection) build(nodes []*Comments, pager *commentsPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Comments
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Comments {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Comments {
			return nodes[i]
		}
	}
	c.Edges = make([]*CommentsEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &CommentsEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// CommentsPaginateOption enables pagination customization.
type CommentsPaginateOption func(*commentsPager) error

// WithCommentsOrder configures pagination ordering.
func WithCommentsOrder(order *CommentsOrder) CommentsPaginateOption {
	if order == nil {
		order = DefaultCommentsOrder
	}
	o := *order
	return func(pager *commentsPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultCommentsOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithCommentsFilter configures pagination filter.
func WithCommentsFilter(filter func(*CommentsQuery) (*CommentsQuery, error)) CommentsPaginateOption {
	return func(pager *commentsPager) error {
		if filter == nil {
			return errors.New("CommentsQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type commentsPager struct {
	order  *CommentsOrder
	filter func(*CommentsQuery) (*CommentsQuery, error)
}

func newCommentsPager(opts []CommentsPaginateOption) (*commentsPager, error) {
	pager := &commentsPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultCommentsOrder
	}
	return pager, nil
}

func (p *commentsPager) applyFilter(query *CommentsQuery) (*CommentsQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *commentsPager) toCursor(c *Comments) Cursor {
	return p.order.Field.toCursor(c)
}

func (p *commentsPager) applyCursors(query *CommentsQuery, after, before *Cursor) *CommentsQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultCommentsOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *commentsPager) applyOrder(query *CommentsQuery, reverse bool) *CommentsQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultCommentsOrder.Field {
		query = query.Order(direction.orderFunc(DefaultCommentsOrder.Field.field))
	}
	return query
}

func (p *commentsPager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultCommentsOrder.Field {
			b.Comma().Ident(DefaultCommentsOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Comments.
func (c *CommentsQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...CommentsPaginateOption,
) (*CommentsConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newCommentsPager(opts)
	if err != nil {
		return nil, err
	}
	if c, err = pager.applyFilter(c); err != nil {
		return nil, err
	}
	conn := &CommentsConnection{Edges: []*CommentsEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = c.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	c = pager.applyCursors(c, after, before)
	c = pager.applyOrder(c, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		c.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := c.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := c.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// CommentsOrderField defines the ordering field of Comments.
type CommentsOrderField struct {
	field    string
	toCursor func(*Comments) Cursor
}

// CommentsOrder defines the ordering of Comments.
type CommentsOrder struct {
	Direction OrderDirection      `json:"direction"`
	Field     *CommentsOrderField `json:"field"`
}

// DefaultCommentsOrder is the default ordering of Comments.
var DefaultCommentsOrder = &CommentsOrder{
	Direction: OrderDirectionAsc,
	Field: &CommentsOrderField{
		field: comments.FieldID,
		toCursor: func(c *Comments) Cursor {
			return Cursor{ID: c.ID}
		},
	},
}

// ToEdge converts Comments into CommentsEdge.
func (c *Comments) ToEdge(order *CommentsOrder) *CommentsEdge {
	if order == nil {
		order = DefaultCommentsOrder
	}
	return &CommentsEdge{
		Node:   c,
		Cursor: order.Field.toCursor(c),
	}
}

// ItemsEdge is the edge representation of Items.
type ItemsEdge struct {
	Node   *Items `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// ItemsConnection is the connection containing edges to Items.
type ItemsConnection struct {
	Edges      []*ItemsEdge `json:"edges"`
	PageInfo   PageInfo     `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

func (c *ItemsConnection) build(nodes []*Items, pager *itemsPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Items
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Items {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Items {
			return nodes[i]
		}
	}
	c.Edges = make([]*ItemsEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &ItemsEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// ItemsPaginateOption enables pagination customization.
type ItemsPaginateOption func(*itemsPager) error

// WithItemsOrder configures pagination ordering.
func WithItemsOrder(order *ItemsOrder) ItemsPaginateOption {
	if order == nil {
		order = DefaultItemsOrder
	}
	o := *order
	return func(pager *itemsPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultItemsOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithItemsFilter configures pagination filter.
func WithItemsFilter(filter func(*ItemsQuery) (*ItemsQuery, error)) ItemsPaginateOption {
	return func(pager *itemsPager) error {
		if filter == nil {
			return errors.New("ItemsQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type itemsPager struct {
	order  *ItemsOrder
	filter func(*ItemsQuery) (*ItemsQuery, error)
}

func newItemsPager(opts []ItemsPaginateOption) (*itemsPager, error) {
	pager := &itemsPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultItemsOrder
	}
	return pager, nil
}

func (p *itemsPager) applyFilter(query *ItemsQuery) (*ItemsQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *itemsPager) toCursor(i *Items) Cursor {
	return p.order.Field.toCursor(i)
}

func (p *itemsPager) applyCursors(query *ItemsQuery, after, before *Cursor) *ItemsQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultItemsOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *itemsPager) applyOrder(query *ItemsQuery, reverse bool) *ItemsQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultItemsOrder.Field {
		query = query.Order(direction.orderFunc(DefaultItemsOrder.Field.field))
	}
	return query
}

func (p *itemsPager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultItemsOrder.Field {
			b.Comma().Ident(DefaultItemsOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Items.
func (i *ItemsQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...ItemsPaginateOption,
) (*ItemsConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newItemsPager(opts)
	if err != nil {
		return nil, err
	}
	if i, err = pager.applyFilter(i); err != nil {
		return nil, err
	}
	conn := &ItemsConnection{Edges: []*ItemsEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = i.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	i = pager.applyCursors(i, after, before)
	i = pager.applyOrder(i, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		i.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := i.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := i.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// ItemsOrderField defines the ordering field of Items.
type ItemsOrderField struct {
	field    string
	toCursor func(*Items) Cursor
}

// ItemsOrder defines the ordering of Items.
type ItemsOrder struct {
	Direction OrderDirection   `json:"direction"`
	Field     *ItemsOrderField `json:"field"`
}

// DefaultItemsOrder is the default ordering of Items.
var DefaultItemsOrder = &ItemsOrder{
	Direction: OrderDirectionAsc,
	Field: &ItemsOrderField{
		field: items.FieldID,
		toCursor: func(i *Items) Cursor {
			return Cursor{ID: i.ID}
		},
	},
}

// ToEdge converts Items into ItemsEdge.
func (i *Items) ToEdge(order *ItemsOrder) *ItemsEdge {
	if order == nil {
		order = DefaultItemsOrder
	}
	return &ItemsEdge{
		Node:   i,
		Cursor: order.Field.toCursor(i),
	}
}
