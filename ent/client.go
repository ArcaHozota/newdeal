// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"newdeal/ent/migrate"

	"newdeal/ent/auth"
	"newdeal/ent/hymn"
	"newdeal/ent/hymnswork"
	"newdeal/ent/role"
	"newdeal/ent/student"

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
	// Auth is the client for interacting with the Auth builders.
	Auth *AuthClient
	// Hymn is the client for interacting with the Hymn builders.
	Hymn *HymnClient
	// HymnsWork is the client for interacting with the HymnsWork builders.
	HymnsWork *HymnsWorkClient
	// Role is the client for interacting with the Role builders.
	Role *RoleClient
	// Student is the client for interacting with the Student builders.
	Student *StudentClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Auth = NewAuthClient(c.config)
	c.Hymn = NewHymnClient(c.config)
	c.HymnsWork = NewHymnsWorkClient(c.config)
	c.Role = NewRoleClient(c.config)
	c.Student = NewStudentClient(c.config)
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
		ctx:       ctx,
		config:    cfg,
		Auth:      NewAuthClient(cfg),
		Hymn:      NewHymnClient(cfg),
		HymnsWork: NewHymnsWorkClient(cfg),
		Role:      NewRoleClient(cfg),
		Student:   NewStudentClient(cfg),
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
		ctx:       ctx,
		config:    cfg,
		Auth:      NewAuthClient(cfg),
		Hymn:      NewHymnClient(cfg),
		HymnsWork: NewHymnsWorkClient(cfg),
		Role:      NewRoleClient(cfg),
		Student:   NewStudentClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Auth.
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
	c.Auth.Use(hooks...)
	c.Hymn.Use(hooks...)
	c.HymnsWork.Use(hooks...)
	c.Role.Use(hooks...)
	c.Student.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Auth.Intercept(interceptors...)
	c.Hymn.Intercept(interceptors...)
	c.HymnsWork.Intercept(interceptors...)
	c.Role.Intercept(interceptors...)
	c.Student.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *AuthMutation:
		return c.Auth.mutate(ctx, m)
	case *HymnMutation:
		return c.Hymn.mutate(ctx, m)
	case *HymnsWorkMutation:
		return c.HymnsWork.mutate(ctx, m)
	case *RoleMutation:
		return c.Role.mutate(ctx, m)
	case *StudentMutation:
		return c.Student.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// AuthClient is a client for the Auth schema.
type AuthClient struct {
	config
}

