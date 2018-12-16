package main

import (
	"fmt"
	"strings"
)

// N-gram
// 任意の文字列や文書を連続したn個の文字で分割するテキスト分割方法
// 特にnが1の場合をユニグラム（uni-gram）、2の場合をバイグラム（bi-gram）、3の場合をトライグラム（tri-gram）と呼ぶ。
// 最初の分割後は1文字ずつ移動して順次分割を行う。
// 「図書館情報学」をバイグラムで分割すると、「図書」「書館」「館情」「情報」「報学」となる。

func main() {
	str := "I am an NLPer"
	charNGram := getCharNGram([]rune(str), 2)
	for _, r := range charNGram {
		fmt.Println("[" + string(r) + "]")
	}
	words := strings.Split(str, " ")
	runeList := make([][]rune, len(words))
	for i, w := range words {
		runeList[i] = []rune(w)
	}
	wordNGram := getWordNGram(runeList, 2)
	for _, w := range wordNGram {
		fmt.Print("[")
		for i, r := range w {
			if i == 1 {
				fmt.Print(string(r))
			} else {
				fmt.Print(string(r) + " ")
			}
		}
		fmt.Println("]")
	}
}

func getCharNGram(r []rune, N int) [][]rune {
	ret := make([][]rune, len(r)-N+1)
	for i := 0; i < len(r)-N+1; i++ {
		ret[i] = r[i : i+N]
	}
	return ret
}

func getWordNGram(r [][]rune, N int) [][][]rune {
	ret := make([][][]rune, len(r)-N+1)
	for i := 0; i < len(r)-N+1; i++ {
		ret[i] = r[i : i+N]
	}
	return ret
}
