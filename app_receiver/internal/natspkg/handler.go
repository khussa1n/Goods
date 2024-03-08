package natspkg

import (
	"encoding/json"
	"github.com/khussa1n/Goods/app_receiver/internal/entity"
	"github.com/khussa1n/Goods/app_receiver/internal/repository/chrepo"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

type NatsHandler struct {
	Chrepo    *chrepo.ClickhouseRepo
	BatchSize int
}

var messageBuffer []entity.Goods

func (nh *NatsHandler) HandleMessage(msg *nats.Msg) {
	log.Printf("Processed message: %s", msg.Data)

	var receivedData entity.GoodsResponce
	err := json.Unmarshal(msg.Data, &receivedData)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return
	}

	goods := &entity.Goods{
		ID:          receivedData.ID,
		ProjectID:   receivedData.ProjectID,
		Name:        receivedData.Name,
		Description: receivedData.Description,
		Priority:    receivedData.Priority,
		Removed:     receivedData.Removed,
		EventTime:   time.Now(),
	}

	messageBuffer = append(messageBuffer, *goods)

	if len(messageBuffer) >= nh.BatchSize {
		nh.insertBatch()
	}
}

func (nh *NatsHandler) insertBatch() {
	if len(messageBuffer) == 0 {
		return
	}

	batch := make([]entity.Goods, len(messageBuffer))
	copy(batch, messageBuffer)
	messageBuffer = nil

	err := nh.Chrepo.InsertBatch(batch)
	if err != nil {
		log.Println("Error inserting data into ClickHouse:", err)
		return
	}

	log.Printf("Data batch inserted into ClickHouse: %d", len(batch))
}
