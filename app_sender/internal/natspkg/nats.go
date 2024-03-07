package natspkg

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"net/url"
)

func Connect(host string) (*nats.Conn, error) {
	q := url.Values{}
	u := url.URL{
		Scheme:   "nats",
		Host:     fmt.Sprintf("%s:%d", host, nats.DefaultPort),
		RawQuery: q.Encode(),
	}

	nc, err := nats.Connect(u.String())
	if err != nil {
		log.Fatal("Error connecting to NATS:", err)
		return nil, err
	}

	return nc, nil
}

func Publish(conn *nats.Conn, topic string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return conn.Publish(topic, jsonData)
}
