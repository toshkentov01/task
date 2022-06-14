package grpcclient

import (
	"fmt"
	"sync"

	"github.com/toshkentov01/task/crud_service/config"
	dataPb "github.com/toshkentov01/task/crud_service/genproto/data_service"
	"google.golang.org/grpc"
)

var (
	onceDataService     sync.Once
	instanceDataService dataPb.DataServiceClient
	cfg                 = config.Get()
)

// DataService ...
func DataService() dataPb.DataServiceClient {
	onceDataService.Do(func() {
		connDataService, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.DataServiceHost, cfg.DataServicePort),
			grpc.WithInsecure())

		if err != nil {
			panic(fmt.Errorf("user service dial host: %s port:%d err: %s",
				cfg.DataServiceHost, cfg.DataServicePort, err))
		}

		instanceDataService = dataPb.NewDataServiceClient(connDataService)
	})

	return instanceDataService
}
