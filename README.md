# hashzip

Calculate and display Go module zip file hash

Example usage:

```
go get github.com/FZambia/hashzip
curl -O https://proxy.golang.org/github.com/uber/jaeger-client-go/@v/v2.26.0+incompatible.zip
hashzip v2.26.0+incompatible.zip

h1:h285ag9YqU5dfE+D2tc2mL93wjg1YLveCCDgm2y4Rsg=
```

In order to get a module checksum with specific version:

```
git clone https://github.com/jaegertracing/jaeger-client-go.git
cd jaeger-client-go
git checkout v2.26.0
rm -rf .git
```

Then:

```go
package main

import (
	"fmt"
	"os"

	"golang.org/x/mod/sumdb/dirhash"
)

func main() {
	sum, err := dirhash.HashDir("/path/to/jaeger-client-go", "github.com/uber/jaeger-client-go@v2.26.0+incompatible", dirhash.DefaultHash)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	fmt.Println(sum)
}
```
