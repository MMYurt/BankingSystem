# Banking System API

This project aims to create a banking services API for:
- Create and manage bank accounts for individuals that contain account owner's name, balance, and currency.
- Record all balance changes of each account. An account entry record will be created.
- Make a transfer between two accounts. It is performed by transactions.

## Installation

- Golang
- Docker Desktop
- TablePlus
- Golang Migrate
- DB Docs
- DBML CLI
- Sqlc
- Gomock
    
## Run Locally

Clone the project

```bash
  git clone https://link-to-project
```

Create postgres container

```bash
  make postgres
```

Create the database

```bash
  make createdb
```

Run db migration up

```bash
  make migrateup
```

Run db migration down

```bash
  make migratedown
```

## Documentation

- Generate DB Documentation
```bash
make db_docs
```

## How to generate code
- Generate schema SQL file with DBML:
```bash
make db_schema
```

- Generate SQL CRUD with sqlc:
```bash
make sqlc
```

## How to run
- Run server
```bash
make server
```
- Run test
```bash
make test
```
