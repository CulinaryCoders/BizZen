# 'sample_data' Package

</br>

## Overview

The 'sample_data' package/directory is intended to be used as a repository for fake sample records that will be loaded into the application for testing and demonstration purposes.

Currently, all sample records have been created and extracted from Mockaroo using custom schemas that match the project's currrent GORM config/definitions. Each DB object's set of sample records is saved in the corresponding 'sample-\*.json' files, with the '\*' corresponding to the DB object's name/alias. The 'sample-data.json' file is a compilation of all the other 'sample-\*.json' files with a plural key for each object type is associated with that object's list of sample JSON records.

During the development of this application's MVP, all database tables are dropped and recreated each time the server is initialized. Then the 'LoadJSONSampleData' function in this package is called to create all of the sample records present in the 'sample-data.json' file within this directory. This is intended to make the initial testing and development efforts for this application more straightforward and easier to troubleshoot across each dev's local environment.

</br>

## Manually triggering sample loads

In the future, if/when the database tables are no longer being dropped/refreshed each time the backend server is initialized, data loads can also be manually performed using either the 'manual-data-load.ps1' (for Windows environments) or the 'manual-data-load.sh' (for UNIX environments) scripts located in this directory.

</br>

* Powershell script example commands (manual-data-load.ps1)

    </br>

    * Load all JSON files defined in script (execute from an open Powershell terminal/instance)

    </br>

    ```
    manual-data-load.ps1
    ```

    </br>

    * Limit load to specific object types (execute from an open Powershell terminal/instance)

    </br>

    ```
    manual-data-load.ps1 -objects Address,Business
    ```

    </br>

* Shell script example commands (manual-data-load.sh)

    </br>

    * N/A  (TBD  --  Still in development)

    </br>

**NOTE**:

As of this time, the 'manual-data-load.sh' shell script is still in development and is not working as intended. Therefore, manual loads are only supported on Windows machines using the 'manual-data-load.ps1' Powershell script.