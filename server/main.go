package main

import (
	"content"
	"log"
	"net"

	pb "github.com/Aadhu23/go-grpc-docker/proto"
	"google.golang.org/grpc"
)
type greeterServer struct {
	pb.UnimplementedGreeeterServer
}

func (s *greeterServer) SayHello(ctx context.Content, req *pb.HelloRequest) (*pb.HelloResponse,error){
	log.Printf("recived request fo namr: %s",req.name)
	return &pb.HelloResponse{Message:"Hello"+req.name +"!"},nil 

}

func main(){
	lis,err := net.Lisent("tcp",":50051")
	if err !=nil{
		log.Fatalf("failed to lisent on port 50051: %v",err)
	}

	s:=grpc.NewServer()
	pb.RegisterGreeterServer(s, &greeterServer{})

	log.Println("grpc serveer listening on port 50051...")
	if err:=s.Server(lis);err !=nil{
		log.Fatalf("failed to serve grpc server: %v",err)
	}
}
