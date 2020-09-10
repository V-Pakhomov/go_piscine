package main

import (
	"fmt"
	// "os"
	// "encoding/xml"
	// "encoding/json"
	"flag"
)

func main() {
	forceFlag := flag.String("f", "", "Flag to read from file")
	flag.Parse()
	fmt.Println(*forceFlag)
}