// NewAuthClient returns a client for the Auth from the given config.
func NewAuthClient(c config) *AuthClient {
	return &AuthClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `auth.Hooks(f(g(h())))`.
func (c *AuthClient) Use(hooks ...Hook) {
	c.hooks.Auth = append(c.hooks.Auth, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `auth.Intercept(f(g(h())))`.
func (c *AuthClient) Intercept(interceptors ...Interceptor) {
	c.inters.Auth = append(c.inters.Auth, interceptors...)
}

// Create returns a builder for creating a Auth entity.
func (c *AuthClient) Create() *AuthCreate {
	mutation := newAuthMutation(c.config, OpCreate)
	return &AuthCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Auth entities.
func (c *AuthClient) CreateBulk(builders ...*AuthCreate) *AuthCreateBulk {
	return &AuthCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *AuthClient) MapCreateBulk(slice any, setFunc func(*AuthCreate, int)) *AuthCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &AuthCreateBulk{err: fmt.Errorf("calling to AuthClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*AuthCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &AuthCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Auth.
func (c *AuthClient) Update() *AuthUpdate {
	mutation := newAuthMutation(c.config, OpUpdate)
	return &AuthUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AuthClient) UpdateOne(a *Auth) *AuthUpdateOne {
	mutation := newAuthMutation(c.config, OpUpdateOne, withAuth(a))
	return &AuthUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AuthClient) UpdateOneID(id int64) *AuthUpdateOne {
	mutation := newAuthMutation(c.config, OpUpdateOne, withAuthID(id))
	return &AuthUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Auth.
func (c *AuthClient) Delete() *AuthDelete {
	mutation := newAuthMutation(c.config, OpDelete)
	return &AuthDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AuthClient) DeleteOne(a *Auth) *AuthDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AuthClient) DeleteOneID(id int64) *AuthDeleteOne {
	builder := c.Delete().Where(auth.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AuthDeleteOne{builder}
}

// Query returns a query builder for Auth.
func (c *AuthClient) Query() *AuthQuery {
	return &AuthQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAuth},
		inters: c.Interceptors(),
	}
}

// Get returns a Auth entity by its id.
func (c *AuthClient) Get(ctx context.Context, id int64) (*Auth, error) {
	return c.Query().Where(auth.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AuthClient) GetX(ctx context.Context, id int64) *Auth {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRoles queries the roles edge of a Auth.
func (c *AuthClient) QueryRoles(a *Auth) *RoleQuery {
	query := (&RoleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(auth.Table, auth.FieldID, id),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, auth.RolesTable, auth.RolesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AuthClient) Hooks() []Hook {
	return c.hooks.Auth
}

// Interceptors returns the client interceptors.
func (c *AuthClient) Interceptors() []Interceptor {
	return c.inters.Auth
}

func (c *AuthClient) mutate(ctx context.Context, m *AuthMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AuthCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AuthUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AuthUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AuthDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Auth mutation op: %q", m.Op())
	}
}

// HymnClient is a client for the Hymn schema.
type HymnClient struct {
	config
}

// NewHymnClient returns a client for the Hymn from the given config.
func NewHymnClient(c config) *HymnClient {
	return &HymnClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `hymn.Hooks(f(g(h())))`.
func (c *HymnClient) Use(hooks ...Hook) {
	c.hooks.Hymn = append(c.hooks.Hymn, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `hymn.Intercept(f(g(h())))`.
func (c *HymnClient) Intercept(interceptors ...Interceptor) {
	c.inters.Hymn = append(c.inters.Hymn, interceptors...)
}

// Create returns a builder for creating a Hymn entity.
func (c *HymnClient) Create() *HymnCreate {
	mutation := newHymnMutation(c.config, OpCreate)
	return &HymnCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Hymn entities.
func (c *HymnClient) CreateBulk(builders ...*HymnCreate) *HymnCreateBulk {
	return &HymnCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *HymnClient) MapCreateBulk(slice any, setFunc func(*HymnCreate, int)) *HymnCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &HymnCreateBulk{err: fmt.Errorf("calling to HymnClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*HymnCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &HymnCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Hymn.
func (c *HymnClient) Update() *HymnUpdate {
	mutation := newHymnMutation(c.config, OpUpdate)
	return &HymnUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *HymnClient) UpdateOne(h *Hymn) *HymnUpdateOne {
	mutation := newHymnMutation(c.config, OpUpdateOne, withHymn(h))
	return &HymnUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *HymnClient) UpdateOneID(id int64) *HymnUpdateOne {
	mutation := newHymnMutation(c.config, OpUpdateOne, withHymnID(id))
	return &HymnUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Hymn.
func (c *HymnClient) Delete() *HymnDelete {
	mutation := newHymnMutation(c.config, OpDelete)
	return &HymnDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *HymnClient) DeleteOne(h *Hymn) *HymnDeleteOne {
	return c.DeleteOneID(h.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *HymnClient) DeleteOneID(id int64) *HymnDeleteOne {
	builder := c.Delete().Where(hymn.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &HymnDeleteOne{builder}
}

// Query returns a query builder for Hymn.
func (c *HymnClient) Query() *HymnQuery {
	return &HymnQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeHymn},
		inters: c.Interceptors(),
	}
}

// Get returns a Hymn entity by its id.
func (c *HymnClient) Get(ctx context.Context, id int64) (*Hymn, error) {
	return c.Query().Where(hymn.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *HymnClient) GetX(ctx context.Context, id int64) *Hymn {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUpdatedBy queries the updated_by edge of a Hymn.
func (c *HymnClient) QueryUpdatedBy(h *Hymn) *StudentQuery {
	query := (&StudentClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := h.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(hymn.Table, hymn.FieldID, id),
			sqlgraph.To(student.Table, student.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, hymn.UpdatedByTable, hymn.UpdatedByColumn),
		)
		fromV = sqlgraph.Neighbors(h.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryWork queries the work edge of a Hymn.
func (c *HymnClient) QueryWork(h *Hymn) *HymnsWorkQuery {
	query := (&HymnsWorkClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := h.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(hymn.Table, hymn.FieldID, id),
			sqlgraph.To(hymnswork.Table, hymnswork.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, hymn.WorkTable, hymn.WorkColumn),
		)
		fromV = sqlgraph.Neighbors(h.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *HymnClient) Hooks() []Hook {
	return c.hooks.Hymn
}

// Interceptors returns the client interceptors.
func (c *HymnClient) Interceptors() []Interceptor {
	return c.inters.Hymn
}

func (c *HymnClient) mutate(ctx context.Context, m *HymnMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&HymnCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&HymnUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&HymnUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&HymnDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Hymn mutation op: %q", m.Op())
	}
}

// HymnsWorkClient is a client for the HymnsWork schema.
type HymnsWorkClient struct {
	config
}

// NewHymnsWorkClient returns a client for the HymnsWork from the given config.
func NewHymnsWorkClient(c config) *HymnsWorkClient {
	return &HymnsWorkClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `hymnswork.Hooks(f(g(h())))`.
func (c *HymnsWorkClient) Use(hooks ...Hook) {
	c.hooks.HymnsWork = append(c.hooks.HymnsWork, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `hymnswork.Intercept(f(g(h())))`.
func (c *HymnsWorkClient) Intercept(interceptors ...Interceptor) {
	c.inters.HymnsWork = append(c.inters.HymnsWork, interceptors...)
}

// Create returns a builder for creating a HymnsWork entity.
func (c *HymnsWorkClient) Create() *HymnsWorkCreate {
	mutation := newHymnsWorkMutation(c.config, OpCreate)
	return &HymnsWorkCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of HymnsWork entities.
func (c *HymnsWorkClient) CreateBulk(builders ...*HymnsWorkCreate) *HymnsWorkCreateBulk {
	return &HymnsWorkCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *HymnsWorkClient) MapCreateBulk(slice any, setFunc func(*HymnsWorkCreate, int)) *HymnsWorkCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &HymnsWorkCreateBulk{err: fmt.Errorf("calling to HymnsWorkClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*HymnsWorkCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &HymnsWorkCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for HymnsWork.
func (c *HymnsWorkClient) Update() *HymnsWorkUpdate {
	mutation := newHymnsWorkMutation(c.config, OpUpdate)
	return &HymnsWorkUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *HymnsWorkClient) UpdateOne(hw *HymnsWork) *HymnsWorkUpdateOne {
	mutation := newHymnsWorkMutation(c.config, OpUpdateOne, withHymnsWork(hw))
	return &HymnsWorkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *HymnsWorkClient) UpdateOneID(id int64) *HymnsWorkUpdateOne {
	mutation := newHymnsWorkMutation(c.config, OpUpdateOne, withHymnsWorkID(id))
	return &HymnsWorkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for HymnsWork.
func (c *HymnsWorkClient) Delete() *HymnsWorkDelete {
	mutation := newHymnsWorkMutation(c.config, OpDelete)
	return &HymnsWorkDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *HymnsWorkClient) DeleteOne(hw *HymnsWork) *HymnsWorkDeleteOne {
	return c.DeleteOneID(hw.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *HymnsWorkClient) DeleteOneID(id int64) *HymnsWorkDeleteOne {
	builder := c.Delete().Where(hymnswork.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &HymnsWorkDeleteOne{builder}
}

// Query returns a query builder for HymnsWork.
func (c *HymnsWorkClient) Query() *HymnsWorkQuery {
	return &HymnsWorkQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeHymnsWork},
		inters: c.Interceptors(),
	}
}

// Get returns a HymnsWork entity by its id.
func (c *HymnsWorkClient) Get(ctx context.Context, id int64) (*HymnsWork, error) {
	return c.Query().Where(hymnswork.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *HymnsWorkClient) GetX(ctx context.Context, id int64) *HymnsWork {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryLinkedHymn queries the linked_hymn edge of a HymnsWork.
func (c *HymnsWorkClient) QueryLinkedHymn(hw *HymnsWork) *HymnQuery {
	query := (&HymnClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := hw.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(hymnswork.Table, hymnswork.FieldID, id),
			sqlgraph.To(hymn.Table, hymn.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, hymnswork.LinkedHymnTable, hymnswork.LinkedHymnColumn),
		)
		fromV = sqlgraph.Neighbors(hw.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *HymnsWorkClient) Hooks() []Hook {
	return c.hooks.HymnsWork
}

// Interceptors returns the client interceptors.
func (c *HymnsWorkClient) Interceptors() []Interceptor {
	return c.inters.HymnsWork
}

func (c *HymnsWorkClient) mutate(ctx context.Context, m *HymnsWorkMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&HymnsWorkCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&HymnsWorkUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&HymnsWorkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&HymnsWorkDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown HymnsWork mutation op: %q", m.Op())
	}
}

// RoleClient is a client for the Role schema.
type RoleClient struct {
	config
}

// NewRoleClient returns a client for the Role from the given config.
func NewRoleClient(c config) *RoleClient {
	return &RoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `role.Hooks(f(g(h())))`.
func (c *RoleClient) Use(hooks ...Hook) {
	c.hooks.Role = append(c.hooks.Role, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `role.Intercept(f(g(h())))`.
func (c *RoleClient) Intercept(interceptors ...Interceptor) {
	c.inters.Role = append(c.inters.Role, interceptors...)
}

// Create returns a builder for creating a Role entity.
func (c *RoleClient) Create() *RoleCreate {
	mutation := newRoleMutation(c.config, OpCreate)
	return &RoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Role entities.
func (c *RoleClient) CreateBulk(builders ...*RoleCreate) *RoleCreateBulk {
	return &RoleCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *RoleClient) MapCreateBulk(slice any, setFunc func(*RoleCreate, int)) *RoleCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &RoleCreateBulk{err: fmt.Errorf("calling to RoleClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*RoleCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &RoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Role.
func (c *RoleClient) Update() *RoleUpdate {
	mutation := newRoleMutation(c.config, OpUpdate)
	return &RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RoleClient) UpdateOne(r *Role) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRole(r))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RoleClient) UpdateOneID(id int64) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRoleID(id))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Role.
func (c *RoleClient) Delete() *RoleDelete {
	mutation := newRoleMutation(c.config, OpDelete)
	return &RoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RoleClient) DeleteOne(r *Role) *RoleDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RoleClient) DeleteOneID(id int64) *RoleDeleteOne {
	builder := c.Delete().Where(role.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RoleDeleteOne{builder}
}

// Query returns a query builder for Role.
func (c *RoleClient) Query() *RoleQuery {
	return &RoleQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRole},
		inters: c.Interceptors(),
	}
}

// Get returns a Role entity by its id.
func (c *RoleClient) Get(ctx context.Context, id int64) (*Role, error) {
	return c.Query().Where(role.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RoleClient) GetX(ctx context.Context, id int64) *Role {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryStudent queries the student edge of a Role.
func (c *RoleClient) QueryStudent(r *Role) *StudentQuery {
	query := (&StudentClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, id),
			sqlgraph.To(student.Table, student.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, role.StudentTable, role.StudentColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAuths queries the auths edge of a Role.
func (c *RoleClient) QueryAuths(r *Role) *AuthQuery {
	query := (&AuthClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, id),
			sqlgraph.To(auth.Table, auth.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, role.AuthsTable, role.AuthsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RoleClient) Hooks() []Hook {
	return c.hooks.Role
}

// Interceptors returns the client interceptors.
func (c *RoleClient) Interceptors() []Interceptor {
	return c.inters.Role
}

func (c *RoleClient) mutate(ctx context.Context, m *RoleMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RoleCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RoleDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Role mutation op: %q", m.Op())
	}
}

// StudentClient is a client for the Student schema.
type StudentClient struct {
	config
}

// NewStudentClient returns a client for the Student from the given config.
func NewStudentClient(c config) *StudentClient {
	return &StudentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `student.Hooks(f(g(h())))`.
func (c *StudentClient) Use(hooks ...Hook) {
	c.hooks.Student = append(c.hooks.Student, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `student.Intercept(f(g(h())))`.
func (c *StudentClient) Intercept(interceptors ...Interceptor) {
	c.inters.Student = append(c.inters.Student, interceptors...)
}

// Create returns a builder for creating a Student entity.
func (c *StudentClient) Create() *StudentCreate {
	mutation := newStudentMutation(c.config, OpCreate)
	return &StudentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Student entities.
func (c *StudentClient) CreateBulk(builders ...*StudentCreate) *StudentCreateBulk {
	return &StudentCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *StudentClient) MapCreateBulk(slice any, setFunc func(*StudentCreate, int)) *StudentCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &StudentCreateBulk{err: fmt.Errorf("calling to StudentClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*StudentCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &StudentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Student.
func (c *StudentClient) Update() *StudentUpdate {
	mutation := newStudentMutation(c.config, OpUpdate)
	return &StudentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StudentClient) UpdateOne(s *Student) *StudentUpdateOne {
	mutation := newStudentMutation(c.config, OpUpdateOne, withStudent(s))
	return &StudentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StudentClient) UpdateOneID(id int64) *StudentUpdateOne {
	mutation := newStudentMutation(c.config, OpUpdateOne, withStudentID(id))
	return &StudentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Student.
func (c *StudentClient) Delete() *StudentDelete {
	mutation := newStudentMutation(c.config, OpDelete)
	return &StudentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *StudentClient) DeleteOne(s *Student) *StudentDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *StudentClient) DeleteOneID(id int64) *StudentDeleteOne {
	builder := c.Delete().Where(student.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StudentDeleteOne{builder}
}

// Query returns a query builder for Student.
func (c *StudentClient) Query() *StudentQuery {
	return &StudentQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeStudent},
		inters: c.Interceptors(),
	}
}

// Get returns a Student entity by its id.
func (c *StudentClient) Get(ctx context.Context, id int64) (*Student, error) {
	return c.Query().Where(student.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StudentClient) GetX(ctx context.Context, id int64) *Student {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUpdatedHymns queries the updated_hymns edge of a Student.
func (c *StudentClient) QueryUpdatedHymns(s *Student) *HymnQuery {
	query := (&HymnClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(student.Table, student.FieldID, id),
			sqlgraph.To(hymn.Table, hymn.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, student.UpdatedHymnsTable, student.UpdatedHymnsColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRoledStudent queries the roled_student edge of a Student.
func (c *StudentClient) QueryRoledStudent(s *Student) *RoleQuery {
	query := (&RoleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(student.Table, student.FieldID, id),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, student.RoledStudentTable, student.RoledStudentColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StudentClient) Hooks() []Hook {
	return c.hooks.Student
}

// Interceptors returns the client interceptors.
func (c *StudentClient) Interceptors() []Interceptor {
	return c.inters.Student
}

func (c *StudentClient) mutate(ctx context.Context, m *StudentMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&StudentCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&StudentUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&StudentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&StudentDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Student mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Auth, Hymn, HymnsWork, Role, Student []ent.Hook
	}
	inters struct {
		Auth, Hymn, HymnsWork, Role, Student []ent.Interceptor
	}
)
