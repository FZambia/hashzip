package main

import (
	"fmt"
	"os"

	"golang.org/x/mod/sumdb/dirhash"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: hashzip ZIP_FILE")
		os.Exit(1)
	}
	sum, err := dirhash.HashZip(os.Args[1], dirhash.DefaultHash)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	fmt.Println(sum)
}
