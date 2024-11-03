package main

import (
	"context"
	"log"
	"net"

	pb "github.com/jokestax/coffeeshop/proto_files"
	grpc "google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCoffeeshopServer
}

func (s *server) GetMenu(menuRequest *pb.MenuRequest, srv grpc.ServerStreamingServer[pb.Menu]) error {
	items := []*pb.Item{
		&pb.Item{
			Id:   "1",
			Name: "Black Coffee",
		},
		&pb.Item{
			Id:   "2",
			Name: "Americano",
		},
	}

	for i, _ := range items {
		srv.Send(&pb.Menu{
			Items: items[0 : i+1],
		})
	}

	return nil

}

func (s *server) PlaceOrder(ctx context.Context, order *pb.Order) (*pb.Receipt, error) {
	return &pb.Receipt{
		Id: "ABC123",
	}, nil

}

func (s *server) GetOrderStatus(ctx context.Context, receipt *pb.Receipt) (*pb.OrderStatus, error) {
	return &pb.OrderStatus{
		OrderId: receipt.Id,
		Status:  "placed",
	}, nil
}

func main() {
	list, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCoffeeshopServer(grpcServer, &server{})
	if err = grpcServer.Serve(list); err != nil {
		log.Fatalf(err.Error())
		return
	}
}
