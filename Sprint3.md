## SOFTWARE ENGINEERING - CAPESTONE PROJECT

<h3>Housekeeper</h3>

<h3>Sprint-1</h3>

<h3>Sprint Window: 18 March 2022 - 1 April 2022</h3>

#### [Link for Demo video](https://drive.google.com/file/d/1GKkM8iadspTaJr-3h-YUN5BQVMffEiOZ/view?usp=sharing)

<h2>FrontEnd Development</h2>

1. Server Login

2. Creating the Profile page and storing customer details

3. Angular Project for Service Provider Dashboad

4. Seeker UI Page completion

5. Documentation

6. Integration of Backend and Frontend

<h2>Backend Development</h2> 

1. Hashed Passwords

2. Implemented JWT

3. Integrated Frontend and Backend using CORS middleware

4. Implemented additional functionality
### API Documentation

**1. Rate Service POST API :**
    This API will store the ratings of the user(seeker) given to the service that was provided.
```
POST /services/:ServiceId/rate_service
```
**Example Request Body**
    ```
    {
        "ServiceId":3
        "ProviderEmail":"alex20@gmail.com"
        "Rating":5
    }

