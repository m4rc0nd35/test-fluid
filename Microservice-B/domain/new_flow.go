package domain

import (
	"encoding/json"
	"log"
	"os"

	"github.com/m4rc0nd35/test-fluid/application/adapter"
	"github.com/m4rc0nd35/test-fluid/application/entity"
)

type config struct {
	amqpx   adapter.Amqp
	recused bool
	logs    adapter.AdapterDataLogger
}

func NewFlow(amqpx adapter.Amqp, logs adapter.AdapterDataLogger) *config {
	return &config{amqpx, false, logs}
}

func (cfg *config) WorkerNewFlow(body *string) bool {

	user := entity.User{}

	// Json to Struct
	if err := json.Unmarshal([]byte(*body), &user); err != nil {
		log.Println("Error structure JSON lead\n[ERROR- ", err)
		return false
	}

	user.StatusFlow = "processing" // Changed status
	log.Println("[ PROCESSING ] ", user.Name.First)

	cfg.logs.LogQueue(user)

	// Struct to Json
	jsonNewLead, _ := json.Marshal(user)                                         // Struct to JSON
	cfg.amqpx.SendToQueu(os.Getenv("QUEUE_RCV_PROCESSING"), string(jsonNewLead)) // Send amqp queue
	cfg.recused = false

	return true
}
