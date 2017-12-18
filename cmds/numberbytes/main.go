package main

import (
	"fmt"
	"os"
)

func main() {
	if numberBytes, err := fmt.Printf("Hello world!\n"); err != nil {
		os.Exit(1)
	} else {
		fmt.Printf("Printed %d bytes\n", numberBytes)
	}
}
