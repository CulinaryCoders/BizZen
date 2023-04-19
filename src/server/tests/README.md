All unit tests for this application are stored in the 'tests' package to avoid issues with package coupling, since certain interdependencies are required for testing various packages and methods.

To run all unit tests, execute the following command from the `/src/server` directory:

    `go test .\tests\ -v`