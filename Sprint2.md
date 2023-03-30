# Sprint 2
## Video Links

</br>

Frontend: https://www.youtube.com/watch?v=5NSP9TG9MUE

Backend: https://youtu.be/HcWRBffzxN4 

</br>

## Frontend Review
For layout work in the front-end, we’ve polished the registration page by adding an onboarding page that it navigates to, and we’ve created a profile page that welcomes the user and is meant as a landing page once the user logs in. As for connecting to the back-end, we’ve created a User service and a Config service that can add a user by posting to the back-end when the user registers on the register component page.

For the unit tests, each component has a “should be created” test that checks that the component is created. For components with more complex functionality, additional unit tests were added. The Login Component includes unit tests for routing using the spyOn function. Any page that includes navigation includes unit tests checking their individual routes. The Register Component contains several unit tests for checking to make sure that valid inputs are placed into the fields. The Onboarding Component has an additional unit test to ensure that selecting the interests causes them to properly toggle. 

Our Cypress tests for this sprint focus on the Login Component. There are two tests. The first checks if the button is disabled while the fields are empty. The second checks to ensure that the button becomes enabled when both the username and password fields are filled.

</br>

## Backend Review
The backend team focused primarily on the following areas for this sprint:

Integrating the backend with the frontend
Adding logging functionality to the backend to record what requests are being made
Adding basic CRUD methods/logic to some of our DB models
Documenting the current and immediate future state of the API
Refactoring our project structure to make it testable

We began this sprint focusing on how to connect the frontend to the backend and were able to successfully setup the necessary CORS, reverse proxy, and handler functionality after struggling for a bit and exploring/testing different options. We had to rewrite some of our server initialization code to be able to incorporate those changes.

We then decided that logging request data to the server terminal would help us to troubleshoot both backend and frontend issues, so we created a new function in our middleware package called “RequestLoggingMiddleware”. This logs the request type and route of the inbound API request to the server terminal and also logs the request body if it’s a PUT or POST request. This helped our testing and troubleshooting efforts tremendously.

We also picked 3 key models to create CRUD methods for, for this sprint: User, Business, and Address. Each of these are key building blocks for our application and we felt they were the best entry point for us to begin the process of fleshing out our database calls/logic. This sprint allowed us to begin figuring out how to structure that code/logic and how best to serve that data back to the frontend.

For our unit testing this sprint, we created the following tests:

</br>

* TestHashPassword  –  Testing password hashing function for user passwords
* TestCheckPassword  –  Testing password confirmation function
* TestParseRequestID  –  Testing function to parse ID from request URL
* TestRespondWithJSON  –  Testing JSON response function that returns json body
* TestRespondWithError  –  Testing JSON response function that returns error body
* CRUD tests for Address, User, Business
    * Create/Get function tests are currently functional
    * Update and Delete function tests are currently just placeholders, and not yet complete

</br>

We encountered difficulties with our original architecture/project structure for the backend for a good portion of this sprint. We originally had the DB logic written into the handler functions for each of our API routes, but couldn’t figure out a way to test those functions as they were written so we opted to refactor our code a bit. We’re now in a spot where we can begin to build out a more comprehensive suite of tests, so overall it was a good learning experience and it should allow us to be a bit more productive moving forward.

Finally, we added various docstrings throughout our code base using both basic comments and the OpenAPI specification to document the backend for this project. Ultimately, we want to have all of our functions documented using the OpenAPI docstrings standard, but for now we opted for expediency and kept some simple comment documentation scattered throughout. We used `godoc` to automatically generate HTML documentation based on our docstrings and also created a “_documentation” folder to store various documentation for the backend, including the database schema we will be using.

As for lessons learned, we realized that we spent too much time thinking about and working on the authentication/authorization for our application this early in the process. It was a larger undertaking than either of us realized and only served to add unnecessary complexity to the project and our testing efforts at this stage. We realized that we should wait to address those requirements until we have all of the basic functionality of our application fleshed out first. We also learned about the importance of branching this sprint since there were a couple of times where one of us pushed breaking changes that should have been addressed in a separate branch until they were ready to merge.


</br>


__ADDITIONAL NOTE REGARDING BACKEND TEAM (as of 3/1):__


