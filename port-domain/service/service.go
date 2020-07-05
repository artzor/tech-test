// Package service implements gRPC server
package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"portdomain/entity"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrMissingID = errors.New("missing id")

// Service implements Save/Get API methods
type Service struct {
	server *grpc.Server
	store  store
	port   string
}

type store interface {
	Save(ctx context.Context, pd entity.PortDetails) error
	Get(ctx context.Context, portID string) (entity.PortDetails, error)
}

// Save stores PortDetails
func (s Service) Save(ctx context.Context, details *PortDetails) (*empty.Empty, error) {
	if details.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, ErrMissingID.Error())
	}

	if err := s.store.Save(ctx, pdFromService(details)); err != nil {
		log.Printf("[error] record save: %v", err)
		return nil, err
	}

	return &empty.Empty{}, nil
}

// Get retrieves PortDetail by port ID
func (s Service) Get(ctx context.Context, request *GetRequest) (*PortDetails, error) {
	if request.Id == "" {
		return nil, ErrMissingID
	}

	pd, err := s.store.Get(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return pdToService(pd), nil
}

// New returns new Service instance
func New(store store, port string) *Service {
	return &Service{store: store, port: port}
}

// Start runs gRPC server
func (s *Service) Start() error {
	server := grpc.NewServer()
	listener, err := net.Listen("tcp", s.port)
	if err != nil {
		return fmt.Errorf("net listen: %v", err)
	}

	RegisterPortDomainServer(server, s)
	log.Printf("[info] starting grpc server on port %s", s.port)
	s.server = server
	if err = server.Serve(listener); err != nil {
		return fmt.Errorf("server start: %v", err)
	}

	return nil
}

func (s *Service) Stop() {
	s.server.Stop()
}

func pdToService(pd entity.PortDetails) *PortDetails {
	return &PortDetails{
		Id:       pd.ID,
		Name:     pd.Name,
		City:     pd.City,
		Country:  pd.Country,
		Alias:    pd.Alias,
		Regions:  pd.Regions,
		Coords:   pd.Coords,
		Province: pd.Province,
		Timezone: pd.Timezone,
		Unlocs:   pd.Unlocs,
		Code:     pd.Code,
	}
}

func pdFromService(servicePD *PortDetails) entity.PortDetails {
	return entity.PortDetails{
		ID:       servicePD.Id,
		Name:     servicePD.Name,
		City:     servicePD.City,
		Country:  servicePD.Country,
		Alias:    servicePD.Alias,
		Coords:   servicePD.Coords,
		Province: servicePD.Province,
		Timezone: servicePD.Timezone,
		Unlocs:   servicePD.Unlocs,
		Code:     servicePD.Code,
		Regions:  servicePD.Regions,
	}
}
