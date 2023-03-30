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

