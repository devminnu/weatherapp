package rpc

import (
	"context"

	pb "github.com/devminnu/weatherapp/pb/location"
)

type server struct {
}

func (s *server) GetLocationByID(context.Context, *pb.LocationID) (*pb.Location, error) {
	return &pb.Location{}, nil
}
