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
	deviceDrop := pb.DeviceDataSet{
		Items: []*pb.DeviceDataSetItem{
			{
				DeviceID: "device1",
				Properties: map[string]*pb.Properties{
					"dev1prop1": {
						Datatype:   pb.DataType_Int,
						Properties: []string{"11", "12", "13", "14", "15"},
					},
					"dev1prop2": {
						Datatype:   pb.DataType_Int,
						Properties: []string{"16", "17", "18", "19", "10"},
					},
				},
			},
			{
				DeviceID: "device2",
				Properties: map[string]*pb.Properties{
					"dev2prop1": {
						Datatype:   pb.DataType_Int,
						Properties: []string{"21", "22", "23", "24", "25"},
					},
					"dev2prop2": {
						Datatype:   pb.DataType_Int,
						Properties: []string{"26", "27", "28", "29", "20"},
					},
				},
			},
		},
	}

	ddd := &pb.DataSet_DeviceDataSet{
		DeviceDataSet: &deviceDrop,
	}
	drop := pb.DataSet{
		DataSet: ddd,
	}
	for i := 0; ; i++ {
		drop.Timestamp = time.Now().Unix()

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
