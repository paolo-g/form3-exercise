# Form3 golang API client exercise

Stack: golang

A client library with integration testing using the provided Form3 stack.

API docs at https://www.api-docs.form3.tech/api/tutorials/getting-started/create-an-account

## Solution architecture

- Module is located in /integrations
- API clients located in the /pkg subdirectory

## Building docker locally for integration tests

```
docker-compose build
docker-compose up
```

### Unit tests

```
cd integrations
go test integrations/pkg/Form3
go test integrations/pkg/Form3/Organisation
```
