package main

import (
	"fmt"
	"reflect"
)

func main() {
	str1 := "paraparaparadise"
	str2 := "paragraph"
	target := "se"
	X := getCharNGram([]rune(str1), 2)
	Y := getCharNGram([]rune(str2), 2)
	fmt.Println(convertToStringSlice(X))
	fmt.Println(convertToStringSlice(Y))
	fmt.Println(convertToStringSlice(getIntersection(X, Y)))
	fmt.Println(convertToStringSlice(getUnion(X, Y)))
	fmt.Println(convertToStringSlice(getDifference(X, Y)))
	SE := getCharNGram([]rune(target), 2)
	fmt.Println(contain(X, SE[0]))
	fmt.Println(contain(Y, SE[0]))
}

func getCharNGram(r []rune, N int) [][]rune {
	ret := make([][]rune, len(r)-N+1)
	for i := 0; i < len(r)-N+1; i++ {
		ret[i] = r[i : i+N]
	}
	return ret
}

func contain(set [][]rune, word []rune) bool {
	flag := false
	for _, v := range set {
		if reflect.DeepEqual(v, word) {
			flag = true
			break
		}
	}
	return flag
}

// X + Y
func getUnion(X, Y [][]rune) [][]rune {
	union := make([][]rune, 0)
	for _, yele := range Y {
		if !contain(X, yele) {
			union = append(union, yele)
		}
	}
	return X
}

// X * Y
func getIntersection(X, Y [][]rune) [][]rune {
	intersection := make([][]rune, 0)
	for _, yele := range Y {
		if contain(X, yele) {
			intersection = append(intersection, yele)
		}
	}
	return intersection
}

// X - Y
func getDifference(X, Y [][]rune) [][]rune {
	difference := make([][]rune, 0)
	for _, xele := range X {
		if !contain(Y, xele) {
			difference = append(difference, xele)
		}
	}
	return difference
}

func convertToStringSlice(rs [][]rune) []string {
	ret := make([]string ,0)
	for _, r := range rs {
		ret = append(ret, string(r))
	}
	return ret
}