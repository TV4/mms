# titleservice

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/TV4/mms/titleservice)

## Installation

    go get -u github.com/TV4/mms/titleservice

## Usage example

```go
package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"

	titleservice "github.com/TV4/mms/titleservice"
)

func main() {
	var username, password string

	flag.StringVar(&username, "user", "", "username")
	flag.StringVar(&password, "pass", "", "password")

	flag.Parse()

	c := titleservice.NewClient(
		username, password,
		titleservice.Simulate,
	)

	resp, err := c.RegisterClip(context.Background(),
		titleservice.MakeClip(
			"123", "Test-title", 456,
			titleservice.Date(2017, 3, 24),
		),
	)
	if err != nil {
		logger := log.New(os.Stderr, "", 0)
		logger.Fatal(resp, "\n", err)
	}

	enc := json.NewEncoder(os.Stdout)

	enc.SetIndent("", "  ")
	enc.Encode(resp)
}
```

### Note
> You need to remove the option `titleservice.Simulate` in order to make requests that are persisted to the MMS database
