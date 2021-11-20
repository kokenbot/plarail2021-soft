package internal

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type RaspberrySpeed struct {
	speed int32
}

var endPoint = "http://localhost:8081"

func NewRaspberrySpeed(speed int32) *RaspberrySpeed {
	return &RaspberrySpeed{speed: speed}
}

func (r *RaspberrySpeed) changeSpeed() error {
	u, err := url.Parse(endPoint)
	if err != nil {
		log.Fatalln("URL Parse Err")
	}
	u.Query().Set("speed", string(r.speed))

	res, err := http.Get(u.String())

	if !(200 <= res.StatusCode && res.StatusCode < 300) {
		return fmt.Errorf("GET Err is `%w` ;HTTP status is `%s`", err, res.Status)
	}
	return err
}
