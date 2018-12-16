package main

import "fmt"

func main() {
	s := "パタトクカシーー"
	t := []rune(s)
	fmt.Println(string(t[1]) + string(t[3]) + string(t[5]) + string(t[7]))
}
