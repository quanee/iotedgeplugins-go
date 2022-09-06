package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/quanee/iotedgeplugins-go/protobuf/function"
	"google.golang.org/grpc"
)

type PorcessorServer struct {
	pb.UnimplementedProcessorServer
}

func (s *PorcessorServer) ProcessDataOnce(ctx context.Context, in *pb.DataSet) (*pb.DataSet, error) {
	println("Process Data Once")

	if in.GetEventDataSet() != nil {
		for _, event := range in.GetEventDataSet().ReportEvents {
			event.Value = "hehehe"
		}
	} else if in.GetDeviceDataSet() != nil {
		for _, item := range in.GetDeviceDataSet().Items {
			for name, props := range item.Properties {
				for i := range props.Properties {
					prop := item.Properties[name].Properties[i]
					fmt.Println(prop)
					if name == "time" {
						prop = string(append([]byte(prop)[:len(prop)-3], []byte("666")...))
					}
					item.Properties[name].Properties[i] = prop
				}
			}
		}
	}

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
