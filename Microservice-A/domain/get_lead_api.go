package domain

import (
	"encoding/json"
	"log"
	"time"

	"github.com/m4rc0nd35/test-fluid/application/adapter"
	"github.com/m4rc0nd35/test-fluid/application/entity"
	"github.com/robfig/cron/v3"
)

type config struct {
	running   bool
	cron      *cron.Cron
	scheduler string
	entryID   cron.EntryID
	amqpx     adapter.Amqp
	getLimit  int
	leadRepo  adapter.LeadRepository
	logs      adapter.AdapterDataLogger
}

var pause = false // Global

func NewLead(amqpx adapter.Amqp, leadRepo adapter.LeadRepository, logs adapter.AdapterDataLogger) *config {
	return &config{
		true,        // Default running
		cron.New(),  // Init instance cron
		"0 * * * *", // Default scheduler
		0,           // Default entry cron ID
		amqpx,       // Adapter amqp
		1,           // Default limit lead by job
		leadRepo,
		logs,
	}
}

func (cfg *config) GetLeadApi() cron.EntryID {
	// Add new scheduler cron
	cfg.cron.Remove(cfg.entryID)
	cfg.entryID, _ = cfg.cron.AddFunc(cfg.scheduler, func() {
		leads := entity.Leads{} // received

		var totalGet = 0
		for totalGet < cfg.getLimit {

			for pause {
				time.Sleep(time.Second)
			}

			// Repository
			response, err := cfg.leadRepo.GetLead()
			if err != nil {
				log.Println("Error repository JSON lead\n[ERROR- ", err)
				return
			}

			// destructure
			if err := json.Unmarshal(*response, &leads); err != nil {
				log.Println("Error structure JSON lead\n[ERROR- ", err)
				return
			}

			// leads by request
			for _, user := range leads.Results {
				user.StatusFlow = "new" // Changed status

				totalGet++ // Increase total lead

				// Data Logger
				cfg.logs.LogQueue(user)

				jsonNewLead, _ := json.Marshal(user)                     // Struct to JSON
				cfg.amqpx.SendToQueu("fluid-new-H", string(jsonNewLead)) // Send amqp queue

				log.Println("[ RUNNNING ] get:", totalGet)

				// Break when limit
				if totalGet >= cfg.getLimit {
					break
				}
			}
		}

		log.Println("Get leads by API random user finished", totalGet)
	})
	return cfg.entryID
}

func (cfg *config) Start() {
	// Start cron
	cfg.cron.Start()
}

func (cfg *config) SetGetLimit(limit int) {
	// Start cron
	cfg.getLimit = limit
}

func (cfg *config) SetScheduler(scheduler string) {
	cfg.scheduler = scheduler
	log.Println("[ SET ] new scheduler", cfg.scheduler)
	cfg.cron.Remove(cfg.entryID)
	cfg.GetLeadApi()
}

func (cfg *config) RemoveScheduler(id cron.EntryID) {
	cfg.cron.Remove(id)
}

func (cfg *config) Pause(p bool) {
	pause = p
}

// * * * * *
// | | | | |
// | | | | |
// | | | | +---- Day of week  (interval: 1-7)
// | | | +------ Month        (interval: 1-12)
// | | +-------- Day          (interval: 1-31)
// | +---------- Hour         (interval: 0-23)
// +------------ Minute       (interval: 0-59)
