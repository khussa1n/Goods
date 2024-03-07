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
	Chrepo *chrepo.ClickhouseRepo
}

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

	err = nh.Chrepo.Insert(goods)
	if err != nil {
		log.Println("Error inserting data into ClickHouse:", err)
		return
	}

	log.Println("Data inserted into ClickHouse:", receivedData)
}
