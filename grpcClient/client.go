package main

import (
	"context"
	pb "grpcclient/proto/gen"
	farewellpb "grpcclient/proto/gen/farewell"
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
	client2:=pb.NewGreeterClient(conn)
    client3:=farewellpb.NewFarewellClient(conn)

	values:=&pb.AddRequest{
       A: 20,
	   B: 20,
	}
	result,err:=client.Add(ctx,values)
	if err!=nil{
		log.Fatal("unable to do the operation",err)
	}
	log.Println("Result is:",result.Result)
	name:=&pb.HelloRequest{
		Name: "Johnathan",
	}


	response,err:=client2.Greet(ctx,name)
	if err!=nil{
		log.Fatal("unable to greet",err)
	}
	log.Println("Response is:",response.Mesage)
     
	request:=&farewellpb.GoodByeRequest{
		Name: "Thomas",
	}
    res,err:=client3.GoodBye(ctx,request)
	if err!=nil{
		log.Fatal("unable to do the operation")
	}
	log.Println(res.Message)
}