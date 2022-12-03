package domain

import (
	"encoding/json"

	"github.com/m4rc0nd35/test-fluid/application/adapter"
	"github.com/m4rc0nd35/test-fluid/application/entity"
)

type leadConfig struct {
	leadRepo adapter.LeadRepository
}

func NewLead(leadRepo adapter.LeadRepository) *leadConfig {
	return &leadConfig{leadRepo}
}

func (d *leadConfig) Lead(body []byte) bool {
	lead := entity.User{}

	// json to struct
	if err := json.Unmarshal(body, &lead); err != nil {
		return false
	}

	// insert lead
	d.leadRepo.Create(lead)
	return true
}

func (d *leadConfig) FindOneLead(uuid string) *entity.User {
	lead, _ := d.leadRepo.FindOneLead(uuid)
	return lead
}

func (d *leadConfig) FindAllLead() []*entity.User {
	leads, _ := d.leadRepo.FindAllLead()
	return leads
}
