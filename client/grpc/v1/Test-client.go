package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"

	v1 "Test-output/proto"
)

func main() {
	// Get configuration
	address := flag.String("server", "localhost:4040", "gRPC server in format host:port")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewTestServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call ReadAll

	req0 := v1.ReadAllRequest{

		V1: 1,
	}
	res0, err := c.ReadAll(ctx, &req0)
	if err != nil {
		log.Fatalf("Call ReadAll failed: %v", err)
	}
	log.Printf("Call ReadAll result: &lt;%+v>\n\n", res0)

	// Call Create
	var toDo1 v1.ToDo
	req1 := v1.CreateRequest{

		V1:   1,
		ToDo: &toDo1,
	}
	res1, err := c.Create(ctx, &req1)
	if err != nil {
		log.Fatalf("Call Create failed: %v", err)
	}
	log.Printf("Call Create result: &lt;%+v>\n\n", res1)

	// Call Update
	var toDo2 v1.ToDo
	req2 := v1.UpdateRequest{

		V1:   1,
		Id:   1,
		ToDo: &toDo2,
	}
	res2, err := c.Update(ctx, &req2)
	if err != nil {
		log.Fatalf("Call Update failed: %v", err)
	}
	log.Printf("Call Update result: &lt;%+v>\n\n", res2)

	// Call Delete

	req3 := v1.DeleteRequest{

		V1: 1,
		Id: 1,
	}
	res3, err := c.Delete(ctx, &req3)
	if err != nil {
		log.Fatalf("Call Delete failed: %v", err)
	}
	log.Printf("Call Delete result: &lt;%+v>\n\n", res3)

	// Call Read

	req4 := v1.ReadRequest{

		V1: 1,
		Id: 1,
	}
	res4, err := c.Read(ctx, &req4)
	if err != nil {
		log.Fatalf("Call Read failed: %v", err)
	}
	log.Printf("Call Read result: &lt;%+v>\n\n", res4)

}
