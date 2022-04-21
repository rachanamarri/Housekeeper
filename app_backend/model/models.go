package model

/* Importing git repository*/

import "github.com/markphelps/flipt/storage/sql/common"

type Seeker struct {

	Name     string `json:"Name"`
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

=======
	Password string `json:"Password"`
	Address  string `json:"Address"`
}

type Login struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type Provider struct {
	ProviderId int64  `json:"ProviderId"`
	Name       string `json:"Name"`
	Email      string `json:"Email"`
	Password   string `json:"Password"`
	Address    string `json:"Address"`
}

type Service struct {
	ServiceId   int64  `json:"ServiceId"`
	ProviderId  int64  `json:"ProviderId"`
	Name        string `json:"Name"`
	Price       int64  `json:"Price"`
	Description string `json:"Description"`

}

type Booking struct {
	ProviderId  int64  `json:"ProviderId"`
	ServiceId   int64  `json:"ServiceId"`
	SeekerName  string `json:"SeekerName"`
	SeekerEmail string `json:"SeekerEmail"`

}

type Store struct {

	*common.Store
}

type Ratings struct {

	ServiceID   int64  `json:"ServiceId"`
	ServiceName string `json:"ServiceName"`
	Rating      int64  `json:"Rating"`
}



















































