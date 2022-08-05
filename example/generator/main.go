package main

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/quanee/iotedgeplugins-go/protobuf/function"
	"google.golang.org/grpc"
)

type GeneratorServer struct {
	pb.UnimplementedGeneratorServer
}

// generator function
func (s *GeneratorServer) SubscribeData(req *pb.GeneratorRequest, srv pb.Generator_SubscribeDataServer) error {
	fmt.Println("enter sub")
	deviceDrop := pb.DeviceDrop{
		Items: []*pb.DeviceDropItem{},
	}
	drop_item := pb.DeviceDropItem{
		Properties: map[string]*pb.Properties{
			"p1": {
				Datatype:   pb.DataType_Int,
				Properties: []string{"1", "2", "3", "4", "5"},
			},
		},
	}
	deviceDrop.Items = append(deviceDrop.Items, &drop_item)

	ddd := &pb.Drop_DeviceDrop{
		DeviceDrop: &deviceDrop,
	}
	drop := pb.Drop{
		Drop: ddd,
	}
	for n := 0; ; n++ {
		drop.Timestamp = int64(n)

		srv.Send(&drop)
		time.Sleep(time.Second)
	}
}

func (s *GeneratorServer) Recv() {}

func main() {
	listen, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterGeneratorServer(srv, &GeneratorServer{})
	log.Printf("server listening at %v", listen.Addr())

	if err := srv.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
