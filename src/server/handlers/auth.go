package handlers

import (
	"encoding/json"
	"net/http"
	"server/models"
	"server/utils"
)

/*
*Description*

type Credentials

Defines the format for user login credentials
*/
type Credentials struct {
	Email    string `json:"email"`    // User's email
	Password string `json:"password"` // User's password
}

/*
*Description*

func Authenticate

Authenticates that the provided user account exists in the database and that the provided password is correct for that account.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:   POST

	Route:  /authenticate

	Body:
		Format: JSON

		Required fields:

			email  <string>

				Email address entered into login form

			password  <string>

				Plain text (unhashed) password entered into login form

		Optional fields:

			N/A

*Example request(s)*

	POST /login
	{
		"email":"johndoe@example.com",
		"password":"ILoveCats!"
	}

*Response format*

	Success:

		HTTP/1.1 200 OK
		Content-Type: application/json

		{
		"ID": "123456",
		"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
		"UpdatedAt": "2020-01-01T01:23:45.6789012-05:00",
		"DeletedAt": null,
		"email": "johndoe@example.com",
		"password": "$2a$14$ITcK9ZosVTZpx3OeJT8qu.I1Qfy31MinvsYvPbOCeIXj2fSxMCh8O",
		"account_type": "Individual",
		"first_name": "John",
		"last_name": "Doe",
		"contact_info_id": 45,
		"business_id": null,
		"permissions_id": 123,
		"user_pref_id": 42,
		"profile_pic_id": 79
		}

	Failure:

		-- Case = User account does not exist in the database
		HTTP/1.1 401 Unauthorized
		Content-Type: application/json

		{
			"error":"Specified user account does not exist in the database"
		}

		-- Case = Bad password
		HTTP/1.1 401 Unauthorized
		Content-Type: application/json

		{
			"error":"Invalid credentials"
		}
*/
func (app *Application) Authenticate(writer http.ResponseWriter, request *http.Request) {
	var credentials Credentials
	user := models.User{}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&credentials); err != nil {
		utils.RespondWithError(writer, http.StatusBadRequest, err.Error())
		return
	}

	defer request.Body.Close()

	returnedUser, err := user.GetUserByEmail(app.AppDB, credentials.Email)
	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusUnauthorized,
			"Specified user account does not exist in the database")

		return
	}

	if err := returnedUser.CheckPassword(credentials.Password); err != nil {
		utils.RespondWithError(
			writer,
			http.StatusUnauthorized,
			"Invalid credentials")

		return
	}

	/*
		PLEASE DO NOT REMOVE
		TODO: Implement authentication logic (func Authenticate)
		session, _ := env.Store.Get(request, "sessionID")
		session.Values["authenticated"] = true
		session.Save(request, writer)
		//validToken, err := utils.GenerateToken(user.Email, user.AccountType, config.AppConfig.GetSigningKey())
		if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusInternalServerError,
			err.Error())
		}
	*/

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		returnedUser)
}

/*
*Description*

func Authorize

A middleware function that uses Gorilla sessions to verify that the user is authenticated before calling the next HTTP handler function in the chain.

First, the session cookie is retrieved from the HTTP request using the Gorilla sessions package. If the session cookie is not found, or if the session does not contain a "sessionID" value, the function responds with a 401 Unauthorized error and stops the chain of HTTP handlers.

If the user is authenticated, the function calls the next HTTP handler function in the chain with the same http.ResponseWriter and *http.Request objects passed to it.

*Parameters*

	next  <http.HandlerFunc>

		The next HTTP handler function in the chain to call if authentication is successful.

*Returns*

	next  <http.HandlerFunc>

		http.HandlerFunc: a new HTTP handler function that performs authentication and calls the next handler function if authentication is successful.

*Response format*

	   Failure:

		   HTTP/1.1 401 Unauthorized
		   Content-Type: application/json

		   {
			   "error":"Invalid or expired session"
		   }
*/
func (app *Application) Authorize(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		session, _ := app.CookieStore.Get(request, "session")
		_, ok := session.Values["sessionID"]
		if !ok {
			utils.RespondWithError(
				writer,
				http.StatusUnauthorized,
				"Invalid or expired session")

			return
		}

		next.ServeHTTP(writer, request)
	}
}
