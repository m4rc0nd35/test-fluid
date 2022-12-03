package repossitory

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type leadApi struct{}

func NewLeadApi() *leadApi {
	return &leadApi{}
}

func (l *leadApi) GetLead() (*[]byte, error) {
	client := &http.Client{}
	// get new leads in API Random User
	req, err := http.NewRequest("GET", "https://randomuser.me/api/", strings.NewReader(""))
	if err != nil {
		log.Println("Error on response.\n[ERROR- ", err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return nil, err
	}

	defer req.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response\n[ERROR- ", err)
		return nil, err
	}

	return &response, nil
}
