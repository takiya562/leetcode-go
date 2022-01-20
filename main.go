package main

import (
	"fmt"
	"path"
)

func main() {
	url := "/path/to/file"
	fmt.Println(path.Base(url))
}
