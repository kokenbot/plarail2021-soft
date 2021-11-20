package internal

import (
	"net/http"
	"ueckoken/plarail2021-soft-speed/pkg/storeSpeed"
	"ueckoken/plarail2021-soft-speed/pkg/train2IP"

	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	pb "ueckoken/plarail2021-soft-speed/spec"
)

type speedStruct struct {
	TrainName string `json:"train_name"`
	Speed     int    `json:"speed"`
}

func (m ClientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := m.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close()

	ctx, cancel := context.WithCancel(context.Background())
	n := make(chan storeSpeed.TrainConf, 64)
	unregister := make(chan struct{})
	notifier := Client{ClientNotifier{n, unregister}}
	m.ClientNotification <- notifier
	c.SetPongHandler(func(string) error {
		c.SetReadDeadline(time.Now().Add(20 * time.Second))
		return nil
	})
	c.SetCloseHandler(func(code int, text string) error {
		fmt.Println("connection closed")
		cancel()
		return nil
	})
	go handleClientPing(ctx, c)
	go handleClientCommand(ctx, c, &m)
	for notification := range notifier.notifier.Notifier {
		sendData := speedStruct{TrainName: notification.GetTrain().Name.String(), Speed: int(notification.GetSpeed())}
		err := c.WriteJSON(sendData)
		if err != nil {
			notifier.notifier.Unregister <- struct{}{}
			cancel()
			break
		}
	}
}

func handleClientCommand(ctx context.Context, c *websocket.Conn, m *ClientHandler) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			r, err := unpackClientSendData(c, m.TrainName2Id)
			if err != nil {
				log.Println(err)
				return
			}
			m.ClientCommand <- *r
		}
	}
}

func handleClientPing(ctx context.Context, c *websocket.Conn) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := c.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(1*time.Second)); err != nil {
				fmt.Println("ping:", err)
			}
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

func unpackClientSendData(c *websocket.Conn, nameDir train2IP.Name2Id) (*storeSpeed.TrainConf, error) {
	_, msg, err := c.ReadMessage()
	if err != nil {
		return nil, fmt.Errorf("websocket read failed: %e", err)
	}
	var ud speedStruct
	err = json.Unmarshal(msg, &ud)
	if err != nil {
		return nil, fmt.Errorf("bad json format: %e", err)
	}
	speed := storeSpeed.NewTrainConf(
		storeSpeed.NewTrain(pb.SendSpeed_Train(pb.SendSpeed_Train_value[ud.TrainName]),
			nameDir.SearchIp(ud.TrainName)),
	)
	speed.SetSpeed(int32(ud.Speed))
	return &speed, nil
}
