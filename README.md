# BizZen

## Project Name
BizZen


## Group Name

-  **Canvas**  --  Project Group 27
-  **GitHub**  --  The Swamp Syndicate


## Project Description

BizZen is a B2B software-as-a-service platform designed to help small and medium-sized businesses manage their appointments and schedule with ease. The platform offers a variety of tools and features that allow businesses to schedule appointments and manage their calendar. It allows customers to book appointments and manage for services offered by businesses on the platform.


## Project Members

**Frontend**
-  Lucinda Quintal
-  Alaine Spade

**Backend**
-  Ryan Brooks
-  <del>Angelina Uriarte-Wilson</del>

## Running the Application

### Backend
## Setup & Config
</br>

# Backend Environmental Variables - Definitions & Configurations

##  Overview of config files
The `config.json` file used to store this project's environmental variables and configuration data is listed in `.gitignore` for security reasons. **DO NOT UNDER ANY CIRCUMSTANCES HARD CODE OR UPLOAD ANY SECRETS, CONNECTION INFO, PASSWORDS, ETC. TO THIS GITHUB.**

We intend to follow the recommendations and conventions laid out in "[The Twelve-Factor App](https://12factor.net/)" to the best of our abilities, so refer to that documentation when questions about architecture and best practices arise. 

Whenever new environmental variables are defined, please update the `config-template.json` file with those entries to avoid uneccessary build errors/issues.

The `config-template.json` file contains all of the keys used to reference environmental variables throughout this codebase. Each key/entry in the `config-template.json` should be assigned a null value for the sake of consistency.

All `config.json` throughout this project utilize JSON formatting.


##  Golang specifics
We are utilizing the 'Viper' package (<github.com/spf13/viper>) to import and manage the environmental variables and config data used in our backend code.

</br>

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

## Unit tests 

</br>

All unit tests for this application are stored in the 'tests' package to avoid issues with package coupling, since certain interdependencies are required for testing various packages and methods.

To run all unit tests, execute the following command from the `/src/server` directory:

    `go test .\tests\ -v`

### Frontend
<i>Prerequisite: backend is running on localhost:8080 (with postgres)</i>

####Steps
1. Navigate to BizZen directory
2. Run `ng serve`
3. In your browser, navigate to <a href="http://localhost:4200/register">http://localhost:4200 </a>
4. The application should be running and functional from here


## Running Tests
### Frontend
#### Unit tests
1. Navigate to the BizZen directory
2. Run `ng test`
3. A browser window should appear with all unit tests


## Versions
### Frontend
* Angular: 15.1.2
* Angular CLI: 15.1.3
* Node: 16.13.0
* Package Manager: npm 8.1.0
* OS: darwin x64
* @ng-bootstrap/ng-bootstrap: version 14.0.0
* ng-bootstrap: version 1.6.3
