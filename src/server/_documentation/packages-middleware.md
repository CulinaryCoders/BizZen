<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# '*middleware*' Package

```go
import "server/middleware"
```

## Index
----

- [func RequestLoggingMiddleware(next http.Handler) http.Handler](<#func-requestloggingmiddleware>)


## func [RequestLoggingMiddleware](<https://github.com/SwampSyndicate/Bizzen/blob/main/src/server/middleware/logging.go#L13>)

```go
func RequestLoggingMiddleware(next http.Handler) http.Handler
```

RequestLoggingMiddleware logs the request type and URL of the inbound request and, if the request type is POST or PUT, also logs the JSON request type for debugging purposes



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
