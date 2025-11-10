package jwt

import (
	jwt "netVServer/AAA/JWT"
	Logging "netVServer/LogPackage"
	"time"

	JWT "github.com/golang-jwt/jwt/v5"
)
const (
	LogFile = "AAA.log"
	PKG     = "AAA::JWT"
)

var (
	log = Logging.Logger{Filename: LogFile, PKG: PKG}
	secret_key = []byte("ThisIsMySecret")                // We can later ensure this is constantly rotated every day and we can have 3 days worth of keys in our cache and DB
)

type JWT_structure struct{
	Username string
	Role string
}

func CreateToken(Username string, Role string) (string, error){
		token := JWT.NewWithClaims(
			JWT.SigningMethodHS256,
			JWT.MapClaims{
				"Username" : Username,                  
				"Role" : Role, 
				"Expiry" : (time.Now().Add(4 * time.Hour)).Unix(),  // Each token is valid for 4 hours
			})
		tokenString, err := token.SignedString(secret_key)
		if err!= nil{
			log.ErrorLog(string(err.Error()))
		}
		return tokenString, err

}

func ValidateToken(token JWT.Token){
	if token.
}