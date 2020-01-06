# Example api in go

* uses go-restful ( https://github.com/emicklei/go-restful )
* generates open-api spec file with the go-restful openapi extension ( github.com/emicklei/go-restful-openapi )
* Uses gorm as storage backend
* Includes a swagger-ui instance served on `/swaggerui`
* API is served under `/api/`
* API spec is served under `/api/apidocs.json`

## Running

To run:

    go run ./cmd/main/go

This sort-of expects the env var `GOFLAGS` set to `-mod=vendor`.

## Disclaimer

This is far from production-ready and is just a POC:

* There is no configuration:
  * SQLite is used as db with a local hardcoded file
  * Starts listening on `localhost:8080`
* no input validation
* no authentication/authorisation at the moment
