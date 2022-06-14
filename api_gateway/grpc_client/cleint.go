package grpcclient

import (
	"fmt"
	"sync"

	"github.com/toshkentov01/task/api_gateway/config"
	crudPb "github.com/toshkentov01/task/api_gateway/genproto/crud_service"
	"google.golang.org/grpc"
)

var (
	onceDataService     sync.Once
	instanceDataService crudPb.CrudServiceClient
	cfg                 = config.Config()
)

// CrudService ...
func CrudService() crudPb.CrudServiceClient {
	onceDataService.Do(func() {
		connDataService, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.CrudServiceHost, cfg.CrudServicePort),
			grpc.WithInsecure())

		if err != nil {
			panic(fmt.Errorf("user service dial host: %s port:%d err: %s",
				cfg.DataServiceHost, cfg.DataServicePort, err))
		}

		instanceDataService = crudPb.NewCrudServiceClient(connDataService)
	})

	return instanceDataService
}
