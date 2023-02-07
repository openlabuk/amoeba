# Amoeba *OpenLab

## About

Amoeba is a extremely fast, lightweight caching service written in Go.  Amoeba can temporarily store frequently accessed data from databases, APIs, or other sources in memory, allowing quicker access to this data than re-fetching it from the original source. This helps to reduce the load on the origin servers and speeds up the performance of dynamic web applications, or other applications in which it is leveraged.

### API

Amoeba consists of 5 endpoints

* `/healthcheck` service healthcheck endpoint
* `/create` create a new record
* `/get` get a single record by key
* `/delete` delete a single record by key
* `/flush` flush all keys

Please see the OpenAPI 3.0 specification in the project `/api` folder for more details.

```tree
...
├── api/openapi.yaml
...
```

## Getting started

### build an executable

From the project root, run the following to build the executable

```
go get
go build
```
### run server

```
./amoeba
```
The service runs on `http://127.0.0.1:7029` by default, but the server and port can be provided as command-line arguments, eg

```
./amoeba -server=127.0.0.1 -port=7030
```

### run developmet code locally

From the project root, run the following to start the server in debug mode

```
go run .
```

## Tests

To run the test suite, execute the following from the project root
```
go test
```
