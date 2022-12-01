package domain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

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
}

func NewLead(amqpx adapter.Amqp) *config {
	return &config{
		true,        // Default running
		cron.New(),  // Init instance cron
		"0 * * * *", // Default scheduler
		0,           // Default entry cron ID
		amqpx,       // Adapter amqp
		1,           // Default limit lead by job
	}
}

func (cfg *config) GetLeadApi() {
	// Add new scheduler
	cfg.entryID, _ = cfg.cron.AddFunc(cfg.scheduler, func() {
		client := &http.Client{}
		leads := entity.Leads{} // received

		go func() {
			fmt.Println("Gorouting")
		}()

		totalGetSuccess := 0
		for totalGetSuccess < cfg.getLimit {

			// get new leads in API Ranbom User
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

			// multiple leads by request
			for _, user := range leads.Results {
				user.StatusFlow = "new"                               // Changed status
				totalGetSuccess++                                     // Increase total lead
				jsonNewLead, _ := json.Marshal(user)                  // Struct to JSON
				cfg.amqpx.SendToQueu("new-lead", string(jsonNewLead)) // Send amqp queue

				log.Println("[ RUNNNING ] get:", totalGetSuccess)

				// Break when limit
				if totalGetSuccess >= cfg.getLimit {
					break
				}
			}
		}

		log.Println("Get leads by API random user total:", totalGetSuccess)
		// send to queue new-lead
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

func (cfg *config) RemoveScheduler() {
	// remove cron by ID
	cfg.cron.Remove(cfg.entryID)
	cfg.entryID = 0
}

// * * * * *
// | | | | |
// | | | | |
// | | | | +---- Day of week  (interval: 1-7)
// | | | +------ Month        (interval: 1-12)
// | | +-------- Day          (interval: 1-31)
// | +---------- Hour         (interval: 0-23)
// +------------ Minute       (interval: 0-59)
