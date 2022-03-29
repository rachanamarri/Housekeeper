## SOFTWARE ENGINEERING PROJECT

<h2> Housekeeper </h2>

<h3>Sprint-2 Demo Video</h3>
<!-- #### [Link for the demo video](https://drive.google.com/file/d/1v5XwKVeVptrQ6rPkIYJn17LWuMmc7mDl/view?usp=sharing)
 -->
<h2>FrontEnd Development</h2>

  <h3>Tech Stack:Angular, Cypress</h3>

1. Integration of Frontend with Backend

2. Used Cypress Testing

3. Run Unit Tests Angular

<h2>Backend Development</h2> 

  <h3>Tech Stack: Golang, SQLite3</h3>
  
1. Modelled the code in a Model-View-Controller Format
2. Database Development
	a. Created the database schema using the "struct"s. Initialising the database using the GORM(ORM library in golang)
	b. Populated the tables to store seekers, service_and_providers, logins and bookings details when the appropriate routes are hit.
3. Run server on port :8080 using gin framework.
4. API Development : Routes for every different functionality
	1. /seeker_registration : Stores the seeker related data into the "seekers" table and the seeker's login credentials into logins table.
	2. /service_registration: Stores the provider and their service related data into the service_and_providers table and the provider's login credentials into 		logins table.
	3. /seeker_login: Authenticates and authorizes login for the seeker (through email id and password
	4. /provider_login: Authenticates and authorizes login for the provider (through email id and password)
	5. /services: List all available services
	6. /services/:ServiceId : To extract details of a particular service
	7. /services/:ServiceId/book : To book a specific service appointment for a user 