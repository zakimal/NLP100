package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("%s: Usage: %s <file> <col>\n", os.Args[0], os.Args[0])
	}
	src, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	col, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(src)
	tbl := make(table, 0)
	rec := record{}
	rtemp := make([]rune, 0)
	index := 0
	for {
		r, _, _ := reader.ReadRune()
		if r == 0 {
			break
		}
		if r == rune('\t') {
			index++
			if index == 3 {
				rec.temp, _ = strconv.ParseFloat(string(rtemp), 64)
			}
		} else {
			switch index {
			case 0:
				rec.pref = append(rec.pref, r)
			case 1:
				rec.city = append(rec.city, r)
			case 2:
				rtemp = append(rtemp, r)
			case 3:
				rec.date = append(rec.date, r)
			}
		}
		if r == rune('\n') {
			tbl = append(tbl, rec)
			rec = record{}
			rtemp = nil
			index = 0
		}
	}
	switch col {
	case 1:
		sort.Sort(ByPref{tbl})
	case 2:
		sort.Sort(ByCity{tbl})
	case 3:
		sort.Sort(ByTemp{tbl})
	case 4:
		sort.Sort(ByDate{tbl})
	}
	for _, r := range tbl {
		fmt.Printf("%s\t", toString(r.pref))
		fmt.Printf("%s\t", toString(r.city))
		fmt.Printf("%.1f\t", r.temp)
		fmt.Printf("%s", toString(r.date))
	}
}

type record struct {
	pref []rune
	city []rune
	temp float64
	date []rune
}
type table []record

func (t table) Len() int      { return len(t) }
func (t table) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

type ByPref struct{ table }

func (byPref ByPref) Less(i, j int) bool {
	return less(byPref.table[i].pref, byPref.table[j].pref)
}

type ByCity struct{ table }

func (byCity ByCity) Less(i, j int) bool {
	return less(byCity.table[i].city, byCity.table[j].city)
}

type ByTemp struct{ table }

func (byTemp ByTemp) Less(i, j int) bool {
	return byTemp.table[i].temp < byTemp.table[j].temp
}

type ByDate struct{ table }

func (byDate ByDate) Less(i, j int) bool {
	return less(byDate.table[i].date, byDate.table[j].date)
}

func less(r0, r1 []rune) bool {
	ret := true
	swapped := false
	if len(r0) > len(r1) {
		r0, r1 = r1, r0
		swapped = true
	}
	for i := 0; i < len(r0); i++ {
		if r0[i] < r1[i] {
			ret = true
			break
		}
		if r1[i] < r0[i] {
			ret = false
			break
		}
		if i == len(r1)-1 {
			ret = true
		}
	}
	if swapped {
		return ret
	} else {
		return !ret
	}
}

func toString(r []rune) string {
	l := len(r)
	ret := make([]string, l)
	for i := range ret {
		ret = append(ret, string(r[i]))
	}
	return strings.Join(ret, "")
}
