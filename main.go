// main.go - Tiny CLI sum tool
package main

import (
	"fmt"
	"os"
	"strconv"
)

// Tool metadata.
const (
	toolName        = "sum-cli"
	toolDescription = "Adds numbers passed as command-line arguments."
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("%s\n", toolName)
		fmt.Println(toolDescription)
		fmt.Println("Usage: sum-cli <n1> <n2> ... <nk>")
		return
	}

	var sum float64
	for _, arg := range os.Args[1:] {
		n, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Printf("error: '%s' is not a valid number\n", arg)
			return
		}
		sum += n
	}

	fmt.Printf("Sum of %d numbers: %g\n", len(os.Args)-1, sum)
}

