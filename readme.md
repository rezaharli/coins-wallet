# GO-KIT DDD WALLET API

Coins.ph Backend Technical Assessment, API with domain driven approach (DDD) using Golang, Go-kit and PostgreSQL

"coins-wallet" is an application that made us possible to transfer several amount of money from one account to another.

## Installation & Run

Download this project

```bash
https://github.com/rezaharli/coins-wallet
```

Start the database service

```bash
#make sure you have docker installed
make docker.start.components
```

Set project environment and run

```bash
# copy and rename .env.example
cp .env.example .env

# set .env to be similiar to your database configuration
# the default
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=coins
DB_NAME=coins
DB_PASSWORD=mysecretpassword

# run golang project
go run main.go

# API Endpoint : http://localhost:8000/v1/
```

#### API Endpoint list
- refer to https://github.com/rezaharli/coins-wallet/blob/main/docs/api.md

## Run Tests

This command run the unit tests

```bash
make run.test
```

## Contributing

Please refer to the project's style for submitting patches and additions. In general, we follow the "fork-and-pull" Git workflow.

1.  **Fork** the repo on GitHub
2.  **Clone** the project to your own machine
3.  **Commit** changes to your own branch
4.  **Push** your work back up to your fork
5.  Submit a **Pull request** so that we can review your changes
