package main

import (
	"log"
	"os"

	"github.com/kubarydz/go-hex/internal/adapters/app/api"
	"github.com/kubarydz/go-hex/internal/adapters/core/arithmetic"
	gRPC "github.com/kubarydz/go-hex/internal/adapters/framework/left/grpc"
	"github.com/kubarydz/go-hex/internal/adapters/framework/right/db"
	"github.com/kubarydz/go-hex/internal/ports"
)

func main() {
	var err error

	// ports
	var dbaseAdapter ports.DbPort
	var coreAdapter ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsName := os.Getenv("DS_NAME")

	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsName)
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	defer dbaseAdapter.CloseDbConnection()

	coreAdapter = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbaseAdapter, coreAdapter)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run()

}
