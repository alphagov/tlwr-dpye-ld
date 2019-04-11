// Echo1 prints it's command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Println("Index:", i, "Value:", s)
	}
	fmt.Println(s)
}
