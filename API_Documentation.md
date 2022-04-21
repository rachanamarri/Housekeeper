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

##### Example Responses:

```
{
    "Name":"Bob", 
    "Email":"Bob@gmail.com",
    "Password":"",
    "Address":"3800 SW 34th ST. GNV FL 32608"
}
```

**2. Create Provider POST API**
This API creates a provider (service provider) with provider details

```
POST /provider_registration
```
**Example Request Body**
```
{   
    "ProviderId" : 0
    "Name":"SPA Treatment",
    "Email":"spa@lakme.com", 
    "Password":"random_password", 
    "Description":"Leave the world behind & slip into a deep state of wellness & relaxation with massage, facials, nails & more.",
    "Address":"400 SW 38th Street"
}
```
##### Example Responses:

```
{   
    "ProviderId" : 0
    "Name":"SPA Treatment",
    "Email":"spa@lakme.com", 
    "Password":"", 
    "Description":"Leave the world behind & slip into a deep state of wellness & relaxation with massage, facials, nails & more.",
    "Address":"400 SW 38th Street"
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
##### Example Responses:

```
{
    "message": "Log in success",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im1pdGFsaTMxMTJAZ21haWwuY29tIiwiZXhwIjoxNjUwNTkzMjkzfQ.mfQsZ0WRCg6mlo5Lbmp1DcPZDIO7NokJ6_qaVreZmbw"
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
##### Example Responses:

```
{
    "message": "Log in success",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im1pdGFsaTMxMTJAZ21haWwuY29tIiwiZXhwIjoxNjUwNTkzMjkzfQ.mfQsZ0WRCg6mlo5Lbmp1DcPZDIO7NokJ6_qaVreZmbw"
}
```

**5. List Providers GET API**
This API list all the services present in the database

```
GET /providers
```
**Example Response**
```
[
    {
        "ProviderId": 0,
        "Name": "Electrician",
        "Email": "electrician@onestop.com",
        "Password": "",
        "Address": "400 SW 38th Street"
    },
    {
        "ProviderId": 0,
        "Name": "Lakme Beauty Services",
        "Email": "beauty@lakme.com",
        "Password": "",
        "Address": "400 SW 38th Street"
    }
]
```
**6. Specific Service Details GET API**
This API lists a particular service with details present in the database

```
GET /services/:id/list
```
**Example Response**
```
{
    "ServiceId": 3,
    "ProviderId": 1,
    "Name": "Fix Sockets",
    "Price": 50,
    "Description": "Fix any wall socket in your house"
}
```
**7. Book Service GET API**
This API books a service for the user

```
GET /services/:id/book
```

##### Header:

    Authorization: Access token

**Example Request Body**
```
{   
    "ServiceId": 1, 
    "SeekerName":"Sheth", 
    "SeekerEmail": "mitali123@gmail.com" 
}
```
**Example Response**
```
{
    "ServiceId": 1,
    "ProviderId": 1,
    "Name": "Fix Sockets",
    "Price": 50,
    "Description": "Fix any wall socket in your house"
}

```

**8. Rate Service POST API**
This API rate a service for the user

```
POST /services/:id/rate_service
```
##### Header:

    Authorization: Access token
    
**Example Request Body**
```
{
    "Rating": 4
}
```
**Example Response**
```
{
    "ServiceId": 1,
    "ServiceName": "Fix a bulb",
    "Rating": 4
}

```

**9. Create Service POST API**
This API creates a service for the user

```
POST /service_registration
```
**Example Request Body**
```
{
    "ServiceId": 0, 
    "ProviderId": 1, 
    "Name":"Fix Sockets", 
    "Price" : 50, 
    "Description":"Fix any wall socket in your house"
}
```
**Example Response**
```
{
    "ServiceId": 1,
    "ProviderId": 1,
    "Name": "Fix Sockets",
    "Price": 50,
    "Description": "Fix any wall socket in your house"
}

```

**10.Create Email GET API**
This API creates a service for the user

```
GET /services/:SeekerName/emailservice
```
##### Header:

    Authorization: Access token
    

**Example Response**
```
{
    "message": "Email sent",
    "result": true
}

```

**Status Codes**
- 200: Status OK
- 400: Bad Request
- 500: Internal Server Error
- 404: Not Found
