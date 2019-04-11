// Echo1 prints it's command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	s, sep := "", ""
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Println("Index:", i, "Value:", s)
		fmt.Println("Time taken: ", time.Since(start).Seconds(), " seconds")
	}
	fmt.Println(s)
}
