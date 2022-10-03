package main

import (
	"context"
	"fmt"
	"log"
	"security-api/controllers"
	"security-api/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	us          services.UserService
	uc          controllers.UserController
	ctx         context.Context
	userc       *mongo.Collection
	mongoclient *mongo.Client
	err         error
)

func init() {
	ctx = context.TODO()
	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("Error while connecting with MongoDB", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Error while trying to ping MongoDB", err)
	}
	fmt.Println("Connected to MongoDB")
	userc = mongoclient.Database("userdb").Collection("users")
	us = services.NewUserService(userc, ctx)
	uc = controllers.New(us)
	server = gin.Default()

}

func main() {
	defer mongoclient.Disconnect(ctx) //disconect if app shuts down

	basePath := server.Group("/security")
	uc.RegisterUserRoutes(basePath)
	log.Fatal(server.Run(":9090"))

}
