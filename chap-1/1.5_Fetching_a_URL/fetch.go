// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] { // single cmdline arg please
		resp, err := http.Get(url) // response or error returned
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err) // v=default fmt
			os.Exit(1)
		}
		println("resp: ", resp, "err: ", err) // DP debug
		b, err := ioutil.ReadAll(resp.Body)   // resp.Body contains server reply
		// b contains server response now, next must close resp function

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b) // s = uninterpreted string
	}
}
