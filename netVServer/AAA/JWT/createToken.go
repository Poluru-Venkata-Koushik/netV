package jwt

import (
	"errors"
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

func IsValidToken(tokenString string) bool{

	claims := JWT.MapClaims{}
	token, err := JWT.ParseWithClaims(tokenString, claims, func(token *JWT.Token) (interface{}, error) {
		if _, ok := token.Method.(*JWT.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid Signing Method")
		}
		return secret_key, nil
	})
	if err != nil {
		log.ErrorLog("Failed to parse")
		log.ErrorLog(err.Error())
		return false
	}

	if !token.Valid {
		return false
	}
	expiry, ok := claims["Expiry"].(float64)
	if !ok {
		return false
	}
	if int64(expiry) < time.Now().Unix() {
		return false
	}
	return true
}



func ExtractToken(tokenString string) (*JWT_structure, error) {

	if IsValidToken(tokenString) {
		claims := JWT.MapClaims{}
		username, _ := claims["Username"].(string)
		role, _ := claims["Role"].(string)
		return &JWT_structure{
			Username : username,
			Role : role,
		}, nil
	}
	
	return nil, errors.New("INVALID TOKEN")

}

