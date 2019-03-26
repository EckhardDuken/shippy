// consignment-service/main.go
package main

import (

	// Import the generated protobuf code
	"fmt"

	proto "github.com/EckhardDuken/shippy/consignment-service/proto/consignment"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

// Repository interface
type Repository interface {
	Create(*proto.Consignment) (*proto.Consignment, error)
	GetAll() []*proto.Consignment
}

// ConsignmentRepository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
type ConsignmentRepository struct {
	consignments []*proto.Consignment
}

// Create method
func (repo *ConsignmentRepository) Create(consignment *proto.Consignment) (*proto.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

// GetAll method
func (repo *ConsignmentRepository) GetAll() []*proto.Consignment {
	return repo.consignments
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	repo Repository
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) CreateConsignment(ctx context.Context, req *proto.Consignment, res *proto.Response) error {

	// Save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *proto.GetRequest, res *proto.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {

	repo := &ConsignmentRepository{}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	proto.RegisterShippingServiceHandler(srv.Server(), &service{repo})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
