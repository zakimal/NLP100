package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	str := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."
	ret := ""
	for _, v := range strings.Split(str, " ") {
		if len(v) > 4 {
			rs := []rune(v)
			v = string(rs[0])
			v += shuffle(rs[1:len(rs)-1])
			v += string(rs[len(rs) - 1])
		}
		ret += v + " "
	}
	ret = ret[:len(ret)-1]
	fmt.Println(str)
	fmt.Println(ret)
}

func shuffle(r []rune) string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(r) - 1; i++ {
		j := random.Intn(len(r) - i - 1)
		r[len(r) - i - 1], r[j] = r[j], r[len(r) - i - 1]
	}
	return  string(r)
}
