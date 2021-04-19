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
	fmt.Println(dirhash.HashZip(os.Args[1], dirhash.DefaultHash))
}
