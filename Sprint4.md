# Sprint 4 Review

## Video Links

</br>

Frontend: https://www.youtube.com/watch?v=6K9OA_iKL58

Backend:  https://youtu.be/EJ49sZGTnAI

</br>

## Frontend Sprint Review

</br>

__Alaine's Work__

The main work I've completed for this sprint involves the service-page component, the view-appointments component, and fully connecting individual parts to the back-end. 

The service-page component has been updated from the previous sprint to include an Edit button. This button can both save and cancel the edits. When the user saves an edit, it communicates with the back-end to update the service using a PUT http request. 

The service-page also includes separate views for both the business and the user. Editing is only allowed on the business view. Businesses can also see what users have signed up by their email addresses. For the user view, users both joining and leaving the service have been connected to the back-end with POST http requests.

Another small change to the service-page was adding an additional variable that the browser state tracks when routing to this component. Since two separate pages route to this component, the back button will return to the appropriate page based on what was passed to the browser state.

As for the view-appointments component, this reuses the scrolling functionality from the find-classes component, whose functionality has been replaced and upgraded by the business-dashboard created by Lucinda. This page uses a back-end route to GET all services attached to a user ID. The user can then reach the service-page component, and join and leave classes from there.

I also updated the login page so that it only logs in when the back-end confirms that the user exists in the database. 


__New unit tests for Sprint 4__

* For the Login Component:
    * it(should pass user when logging in successfully)

        * This tests the new “successfulLogin” function, which is called when the back-end successfully retrieves a user from the database. It checks to make sure that a test user is passed.

    * it(should say user does not exist when unsuccessfully logging in)

        * This ensures that a boolean is set to false when the user does not exist, meaning that HTML elements will appear to tell the user that their log in failed.

* For the Profile Component:
    * it(should navigate to classes) makes sure that it properly navigates to /home when the “Go” beneath “Find Classes” is clicked.

    * it(should navigate to view appointments) makes sure that it properly navigates to /view-appointments.

    * it(should give the testUser’s first name) makes sure that the “Welcome” message at the top of the page shows the inputted user’s name.

* For the View Appointments Component:
    * it(should navigate to profile) checks to make sure that the back button routes properly.
    
    * it(should navigate to a passed in service) checks to make sure that it correctly routes to the service-page component with the correct browser state passed in.

* For the Service Page Component:
    * it(should update service) checks to make sure that an http PUT request is made upon saving an edit.
    
    * it(should leave class after joining) checks to make sure the boolean that tracks whether the user has joined the class changes appropriately. It also makes sure that there is a POST request when the user joins the class.

    * it(should edit) makes sure the boolean for editing is appropriately updated.
    
    * it (should navigate to find classes) makes sure it routes properly back to the business dashboard.
    
    * it(should cancel edit) makes sure the boolean is updated when cancel edit is clicked.
    
    * it(should join class) makes sure there is a POST http request when the user joins the service.


* For end-to-end Cypress testing:
business-view.cy.ts
it(should edit classes)
This test logs in as a business, navigates to the service “Test Service,” clicks on it, edits the description of the service, and saves the edit.
it(should add a new class)
This test logs in as a business, navigates to the add service page through the business dashboard, fills in the appropriate data, and adds a new service.
register-login.cy.ts
it(should register and log in as a user)
This test fills in all the fields to register a user account, then logs in using that data.
it(should register and log in as a business)
This test registers a business account, including clicking the checkbox for businesses, and then logs in using that data.
view-classes.cy.ts
it(should log in as a user and join a class)
This test logs in as a user and navigates to the business dashboard, where it clicks the “more info” button to see the details of the “Test Service” class. It then clicks the “Join” button, navigates back to the profile, and goes to the “View Appointments” page to see the new service that it joined.

Lucinda’s Work
Redesigned several pages for consistent user interface, including Create Service page.
Created the Business Dashboard page which combined several elements, including getting all services from the database, formatting them into a readable list, and feeding them into a calendar.
Several sub-features were implemented such as a search feature and the ability to go to a service’s info page by clicking it on the calendar.
The calendar component was created from the angular-calendar library, which takes in CalendarEvent objects and displays them in a calendar format with features such as month/week/day view, and movement between months, weeks, and days. This requires the database data to be formatted specifically to fit this object model so that the calendar data is always current with the list of Services displayed.
Major work connecting frontend pages to backend endpoints, which included debugging and communicating minor adjustments to the database and the endpoints.
New tests implemented:
Business Dashboard
  it('Completes GET request to fetch list of services from db')
  it('Filters services by ascending date')
  it('Should navigate to create service')
  it('Returns correct Date-Time from start date and duration')
  it('Should navigate to service info page')
  it('Updates the view date range for services to be filtered by')
