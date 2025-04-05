# An example app for traceme

See [traceme](https://github.com/wazzaps/traceme) for context.

This example app implements a counter. `GET /` returns the current count, and `POST /increment` increments it.

## Usage

- Assumptions:
  - `go` and `docker` are installed
  - The user can run `docker` commands without sudo
  - The `wazzaps/traceme` image exists locally (see main repo)

```bash
make start-webcounter

curl -X POST http://127.0.0.1:8080/increment
curl http://127.0.0.1:8080/

make stop-webcounter
```


