# Go Backend

Intended to be a starting point for building backend servers with Golang

Uses build tags to compile with different database drivers

## Setup

In your editor, make sure to add build tags to the LSP.
For me in VSCode it looks like:

```json
  "gopls": {
    "build.buildFlags": ["-tags=postgres,mysql,sqlite"]
  }
```

## TODO:

-   [x] Makefile
-   [ ] Setup from config.json
-   [ ] Database connection
-   [ ] Support different databases
    -   [ ] Postgres
    -   [ ] Sqlite
    -   [ ] Mysql
-   [ ] Admin dashboard UI
-   [ ] REST API
-   [ ] Authentication
-   [ ] Realtime subscriptions
-   [ ] File storage
