package entity

type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

// Begin Location
type Street struct {
	Number int64  `json:"number"`
	Name   string `json:"name"`
}
type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}
type TimeZone struct {
	OffSet      string `json:"offset"`
	Description string `json:"description"`
}

type Location struct {
	Street      `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Postcode    any    `json:"postcode"` // Random User string | int64
	Coordinates `json:"coordinates"`
	TimeZone    `json:"timezone"`
}

// End Location

type Login struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Md5      string `json:"md5"`
	Sha1     string `json:"sha1"`
	Sha256   string `json:"sha256"`
}

type Dob struct {
	Date string `json:"date"`
	Age  int    `json:"age"`
}

type Registered struct {
	Date string `json:"date"`
	Age  int    `json:"age"`
}

type Id struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Picture struct {
	Large     string `json:"large"`
	Medium    string `json:"madium"`
	Thumbnail string `json:"thumbnail"`
}

type User struct {
	Gender     string `json:"gender"`
	Name       `json:"name"`
	StatusFlow string `json:"statusFlow"` // new, processing, processed, declined
	Location   `json:"location"`
	Email      string `json:"email"`
	Login      `json:"login"`
	Dob        `json:"dob"`
	Registered `json:"registered"`
	Phone      string `json:"phone"`
	Cell       string `json:"cell"`
	Id         `json:"id"`
	Picture    `json:"picture"`
	Nat        string `json:"nat"`
}

type Leads struct {
	Results []User `json:"results"`
}
