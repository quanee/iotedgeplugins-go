package main

import (
	"context"
	"log"
	"net"

	pb "github.com/quanee/iotedgeplugins-go/protobuf/function"
	"google.golang.org/grpc"
)

type PorcessorServer struct {
	pb.UnimplementedProcessorServer
}

func (s *PorcessorServer) ProcessDataOnce(ctx context.Context, in *pb.Drop) (*pb.Drop, error) {
	in.Timestamp = 67890

	return in, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8091")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterProcessorServer(srv, &PorcessorServer{})
	log.Printf("server listening at %v", listen.Addr())

	if err := srv.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
