package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("%s: Usage: %s <file> <n>\n", os.Args[0])
	}
	src, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(src)
	counter := 0
	for counter < n {
		l, _, err := reader.ReadLine()
		counter++
		if err == io.EOF {
			break
		}
		fmt.Printf("%d\t| %s\n", counter, string(l))
	}
}
