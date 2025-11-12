package main

import (
	"context"
	"fmt"
	pb "netVServer/AAA/pb"
	Logging "netVServer/LogPackage"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
const(
	ContextTimeout = 2
	address = "localhost:9998"
	LogFile = "Client.log"
	PKG     = "Client"
)

var (
	log = Logging.Logger{Filename: LogFile, PKG: PKG}
)


func main(){
	ctx, cancel := context.WithTimeout(context.Background(), ContextTimeout * time.Second)
	defer cancel()

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.ErrorLog("Cannot Start a new GRPC client")
		log.ErrorLog(err.Error())
	}
	defer conn.Close()
	client := pb.NewAAAClient(conn)

	val, err := client.ExtractToken(ctx, &pb.JWT{JWT: "Koushik"})
	if err!=nil{
		log.ErrorLog(err.Error())
	}
	log.PrintLog(fmt.Sprintf(" User : %v, Role : %v", val.Name, val.Role))

	val, err = client.CreateToken(ctx, &pb.Username{Name: "User1", Role: "ADMIN"})
}