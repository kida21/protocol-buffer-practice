package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	port:=":50051"
	listen,err:=net.Listen("tcp",port)
	if err!=nil{
		log.Fatal("failed to listen",err)
	}
	grpcServer:=grpc.NewServer()
	log.Println("Server running on..",port)
	err=grpcServer.Serve(listen)
	if err!=nil{
		log.Fatal("Failed to serve")
	}
	

}