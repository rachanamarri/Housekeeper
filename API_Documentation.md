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