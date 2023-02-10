#!/bin/sh
set -e

# Use `main.go` to house the entrypoint check script
go build main.go

# Wait for the API to be up
until ./main $API_PROTOCOL://$API_HOST:$API_PORT/; do
  >&2 echo "API at $API_PROTOCOL://$API_HOST:$API_PORT/ is unavailable; sleeping for now."
  sleep 5
done

>&2 echo "API is up; running integration tests."

exec "$@"
