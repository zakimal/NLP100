package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	str = strings.Replace(str, ".", "", -1)
	dict := make(map[string]int)
	for i, w := range strings.Split(str, " ") {
		switch i + 1 {
		case 1, 5, 6, 7, 8, 9, 15, 16, 19:
			dict[string([]rune(w)[0])] = i
		default:
			dict[string([]rune(w)[0:2])] = i
		}
	}
	fmt.Println(dict)
}
