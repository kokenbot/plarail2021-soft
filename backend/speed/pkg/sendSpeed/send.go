package sendSpeed

import (
	"context"
	"fmt"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"time"
	"ueckoken/plarail2021-soft-speed/pkg/storeSpeed"
	pb "ueckoken/plarail2021-soft-speed/spec"
)

type sendSpeed struct {
	Train storeSpeed.TrainConf
}

func (s *sendSpeed) Send() error {
	return trapSendErr(s.sendRaw())
}

func (s *sendSpeed) sendRaw() (*pb.StatusCode, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(
		ctx,
		s.Train.GetTrain().Addr.String(),
		grpc.WithInsecure(), grpc.WithBlock(),
		grpc.WithUnaryInterceptor(grpcPrometheus.UnaryClientInterceptor),
	)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewSpeedClient(conn)
	ctx, cancel1 := context.WithTimeout(ctx, 2*time.Second)
	defer cancel1()
	r, err := c.ControlSpeed(ctx, s.convert2pb())
	if err != nil {
		return nil, err
	}
	return r, nil
}
func trapSendErr(code *pb.StatusCode, err error) error {
	switch code.GetCode() {
	case pb.StatusCode_SUCCESS:
		if err != nil {
			return nil
		} else {
			return err
		}
	default:
		return fmt.Errorf("gRPC error `%w`,status code is `%d`", err, code.GetCode())
	}
}
func (s *sendSpeed) convert2pb() *pb.SendSpeed {
	pss := pb.SendSpeed{
		Speed: s.Train.GetSpeed(),
		Train: s.Train.GetTrain().Name,
	}
	return &pss
}