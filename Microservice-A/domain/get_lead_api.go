package domain

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
	logs      adapter.AdapterDataLogger
}

var pause = false

func NewLead(amqpx adapter.Amqp, logs adapter.AdapterDataLogger) *config {
	return &config{
		true,        // Default running
		cron.New(),  // Init instance cron
		"0 * * * *", // Default scheduler
		0,           // Default entry cron ID
		amqpx,       // Adapter amqp
		1,           // Default limit lead by job
		logs,
	}
}

func (cfg *config) GetLeadApi() {
	// Add new scheduler cron
	cfg.cron.Remove(cfg.entryID)
	cfg.entryID, _ = cfg.cron.AddFunc(cfg.scheduler, func() {
		client := &http.Client{}
		leads := entity.Leads{} // received

		var totalGet = 0
		for totalGet < cfg.getLimit {

			for pause {
				time.Sleep(time.Second)
			}

			// get new leads in API Random User
			req, err := http.NewRequest("GET", "https://randomuser.me/api/", strings.NewReader(""))
			if err != nil {
				log.Println("Error on response.\n[ERROR- ", err)
				return
			}

			req.Header.Add("Content-Type", "application/json")
			resp, err := client.Do(req)
			if err != nil {
				log.Println("Error on response.\n[ERROR] -", err)
				return
			}

			defer req.Body.Close()
			response, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println("Error while reading the response\n[ERROR- ", err)
				return
			}

			// destructure
			if err := json.Unmarshal(response, &leads); err != nil {
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

func (cfg *config) RemoveScheduler(p bool) {
	// remove cron by ID
	// cfg.cron.Remove(cfg.entryID)
	// cfg.entryID = 0
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
