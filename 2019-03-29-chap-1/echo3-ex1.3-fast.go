// Echo3 prints it's command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], ""))
	fmt.Println("Time taken: ", time.Since(start).Seconds(), " seconds")
}
