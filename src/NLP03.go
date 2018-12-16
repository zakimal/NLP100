package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	str = strings.Replace(str, ".", "", -1)
	str = strings.Replace(str, ",", "", -1)
	for _, n := range strings.Split(str, " ") {
		fmt.Println(len(n))
	}
}
