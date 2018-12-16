package main

import "fmt"

func main() {
	s := "stressed"
	fmt.Println(reverse(s))
}

func reverse(str string) string {
	n := len(str)
	ret := make([]rune, n)
	for i, ch := range str {
		ret[n-i-1] = ch
	}
	return string(ret)
}