My (Ryan Brooks) 2nd child was born on 2/9 (a full month earlier than expected), so my plans for this sprint kind of went out the window. Her and my wife are both healthy now, but we had a few challenging circumstances and changes to deal with on short notice. I was still able to participate and contribute for this sprint, but it definitely had a significant impact on my availability and productivity.

Additionally, our team received a message from Angelina Uriarte-Wilson at 5PM on Wednesday (3/1) letting us know that she is dropping the class. None of us were aware that she was considering dropping the course prior to today, but that may impact our approach to and plans for this project moving forward. We had a few other candidate ideas for apps and, depending on our discussion in the next few days, we may decide to pivot a bit since Angelina was the main advocate for the business app to begin with.

</br>


# Backend documentation

The documentation we have created for the backend can also be found in the `src/server/_documentation` directory of this repo and within the various 'README.md' docs of each package directory.


</br>

###  __Unit Tests__
| **Function Name**    | **Description**                                                                                                                                                                                                   | **Implementation Status** | **Latest Test Status** |
|----------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------|------------------------|
| TestHashPassword     | Tests the HashPassword function to ensure that user passwords are being hashed as expected                                                                                                                        | Complete                  | Pass                   |
| TestCheckPassword    | Tests the CheckPassword function to ensure the function is correctly confirming when passwords match and throwing the appropriate error/response when they don't                                                  | Complete                  | Pass                   |
| TestParseRequestID   | Tests the ParseRequestID function to confirm that the ID field from the request URL is parsed into uint format and that the appropriate error is returned if the ID is missing or formatted incorrectly.          | Complete                  | Pass                   |
| TestRespondWithJSON  | Tests the RespondWithJSON function and ensures that the response being returned by the function is formatted correctly and returns what is expected                                                               | Complete                  | Pass                   |
| TestRespondWithError | Tests the RespondWithError function and ensures that the response being returned by the function is formatted correctly and returns what is expected                                                              | Complete                  | Pass                   |
| TestCreateBusiness   | Tests that the CreateBusiness function returns the created Business object and that the record is created in the application  database                                                                            | Complete                  | Pass                   |
| TestGetBusiness      | Tests that the GetBusiness function returns the appropriate Business object from the database or throws the appropriate error if the record doesn't exist in the database                                         | Placeholder               | N/A                    |
| TestUpdateBusiness   | Tests that the UpdateBusiness function returns the updated Business object from the database or throws the appropriate error if the record doesn't exist in the database                                          | Placeholder               | N/A                    |
| TestDeleteBusiness   | Tests that the DeleteBusiness function returns the deleted Business object from the database, confirms that the record has been deleted from the DB, or throws the appropriate error if the record doesn't exist. | Placeholder               | N/A                    |
| TestCreateAddress    | Tests that the CreateAddress function returns the created Address object and that the record is created in the application  database                                                                              | Complete                  | Pass                   |
| TestGetAddress       | Tests that the GetAddress function returns the appropriate Address object from the database or throws the appropriate error if the record doesn't exist in the database                                           | Complete                  | Pass                   |
| TestUpdateAddress    | Tests that the UpdateAddress function returns the updated Address object from the database or throws the appropriate error if the record doesn't exist in the database                                            | Placeholder               | N/A                    |
| TestDeleteAddress    | Tests that the DeleteAddress function returns the deleted Address object from the database, confirms that the record has been deleted from the DB, or throws the appropriate error if the record doesn't exist.   | Placeholder               | N/A                    |
| TestCreateUser       | Tests that the CreateUser function returns the created User object and that the record is created in the application  database                                                                                    | Complete                  | Pass                   |
| TestGetUser          | Tests that the GetUser function returns the appropriate User object from the database or throws the appropriate error if the record doesn't exist in the database                                                 | Placeholder               | N/A                    |
| TestUpdateUser       | Tests that the UpdateUser function returns the updated User object from the database or throws the appropriate error if the record doesn't exist in the database                                                  | Placeholder               | N/A                    |
| TestDeleteUser       | Tests that the DeleteUser function returns the deleted User object from the database, confirms that the record has been deleted from the DB, or throws the appropriate error if the record doesn't exist.         | Placeholder               | N/A                    |
| TestAuthenticate     | Tests the Authenticate function is working correctly and updates the session attributes appropriately                                                                                                             | Placeholder               | N/A                    |
