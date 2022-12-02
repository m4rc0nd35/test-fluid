package entity

type Log struct {
	Uuid       string `json:"uuid"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	StatusFlow string `json:"statusFlow"`
}
