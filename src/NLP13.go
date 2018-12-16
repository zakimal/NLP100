package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	col1, err := os.Open("/Users/ozaki/go/src/NLP100/data/col1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer col1.Close()
	col2, err := os.Open("/Users/ozaki/go/src/NLP100/data/col2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer col2.Close()
	reader1 := bufio.NewReader(col1)
	reader2 := bufio.NewReader(col2)
	dstName := "/Users/ozaki/go/src/NLP100/data/merged.txt"
	dst, err := os.Create(dstName)
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()
	for {
		l1, _, err := reader1.ReadLine()
		if err == io.EOF {
			break
		}
		l2, _, err := reader2.ReadLine()
		if err == io.EOF {
			break
		}
		if _, err := dst.Write(l1); err != nil {
			log.Fatal(err)
		}
		if _, err := dst.Write([]byte("\t")); err != nil {
			log.Fatal(err)
		}
		if _, err := dst.Write(l2); err != nil {
			log.Fatal(err)
		}
		if _, err := dst.Write([]byte("\n")); err != nil {
			log.Fatal(err)
		}
	}
}
