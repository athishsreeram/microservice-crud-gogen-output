package v1

import (
	"Test-output/domain"
	"Test-output/proto"
	"context"
	"log"
	"strconv"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// testServer is implementation of proto.TestServer proto interface
type testServer struct {
}

// NewTestServiceServer creates Test service
func NewTestServiceServer() proto.TestServiceServer {
	return &testServer{}
}

// Create new todo task
func (s *testServer) ReadAll(ctx context.Context, req *proto.ReadAllRequest) (*proto.ReadAllResponse, error) {
	// TO-DO

	var toDos []*proto.ToDo
	for _, toDoDomains := range domain.ReadAllToDoDomains() {
		toDo := domain.ConvertToDoDomains2ToDo(domain.ReadToDoDomains(toDoDomains.Sno))
		toDos = append(toDos, toDo)
	}

	list := []*proto.ToDo{}
	resp := &proto.ReadAllResponse{
		V1: 1,

		ToDos: list,
	}
	return resp, nil
}

// Create new todo task
func (s *testServer) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	// TO-DO

	var toDo *proto.ToDo
	log.Println("DTO Data", req.ToDo)
	toDoDomains := domain.ConvertToDo2ToDoDomains(req.ToDo)
	log.Println("Domain Data ", toDoDomains)
	domain.CreateToDoDomains(toDoDomains)
	toDo = domain.ConvertToDoDomains2ToDo(domain.ReadToDoDomains(toDoDomains.Sno))
	log.Println("Response DTO Data", toDo)

	resp := &proto.CreateResponse{
		V1: 1,

		Id: 1,
	}
	return resp, nil
}

// Create new todo task
func (s *testServer) Update(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	// TO-DO

	var toDo *proto.ToDo
	log.Println("DTO Data", req.Id)
	log.Println("DTO Data", req.ToDo)
	toDoDomains := domain.ConvertToDo2ToDoDomains(req.ToDo)
	log.Println("Domain Data ", toDoDomains)
	domain.UpdateToDoDomains(toDoDomains.Sno, toDoDomains)
	toDo = domain.ConvertToDoDomains2ToDo(domain.ReadToDoDomains(toDoDomains.Sno))
	log.Println("Response DTO Data", toDo)

	resp := &proto.UpdateResponse{
		V1: 1,

		Updated: 1,
	}
	return resp, nil
}

// Create new todo task
func (s *testServer) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	// TO-DO

	log.Println("DTO Data", req.Id)
	n, err := strconv.Atoi(strconv.FormatInt(req.Id, 10))
	if err == nil {
		domain.DeleteToDoDomains(n)
	}

	resp := &proto.DeleteResponse{
		V1: 1,

		Deleted: true,
	}
	return resp, nil
}

// Create new todo task
func (s *testServer) Read(ctx context.Context, req *proto.ReadRequest) (*proto.ReadResponse, error) {
	// TO-DO

	var toDo *proto.ToDo
	n, err := strconv.Atoi(strconv.FormatInt(req.Id, 10))
	if err == nil {
		toDo = domain.ConvertToDoDomains2ToDo(domain.ReadToDoDomains(n))
	}

	resp := &proto.ReadResponse{
		V1: 1,

		ToDo: toDo,
	}
	return resp, nil
}
