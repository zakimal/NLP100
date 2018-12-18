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
		log.Fatalf("%s: Usage: %s <file> <n>\n", os.Args[0], os.Args[0])
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
	output := make([][]byte, n)
	for {
		l, _, err := reader.ReadLine()
		counter++
		if err == io.EOF {
			break
		}
		output[(counter-1)%n] = l
	}
	if counter < n {
		output = output[:counter - 1]
	}
	start := (counter - 1)%n
	output = append(output[start:], output[:start]...)
	for _, l := range output {
		fmt.Println(string(l))
	}
}
