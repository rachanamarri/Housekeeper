## SOFTAWRE ENGINEERING - CAPSTONE PROJECT

<h2> Housekeeper </h2>

<h3>Sprint-1</h3>

<h3>Sprint Window : 24 January 2022 - 3 february 2022</h3>

#### [Link for the demo video](https://drive.google.com/file/d/1v5XwKVeVptrQ6rPkIYJn17LWuMmc7mDl/view?usp=sharing)

<h2>FrontEnd Development</h2>

1. Created a Landing Page for the webapp.

2. Created a flexbox to display the services provided by the app in Landing page.

3. Added a router module to be able to route through the modules efficiently.

4. Added a resuable component to show the service description and price.

5. Created an angular service to mock services data and display the information about different services.

6. Added mock for book appointment functionality

<h2>Backend Development</h2> 

  <h3>Tech Stack: Golang, SQLite3</h3>
  
1. Database Development
	a. Created the database schema using the "struct"s. Initialising the database using the GORM(ORM library in golang)
	b. Populated the tables to store seekers, service_and_providers, logins and bookings details when the appropriate routes are hit.

2. Run the server on port :8080 using gin framework.

3. API Development : Routes for every different functionality
	1. /seeker_registration : Stores the seeker related data into the "seekers" table and the seeker's login credentials into logins table.
	2. /service_registration: Stores the provider and their service related data into the service_and_providers table and the provider's login credentials into 		logins table.
	3. /seeker_login: Authenticates and authorizes login for the seeker (through email id and password
	4. /provider_login: Authenticates and authorizes login for the provider (through email id and password)
	5. /services: List all available services
	6. /services/:ServiceId : To extract details of a particular service
	7. /services/:ServiceId/book : To book a specific service appointment for a user 

<h3>The whole project is accomplished as implementing user-stories and fixing the issues as mentioned in the github issues</h3>

