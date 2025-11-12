package pb

import (
	"context"
	jwt "netVServer/AAA/JWT"
	Logging "netVServer/LogPackage"
)

type Server struct {
	UnimplementedAAAServer
	Map map[string]string
}

const (
	LogFile = "AAA.log"
	PKG     = "AAA: gRPC"
)

var (
	log = Logging.Logger{Filename: LogFile, PKG: PKG}
)

func NewUserServer() *Server {
	return &Server{
		Map: make(map[string]string),
	}
}

func (S *Server) CreateToken(ctx context.Context, Username *Username) (*JWT, error) {
	log.InfoLog("Create Token Invoked.")
	if val, ok := S.Map[Username.Name]; ok {
		log.ErrorLog(" Token exists for the user ")
		if jwt.IsValidToken(val) {
			log.InfoLog("Token is still valid")
			return &JWT{JWT: val}, nil
		}
	}
	string, error := jwt.CreateToken(Username.Name, Username.Role)
	if error != nil {
		return &JWT{JWT: ""}, error
	}
	return &JWT{JWT: string}, nil
}