Landing Component
  it('Should navigate to Register page')
  it('Should navigate to Login page')
Navbar Component
  it('Should navigate to Profile page')
  it('Should navigate to Landing page')
  it('Should navigate to Business Dashboard page')
Register component (modified/updated)
  it('should navigate to home')
  it('should check that all fields are filled in')
  it('should catch when not all fields filled in'
  it('check for matching passwords')
  it('check for MISmatching passwords')
Service Calendar
  it('Should set calendar view to given size')


## Backend Sprint Review

</br>

As stated in the previous 2 submission docs, all of the backend work for this application is being done by me (Ryan Brooks), since Angelina Uriarte-Wilson withdrew from the class just before the end of Sprint 2.

I focused primarily on the following areas for Sprint 4:

* Added additional routes and methods to the backend:

    * Routes added this sprint:

    /users  -  Get a list of all User records in the database

    /user/{id}/service-appointments  -  Get all appointments a user has made (along with their associated Service object)

    /business/{id}/services  -  Get a list of all the services that a business has created.

    /business/{id}/service-appointments  –  Get a list of all the appointments a business has for the services that they’ve created

    /services  –  Get a list of all Service records in the database

    /service/{id}/users  –  Get a list of all the users with an active appointment for the specified service 

    /service/{id}/user-count –  Get a count of all the users with an active appointment for the specified service 

    /service/{service-id}/user/{user-id}  -  Check whether User already has an appointment for the specified Service.

    /service/{id}/appointments
    /service/{id}/appointments/active  –  Get a list of all active Appointment records for the specified service.

    /service/{id}/appointments/all   –  Get a list of all (active and inactive) Appointments records in the database for the specified service.

    /appointment/{id}/cancel  –  Cancel an appointment

    /appointments
    /appointments/active   –  Get a list of all active Appointment records in the database

    /appointments/all  –  Get a list of all (active and inactive) Appointments records in the database

    /invoice  –  Basic CRUD methods

    /invoices  – Get a list of all Invoice records in the database

    * Adding hooks and data standardization methods to various objects. Utilized the standard hooks in GORM to automatically perform database operations on related objects during certain transactions. This included deleting all of the Invoice records tied to an Appointment when the Appointment record was deleted, Appointment records that were tied to a Service being deleted, and deleting all of the Service records associated with a Business being deleted. Additionally, I added logic to keep certain attribute fields updated when certain actions are performed. This included updating the Appointment count attribute and boolean flag for each Service whenever an Appointment was added; creating a Business object when a User with a Business account type is created; and updating the Status attribute of an invoice based on the remaining balance field anytime an Invoice record was created or updated.

    * Revamping scope of work for final MVP and pivoting based on the needs and scope of work completed by the frontend team. Our initial scope was pretty large for this project at the beginning of the semester, so we pivoted towards the end of Sprint 3 and for all of Sprint 4 to focus on the core functionality that we felt gave us a useful application. This involved retailoring the database schema a bit and focusing on 5 main objects (Users, Businesses, Services, Appointments, and Invoices).

    * Continued documentation and testing efforts. Continued documenting new and existing code with docstrings. Markdown pages can be easily viewed within VS Code using various extensions and also displays cleanly within GitHub, circumventing the need to launch/host a separate documentation site for the backend.

    * Updated test schemas in Mockaroo and added new logic to produce more clean/appropriate test data (rounded dollar amounts, appointment lengths in increments of 30 minutes, etc.). 

    * Added routes to serve a basic table of contents page with reference links to documentation and helpful resources used to build and configure the backend.

    * Worked with frontend team to identify and troubleshoot various bugs and discuss various behaviors and error handling, so that they could be appropriately implemented or updated on the backend


This sprint was spent focusing on finalizing our MVP and recentering our focus on the most important backlog items in our remaining scope. Overall, I think this sprint was our most productive one as a team and it was really neat to see our app come together and our efforts start to materialize with a functional piece of software. I continued to document and refine the codebase and completed the backend logic for our core MVP. 

If this were a real application, I think we’d have a pretty good proof of concept to present and to decide if the project was worth pursuing further. Overall, this process has definitely been a humbling and valuable learning experience for me. It was interesting to see the struggles that my team had with communication and I have a much greater appreciation for how difficult it is to manage a greenfield project in its initial stages.
