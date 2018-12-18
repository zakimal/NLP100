package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("%s: Usage: %s <file>\n", os.Args[0])
	}
	src, err := os.Open("/Users/ozaki/go/src/NLP100/date/hightemp.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	reader := bufio.NewReader(src)
	col, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	dn := "/Users/ozaki/go/src/NLP100/date/col" + os.Args[1] + ".txt"
	dst, err := os.Create(dn)
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()
	now := 1
	buf := make([]rune, 0)
	for {
		r, _, _ := reader.ReadRune()
		if r == 0 {
			break
		}
		if r == '\t' {
			now++
		} else if now == col {
			buf = append(buf, r)
		}
		if r == '\n' {
			_, err := dst.Write([]byte(string(buf)))
			if err != nil {
				log.Fatal(err)
			}
			_, err = dst.Write([]byte(string('\n')))
			if err != nil {
				log.Fatal(err)
			}
			now = 1
			buf = nil
		}
	}
}
