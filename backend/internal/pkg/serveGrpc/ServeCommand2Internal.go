package serveGrpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"ueckoken/plarail2021-soft-internal/pkg/msg2Esp"
	"ueckoken/plarail2021-soft-internal/pkg/station2espIp"
	pb "ueckoken/plarail2021-soft-internal/spec"
)

type ControlServer struct {
	pb.UnimplementedControlServer
	Stations station2espIp.Stations
	client   *http.Client
}

func (c *ControlServer) Command2Internal(_ context.Context, req *pb.RequestSync) (*pb.ResponseSync, error) {
	sta, err := c.unpackStations(req.GetStation())
	if err != nil {
		return nil, err
	}
	angle := 0
	if sta.IsAngleDefined() {
		angle, err = sta.GetAngle(req.GetState())
		if err != nil {
			return nil, err
		}
	}
	s2n := msg2Esp.NewSend2Node(c.client, sta, c.unpackState(req.GetState()), angle)
	err = s2n.Send()

	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "sender err %s; not connected to Node", err.Error())
	}
	return &pb.ResponseSync{Response: pb.ResponseSync_SUCCESS}, nil
}

func (c *ControlServer) unpackStations(req *pb.Stations) (*station2espIp.StationDetail, error) {
	s, ok := pb.Stations_StationId_name[int32(req.GetStationId())]
	if !ok {
		return nil, fmt.Errorf("station: %s do not define in proto file\n", req.String())
	}
	sta, err := c.Stations.Detail(s)
	if err != nil {
		return nil, fmt.Errorf("%w; station `%s` is not defined in yaml file\n", err, s)
	}
	return sta, nil
}

func (*ControlServer) unpackState(state pb.RequestSync_State) string {
	return state.String()
}
