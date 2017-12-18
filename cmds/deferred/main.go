package main

import "fmt"

func printer(msg string) error {
	defer fmt.Printf("\n")
	_, err := fmt.Printf(msg)
	return err
}

func main() {
	printer("vim-go")
}
