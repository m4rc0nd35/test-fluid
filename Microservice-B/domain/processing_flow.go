package domain

import (
	"encoding/json"
	"log"
	"time"

	"github.com/m4rc0nd35/test-fluid/application/adapter"
	"github.com/m4rc0nd35/test-fluid/application/entity"
)

type configProcessing struct {
	amqpx   adapter.Amqp
	recused bool
	logs    adapter.AdapterDataLogger
}

func NewProcessingFlow(amqpx adapter.Amqp, logs adapter.AdapterDataLogger) *configProcessing {
	return &configProcessing{amqpx, false, logs}
}

func (cfg *configProcessing) WorkerProcessingFlow(body *string) bool {

	user := entity.User{}

	// Json to Struct
	if err := json.Unmarshal([]byte(*body), &user); err != nil {
		log.Println("Error structure JSON lead\n[ERROR- ", err)
		return false
	}

	go func() {
		log.Println("[ PROCESSING ] ", user.Name.First)
		time.Sleep(time.Second * 10) // T

		if cfg.recused {
			user.StatusFlow = "recused" // Changed status
			log.Println("[ RECUSED ] ", user.Name.First)
		}

		if !cfg.recused {
			user.StatusFlow = "processed" // Changed status
			log.Println("[ PROCESSED ] ", user.Name.First)
		}

		cfg.logs.LogQueue(user)

		// Struct to Json
		jsonNewLead, _ := json.Marshal(user)                           // Struct to JSON
		cfg.amqpx.SendToQueu("fluid-processed-K", string(jsonNewLead)) // Send amqp queue
		cfg.recused = false
	}()

	return true
}

func (cfg *configProcessing) Recused() {
	cfg.recused = true
}
