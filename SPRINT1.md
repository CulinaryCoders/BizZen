# Sprint 1

## Video Links
Frontend: https://www.youtube.com/watch?v=o8U-HqkyJN4
Backend: https://www.youtube.com/watch?v=mbtpTFUQO1Y

## User Stories
* **Create and manage user accounts**
  * As a **user**, I want to be able to create a new user account or login to an existing user account in order to access the application.
* **User / Account Profiles**
  * As a **client** user, I should be able to update, manage, and view my personal profile so that I can update my preferences, contact information, and monitor my various activities within the application.
  * As a **business owner**, I should be able to update, manage, and view the profile for my business on the application so that I can present/market my business and manage various aspects of my operations and administration.

* **Online Scheduling**
  * As a **customer**, I want to be able to schedule appointments with businesses online, so that I can save time.
  * As a **customer**, I want to be able to see my past and upcoming appointments, so that I can easily manage my schedule.
  * As a **business owner**, I want customers to be able to schedule appointments online, so that I can manage my appointments better.
  * As a **business owner**, I want to be able to see my past and upcoming appointments, so that I can easily manage my schedule and services.

* **Reviews and Feedback**
  * As a **customer**, I want to be able to leave reviews or feedback in order to let businesses know what parts of their services might need adjusting.
  * As a **business owner**, I want to be able to see customer reviews so I know what my customers think about the business, and adjust things if possible.

* **Service Management**
  * As a **business owner**, I want to be able to create and edit an appointment or class so that participants can sign up for my classes that they are interested in.


## Issues Completed
For Sprint 1, we focused on the Create and manage user accounts story. In the backend, we started by setting up basic outline for the backend and setting up boilerplate code for the server. We successfully implemented code to register a user by saving their credentials into a PostgreSQL database hosted locally. A user object was created with the help of GORM to map the user object into Postgres. The bcrypt package was used to hash the new user’s password and store it securely in the database. We used Gorilla Mux for handling routing and creating a basic REST API to handle CRUD operations on user objects. For logging a user in, we implemented an endpoint that checks that the user’s login credentials match what is stored in the local instance of Postgre. The login endpoint also generates and responds to the client with a signed JSON web token, which could be used for authorization. We are still currently researching other authorization and session management practices and may move away from JWTs moving forward. In addition, we worked on creating potential database schemas for different types of user accounts and started to implement functionality for creating secured routes and role-based access.
<br/><br/>
For the front-end, we successfully created a landing page, a registration page, and a log-in page. Each of these pages can route to one another, and are all designed using bootstrap. The login page contains fields for username and password, and a dropdown that allows the user to select whether they are a customer or a business. The registration page allows the user to enter in their basic information, including first name, last name, a password, and a confirmation of their password. The next step is to create a welcome sequence to personalize their account.


## Issues Not Completed
Many of the user stories that we wrote we have yet to implement; we planned several user stories as long-term ideas to focus on what we’re working on and give each team member something to work on. 
<br/><br/>
Issues we were not able to address include connecting the front and back end and creating separate registration for customers and businesses. While these are still planned to be implemented, we were unable to implement them in Sprint 1 due to time and planning constraints. We have since mapped out the database schema, so the front-end and back-end should be able to begin working on connecting to one another while still mostly working independently. 
