# Preparation:
- install Docker
- install PostgreSQL

## install PostgreSQL
- ```docker pull postgres:16-alpine```
- ```docker run --name postgres16 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:16-alpine```

## Go Migrate CLI
- source: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
- ```curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-386.tar.gz | tar xvz```
- ```mv migrate $GOPATH/bin/migrate```

## Access DB docker
- ```docker exec -it postgres16 /bin/sh```
  - ```createdb --username=root --owner=root simple_bank```
  - ```psql simple_bank```
- OR directly
- ```docker exec -it postgres16 createdb --username=root --owner=root simple_bank```

## Migrate first migration
- ```migrate --path db/migrations --database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" --verbose up```