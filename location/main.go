package main

import (
	"fmt"
	"net"
	"strconv"

	"github.com/devminnu/weatherapp/location/config"
	"github.com/devminnu/weatherapp/location/db"
	"github.com/devminnu/weatherapp/location/logger"
	"github.com/devminnu/weatherapp/location/router"
	"github.com/devminnu/weatherapp/location/service"
	"google.golang.org/grpc"
)

func main() {
	config.Load()
	logger.Init()
	db.Init()
	router.Init()
	service.Init()
	router := router.Get()
	port := config.AppPort() // This should be changed to the service port number via argument or environment variable.

	addr := fmt.Sprintf(":%s", strconv.Itoa(port))

	// @grpc
	go GRPCServe()
	router.Run(addr)
}

func GRPCServe() {
	host := config.ReadEnvString("GRPC_HOST")
	port := config.ReadEnvInt("GRPC_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		logger.Get()
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	logger.Get()

	grpcServer.Serve(lis)
}
