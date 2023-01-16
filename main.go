package main

import (
	"fmt"
	"github.com/arsham/rainbow/rainbow"
	"github.com/hashicorp/go-getter"
	"os"
)

func main() {
	fmt.Println("Hello va scans")

	// ...
	l := rainbow.Light{
		Reader: os.Stdin,  // to read from
		Writer: os.Stdout, // to write to
	}
	getter.Get("", "")
	l.Paint() // will rainbow everything it reads from reader to writer.
}
