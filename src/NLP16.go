package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
	lnum := lineCounter(os.Args[1])
	lnum = lnum/n+1
	if lnum%n == 0 {
		lnum--
	}
	dir, file := filepath.Split(os.Args[1])
	prefix := strings.Split(file, ".")[0]
	splited := make([]*os.File, n)
	for i := range splited {
		fname := prefix + "_div_" + strconv.Itoa(i) + ".txt"
		splited[i], err = os.Create(dir + fname)
		if err != nil {
			log.Fatal(err)
		}
		defer splited[i].Close()
	}
	index, lnow := 0, 0
	for {
		l, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		lnow++
		_, err = splited[index].Write(l)
		if err != nil {
			log.Fatal(err)
		}
		_, err = splited[index].Write([]byte(string('\n')))
		if err != nil {
			log.Fatal(err)
		}
		if lnow == lnum {
			index++
			lnow = 0
		}
	}
}

func lineCounter(name string) int {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	ret := 0
	for {
		_, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		ret++
	}
	return ret
}
