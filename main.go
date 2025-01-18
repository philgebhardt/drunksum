package main

import (
	"bufio"
	"drunksum/checksum"
	"drunksum/drunk"
	"flag"
	"fmt"
	"os"
)

func main() {
	var algorithm string
	flag.StringVar(&algorithm, "a", "sha256", "the hash algoritm (default: sha256")
	flag.Parse()

	data, err := checksum.NewChecksum(algorithm, bufio.NewReader(os.Stdin))
	if err != nil {
		fmt.Fprintf(os.Stderr, "checksum: %v\n", err)
		os.Exit(1)
	}

	b, err := drunk.NewBishop(17, 9, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "print: %v", err)
		os.Exit(1)

	}
	b.Print()
	fmt.Println()
	fmt.Printf("%s: %x\n", algorithm, data)
}
