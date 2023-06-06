package main

import (
	"log"

	"github.com/bishtpramod19/microservices/order/order/config"
	"github.com/bishtpramod19/microservices/order/order/internal/adapters/db"
	"github.com/bishtpramod19/microservices/order/order/internal/adapters/grpc"
	"github.com/bishtpramod19/microservices/order/order/internal/application/core/api"
)

func main() {
	dbadapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("failed to connect to database url : %v", err)
	}

	application := api.NewApplication(dbadapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()

}
