// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	prefix = "http://"
)

func main() {

	// single cmdline arg please
	for _, url := range os.Args[1:] {

		var fullurl string

		// Is there a prefix? If not add our prefix
		if !strings.HasPrefix(url, prefix) {
			fullurl = prefix + url
		} else {
			fullurl = url
		}

		// Debug
		// fmt.Printf("URL", url)
		// fmt.Printf("FULLURL", fullurl)

		resp, err := http.Get(fullurl)

		if err != nil {
			log.Fatal(err)
			// func {
			//	Error(w ResponseWriter, error err, code int)
			// }
			// fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
			//	os.Exit(1)
		}

		fmt.Println("HTTP Status Code:", resp.StatusCode)

		b, err := io.Copy(os.Stdout, resp.Body)
		defer resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
