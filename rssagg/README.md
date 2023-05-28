# RSS Aggregator

Build a blog aggregator service in Go.

In this project you'll practice building a RESTful API in Go, and you'll use production-ready database tools like PostgreSQL, SQLc, Goose, and pgAdmin. This won't just be another CRUD app, but a service that has a long-running service worker that reaches out over the internet to fetch data from remote locations.

## Getting Started

- Copy `.env` from `.env.example`, then install and run it

```shell
$ go install github.com/hwangblood/fcc-learn-golang-assets/rssagg
$ export PORT=8000
$ rssagg
```

After running the `rssagg` command, you'll see the outpu like this:

```shell
Hello, Welcome to RSS Aggregator!
Port:  8000
```
