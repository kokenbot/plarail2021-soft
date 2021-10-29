package internal

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	pb "ueckoken/plarail2021-soft-external/spec"
)

const (
	UNKNOWN = "UNKNOWN"
	SUCCESS = "SUCCESS"
	FAILED  = "FAILED"
)

type Command2Internal struct {
	station *StationState
	env     *Env
}

// NewCommand2Internal is Constructor of CommandInternal.
// CommandInternal Struct has a method to talk to Internal server with gRPC.
func NewCommand2Internal(state StationState, e *Env) *Command2Internal {
	return &Command2Internal{station: &state, env: e}
}

// sendRaw is making a connection to internal server and talk with internal server.
// This method will return gRPC response and gRPC error val.
// If you want join gRPC response Status Code and gRPC error msg, please use Command2Internal.trapResponseGrpcErr method.
func (c2i *Command2Internal) sendRaw() (*pb.ResponseSync, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(c2i.env.InternalServer.Addr.String(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewControlClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), c2i.env.InternalServer.TimeoutSec)
	defer cancel()
	r, err := c.Command2Internal(ctx, c2i.convert2pb())
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c2i *Command2Internal) Send() error {
	return trapResponseGrpcErr(c2i.sendRaw())
}

func (c2i *Command2Internal) convert2pb() *pb.RequestSync {
	return &pb.RequestSync{
		Station: &pb.Stations{StationId: pb.Stations_StationId(c2i.station.StationID)},
		State:   pb.RequestSync_State(c2i.station.State),
	}
}

func trapResponseGrpcErr(rs *pb.ResponseSync, grpcErr error) error {
	sta, ok := status.FromError(grpcErr)
	if (sta != nil && ok) || rs == nil { // gRPC error occur
		return fmt.Errorf("gRPC Err: %w", grpcErr)
	}
	// check Response Status
	switch rs.Response.String() {
	case UNKNOWN:
		return fmt.Errorf("gRPC Err: %w; gRPC Response status is %s", grpcErr, UNKNOWN)
	case SUCCESS:
		if grpcErr != nil {
			return fmt.Errorf("gRPC Err: %w; gRPC Response status is %s", grpcErr, SUCCESS)
		} else {
			return nil
		}
	case FAILED:
		return fmt.Errorf("gRPC Err: %w; gRPC Response status is %s", grpcErr, FAILED)
	default:
		return fmt.Errorf("gRPC Err: %w; Unknown error is occured", grpcErr)
	}
}
