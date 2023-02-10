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

Integration and unit tests run on the test container.

## Using the client

Provide an http.client and url to the constructor:

```
import (
	"integrations/pkg/Form3"
	"net/http"
)

HttpClient := http.Client{}
url = <SERVER_URL>

client, _ := Form3.New(url, HttpClient)
```

Now that you have an instance of the client, you can use the create, fetch and delete fucntions.

For example, to fetch an account:
```
fetched_account, _ := client.FetchAccount(<ACCOUNT_ID>)
```

More examples available in Client_test.go
