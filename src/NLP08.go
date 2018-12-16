package main

import "fmt"

func main() {
	plain := "Just do it."
	fmt.Println(encrypt(plain))
}

func encrypt(plain string) string {
	runes := []rune(plain)
	ret := make([]rune, 0)
	for _, r := range runes {
		var u rune
		if []rune("a")[0] <= r && r <= []rune("z")[0] {
			u = 219 - r
		} else {
			u = r
		}
		ret = append(ret, u)
	}
	return string(ret)
}