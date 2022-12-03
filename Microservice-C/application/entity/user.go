package entity

type Name struct {
	Title string `json:"title" bson:"title"`
	First string `json:"first" bson:"first"`
	Last  string `json:"last" bson:"last"`
}

// Begin Location
type Street struct {
	Number int64  `json:"number" bson:"number"`
	Name   string `json:"name" bson:""name`
}
type Coordinates struct {
	Latitude  string `json:"latitude" bson:"latitude"`
	Longitude string `json:"longitude" bson:"longitude"`
}
type TimeZone struct {
	OffSet      string `json:"offset" bson:"offset"`
	Description string `json:"description" bson:"description"`
}

type Location struct {
	Street      `json:"street" bson:"street"`
	City        string `json:"city" bson:"city"`
	State       string `json:"state" bson:"state"`
	Country     string `json:"country" bson:"country"`
	Postcode    any    `json:"postcode" bson:"postCode"` // Random User string | int64
	Coordinates `json:"coordinates" bson:"coordinates"`
	TimeZone    `json:"timezone" bson:"timeZone"`
}

// End Location

type Login struct {
	Uuid     string `json:"uuid" bson:"uuid"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Salt     string `json:"salt" bson:"salt"`
	Md5      string `json:"md5" bson:"md5"`
	Sha1     string `json:"sha1" bson:"sha1"`
	Sha256   string `json:"sha256" bson:"sha256"`
}

type Dob struct {
	Date string `json:"date" bson:"date"`
	Age  int    `json:"age" bson:"age"`
}

type Registered struct {
	Date string `json:"date" bson:"date"`
	Age  int    `json:"age" bson:"age"`
}

type Id struct {
	Name  string `json:"name" bson:"name"`
	Value string `json:"value" bson:"value"`
}

type Picture struct {
	Large     string `json:"large" bson:"large"`
	Medium    string `json:"medium" bson:"medium"`
	Thumbnail string `json:"thumbnail" bson:"thumbnail"`
}

type User struct {
	Gender     string `json:"gender" bson:"gender"`
	Name       `json:"name" bson:"name"`
	StatusFlow string `json:"statusFlow" bson:"statusFlow"` // new, processing, processed, recused
	Location   `json:"location" bson:"location"`
	Email      string `json:"email" bson:"email"`
	Login      `json:"login" bson:"login"`
	Dob        `json:"dob" bson:"dob"`
	Registered `json:"registered" bson:"registered"`
	Phone      string `json:"phone" bson:"phone"`
	Cell       string `json:"cell" bson:"cell"`
	Id         `json:"id" bson:"id"`
	Picture    `json:"picture" bson:"picture"`
	Nat        string `json:"nat" bson:"nat"`
	CreatedAt  string `json:"createdAt" bson:"createdAt"`
}

type Leads struct {
	Results []User `json:"results"`
}
