# Golang Github fetch user repositories

Implement an API server that lists all public GitHub repositories of the user specified.

## How to run this server

* Fork and clone this repo
* From root folder, run ```go run server.go```

### Run API examples

To get user's github repositories, make a request to ```http://localhost:5000/{userId}/repositories```

Where:
    
* Params: {userId} is Github username

Examples:

```batch
curl http://localhost:5000/digroad/repositories
```

response:

```json
[{
	"id": 369084304,
	"name": "backend-coding-test",
	"description": "",
	"stargazersCount": 0,
	"ownerId": 13029511,
	"ownerLogin": "digroad"
}, {
	"id": 354199537,
	"name": "go-scim-client",
	"description": "SCIM 2.0 API client auto-generated using RingCentral SCIM OpenAPI / Swagger 2.0 specification.",
	"stargazersCount": 1,
	"ownerId": 13029511,
	"ownerLogin": "digroad"
}, {
	"id": 261637519,
	"name": "graphql",
	"description": "Simple low-level GraphQL HTTP client for Go",
	"stargazersCount": 0,
	"ownerId": 13029511,
	"ownerLogin": "digroad"
}, {
	"id": 366303541,
	"name": "oapi-codegen",
	"description": "Generate Go client and server boilerplate from OpenAPI 3 specifications",
	"stargazersCount": 0,
	"ownerId": 13029511,
	"ownerLogin": "digroad"
}]
```
## Test

To run test, just run ```go test``` from root folder
