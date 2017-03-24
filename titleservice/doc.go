/*

Package titleservice is a client for the MMS TitleService API

Documentation

https://godoc.org/github.com/TV4/mms/titleservice

Installation

Just go get the package:

    go get -u github.com/TV4/mms/titleservice

Usage

A small usage example:

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

        	var (
        		logger = log.New(os.Stderr, "", 0)
        		enc    = json.NewEncoder(os.Stdout)
        		ctx    = context.Background()
        		c      = titleservice.NewClient(username, password, titleservice.Simulate)

        		clip = titleservice.Clip{
        			TitleCode:   "123",
        			Title:       "Test-title",
        			Length:      1,
        			PublishedAt: "20170324",
        		}
        	)

        	resp, err := c.RegisterClip(ctx, clip)
        	if err != nil {
        		logger.Fatal(resp, "\n", err)
        	}

        	enc.SetIndent("", "  ")
        	enc.Encode(resp)
				}

Output:

        {
          "StatusCode": 200,
          "StatusDescription": "",
          "Errors": null
        }

*/
package titleservice
