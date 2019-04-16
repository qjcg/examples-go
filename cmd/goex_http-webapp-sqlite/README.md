# Build

`CGO_ENABLED=1 go build`

As the sqlite3 library uses CGO, your binary will be dynamically, NOT statically
linked.


# Load test

Use [hey](https://github.com/rakyll/hey), for example:

`hey -c 200 -n 1000 http://localhost:9999/`
