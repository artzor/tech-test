package service

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	"github.com/artzor/tech-test/port-domain/api"
	"github.com/artzor/tech-test/port-domain/entity"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) Save(ctx context.Context, pd entity.PortDetails) error {
	args := m.Called(ctx, pd)
	return args.Error(0)
}

func (m *MockStore) Get(ctx context.Context, portID string) (entity.PortDetails, error) {
	args := m.Called(ctx, portID)
	return args.Get(0).(entity.PortDetails), args.Error(1)
}

func TestService_Get(t *testing.T) {

}

func TestService_Save(t *testing.T) {
	store := &MockStore{}

	dialerCtx, srv := start(store)
	defer srv.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(dialerCtx), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()

	client := api.NewPortDomainClient(conn)
	tt := []struct {
		pd      entity.PortDetails
		wantErr string
	}{
		{pd: entity.PortDetails{}, wantErr: "rpc error: code = InvalidArgument desc = missing id"},
		{pd: entity.PortDetails{ID: "ID", Name: "Name"}},
	}

	for _, test := range tt {
		store.On("Save", mock.Anything, mock.Anything).Return(nil)
		_, err := client.Save(context.Background(), pdToService(test.pd))

		if test.wantErr != "" {
			require.EqualError(t, err, test.wantErr)
			continue
		}

		require.NoError(t, err)
	}
}

const bufSize = 1024 * 1024

type ctxFunc = func(context.Context, string) (net.Conn, error)

func start(store store) (ctxFunc, *grpc.Server) {
	var lis *bufconn.Listener

	bufDialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	svc := &Service{
		store: store,
	}
	api.RegisterPortDomainServer(s, svc)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	return bufDialer, s
}
