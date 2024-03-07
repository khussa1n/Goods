package natspkg

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"net/url"
)

type NatsConn struct {
	Host  string
	Topic string
}

func New(host string, topic string) *NatsConn {
	return &NatsConn{
		Host:  host,
		Topic: topic,
	}
}

func (n *NatsConn) Connect() (*nats.Conn, error) {
	q := url.Values{}
	u := url.URL{
		Scheme:   "nats",
		Host:     fmt.Sprintf("%s:%d", n.Host, nats.DefaultPort),
		RawQuery: q.Encode(),
	}

	nc, err := nats.Connect(u.String())
	if err != nil {
		log.Fatal("Error connecting to NATS:", err)
		return nil, err
	}

	return nc, nil
}
