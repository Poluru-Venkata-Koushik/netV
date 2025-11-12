package main

import (
	"fmt"
	"net"
	pb "netVServer/AAA/pb"
	Logging "netVServer/LogPackage"

	"google.golang.org/grpc"
)

const(
	Protocol = "tcp"
	port = ":9876"
	LogFile = "Server.log"
	PKG     = "SERVER"
)

var (
	log = Logging.Logger{Filename: LogFile, PKG: PKG}
)

func main() {
	list, err := net.Listen(Protocol, port)
	if err!=nil{
		log.ErrorLog(fmt.Sprintf("Cannot listen on port %v", port))
	}else{
		log.PrintLog(fmt.Sprintf("Started listening on port %v", port))
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAAAServer(grpcServer, pb.NewUserServer())
	log.InfoLog("Starting gRPC Server")
	if err:= grpcServer.Serve(list); err!=nil{
		log.ErrorLog(err.Error())
	}


}