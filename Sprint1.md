## SOFTAWRE ENGINEERING - CAPESTONE PROJECT

<h2> Housekeeper </h2>

<h3>Sprint-1</h3>

<h3>Sprint Window : 24 January 2022 - 3 february 2022</h3>



FrontEnd

1. Created a Landing Page for the webapp.

2. Created a flexbox to display the services provided by the app in Landing page.

3. Added a router module to be able to route through the modules efficiently.

4. Added a resuable component to show the service description and price.

5. Created an angular service to mock services data and display the information about different services.

6. Added mock for book appointment functionality

Backend Development 

  Tech Stack: Golang, SQLite3

1. Database Development - Created database with tables to store Seeker, Provider, Login and Booking Details

2. API Development : Routes for every different functionality
	1. /seeker_registration : Stores the seeker related data into the database
	2. /service_registration: Stores the provider and their service related data into the database
	3. /seeker_login: Authenticates and authorizes login for the seeker (through email id and password
	4. /provider_login: Authenticates and authorizes login for the provider (through email id and password)
	5. /services: List all available services
	6. /services/:ServiceId : To extract details of a particular service
	7. /services/:ServiceId/book : To book a specific service appointment for a user 

