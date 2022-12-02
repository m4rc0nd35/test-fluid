package entity

type DataLogger struct {
	Uuid       string `json:"uuid" bson:"uuid"`
	Username   string `json:"username" bson:"username"`
	Email      string `json:"email" bson:"email"`
	StatusFlow string `json:"statusFlow" bson:"statusFlow"`
	CreatedAt  string `json:"createdAt" bson:"createdAt"`
}
