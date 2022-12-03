package adapter

import "github.com/m4rc0nd35/test-fluid/application/entity"

type LeadRepository interface {
	Create(data entity.User) string
	FindOneLead(uuid string) (*entity.User, error)
	FindAllLead() ([]*entity.User, error)
}

type LeadDomain interface {
	FindOneLead(uuid string) *entity.User
	FindAllLead() []*entity.User
}
