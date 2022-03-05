## SOFTWARE ENGINEERING PROJECT

## Sprint 2


## FrontEnd Development
**1. created and implemented cypress test case for landing page.**
**2. created and implemented a cypress test for checking reusable componenet service**
**3. created and implemented a cypress test for booking an appointment sevice**
**4. created the about us(contact) page.**


## Backend Development

### API Documentation

**1. Create Seeker POST API**
This API creates a seeker (user) with user details

```
POST /seeker_registration
```
**Example Request Body**
```
{
    "Name":"Bob", 
    "Email":"Bob@gmail.com",
    "Password":"random_password",
    "Address":"3800 SW 34th ST. GNV FL 32608"

}
```

**2. Create Provider POST API**
This API creates a provider (service provider) with service details

```
POST /service_registration
```
**Example Request Body**
```
{  
    "ServiceName":"SPA Treatment",
    "ProviderEmail":"spa@lakme.com", 
    "ProviderPassword":"random_password", 
    "ServicePrice":500,
    "ServiceDescription":"Leave the world behind & slip into a deep state of wellness & relaxation with massage, facials, nails & more."
}
```

**3. Seeker Login POST API**
This API checks login authentication of a seeker (user)

```
POST /seeker_login
```
**Example Request Body**
```
{
    "Email":"Bob@gmail.com",
    "Password":"random_password",
}
```

**4. Provider Login POST API**
This API checks login authentication of a service provider

```
POST /provider_login
```
**Example Request Body**
```
{
    "Email":"spa@lakme.com", 
    "Password":"random_password",
}
```

**5. List Services GET API**
This API list all the services present in the database

```
GET /services
```
**Example Response**
```
[
    {
        "ServiceId": 0,
        "ServiceName":"SPA Treatment",
        "ProviderEmail":"spa@lakme.com", 
        "ProviderPassword":"random_password", 
        "ServicePrice":500,
        "ServiceDescription":"Leave the world behind & slip into a deep state of wellness & relaxation with massage, facials, nails & more."
    },
    {
        "ServiceId": 1,
        "ServiceName": "Technician",
        "ProviderEmail": "support@quickfix.com",
        "ProviderPassword": "random_password",
        "ServicePrice": 500,
        "ServiceDescription": "Fix all gadgets of any kind without warranty and in 72 hours"
    },
]
```
**6. List Specific Service Details GET API**
This API lists a particular service with details present in the database

```
GET /services/:id
```
**Example Response**
```
{
    "ServiceId": 0,
    "ServiceName":"SPA Treatment",
    "ProviderEmail":"spa@lakme.com", 
    "ProviderPassword":"random_password", 
    "ServicePrice":500,
    "ServiceDescription":"Leave the world behind & slip into a deep state of wellness & relaxation with massage, facials, nails & more."
}
```
**7. Book Service POST API**
This API books a service for the user

```
POST /services/:id/book
```
**Example Request Body**
```
{
    "ServiceId": 0,
    "SeekerName": "Amy Santiago",
    "SeekerEmail": "amy@brooklyn99.com"
}
```


**Status Codes**
- 200: Status OK
- 400: Bad Request
- 500: Internal Server Error
- 404: Not Found

#### Backend Testing (Using Go Test)
To run the unit test, following command to be used - 
```
go test
```

  <h3>Tech Stack: Golang, SQLite3</h3>
  
1. Modelled the code in a Model-View-Controller Format
2. Unit Tests for backend using Go Test
3. Database Development
	a. Created the database schema using the "struct"s. Initialising the database using the GORM(ORM library in golang)
	b. Populated the tables to store seekers, service_and_providers, logins and bookings details when the appropriate routes are hit.

4. Run the server on port :8080 using gin framework.

5. API Development : Routes for every different functionality
	1. /seeker_registration : Stores the seeker related data into the "seekers" table and the seeker's login credentials into logins table.
	2. /service_registration: Stores the provider and their service related data into the service_and_providers table and the provider's login credentials into 		logins table.
	3. /seeker_login: Authenticates and authorizes login for the seeker (through email id and password
	4. /provider_login: Authenticates and authorizes login for the provider (through email id and password)
	5. /services: List all available services
	6. /services/:ServiceId : To extract details of a particular service
	7. /services/:ServiceId/book : To book a specific service appointment for a user 

