package main

import (
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/toshkentov01/task/crud_service/config"
	"github.com/toshkentov01/task/crud_service/pkg/logger"
	"github.com/toshkentov01/task/crud_service/service"
	"google.golang.org/grpc"

	crudPb "github.com/toshkentov01/task/crud_service/genproto/crud_service"
)

func main() {
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

	DataService := service.NewCrudService(log)
	server := grpc.NewServer()
	crudPb.RegisterCrudServiceServer(server, DataService)

	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := server.Serve(listen); err != nil {
		log.Fatal("error listening: %v", logger.Error(err))
	}
}
