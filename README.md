browser
=======

A helper library for Go, which parses HTML and CSS documents for required links.

Example:
```go
package main

import (
	"github.com/SlyMarbo/browser"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error.
	}
	
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error.
	}
	
	links, err := browser.Links(string(body))
	if err != nil {
		// handle error.
	}
	
	// links is a []string of URIs.
}
```
