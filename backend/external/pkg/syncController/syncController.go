package syncController

import (
	"errors"
	"fmt"
	"github.com/prometheus/common/log"
	"sync"
	"time"
	"ueckoken/plarail2021-soft-external/pkg/envStore"
	"ueckoken/plarail2021-soft-external/pkg/servo"
	"ueckoken/plarail2021-soft-external/pkg/stationNameId"
)

type StationState struct {
	servo.StationState
}

type stationKVS struct {
	stations  []StationState
	mtx       sync.Mutex
	validator IValidator
}

func newStationKvs() *stationKVS {
	v := NewRouteValidator()
	skvs := stationKVS{validator: v}
	return &skvs
}
func (skvs *stationKVS) update(u StationState) error {
	skvs.mtx.Lock()
	defer skvs.mtx.Unlock()
	err := skvs.validator.Validate(u, skvs.stations)
	if err != nil {
		return err
	}
	for i, s := range skvs.stations {
		if s.StationID == u.StationID {
			skvs.stations[i].State = u.State
			return nil
		}
	}
	skvs.stations = append(skvs.stations, u)
	return nil
}

// forceUpdate differs from update for ignore route validation.
func (skvs *stationKVS) forceUpdate(u StationState) error {
	skvs.mtx.Lock()
	defer skvs.mtx.Unlock()
	for i, s := range skvs.stations {
		if s.StationID == u.StationID {
			skvs.stations[i].State = u.State
			return nil
		}
	}
	skvs.stations = append(skvs.stations, u)
	return nil
}
func (skvs *stationKVS) get(stationID int32) (station StationState, err error) {
	skvs.mtx.Lock()
	defer skvs.mtx.Unlock()
	for _, s := range skvs.stations {
		if s.StationID == stationID {
			return s, nil
		}
	}
	return StationState{}, errors.New("Not found")
}

func (skvs *stationKVS) retrieve() []StationState {
	return skvs.stations
}

type SyncController struct {
	ClientHandler2syncController chan StationState
	SyncController2clientHandler chan StationState
	InitServoRoute               chan StationState
	Environment                  *envStore.Env
}

func (s *SyncController) StartSyncController() {
	kvs := newStationKvs()

	go s.Init(NewInitializeRule())
	s.initNode(s.Environment, kvs)

	go s.periodicallySync(kvs)
	s.triggeredSync(s.Environment, kvs)
}

func (s *SyncController) initNode(e *envStore.Env, kvs *stationKVS) {
	for c := range s.InitServoRoute {
		err := kvs.forceUpdate(c)
		if err != nil {
			log.Errorln(err)
			return
		}
		c2i := servo.NewCommand2Internal(c.StationState, e)
		err = c2i.Send()
		if err != nil {
			log.Errorln(err)
		}
	}
}

func (s *SyncController) triggeredSync(e *envStore.Env, kvs *stationKVS) {
	for c := range s.ClientHandler2syncController {
		kvs.update(c)
		//todo; handle update error
		c2i := servo.NewCommand2Internal(c.StationState, e)
		err := c2i.Send()
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@:", err)
		s.SyncController2clientHandler <- c
	}
}

func (s *SyncController) periodicallySync(kvs *stationKVS) {
	ch := time.Tick(2 * time.Second)
	for range ch {
		fmt.Println("locking")
		kvs.mtx.Lock()
		fmt.Println("locked")
		for _, st := range kvs.retrieve() {
			s.SyncController2clientHandler <- st
		}
		kvs.mtx.Unlock()
	}
}

func (s SyncController) Init(r *InitRule) {
	for _, sta := range r.Stations {
		id, err := stationNameId.Name2Id(sta.Name)
		if err != nil {
			log.Fatalln(err)
		}
		s.InitServoRoute <- StationState{
			struct {
				StationID int32
				State     int32
			}{
				StationID: id,
				State:     int32(sta.State),
			},
		}
		time.Sleep(500 * time.Millisecond)
	}
	close(s.InitServoRoute)
}
