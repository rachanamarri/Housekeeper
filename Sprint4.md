## SOFTWARE ENGINEERING PROJECT

<h2>HOUSEKEEPER</h2>
<h2>Sprint-4 Deliverables</h2>

<h3>Sprint Window: 1 April 2022 - 20 April 2022</h3>

## Frontend Development
- Added an about us Page to contact in case a user cannot book an appointment
- Created cypress test to register as seeker.
- Created cypress test to register as provider.
- Created cypress test to login as seeker.
- Added a date picker and time slot picker for ease of booking an appointment
- Added component tests for new components.
- fixed any bugs present.

## Backend Development

#### 1. Create Service API

This API is used for creating a new service.

```
POST  /service_registration
```

#####  Example Request Body:

```
{
    "ServiceId": 0, 
    "ProviderId": 1, 
    "Name":"Fix Sockets", 
    "Price" : 50,
    "Description":"Fix any wall socket in your house"
}

```

##### Example Responses:

```
{
    "ServiceId": 24,
    "ProviderId": 1,
    "Name": "Fix Sockets",
    "Price": 50,
    "Description": "Fix any wall socket in your house"
}
```

##### Status Codes:

-   **200**: Status OK
-   **400**: Bad Request
-   **500**: Internal Server Error
-   **503**: Service Unavailable

#### 2. Post Ratings API

This API is used to give ratings on any service.

```
POST  /services/:ServiceId/rate_service
```
#####  Example Request Body:

```
{
    "Rating": 4
}
```
##### Example Responses:

```
{
    "ServiceId": 1,
    "ServiceName": "Fix a bulb",
    "Rating": 4
}
```

##### Status Codes:

-   **200**: Status OK
-   **400**: Bad Request
-   **500**: Internal Server Error
-   **503**: Service Unavailable

#### 3.Create Email GET API

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

##### Status Codes:

-   **200**: Status OK
-   **400**: Bad Request
-   **500**: Internal Server Error
-   **503**: Service Unavailable

## Back-end tests

To run the unit test, following command is to be used -

    go test -v

1. CreateServiceAPI - To test creating a new service by the provider/admin
2. CreateRatingsAPI - To test posting ratings for a service by the user


## Video to demonstrate HOUSEKEEPER functionality
https://www.youtube.com/watch?v=vT3bdijZuNc

## Video to demonstrate Cypress-Tests on the Frontend
- cypress test to login:
- https://youtu.be/jOmd3ZZmMPc
- cypress test to register:
- https://youtu.be/JNo8rSxO9Zg
- https://youtu.be/OT-BKOpcpFs

## Video to demonstrate Unit-Tests on the Backend

## BONUS
- Deployed App Url: [https://lbarad.github.io/Housekeeper/](https://lbarad.github.io/Housekeeper/)
#### [API Documentation]()

#### [Project Board](https://github.com/mitali3112/Housekeeper/projects)

#### [Sprint-4 Deliverables](https://github.com/mitali3112/Housekeeper/blob/main/Sprint4.md)

## Team Members
### Frontend
1. Rachana Reddy Marri
2. Yasaswini Valisetty
### Backend
1. Nitheesha Reddy Beereddy
2. Mitali Sheth



