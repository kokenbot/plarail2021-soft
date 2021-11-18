package storeSpeed

import (
	"errors"
	"net"
	"sync"
	pb "ueckoken/plarail2021-soft-speed/spec"
)

type Train struct {
	Name pb.SendSpeed_Train
	Addr net.Addr
}

func NewTrain(name pb.SendSpeed_Train) Train {
	return Train{Name: name}
}

type TrainConf interface {
	GetTrain() Train
	SetSpeed(speed int32) error
	GetSpeed() int32
}
type trainConf struct {
	Train Train
	Speed int32
}

func NewTrainConf(train Train) TrainConf {
	return &trainConf{Train: train}
}

func (t *trainConf) SetSpeed(speed int32) error {
	if !(-100 <= speed && speed <= 100) {
		return errors.New("speed range error")
	}
	t.Speed = speed
	return nil
}

func (t *trainConf) GetTrain() Train {
	return t.Train
}

func (t *trainConf) GetSpeed() int32 {
	return t.Speed
}

type SpeedStore interface {
	Get(t *Train) (speed TrainConf, err error)
	Update(t TrainConf)
	GetAll() []TrainConf
}
type speedStore struct {
	speedList []TrainConf
	mutex     sync.Mutex
}

func NewStore() SpeedStore {
	return &speedStore{}
}

func (s *speedStore) Get(t *Train) (speed TrainConf, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, r := range s.speedList {
		if r.GetTrain() == *t {
			return r, nil
		}
	}
	return nil, errors.New("Speed Record Not Found\n")
}

func (s *speedStore) Update(t TrainConf) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for i, r := range s.speedList {
		if r.GetTrain() == t.GetTrain() {
			s.speedList[i] = t
			return
		}
	}
	s.speedList = append(s.speedList, t)
	return
}

func (s *speedStore) GetAll() []TrainConf {
	return s.speedList
}
