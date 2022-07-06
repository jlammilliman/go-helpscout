package api

import "time"

type Customer struct {
	Id           int
	FirstName    string
	LastName     string
	PhotoUrl     string
	PhotoType    string
	Gender       string
	Age          string
	Organization string
	JobTitle     string
	Location     string
	CreatedAt    time.Time
	ModifiedAt   time.Time

	Background     string
	Address        CustomerAddress
	SocialProfiles []CustomerSocialProfile
	Emails         []CustomerEmail
	Phones         []CustomerPhone
	Chats          []CustomerChat
	Websites       []CustomerWebsite
}

type CustomerWebsite struct {
	Id    int
	Value string
}

type CustomerSocialProfile struct {
	Id     int
	Value  string
	TypeOf string
}

type CustomerPhone struct {
	Id       int
	Value    string
	Location string
}

type CustomerEmail struct {
	Id       int
	Value    string
	Location string
}

type CustomerChat struct {
	Id     int
	Value  string
	TypeOf string
}

type CustomerAddress struct {
	Id         int
	Lines      []string
	City       string
	State      string
	PostalCode string
	Country    string
	CreatedAt  time.Time
	ModifiedAt time.Time
}
