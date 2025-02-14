<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# '*utils*' Package

```go
import "server/utils"
```

## Index
----

- [func DecodeJSON(request *http.Request, modelObj interface{}) (interface{}, error)](<#func-decodejson>)
- [func ParseRequestID(request *http.Request) (uint, error)](<#func-parserequestid>)
- [func ParseRequestIDField(request *http.Request, idFieldKey string) (uint, error)](<#func-parserequestidfield>)
- [func RespondWithError(writer http.ResponseWriter, code int, message string)](<#func-respondwitherror>)
- [func RespondWithJSON(writer http.ResponseWriter, code int, payload interface{})](<#func-respondwithjson>)


## func [DecodeJSON](<https://github.com/SwampSyndicate/Bizzen/blob/main/src/server/utils/responses.go#L177>)

```go
func DecodeJSON(request *http.Request, modelObj interface{}) (interface{}, error)
```

<br />

***Description***

func DecodeJSON

Helper method to unmarshal the request body into the provided gorm.Model object.

<br />

***Parameters***

```
request  <*http.Request>

	The HTTP request.

modelObj  <interface{}>

	The gorm.Model object that the request body JSON will be unmarshalled into.
```

<br />

***Returns***

```
_  <uint>

	The updated gorm.Model object that the request's JSON body was unmarshalled into.

_  <error>

	Encountered error (nil if no errors are encountered).
```

## func [ParseRequestID](<https://github.com/SwampSyndicate/Bizzen/blob/main/src/server/utils/responses.go#L108>)

```go
func ParseRequestID(request *http.Request) (uint, error)
```

<br />

***Description***

func ParseRequestID

Helper method to parse the "id" variable present in the request and convert it to an unsigned integer.

<br />

***Parameters***

```
request  <*http.Request>

	The HTTP request.
```

<br />

***Returns***

```
_  <uint>

	The "id" variable value that has been converted to uint format.

_  <error>

	Encountered error (nil if no errors are encountered).
```

## func [ParseRequestIDField](<https://github.com/SwampSyndicate/Bizzen/blob/main/src/server/utils/responses.go#L144>)

```go
func ParseRequestIDField(request *http.Request, idFieldKey string) (uint, error)
```

<br />

***Description***

func ParseRequestIDField

Helper method to parse the specified ID variable present in the request and convert it to an unsigned integer.

Offers a more generic alternative to the ParseRequestID method, in instances where there are multiple ID variables in the route or when a route has an ID variable with a key other than "id".

<br />

***Parameters***

```
request  <*http.Request>

	The HTTP request.

idFieldKey  <string>

	The ID variable key used for the request's route.
```

<br />

***Returns***

```
_  <uint>

	The specified ID variable's value that has been converted to uint format.

_  <error>

	Encountered error (nil if no errors are encountered).
```

## func [RespondWithError](<https://github.com/SwampSyndicate/Bizzen/blob/main/src/server/utils/responses.go#L77>)

```go
func RespondWithError(writer http.ResponseWriter, code int, message string)
```

<br />

***Description***

func RespondWithJSON

Takes a http.ResponseWriter, a HTTP status code and a message as input parameters. It formats the error message as a JSON object with a "error" field containing the message, and writes it to the ResponseWriter with the given status code.

<br />

***Parameters***

```
writer  <http.ResponseWriter>

	The HTTP response writer

code  <int>

	The HTTP status code to be used in the error response.

message  <string>

	Error message to be sent in the response.
```

<br />

***Returns***

```
None
```

## func [RespondWithJSON](<https://github.com/SwampSyndicate/Bizzen/blob/main/src/server/utils/responses.go#L40>)

```go
func RespondWithJSON(writer http.ResponseWriter, code int, payload interface{})
```

<br />

***Description***

func RespondWithJSON

Marshals the payload into JSON format and returns the HTTP response.

The method uses the json.Marshal method to encode the payload to a JSON byte slice. If there is an error during the encoding process, the method will return an HTTP 500 Internal Server Error response with an error message.

The method sets the response header's content type to application/json and writes the serialized JSON payload to the response writer.

<br />

***Parameters***

```
writer  <http.ResponseWriter>

	The HTTP response writer

request  <*http.Request>

	The HTTP request

payload <interface{}>

	Payload used to serialize to JSON format and send to the HTTP client as the response body.
```

<br />

***Returns***

```
None
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
