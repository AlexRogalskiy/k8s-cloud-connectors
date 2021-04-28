// Code generated by sdkgen. DO NOT EDIT.

//nolint
package postgresql

import (
	"context"

	"google.golang.org/grpc"

	postgresql "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/postgresql/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// DatabaseServiceClient is a postgresql.DatabaseServiceClient with
// lazy GRPC connection initialization.
type DatabaseServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements postgresql.DatabaseServiceClient
func (c *DatabaseServiceClient) Create(ctx context.Context, in *postgresql.CreateDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return postgresql.NewDatabaseServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements postgresql.DatabaseServiceClient
func (c *DatabaseServiceClient) Delete(ctx context.Context, in *postgresql.DeleteDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return postgresql.NewDatabaseServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements postgresql.DatabaseServiceClient
func (c *DatabaseServiceClient) Get(ctx context.Context, in *postgresql.GetDatabaseRequest, opts ...grpc.CallOption) (*postgresql.Database, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return postgresql.NewDatabaseServiceClient(conn).Get(ctx, in, opts...)
}

// List implements postgresql.DatabaseServiceClient
func (c *DatabaseServiceClient) List(ctx context.Context, in *postgresql.ListDatabasesRequest, opts ...grpc.CallOption) (*postgresql.ListDatabasesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return postgresql.NewDatabaseServiceClient(conn).List(ctx, in, opts...)
}

type DatabaseIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *DatabaseServiceClient
	request *postgresql.ListDatabasesRequest

	items []*postgresql.Database
}

func (c *DatabaseServiceClient) DatabaseIterator(ctx context.Context, req *postgresql.ListDatabasesRequest, opts ...grpc.CallOption) *DatabaseIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &DatabaseIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *DatabaseIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Databases
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *DatabaseIterator) Take(size int64) ([]*postgresql.Database, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*postgresql.Database

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *DatabaseIterator) TakeAll() ([]*postgresql.Database, error) {
	return it.Take(0)
}

func (it *DatabaseIterator) Value() *postgresql.Database {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *DatabaseIterator) Error() error {
	return it.err
}

// Update implements postgresql.DatabaseServiceClient
func (c *DatabaseServiceClient) Update(ctx context.Context, in *postgresql.UpdateDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return postgresql.NewDatabaseServiceClient(conn).Update(ctx, in, opts...)
}