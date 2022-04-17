package model

import "github.com/markphelps/flipt/storage/sql/common"

type Seeker struct {
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Address  string `json:"Address"`
}

type Login struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type Provider struct {
	ProviderId  int64  `json:"ProviderId"`
	Name        string `json:"Name"`
	Email       string `json:"Email"`
	Password    string `json:"Password"`
	Address     string `json:"Address"`
	PhoneNumber string `json:"PhoneNumber"`
}

type Service struct {
	ServiceId   int64  `json:"ServiceId"`
	ProviderId  int64  `json:"ProviderId"`
	Name        string `json:"Name"`
	Price       int64  `json:"Price"`
	Description string `json:"Description"`
}

type Booking struct {
	ServiceId   int64  `json:"ServiceId"`
	SeekerName  string `json:"SeekerName"`
	SeekerEmail string `json:"SeekerEmail"`
}

type Store struct {
	*common.Store
}

type Ratings struct {
	ServiceID     int64  `json:"ServiceId"`
	ProviderEmail string `json:"ProviderEmail"`
	Rating        int64  `json:"Rating"`
}
