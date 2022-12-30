// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Pacerino/pr0music/ent/migrate"

	"github.com/Pacerino/pr0music/ent/comments"
	"github.com/Pacerino/pr0music/ent/items"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Comments is the client for interacting with the Comments builders.
	Comments *CommentsClient
	// Items is the client for interacting with the Items builders.
	Items *ItemsClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Comments = NewCommentsClient(c.config)
	c.Items = NewItemsClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Comments: NewCommentsClient(cfg),
		Items:    NewItemsClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Comments: NewCommentsClient(cfg),
		Items:    NewItemsClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Comments.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Comments.Use(hooks...)
	c.Items.Use(hooks...)
}

// CommentsClient is a client for the Comments schema.
type CommentsClient struct {
	config
}

// NewCommentsClient returns a client for the Comments from the given config.
func NewCommentsClient(c config) *CommentsClient {
	return &CommentsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `comments.Hooks(f(g(h())))`.
func (c *CommentsClient) Use(hooks ...Hook) {
	c.hooks.Comments = append(c.hooks.Comments, hooks...)
}

// Create returns a builder for creating a Comments entity.
func (c *CommentsClient) Create() *CommentsCreate {
	mutation := newCommentsMutation(c.config, OpCreate)
	return &CommentsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Comments entities.
func (c *CommentsClient) CreateBulk(builders ...*CommentsCreate) *CommentsCreateBulk {
	return &CommentsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Comments.
func (c *CommentsClient) Update() *CommentsUpdate {
	mutation := newCommentsMutation(c.config, OpUpdate)
	return &CommentsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CommentsClient) UpdateOne(co *Comments) *CommentsUpdateOne {
	mutation := newCommentsMutation(c.config, OpUpdateOne, withComments(co))
	return &CommentsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CommentsClient) UpdateOneID(id int) *CommentsUpdateOne {
	mutation := newCommentsMutation(c.config, OpUpdateOne, withCommentsID(id))
	return &CommentsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Comments.
func (c *CommentsClient) Delete() *CommentsDelete {
	mutation := newCommentsMutation(c.config, OpDelete)
	return &CommentsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CommentsClient) DeleteOne(co *Comments) *CommentsDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CommentsClient) DeleteOneID(id int) *CommentsDeleteOne {
	builder := c.Delete().Where(comments.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CommentsDeleteOne{builder}
}

// Query returns a query builder for Comments.
func (c *CommentsClient) Query() *CommentsQuery {
	return &CommentsQuery{
		config: c.config,
	}
}

// Get returns a Comments entity by its id.
func (c *CommentsClient) Get(ctx context.Context, id int) (*Comments, error) {
	return c.Query().Where(comments.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CommentsClient) GetX(ctx context.Context, id int) *Comments {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryItems queries the items edge of a Comments.
func (c *CommentsClient) QueryItems(co *Comments) *ItemsQuery {
	query := &ItemsQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comments.Table, comments.FieldID, id),
			sqlgraph.To(items.Table, items.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, comments.ItemsTable, comments.ItemsColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CommentsClient) Hooks() []Hook {
	return c.hooks.Comments
}

// ItemsClient is a client for the Items schema.
type ItemsClient struct {
	config
}

// NewItemsClient returns a client for the Items from the given config.
func NewItemsClient(c config) *ItemsClient {
	return &ItemsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `items.Hooks(f(g(h())))`.
func (c *ItemsClient) Use(hooks ...Hook) {
	c.hooks.Items = append(c.hooks.Items, hooks...)
}

// Create returns a builder for creating a Items entity.
func (c *ItemsClient) Create() *ItemsCreate {
	mutation := newItemsMutation(c.config, OpCreate)
	return &ItemsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Items entities.
func (c *ItemsClient) CreateBulk(builders ...*ItemsCreate) *ItemsCreateBulk {
	return &ItemsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Items.
func (c *ItemsClient) Update() *ItemsUpdate {
	mutation := newItemsMutation(c.config, OpUpdate)
	return &ItemsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ItemsClient) UpdateOne(i *Items) *ItemsUpdateOne {
	mutation := newItemsMutation(c.config, OpUpdateOne, withItems(i))
	return &ItemsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ItemsClient) UpdateOneID(id int) *ItemsUpdateOne {
	mutation := newItemsMutation(c.config, OpUpdateOne, withItemsID(id))
	return &ItemsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Items.
func (c *ItemsClient) Delete() *ItemsDelete {
	mutation := newItemsMutation(c.config, OpDelete)
	return &ItemsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ItemsClient) DeleteOne(i *Items) *ItemsDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ItemsClient) DeleteOneID(id int) *ItemsDeleteOne {
	builder := c.Delete().Where(items.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ItemsDeleteOne{builder}
}

// Query returns a query builder for Items.
func (c *ItemsClient) Query() *ItemsQuery {
	return &ItemsQuery{
		config: c.config,
	}
}

// Get returns a Items entity by its id.
func (c *ItemsClient) Get(ctx context.Context, id int) (*Items, error) {
	return c.Query().Where(items.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ItemsClient) GetX(ctx context.Context, id int) *Items {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryComments queries the comments edge of a Items.
func (c *ItemsClient) QueryComments(i *Items) *CommentsQuery {
	query := &CommentsQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(items.Table, items.FieldID, id),
			sqlgraph.To(comments.Table, comments.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, items.CommentsTable, items.CommentsColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ItemsClient) Hooks() []Hook {
	return c.hooks.Items
}
