package main

import (
	"net"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"github.com/toshkentov01/task/data_service/config"
	"github.com/toshkentov01/task/data_service/pkg/logger"
	"github.com/toshkentov01/task/data_service/pkg/migration"
	"github.com/toshkentov01/task/data_service/service"
	"github.com/toshkentov01/task/data_service/storage/post"
	"google.golang.org/grpc"

	dataPb "github.com/toshkentov01/task/data_service/genproto/data_service"
)

func main() {
	// Migations Up
	migration.Up()

	if info, err := os.Stat(".env"); !os.IsNotExist(err) {
		if !info.IsDir() {
			godotenv.Load(".env")
		}
	}

	var cfg = config.Get()

	log := logger.New(cfg.LogLevel, "data_service")
	defer logger.Cleanup(log)

	listen, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("error listening tcp port: ", logger.Error(err))
	}

	DataService := service.NewDataService(log)
	server := grpc.NewServer()
	dataPb.RegisterDataServiceServer(server, DataService)

	go func() {
		// Getting posts from open api and inserting them
		// It may take 10 second to get and insert posts to database
		postRepo := post.NewPostRepo()
		err = postRepo.InsertPosts()

		if err != nil {
			log.Error("Error while getting posts from open api and inserting them: " + err.Error())
			return
		}
	}()

	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := server.Serve(listen); err != nil {
		log.Fatal("error listening: %v", logger.Error(err))
	}

}
