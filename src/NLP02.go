package main

import "fmt"

func main() {
	str1, str2 := "パトカー", "タクシー"
	fmt.Println(alternativeConcatenate(str1, str2))
}

func alternativeConcatenate(str1, str2 string) string {

	// << stringとrune >>
	// 参考: http://text.baldanders.info/golang/string-and-rune/
	// fmt.Printf("%d\n", len(str1)) -> 12
	// fmt.Printf("%d\n", len(r1))   ->  4

	r1, r2 := []rune(str1), []rune(str2)
	n1, n2 := len(r1), len(r2)
	ret := make([]rune, n1+n2)
	for i := 0; i < n1; i++ {
		ret[2*i], ret[2*i+1] = r1[i], r2[i]
	}
	return string(ret)
}
