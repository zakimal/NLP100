package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("%s: Usage: %s <file>\n", os.Args[0])
	}
	fp, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	for {
		r, _, _ := reader.ReadRune()
		if r == '\t' {
			r = ' '
		}
		if r == 0 {
			break
		}
		fmt.Printf(string(r))
	}
	fmt.Println()
}
