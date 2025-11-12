package main

import (
	"context"
	"fmt"
	Logging "netVServer/LogPackage"
	"netVServer/NecessaryStructs"
	"time"

	_ "go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)


const (
	LogFile 		= "MongoClient.log"
	PKG     		= "Database"
    URI 			= "mongodb://root:netvroot@localhost:27018/?authSource=admin"
	Context_timeout = 3 // 3 seconds
)

var (
	log = Logging.Logger{Filename: LogFile, PKG: PKG}
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), Context_timeout * time.Second)
	defer cancel()

	MongoOpt := options.Client().ApplyURI(URI)
	Client,err := mongo.Connect(MongoOpt)
	if err!=nil{
		log.ErrorLog(err.Error())
	}else{
		log.PrintLog("Connected to Mongo")
	}

	collection := Client.Database("Testing").Collection("TestCollection")
	id , err := collection.InsertOne(ctx, necessarystructs.StructUser{
		Username: "Admin",
		Role: "Koushik",
		Token: "JWT here",
		Validity: 1234,
	})
	if err!=nil{
		log.ErrorLog(err.Error())
	}else{
		log.InfoLog(fmt.Sprintf("Inserted an Entry %v", id.InsertedID))
	}
	

}
