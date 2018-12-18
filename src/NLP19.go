package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
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
	if !(col == 1 || col == 2 || col == 3 || col == 4) {
		os.Exit(0)
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
		sort.Sort(byPref{tbl})
	case 2:
		sort.Sort(byCity{tbl})
	case 3:
		sort.Sort(byTemp{tbl})
	case 4:
		sort.Sort(byDate{tbl})
	}
	tbl = adventCounter(tbl, col)
	sort.Sort(byAdvent{tbl})

	for i := 0; i < len(tbl); i++ {
		fmt.Printf("%d\t", tbl[i].advent)
		fmt.Printf("%s\t", toString(tbl[i].pref))
		fmt.Printf("%s\t", toString(tbl[i].city))
		fmt.Printf("%.1f\t", tbl[i].temp)
		fmt.Printf("%s", toString(tbl[i].date))
	}
}

type record struct {
	pref   []rune
	city   []rune
	temp   float64
	date   []rune
	advent int
}
type table []record

func (t table) Len() int      { return len(t) }
func (t table) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

type byPref struct{ table }

func (byPref byPref) Less(i, j int) bool {
	return less(byPref.table[i].pref, byPref.table[j].pref)
}

type byCity struct{ table }

func (byCity byCity) Less(i, j int) bool {
	return less(byCity.table[i].city, byCity.table[j].city)
}

type byTemp struct{ table }

func (byTemp byTemp) Less(i, j int) bool {
	return byTemp.table[i].temp < byTemp.table[j].temp
}

type byDate struct{ table }

func (byDate byDate) Less(i, j int) bool {
	return less(byDate.table[i].date, byDate.table[j].date)
}

type byAdvent struct{ table }

func (byAdvent byAdvent) Less(i, j int) bool {
	return byAdvent.table[i].advent > byAdvent.table[j].advent
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

func adventCounter(tbl table, col int) table {
	count := 0
	index := 0
	for i := 0; i < len(tbl); i++ {
		switch col {
		case 1:
			if reflect.DeepEqual(tbl[index].pref, tbl[i].pref) {
				count++
			} else {
				for j := index; j < i; j++ {
					tbl[j].advent = count
				}
				count = 1
				index = i
			}
		case 2:
			if reflect.DeepEqual(tbl[index].city, tbl[i].city) {
				count++
			} else {
				for j := index; j < i; j++ {
					tbl[j].advent = count
				}
				count = 1
				index = i
			}
		case 3:
			if tbl[index].temp == tbl[i].temp {
				count++
			} else {
				for j := index; j < i; j++ {
					tbl[j].advent = count
				}
				count = 1
				index = i
			}
		case 4:
			if reflect.DeepEqual(tbl[index].date, tbl[i].date) {
				count++
			} else {
				for j := index; j < i; j++ {
					tbl[j].advent = count
				}
				count = 1
				index = i
			}
		}
		if i == len(tbl)-1 {
			tbl[i].advent = count
		}
	}
	return tbl
}

func toString(r []rune) string {
	l := len(r)
	ret := make([]string, l)
	for i := range ret {
		ret = append(ret, string(r[i]))
	}
	return strings.Join(ret, "")
}
