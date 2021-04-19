package main

import (
	"fmt"
	"os"

	"golang.org/x/mod/sumdb/dirhash"
)

func main() {
	fmt.Println(dirhash.HashZip(os.Args[0], dirhash.DefaultHash))
}
