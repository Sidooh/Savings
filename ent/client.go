// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"Savings/ent/migrate"

	"Savings/ent/personalaccount"
	"Savings/ent/personalaccounttransaction"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// PersonalAccount is the client for interacting with the PersonalAccount builders.
	PersonalAccount *PersonalAccountClient
	// PersonalAccountTransaction is the client for interacting with the PersonalAccountTransaction builders.
	PersonalAccountTransaction *PersonalAccountTransactionClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.PersonalAccount = NewPersonalAccountClient(c.config)
	c.PersonalAccountTransaction = NewPersonalAccountTransactionClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                        ctx,
		config:                     cfg,
		PersonalAccount:            NewPersonalAccountClient(cfg),
		PersonalAccountTransaction: NewPersonalAccountTransactionClient(cfg),
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
		ctx:                        ctx,
		config:                     cfg,
		PersonalAccount:            NewPersonalAccountClient(cfg),
		PersonalAccountTransaction: NewPersonalAccountTransactionClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		PersonalAccount.
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
	c.PersonalAccount.Use(hooks...)
	c.PersonalAccountTransaction.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.PersonalAccount.Intercept(interceptors...)
	c.PersonalAccountTransaction.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *PersonalAccountMutation:
		return c.PersonalAccount.mutate(ctx, m)
	case *PersonalAccountTransactionMutation:
		return c.PersonalAccountTransaction.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// PersonalAccountClient is a client for the PersonalAccount schema.
type PersonalAccountClient struct {
	config
}

