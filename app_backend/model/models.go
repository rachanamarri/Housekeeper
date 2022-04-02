package model

import "github.com/markphelps/flipt/storage/sql/common"

type Seeker struct {
	Name     string `json:"name"`
	Email    string `json:"Email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ServiceAndProvider struct {
	ServiceId          int64  `json:"ServiceId"`
	ServiceName        string `json:"ServiceName"`
	ProviderEmail      string `json:"ProviderEmail"`
	ProviderPassword   string `json:"ProviderPassword"`
	ServicePrice       int64  `json:"ServicePrice"`
	ServiceDescription string `json:"ServiceDescription"`
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
