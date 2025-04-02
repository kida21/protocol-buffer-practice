package main

import (
	"context"
	"log"
	"net"
    "google.golang.org/grpc"
	pb"grpcserver/proto/gen"

	
)
type server struct{
 pb.UnimplementedCalculateServer
}
func (s *server) Add(ctx context.Context,req *pb.AddRequest)(*pb.AddResponse,error){

	return &pb.AddResponse{
		Result: req.A+req.B,
	},nil
}
func main() { 
	port:=":50051"
	listen,err:=net.Listen("tcp",port)
	if err!=nil{
		log.Fatal("failed to listen",err)
	}
	grpcServer:=grpc.NewServer()
	pb.RegisterCalculateServer(grpcServer,&server{})
	log.Println("Server running on..",port)
	err=grpcServer.Serve(listen)
	if err!=nil{
		log.Fatal("Failed to serve")
	}
	

}