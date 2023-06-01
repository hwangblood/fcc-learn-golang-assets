# RSS Aggregator

Build a blog aggregator service in Go.

In this project you'll practice building a RESTful API in Go, and you'll use production-ready database tools like PostgreSQL, SQLc, Goose, and pgAdmin. This won't just be another CRUD app, but a service that has a long-running service worker that reaches out over the internet to fetch data from remote locations.

## Getting Started

### Setup Database

```shell
docker volume create rssagg_data

docker run -d \
        --name rssagg_data \
        -e POSTGRES_PASSWORD=foobarbaz \
        -v rssagg-db:/var/lib/postgresql/data \
        -p 5432:5432 \
        postgres:15.1-alpine

cd sql/schema
goose postgres postgres://postgres:foobarbaz@localhost:5432/ up

# run at project root folder
sqlc generate
```

### Run Server

- Copy `.env` from `.env.example`, then install and run it

```shell
$ go install github.com/hwangblood/fcc-learn-golang-assets/rssagg
$ export PORT=8000
$ rssagg
```

After running the `rssagg` command, you'll see the outpu like this:

```shell
Hello, Welcome to RSS Aggregator!
Server starting at port: 8000
```

## Tools

- [sqlc.dev | Compile SQL to type-safe Go](https://sqlc.dev/)

- [pressly/goose: A database migration tool. Supports SQL migrations and Go functions.](https://github.com/pressly/goose)
