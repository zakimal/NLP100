package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	count := 0
	if len(os.Args) != 2{
		log.Fatalf("%s: Usage: %s <file>\n", os.Args[0])
	}
	fp, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	for {
		_, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}
		count++
	}
	fmt.Println(count)
}
