package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strings"
)

const WANT = 1

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("%s: Usage: %s <file>\n", os.Args[0], os.Args[0])
	}
	src, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	reader := bufio.NewReader(src)
	word := make([]rune, 0)
	words := make(runes, 0)
	ncol := 0
	for {
		r, _, _ := reader.ReadRune()
		if r == 0 {
			break
		}
		if r == rune('\t') {
			ncol++
		} else if ncol == WANT {
				word = append(word, r)
		}
		if r == rune('\n') {
			words = append(words, word)
			ncol = 1
			word = nil
		}
	}
	sort.Sort(words)
	uwords := getUniq(words)
	for _, w := range uwords {
		fmt.Println(toString(w))
	}
}

type runes [][]rune
func (rs runes) Len() int { return len(rs) }
func (rs runes) Less(i, j int) bool {
	for k := 0; k < len(rs[i]) && k < len(rs[j]); k++ {
		if rs[i][k] < rs[j][k] {
			return true
		}
		if rs[i][k] > rs[j][k] {
			return false
		}
	}
	return false
}
func (rs runes) Swap(i, j int) { rs[i], rs[j] = rs[j], rs[i] }

func toString(r []rune) string {
	l := len(r)
	ret := make([]string, l)
	for i :=range ret {
		ret = append(ret, string(r[i]))
	}
	return strings.Join(ret, "")
}

func getUniq(rs runes) runes {
	i := 0
	for {
		if  reflect.DeepEqual(rs[i],rs[i+1]) {
			copy(rs[i:], rs[i+1:])
			rs[len(rs)-1] = nil
			rs = rs[:len(rs) - 1]
		} else {
			i++
		}
		if i == len(rs)-1 {
			break
		}
	}
	return rs
}