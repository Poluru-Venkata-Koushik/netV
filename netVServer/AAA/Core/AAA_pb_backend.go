package pb

import (
	"fmt"
	jwt "netVServer/AAA/JWT"
	pb "netVServer/AAA/pb"
	Logging "netVServer/LogPackage"
)

type Server struct{
	pb.UnimplementedAAAServer
	Map  map[string]string
}

const (
	LogFile = "AAA.log"
	PKG     = "AAA: gRPC"
)

var (
	log = Logging.Logger{Filename: LogFile, PKG: PKG}
)

func (S *Server) CreateToken (Username *pb.Username)(pb.JWT, error) {
	log.InfoLog("Create Token Invoked.")
	if val, ok := S.Map[Username.Name]; ok {
		log.ErrorLog(" Token exists for the user ")
	}
	string, error := jwt.CreateToken(Username.Name, Username.Role)
	if error!=nil{
		return pb.JWT{JWT: ""}, error
	}
	return pb.JWT{JWT: string}, nil

}