package main

import (
	"bufio"
	"flag"
	"fmt"
	"gocksum"
	"os"
)

func main() {
	flag.Parse()

	// If reading from stdin
	if flag.NArg() == 0 {
		s, sz, _ := gocksum.Cksum(bufio.NewReader(os.Stdin))
		fmt.Printf("%d %d\n", s, sz)
		return
	}

	// If given a list of 1 or more files on the command line
	for _, filename := range flag.Args() {
		f, err := os.Open(filename)
		defer f.Close()
		reader := bufio.NewReader(f)

		if f == nil {
			fmt.Println("cksum: ", err)
			continue
		}
		s, sz, err := gocksum.Cksum(reader)
		if err != nil {
			fmt.Printf("%s\n", err)
			continue
		}
		fmt.Printf("%d %d %s\n", s, sz, filename)
	}
}
