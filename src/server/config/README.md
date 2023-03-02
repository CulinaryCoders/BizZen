# Backend Environmental Variables - Definitions & Configurations

##  Overview of config files
The `config.json` file used to store this project's environmental variables and configuration data is listed in `.gitignore` for security reasons. **DO NOT UNDER ANY CIRCUMSTANCES HARD CODE OR UPLOAD ANY SECRETS, CONNECTION INFO, PASSWORDS, ETC. TO THIS GITHUB.**

We intend to follow the recommendations and conventions laid out in "[The Twelve-Factor App](https://12factor.net/)" to the best of our abilities, so refer to that documentation when questions about architecture and best practices arise. 

Whenever new environmental variables are defined, please update the `config-template.json` file with those entries to avoid uneccessary build errors/issues.

The `config-template.json` file contains all of the keys used to reference environmental variables throughout this codebase. Each key/entry in the `config-template.json` should be assigned a null value for the sake of consistency.

All `config.json` throughout this project utilize JSON formatting.


##  Golang specifics
We are utilizing the 'Viper' package (<github.com/spf13/viper>) to import and manage the environmental variables and config data used in our backend code.