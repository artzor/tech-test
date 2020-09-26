// Package portdomain provides a client for external service
package portdomain

import (
	"context"

	portDomainApi "github.com/artzor/tech-test/client-api/portdomain/api"

	"github.com/artzor/tech-test/client-api/entity"

	"google.golang.org/grpc"
)

// Client exposes external service APIs
type Client struct {
	client     portDomainApi.PortDomainClient
	connection *grpc.ClientConn
	serviceURL string
}

// New returns new Client instance
func New(serviceURL string) *Client {
	return &Client{serviceURL: serviceURL}
}

// Connect establishes connection with remote server
func (c *Client) Connect() error {
	conn, err := grpc.Dial(c.serviceURL, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}

	c.client = portDomainApi.NewPortDomainClient(conn)
	c.connection = conn
	return nil
}

// Disconnect closes connection to gRPC server
func (c *Client) Disconnect() error {
	return c.connection.Close()
}

// Save calls service's Save method
func (c *Client) Save(ctx context.Context, portDetails entity.PortDetails) error {
	_, err := c.client.Save(ctx, pdToService(portDetails))
	return err
}

// Get calls service's Get method
func (c *Client) Get(ctx context.Context, portID string) (entity.PortDetails, error) {
	servicePD, err := c.client.Get(ctx, &portDomainApi.GetRequest{Id: portID})
	if err != nil {
		return entity.PortDetails{}, err
	}
	return pdFromService(servicePD), nil
}

func pdToService(pd entity.PortDetails) *portDomainApi.PortDetails {
	return &portDomainApi.PortDetails{
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

func pdFromService(servicePD *portDomainApi.PortDetails) entity.PortDetails {
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
