package main

import (
	"context"
	pb "grpcserver/proto/gen"
	farewellpb "grpcserver/proto/gen/farewell"
	"log"
	"net"

	"google.golang.org/grpc"
)
type server struct{
 pb.UnimplementedCalculateServer
 pb.UnimplementedGreeterServer
 farewellpb.UnimplementedFarewellServer
}
func (s *server) Add(ctx context.Context,req *pb.AddRequest)(*pb.AddResponse,error){
     sum:=req.A + req.B
	 log.Println("sum is:",sum)
	return &pb.AddResponse{
		Result: req.A+req.B,
	},nil
}
func (s *server) Greet(ctx context.Context,req *pb.HelloRequest)(*pb.HelloResponse,error){
	return &pb.HelloResponse{
		Mesage: req.Name,
	},nil
}

func (s *server) GoodBye(ctx context.Context,req *farewellpb.GoodByeRequest)(*farewellpb.GoodByeResponse,error){
	goodbye:="Good Bye :"+" "+req.Name
	return &farewellpb.GoodByeResponse{
		Message: goodbye,
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
	pb.RegisterGreeterServer(grpcServer,&server{})
	farewellpb.RegisterFarewellServer(grpcServer,&server{})
	log.Println("Server running on..",port)
	err=grpcServer.Serve(listen)
	if err!=nil{
		log.Fatal("Failed to serve")
	}
	

}