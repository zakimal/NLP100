package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	articles := make([]*article, 0)
	src, err := os.Open("/Users/ozaki/go/src/NLP100/data/jawiki-country.json")
	defer src.Close()
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(src)
	for {
		// ReadBytes reads until the first occurrence of delim in the input,
		// returning a slice containing the data up to and including the delimiter.
		reading, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		a := new(article)
		// Unmarshal parses the JSON-encoded data and stores the result
		// in the value pointed to by v. If v is nil or not a pointer,
		// Unmarshal returns an InvalidUnmarshalError.
		err = json.Unmarshal(reading, a)
		if err != nil {
			log.Fatal(err)
		}
		articles = append(articles, a)
	}
	for i := range articles {
		if strings.Contains(articles[i].title, "イギリス") {
			fmt.Println(articles[i].text)
		}
	}
}

type article struct {
	title string `json:"title"`
	text string `json:"text"`
}