// NewPersonalAccountClient returns a client for the PersonalAccount from the given config.
func NewPersonalAccountClient(c config) *PersonalAccountClient {
	return &PersonalAccountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `personalaccount.Hooks(f(g(h())))`.
func (c *PersonalAccountClient) Use(hooks ...Hook) {
	c.hooks.PersonalAccount = append(c.hooks.PersonalAccount, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `personalaccount.Intercept(f(g(h())))`.
func (c *PersonalAccountClient) Intercept(interceptors ...Interceptor) {
	c.inters.PersonalAccount = append(c.inters.PersonalAccount, interceptors...)
}

// Create returns a builder for creating a PersonalAccount entity.
func (c *PersonalAccountClient) Create() *PersonalAccountCreate {
	mutation := newPersonalAccountMutation(c.config, OpCreate)
	return &PersonalAccountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PersonalAccount entities.
func (c *PersonalAccountClient) CreateBulk(builders ...*PersonalAccountCreate) *PersonalAccountCreateBulk {
	return &PersonalAccountCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *PersonalAccountClient) MapCreateBulk(slice any, setFunc func(*PersonalAccountCreate, int)) *PersonalAccountCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PersonalAccountCreateBulk{err: fmt.Errorf("calling to PersonalAccountClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PersonalAccountCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PersonalAccountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PersonalAccount.
func (c *PersonalAccountClient) Update() *PersonalAccountUpdate {
	mutation := newPersonalAccountMutation(c.config, OpUpdate)
	return &PersonalAccountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PersonalAccountClient) UpdateOne(pa *PersonalAccount) *PersonalAccountUpdateOne {
	mutation := newPersonalAccountMutation(c.config, OpUpdateOne, withPersonalAccount(pa))
	return &PersonalAccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PersonalAccountClient) UpdateOneID(id uint64) *PersonalAccountUpdateOne {
	mutation := newPersonalAccountMutation(c.config, OpUpdateOne, withPersonalAccountID(id))
	return &PersonalAccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PersonalAccount.
func (c *PersonalAccountClient) Delete() *PersonalAccountDelete {
	mutation := newPersonalAccountMutation(c.config, OpDelete)
	return &PersonalAccountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PersonalAccountClient) DeleteOne(pa *PersonalAccount) *PersonalAccountDeleteOne {
	return c.DeleteOneID(pa.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PersonalAccountClient) DeleteOneID(id uint64) *PersonalAccountDeleteOne {
	builder := c.Delete().Where(personalaccount.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PersonalAccountDeleteOne{builder}
}

// Query returns a query builder for PersonalAccount.
func (c *PersonalAccountClient) Query() *PersonalAccountQuery {
	return &PersonalAccountQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePersonalAccount},
		inters: c.Interceptors(),
	}
}

// Get returns a PersonalAccount entity by its id.
func (c *PersonalAccountClient) Get(ctx context.Context, id uint64) (*PersonalAccount, error) {
	return c.Query().Where(personalaccount.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PersonalAccountClient) GetX(ctx context.Context, id uint64) *PersonalAccount {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTransactions queries the transactions edge of a PersonalAccount.
func (c *PersonalAccountClient) QueryTransactions(pa *PersonalAccount) *PersonalAccountTransactionQuery {
	query := (&PersonalAccountTransactionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pa.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(personalaccount.Table, personalaccount.FieldID, id),
			sqlgraph.To(personalaccounttransaction.Table, personalaccounttransaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, personalaccount.TransactionsTable, personalaccount.TransactionsColumn),
		)
		fromV = sqlgraph.Neighbors(pa.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PersonalAccountClient) Hooks() []Hook {
	return c.hooks.PersonalAccount
}

// Interceptors returns the client interceptors.
func (c *PersonalAccountClient) Interceptors() []Interceptor {
	return c.inters.PersonalAccount
}

func (c *PersonalAccountClient) mutate(ctx context.Context, m *PersonalAccountMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PersonalAccountCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PersonalAccountUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PersonalAccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PersonalAccountDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown PersonalAccount mutation op: %q", m.Op())
	}
}

// PersonalAccountTransactionClient is a client for the PersonalAccountTransaction schema.
type PersonalAccountTransactionClient struct {
	config
}

// NewPersonalAccountTransactionClient returns a client for the PersonalAccountTransaction from the given config.
func NewPersonalAccountTransactionClient(c config) *PersonalAccountTransactionClient {
	return &PersonalAccountTransactionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `personalaccounttransaction.Hooks(f(g(h())))`.
func (c *PersonalAccountTransactionClient) Use(hooks ...Hook) {
	c.hooks.PersonalAccountTransaction = append(c.hooks.PersonalAccountTransaction, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `personalaccounttransaction.Intercept(f(g(h())))`.
func (c *PersonalAccountTransactionClient) Intercept(interceptors ...Interceptor) {
	c.inters.PersonalAccountTransaction = append(c.inters.PersonalAccountTransaction, interceptors...)
}

// Create returns a builder for creating a PersonalAccountTransaction entity.
func (c *PersonalAccountTransactionClient) Create() *PersonalAccountTransactionCreate {
	mutation := newPersonalAccountTransactionMutation(c.config, OpCreate)
	return &PersonalAccountTransactionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PersonalAccountTransaction entities.
func (c *PersonalAccountTransactionClient) CreateBulk(builders ...*PersonalAccountTransactionCreate) *PersonalAccountTransactionCreateBulk {
	return &PersonalAccountTransactionCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *PersonalAccountTransactionClient) MapCreateBulk(slice any, setFunc func(*PersonalAccountTransactionCreate, int)) *PersonalAccountTransactionCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PersonalAccountTransactionCreateBulk{err: fmt.Errorf("calling to PersonalAccountTransactionClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PersonalAccountTransactionCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PersonalAccountTransactionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PersonalAccountTransaction.
func (c *PersonalAccountTransactionClient) Update() *PersonalAccountTransactionUpdate {
	mutation := newPersonalAccountTransactionMutation(c.config, OpUpdate)
	return &PersonalAccountTransactionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PersonalAccountTransactionClient) UpdateOne(pat *PersonalAccountTransaction) *PersonalAccountTransactionUpdateOne {
	mutation := newPersonalAccountTransactionMutation(c.config, OpUpdateOne, withPersonalAccountTransaction(pat))
	return &PersonalAccountTransactionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PersonalAccountTransactionClient) UpdateOneID(id uint64) *PersonalAccountTransactionUpdateOne {
	mutation := newPersonalAccountTransactionMutation(c.config, OpUpdateOne, withPersonalAccountTransactionID(id))
	return &PersonalAccountTransactionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PersonalAccountTransaction.
func (c *PersonalAccountTransactionClient) Delete() *PersonalAccountTransactionDelete {
	mutation := newPersonalAccountTransactionMutation(c.config, OpDelete)
	return &PersonalAccountTransactionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PersonalAccountTransactionClient) DeleteOne(pat *PersonalAccountTransaction) *PersonalAccountTransactionDeleteOne {
	return c.DeleteOneID(pat.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PersonalAccountTransactionClient) DeleteOneID(id uint64) *PersonalAccountTransactionDeleteOne {
	builder := c.Delete().Where(personalaccounttransaction.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PersonalAccountTransactionDeleteOne{builder}
}

// Query returns a query builder for PersonalAccountTransaction.
func (c *PersonalAccountTransactionClient) Query() *PersonalAccountTransactionQuery {
	return &PersonalAccountTransactionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePersonalAccountTransaction},
		inters: c.Interceptors(),
	}
}

// Get returns a PersonalAccountTransaction entity by its id.
func (c *PersonalAccountTransactionClient) Get(ctx context.Context, id uint64) (*PersonalAccountTransaction, error) {
	return c.Query().Where(personalaccounttransaction.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PersonalAccountTransactionClient) GetX(ctx context.Context, id uint64) *PersonalAccountTransaction {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAccount queries the account edge of a PersonalAccountTransaction.
func (c *PersonalAccountTransactionClient) QueryAccount(pat *PersonalAccountTransaction) *PersonalAccountQuery {
	query := (&PersonalAccountClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pat.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(personalaccounttransaction.Table, personalaccounttransaction.FieldID, id),
			sqlgraph.To(personalaccount.Table, personalaccount.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, personalaccounttransaction.AccountTable, personalaccounttransaction.AccountColumn),
		)
		fromV = sqlgraph.Neighbors(pat.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PersonalAccountTransactionClient) Hooks() []Hook {
	return c.hooks.PersonalAccountTransaction
}

// Interceptors returns the client interceptors.
func (c *PersonalAccountTransactionClient) Interceptors() []Interceptor {
	return c.inters.PersonalAccountTransaction
}

func (c *PersonalAccountTransactionClient) mutate(ctx context.Context, m *PersonalAccountTransactionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PersonalAccountTransactionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PersonalAccountTransactionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PersonalAccountTransactionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PersonalAccountTransactionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown PersonalAccountTransaction mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		PersonalAccount, PersonalAccountTransaction []ent.Hook
	}
	inters struct {
		PersonalAccount, PersonalAccountTransaction []ent.Interceptor
	}
)
