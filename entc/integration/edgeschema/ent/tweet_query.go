// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/tag"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweet"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweetlike"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweettag"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
	"entgo.io/ent/entc/integration/edgeschema/ent/usertweet"
	"entgo.io/ent/schema/field"
)

// TweetQuery is the builder for querying Tweet entities.
type TweetQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Tweet
	// eager-loading edges.
	withLikedUsers *UserQuery
	withUser       *UserQuery
	withTags       *TagQuery
	withLikes      *TweetLikeQuery
	withTweetUser  *UserTweetQuery
	withTweetTags  *TweetTagQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TweetQuery builder.
func (tq *TweetQuery) Where(ps ...predicate.Tweet) *TweetQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit adds a limit step to the query.
func (tq *TweetQuery) Limit(limit int) *TweetQuery {
	tq.limit = &limit
	return tq
}

// Offset adds an offset step to the query.
func (tq *TweetQuery) Offset(offset int) *TweetQuery {
	tq.offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TweetQuery) Unique(unique bool) *TweetQuery {
	tq.unique = &unique
	return tq
}

// Order adds an order step to the query.
func (tq *TweetQuery) Order(o ...OrderFunc) *TweetQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryLikedUsers chains the current query on the "liked_users" edge.
func (tq *TweetQuery) QueryLikedUsers() *UserQuery {
	query := &UserQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tweet.Table, tweet.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tweet.LikedUsersTable, tweet.LikedUsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (tq *TweetQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tweet.Table, tweet.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tweet.UserTable, tweet.UserPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTags chains the current query on the "tags" edge.
func (tq *TweetQuery) QueryTags() *TagQuery {
	query := &TagQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tweet.Table, tweet.FieldID, selector),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tweet.TagsTable, tweet.TagsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLikes chains the current query on the "likes" edge.
func (tq *TweetQuery) QueryLikes() *TweetLikeQuery {
	query := &TweetLikeQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tweet.Table, tweet.FieldID, selector),
			sqlgraph.To(tweetlike.Table, tweetlike.TweetColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, tweet.LikesTable, tweet.LikesColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTweetUser chains the current query on the "tweet_user" edge.
func (tq *TweetQuery) QueryTweetUser() *UserTweetQuery {
	query := &UserTweetQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tweet.Table, tweet.FieldID, selector),
			sqlgraph.To(usertweet.Table, usertweet.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, tweet.TweetUserTable, tweet.TweetUserColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTweetTags chains the current query on the "tweet_tags" edge.
func (tq *TweetQuery) QueryTweetTags() *TweetTagQuery {
	query := &TweetTagQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tweet.Table, tweet.FieldID, selector),
			sqlgraph.To(tweettag.Table, tweettag.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, tweet.TweetTagsTable, tweet.TweetTagsColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Tweet entity from the query.
// Returns a *NotFoundError when no Tweet was found.
func (tq *TweetQuery) First(ctx context.Context) (*Tweet, error) {
	nodes, err := tq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tweet.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TweetQuery) FirstX(ctx context.Context) *Tweet {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Tweet ID from the query.
// Returns a *NotFoundError when no Tweet ID was found.
func (tq *TweetQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tweet.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TweetQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Tweet entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Tweet entity is found.
// Returns a *NotFoundError when no Tweet entities are found.
func (tq *TweetQuery) Only(ctx context.Context) (*Tweet, error) {
	nodes, err := tq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tweet.Label}
	default:
		return nil, &NotSingularError{tweet.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TweetQuery) OnlyX(ctx context.Context) *Tweet {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Tweet ID in the query.
// Returns a *NotSingularError when more than one Tweet ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TweetQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tweet.Label}
	default:
		err = &NotSingularError{tweet.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TweetQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Tweets.
func (tq *TweetQuery) All(ctx context.Context) ([]*Tweet, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tq *TweetQuery) AllX(ctx context.Context) []*Tweet {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Tweet IDs.
func (tq *TweetQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := tq.Select(tweet.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TweetQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TweetQuery) Count(ctx context.Context) (int, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TweetQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TweetQuery) Exist(ctx context.Context) (bool, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TweetQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TweetQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TweetQuery) Clone() *TweetQuery {
	if tq == nil {
		return nil
	}
	return &TweetQuery{
		config:         tq.config,
		limit:          tq.limit,
		offset:         tq.offset,
		order:          append([]OrderFunc{}, tq.order...),
		predicates:     append([]predicate.Tweet{}, tq.predicates...),
		withLikedUsers: tq.withLikedUsers.Clone(),
		withUser:       tq.withUser.Clone(),
		withTags:       tq.withTags.Clone(),
		withLikes:      tq.withLikes.Clone(),
		withTweetUser:  tq.withTweetUser.Clone(),
		withTweetTags:  tq.withTweetTags.Clone(),
		// clone intermediate query.
		sql:    tq.sql.Clone(),
		path:   tq.path,
		unique: tq.unique,
	}
}

// WithLikedUsers tells the query-builder to eager-load the nodes that are connected to
// the "liked_users" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TweetQuery) WithLikedUsers(opts ...func(*UserQuery)) *TweetQuery {
	query := &UserQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withLikedUsers = query
	return tq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TweetQuery) WithUser(opts ...func(*UserQuery)) *TweetQuery {
	query := &UserQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withUser = query
	return tq
}

// WithTags tells the query-builder to eager-load the nodes that are connected to
// the "tags" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TweetQuery) WithTags(opts ...func(*TagQuery)) *TweetQuery {
	query := &TagQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withTags = query
	return tq
}

// WithLikes tells the query-builder to eager-load the nodes that are connected to
// the "likes" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TweetQuery) WithLikes(opts ...func(*TweetLikeQuery)) *TweetQuery {
	query := &TweetLikeQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withLikes = query
	return tq
}

// WithTweetUser tells the query-builder to eager-load the nodes that are connected to
// the "tweet_user" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TweetQuery) WithTweetUser(opts ...func(*UserTweetQuery)) *TweetQuery {
	query := &UserTweetQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withTweetUser = query
	return tq
}

// WithTweetTags tells the query-builder to eager-load the nodes that are connected to
// the "tweet_tags" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TweetQuery) WithTweetTags(opts ...func(*TweetTagQuery)) *TweetQuery {
	query := &TweetTagQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withTweetTags = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Tweet.Query().
//		GroupBy(tweet.FieldText).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tq *TweetQuery) GroupBy(field string, fields ...string) *TweetGroupBy {
	grbuild := &TweetGroupBy{config: tq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tq.sqlQuery(ctx), nil
	}
	grbuild.label = tweet.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//	}
//
//	client.Tweet.Query().
//		Select(tweet.FieldText).
//		Scan(ctx, &v)
//
func (tq *TweetQuery) Select(fields ...string) *TweetSelect {
	tq.fields = append(tq.fields, fields...)
	selbuild := &TweetSelect{TweetQuery: tq}
	selbuild.label = tweet.Label
	selbuild.flds, selbuild.scan = &tq.fields, selbuild.Scan
	return selbuild
}

func (tq *TweetQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tq.fields {
		if !tweet.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TweetQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Tweet, error) {
	var (
		nodes       = []*Tweet{}
		_spec       = tq.querySpec()
		loadedTypes = [6]bool{
			tq.withLikedUsers != nil,
			tq.withUser != nil,
			tq.withTags != nil,
			tq.withLikes != nil,
			tq.withTweetUser != nil,
			tq.withTweetTags != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Tweet).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Tweet{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := tq.withLikedUsers; query != nil {
		edgeids := make([]driver.Value, len(nodes))
		byid := make(map[int]*Tweet)
		nids := make(map[int]map[*Tweet]struct{})
		for i, node := range nodes {
			edgeids[i] = node.ID
			byid[node.ID] = node
			node.Edges.LikedUsers = []*User{}
		}
		query.Where(func(s *sql.Selector) {
			joinT := sql.Table(tweet.LikedUsersTable)
			s.Join(joinT).On(s.C(user.FieldID), joinT.C(tweet.LikedUsersPrimaryKey[0]))
			s.Where(sql.InValues(joinT.C(tweet.LikedUsersPrimaryKey[1]), edgeids...))
			columns := s.SelectedColumns()
			s.Select(joinT.C(tweet.LikedUsersPrimaryKey[1]))
			s.AppendSelect(columns...)
			s.SetDistinct(false)
		})
		neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]interface{}, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]interface{}{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []interface{}) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Tweet]struct{}{byid[outValue]: struct{}{}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byid[outValue]] = struct{}{}
				return nil
			}
		})
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "liked_users" node returned %v`, n.ID)
			}
			for kn := range nodes {
				kn.Edges.LikedUsers = append(kn.Edges.LikedUsers, n)
			}
		}
	}

	if query := tq.withUser; query != nil {
		edgeids := make([]driver.Value, len(nodes))
		byid := make(map[int]*Tweet)
		nids := make(map[int]map[*Tweet]struct{})
		for i, node := range nodes {
			edgeids[i] = node.ID
			byid[node.ID] = node
			node.Edges.User = []*User{}
		}
		query.Where(func(s *sql.Selector) {
			joinT := sql.Table(tweet.UserTable)
			s.Join(joinT).On(s.C(user.FieldID), joinT.C(tweet.UserPrimaryKey[0]))
			s.Where(sql.InValues(joinT.C(tweet.UserPrimaryKey[1]), edgeids...))
			columns := s.SelectedColumns()
			s.Select(joinT.C(tweet.UserPrimaryKey[1]))
			s.AppendSelect(columns...)
			s.SetDistinct(false)
		})
		neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]interface{}, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]interface{}{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []interface{}) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Tweet]struct{}{byid[outValue]: struct{}{}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byid[outValue]] = struct{}{}
				return nil
			}
		})
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "user" node returned %v`, n.ID)
			}
			for kn := range nodes {
				kn.Edges.User = append(kn.Edges.User, n)
			}
		}
	}

	if query := tq.withTags; query != nil {
		edgeids := make([]driver.Value, len(nodes))
		byid := make(map[int]*Tweet)
		nids := make(map[int]map[*Tweet]struct{})
		for i, node := range nodes {
			edgeids[i] = node.ID
			byid[node.ID] = node
			node.Edges.Tags = []*Tag{}
		}
		query.Where(func(s *sql.Selector) {
			joinT := sql.Table(tweet.TagsTable)
			s.Join(joinT).On(s.C(tag.FieldID), joinT.C(tweet.TagsPrimaryKey[0]))
			s.Where(sql.InValues(joinT.C(tweet.TagsPrimaryKey[1]), edgeids...))
			columns := s.SelectedColumns()
			s.Select(joinT.C(tweet.TagsPrimaryKey[1]))
			s.AppendSelect(columns...)
			s.SetDistinct(false)
		})
		neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]interface{}, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]interface{}{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []interface{}) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Tweet]struct{}{byid[outValue]: struct{}{}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byid[outValue]] = struct{}{}
				return nil
			}
		})
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "tags" node returned %v`, n.ID)
			}
			for kn := range nodes {
				kn.Edges.Tags = append(kn.Edges.Tags, n)
			}
		}
	}

	if query := tq.withLikes; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Tweet)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Likes = []*TweetLike{}
		}
		query.Where(predicate.TweetLike(func(s *sql.Selector) {
			s.Where(sql.InValues(tweet.LikesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.TweetID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "tweet_id" returned %v for node %v`, fk, n)
			}
			node.Edges.Likes = append(node.Edges.Likes, n)
		}
	}

	if query := tq.withTweetUser; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Tweet)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.TweetUser = []*UserTweet{}
		}
		query.Where(predicate.UserTweet(func(s *sql.Selector) {
			s.Where(sql.InValues(tweet.TweetUserColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.TweetID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "tweet_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.TweetUser = append(node.Edges.TweetUser, n)
		}
	}

	if query := tq.withTweetTags; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Tweet)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.TweetTags = []*TweetTag{}
		}
		query.Where(predicate.TweetTag(func(s *sql.Selector) {
			s.Where(sql.InValues(tweet.TweetTagsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.TweetID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "tweet_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.TweetTags = append(node.Edges.TweetTags, n)
		}
	}

	return nodes, nil
}

func (tq *TweetQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.fields
	if len(tq.fields) > 0 {
		_spec.Unique = tq.unique != nil && *tq.unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TweetQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tq *TweetQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tweet.Table,
			Columns: tweet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tweet.FieldID,
			},
		},
		From:   tq.sql,
		Unique: true,
	}
	if unique := tq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tweet.FieldID)
		for i := range fields {
			if fields[i] != tweet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TweetQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(tweet.Table)
	columns := tq.fields
	if len(columns) == 0 {
		columns = tweet.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.unique != nil && *tq.unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TweetGroupBy is the group-by builder for Tweet entities.
type TweetGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TweetGroupBy) Aggregate(fns ...AggregateFunc) *TweetGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tgb *TweetGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tgb.path(ctx)
	if err != nil {
		return err
	}
	tgb.sql = query
	return tgb.sqlScan(ctx, v)
}

func (tgb *TweetGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tgb.fields {
		if !tweet.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tgb *TweetGroupBy) sqlQuery() *sql.Selector {
	selector := tgb.sql.Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tgb.fields)+len(tgb.fns))
		for _, f := range tgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tgb.fields...)...)
}

// TweetSelect is the builder for selecting fields of Tweet entities.
type TweetSelect struct {
	*TweetQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TweetSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	ts.sql = ts.TweetQuery.sqlQuery(ctx)
	return ts.sqlScan(ctx, v)
}

func (ts *TweetSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ts.sql.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}