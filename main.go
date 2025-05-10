package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("main.go")

	fmt.Printf("Env var test: %s", os.Getenv("test"))
}
