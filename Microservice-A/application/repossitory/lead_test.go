package repossitory

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/m4rc0nd35/test-fluid/application/entity"
	"github.com/stretchr/testify/assert"
)

func TestGetLead_return_json(t *testing.T) {
	leads := entity.Leads{}
	tl := NewLeadApi()

	response, _ := tl.GetLead()

	// destructure
	if err := json.Unmarshal(*response, &leads); err != nil {
		log.Println("Error structure JSON lead\n[ERROR- ", err)
		t.Fail()
	}
	jsonR, err := json.Marshal(leads)
	if err != nil {
		t.Fail()
	}

	assert.JSONEq(t, string(jsonR), string(jsonR))                                             // Json validation
	assert.Regexp(t, "^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+.[a-zA-Z0-9-.]+$", leads.Results[0].Email) // Email validation
}
