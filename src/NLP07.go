package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(getTemplate(12, "気温" , 22.4))
}

func getTemplate(x, y, z interface{}) string {
	return convert2string(x) + "時の" + convert2string(y) + "は" + convert2string(z) + "です。"
}

func convert2string(something interface{}) string {
	switch casted := something.(type) {
	case int:
		return strconv.Itoa(casted)
	case float64:
		return strconv.FormatFloat(casted, 'f', -1, 64)
	case string:
		return casted
	default:
		return "OMG!"
	}
}
