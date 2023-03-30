# Sprint 2
## Video Links

</br>

Frontend:
* Alaine  --  https://www.youtube.com/watch?v=EU7hBra51zk
* Lucinda --  https://youtu.be/EAqMi8WJ01s

Backend: https://youtu.be/1x1grRPc5sM

</br>

## Frontend Review

</br> 

__Alaine's Work__

The main work I've completed for this sprint involves the Find-Classes and Service-Page components. 

The Find-Classes component uses cdk-scrolling-viewport to let the user scroll through available classes without the back button at the bottom scrolling. At the moment, there are only 2 hard-coded classes in the Find Classes component, but it will pull from the backend in the future. 

The Service-Page component shows the details of the class that the user clicked on the Find Classes page. It then gives an option to click join or leave. This class is then added to the user's object, if the user isn’t already in the class, or removed if the user clicks leave.

In order to make these features work, I have also adjusted the routing to pass the user object around using the browser state, once the user logs in. The components can then check fields of the user object, such as username or the list of classes the user is signed up for.

The following unit tests were implemented:

For the Find-Classes component:

* it(should create) tests to make sure the component is created

* it(should navigate to summary) tests routing while passing user and service data.

* it (should navigate to profile) tests routing while passing user data.

For the service-page component:

* it(should create) tests the creation of the component.

* it(should navigate to find classes) tests routing to the Find Classes component while passing user data through the browser state.

* it(should join class) tests the "join" button. It checks that the component's service is added to the user's classes field.

* it(should leave class after joining) tests the "leave" button. It first calls join class, then leave class. It then makes sure that it cannot find that service in the user's classes field.

</br>

__Lucinda’s Work__

Implemented two new key pages in addition to improvements to the Onboarding component and the Registration component. Implemented the Business Onboarding component which tailors the welcome experience to what a business account needs (ie business hours, name, description, etc). This included adding unique checks such as time input checks. 

The Create Service component is a page accessible from the business owner’s dashboard which allows them to create a new class, lecture, appointment, etc. There are various input checks for this page which all merit their individual tests. 

New tests implemented:
* CreateServiceComponent

    * it('verifies that all fields are entered’)

    * it('adds to error message when not all fields filled in’)

    * it('checks that the specified start is before the end’)

    * it('returns error if end time is before start’)

* BusinessOnboardingComponent

    * it('changes which business tag is selected’)

    * it('checks that the stores opening time is before closing’)

    * it('returns error if closing time is before opening’)


</br>

## Backend Review

</br>

As stated in the submission doc for Sprint 2, all of the backend work for this application is being done by me (Ryan Brooks), since Angelina Uriarte-Wilson withdrew from the class just before the end of Sprint 2.

I focused primarily on the following areas for this sprint:

* Creating schemas and test records in Mockaroo for the various objects that are currently implemented, so that test data can be easily loaded and refreshed in the application so that test scenarios and user data is available to showcase and test both the frontend and backend functionality of the application.

* Created a generic ‘Model’ interface for all DB objects to implement in order to simplify testing efforts and automated data loads, and reduce the need for duplicate code/logic by utilizing generic function definitions.

* Continued refactoring code structure to make it more testable and easier to implement new object models for the database. This included updating existing functions for all previously implemented database objects in the ‘models’ and ‘handlers’ packages to implement the new generic ‘Model’ interface. 

* Continued documenting existing code with docstrings and updated process to automatically generate and format the formal documentation for the backend using markdown. This update makes the documentation easier to maintain on an ongoing basis and improves the readability / ease of access  for anyone reviewing the API or backend codebase. Markdown pages can be easily viewed within VS Code using various extensions and also displays cleanly within GitHub, circumventing the need to launch/host a separate documentation site for the backend.

* Explored the possibility of configuring a docker container for the backend server so that it’s easier to deploy on each dev’s local machine. I’m still uncertain whether this will be beyond the scope of the next sprint, but it’s something that I felt was worth exploring since we’ve already had to troubleshoot the frontend team’s local environments due to basic configuration issues.

* Implemented the following database object types in the ‘models’, ‘handlers’, ‘sample_data’ packages.

    * Office

    * Service

    * ServiceOffering

</br>

I attempted to alleviate the issues we were encountering last sprint with our backend architecture/project structure. I believe I managed to make a good deal of progress on this and that I'll be set up to be much more effective in the coming sprint with the new foundation I built. Loading sample records and testing will be significantly easier and managing the complexity as a single dev for the backend should be a bit more straightforward.

As for lessons learned, I learned a lot about implementing generic interface types and functions in Golang, but I also may have focused too much on pre-maturely optimizing code. It's really hard to say since it was imperative for me as a single-person dev team to be able to effectively implement new code, tests, and data loads given the scope of this application. Spending too much time debugging duplicative logic and managing spaghetti code for the sake of speed would likely quickly catch up to me. 


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
