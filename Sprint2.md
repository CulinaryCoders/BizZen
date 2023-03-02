# Sprint 2
## Frontend Review
For layout work in the front-end, we’ve polished the registration page by adding an onboarding page that it navigates to, and we’ve created a profile page that welcomes the user and is meant as a landing page once the user logs in. As for connecting to the back-end, we’ve created a User service and a Config service that can add a user by posting to the back-end when the user registers on the register component page.

For the unit tests, each component has a “should be created” test that checks that the component is created. For components with more complex functionality, additional unit tests were added. The Login Component includes unit tests for routing using the spyOn function. Any page that includes navigation includes unit tests checking their individual routes. The Register Component contains several unit tests for checking to make sure that valid inputs are placed into the fields. The Onboarding Component has an additional unit test to ensure that selecting the interests causes them to properly toggle. 

Our Cypress tests for this sprint focus on the Login Component. There are two tests. The first checks if the button is disabled while the fields are empty. The second checks to ensure that the button becomes enabled when both the username and password fields are filled.

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

* TestHashPassword  –  Testing password hashing function for user passwords
* TestCheckPassword  –  Testing password confirmation function
* TestParseRequestID  –  Testing function to parse ID from request URL
* TestRespondWithJSON  –  Testing JSON response function that returns json body
* TestRespondWithError  –  Testing JSON response function that returns error body
* CRUD tests for Address, User, Business
    * Create/Get function tests are currently functional
    * Update and Delete function tests are currently just placeholders, and not yet complete

We encountered difficulties with our original architecture/project structure for the backend for a good portion of this sprint. We originally had the DB logic written into the handler functions for each of our API routes, but couldn’t figure out a way to test those functions as they were written so we opted to refactor our code a bit. We’re now in a spot where we can begin to build out a more comprehensive suite of tests, so overall it was a good learning experience and it should allow us to be a bit more productive moving forward.

Finally, we added various docstrings throughout our code base using both basic comments and the OpenAPI specification to document the backend for this project. Ultimately, we want to have all of our functions documented using the OpenAPI docstrings standard, but for now we opted for expediency and kept some simple comment documentation scattered throughout. We used `godoc` to automatically generate HTML documentation based on our docstrings and also created a “_documentation” folder to store various documentation for the backend, including the database schema we will be using.

As for lessons learned, we realized that we spent too much time thinking about and working on the authentication/authorization for our application this early in the process. It was a larger undertaking than either of us realized and only served to add unnecessary complexity to the project and our testing efforts at this stage. We realized that we should wait to address those requirements until we have all of the basic functionality of our application fleshed out first. We also learned about the importance of branching this sprint since there were a couple of times where one of us pushed breaking changes that should have been addressed in a separate branch until they were ready to merge.



__ADDITIONAL NOTE REGARDING BACKEND TEAM (as of 3/1):__


My (Ryan Brooks) 2nd child was born on 2/9 (a full month earlier than expected), so my plans for this sprint kind of went out the window. Her and my wife are both healthy now, but we had a few challenging circumstances and changes to deal with on short notice. I was still able to participate and contribute for this sprint, but it definitely had a significant impact on my availability and productivity.

Additionally, our team received a message from Angelina Uriarte-Wilson at 5PM on Wednesday (3/1) letting us know that she is dropping the class. None of us were aware that she was considering dropping the course prior to today, but that may impact our approach to and plans for this project moving forward. We had a few other candidate ideas for apps and, depending on our discussion in the next few days, we may decide to pivot a bit since Angelina was the main advocate for the business app to begin with.


# Backend documentation

The documentation we have created can also be found in the `src/server/_documentation` directory of this repo.

### __Initial Setup Instructions__
1) Install the following software on your local machine:
    * [Git](https://git-scm.com/downloads)
    * [Postgres](https://www.postgresql.org/download/)
        * Confirm Postgres is running on your machine after installing and create 2 new databases in your Postgres instance. `bizzen` for the production database name and `bizzen_test` for the test database name are recommended, but any names will work as long as you assign them to the appropriate keys in the `src/server/config/config.json` file.
    * [Visual Studio Code](https://code.visualstudio.com/)
2) Clone this repo to your local machine
    * `git clone --branch <branch-name-here> https://github.com/SwampSyndicate/BizZen`
3) In the repo directory, navigate to the `src/server` directory on your chosen CLI
4) Install all of the package dependencies for this project
    *  `go mod tidy`
5) Create a copy of the `config-template.json` file name `config.json` in the `src/server/config` directory and fill out the new `config.json` with appropriate values. Message one of the managing members for this repo if you'd like them to provide you with an example copy of that file. Following best security practices, the `config.json` file is listed in the `.gitignore` file for this project.
6) You can run the `server.exe` in the `src/server` directory on your local machine; however the most recent build may not always be pushed to the repo.
    * To build a new executable with the repo you have on your local machine, make sure you are still in the `src/server` directory on your CLI of choice and run `go build`


### __Running automated unit tests__
All automated unit tests are saved in the `tests` package/directory of this repo. To run these tests, run the following command from the `src/server` directory:

    `go test .\tests\ -v`

Each individual test will complete with either a 'PASS' or 'FAIL' status and the package will show with a status of 'ok' if all tests pass or 'fail' if any tests fail.

###  __Generating HTML documentation site from docstrings__
The docstrings for this project mostly will follow the OpenAPI specification in order to be able to generate the most current and accurate documentation from our source code.

We use and recommend using the `godoc` package to generate this documentation.

First make sure `godoc` is install on your machine by running the following command:

    `go install golang.org/x/tools/cmd/godoc@latest`

You can then generate/host the documentation site for this repo by running:

    `godoc -http=localhost:6060`