// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io"
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
		// alt position for defer resp.Body.Close
		// println("resp: ", resp, "err: ", err) // DP debug
		// orig line b, err := ioutil.ReadAll(resp.Body)   // resp.Body contains server reply
		// b contains server response now, next must close resp function

		b, err := io.Copy(os.Stdout, resp.Body)
		defer resp.Body.Close() // from Daniel - defers error handling to end of loop
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b) // s = uninterpreted string
	}
}
