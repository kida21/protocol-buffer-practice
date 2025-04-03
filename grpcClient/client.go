package main

import (
	"context"
	pb "grpcclient/proto/gen"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn,err:=grpc.NewClient("localhost:50051",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Fatal("Failed to connect",err)
	}
	defer conn.Close()
	ctx,cancel:=context.WithTimeout(context.Background(),3*time.Second)
	defer cancel()
	client:=pb.NewCalculateClient(conn)
	values:=&pb.AddRequest{
       A: 10,
	   B: 20,
	}
	result,err:=client.Add(ctx,values)
	if err!=nil{
		log.Fatal("unable to do the operation",err)
	}
	log.Println("Result is:",result.Result)
}