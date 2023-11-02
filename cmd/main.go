package main

import (
	"os"

	"github.com/akashabbasi/hex/internal/adapters/app/api"
	"github.com/akashabbasi/hex/internal/adapters/core/arithmetic"
	gRPC "github.com/akashabbasi/hex/internal/adapters/framework/left/grpc"
	"github.com/akashabbasi/hex/internal/adapters/framework/right/db"
	"github.com/akashabbasi/hex/internal/ports"
)

func main() {
	var err error
	// ports
	var dbaseAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbaseAdapter = db.NewAdapter(dbaseDriver, dsourceName)
	defer dbaseAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbaseAdapter, core)
	gRPCAdapter = gRPC.NewAdapter(appAdapter)

	gRPCAdapter.Run()
}
