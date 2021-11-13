package Send2Json

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func SendJson(url string, json []byte, timeout time.Duration) error {
	client := &http.Client{Timeout: timeout}
	res, err := client.Post(url, "application/json", bytes.NewBuffer(json))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if !(200 <= res.StatusCode && res.StatusCode < 300) {
		return fmt.Errorf("http status is not succeed. status is %s\n", res.Status)
	}
	return nil
}