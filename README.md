# Project Repo Title

A little introduction here...

<hr />

## **Requirements**
- Golang 1.20 or higher
- Openssl3

<hr />

## **Dependencies**
- [PymigrateDB](https://pypi.org/project/pymigratedb/)
- [Gcloud CLI](https://cloud.google.com/sdk/docs/install)

<hr />

## Setup
- To install project run ```make```.
#- It will create two new files in the project root, called: ".env" and "config.yaml"
#- Fill it with correct values to procceed with development

<hr />

## Run
- To run project execute ```make run``` into the terminal. It will start the API and serve the requests connected to the resources filled in the config files.

<hr />

## Tests
- To run and perform test cases, run the following command: ```make test```. It will begin the tests execution.

<hr />

## Accepted Methods and Content-Types

| Method | Content-Type |
|:------:|:------------:|
|POST    |application/json|
|GET     |
|OPTIONS |

<hr />


## API Structure

```bash
.
├── adapters
│   ├── database
│   │   └── database.go
│   ├── logging
│   │   └── logging.go
│   ├── storage
│   │   └── storage.go
│   └── totp
│       └── totp.go
├── app
│   ├── errors
│   │   └── errors.go
│   ├── repository
│   │   └── .gitkeep
│   └── usecases
│       └── .gitkeep
├── clients
│   ├── google
│   │   ├── logging
│   │   │   └── logging.go
│   │   └── storage
│   │       └── storage.go
│   ├── iam
│   │   └── client.go
│   └── postgres
│       └── postgres.go
├── config
│   ├── constants
│   │   └── constants.go
│   └── config.go
├── entity
│   ├── http_response.go
│   └── log.go
├── handler
│   └── health
│       └── health.go
├── middleware
│   ├── auth.go
│   ├── content.go
│   └── security.go
├── pkg
│   ├── app
│   │   └── app.go
│   ├── crypt
│   │   └── crypt.go
│   ├── helpers
│   │   ├── http.go
│   │   ├── json.go
│   │   └── utils.go
│   └── jwt
│       └── jwt.go
├── .dockerignore
├── .env.example
├── .gitignore
├── .marincor_env
├── config.example.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── README.md
└── route.go
```