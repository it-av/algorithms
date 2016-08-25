// TASK:
// https://www.hackerrank.com/contests/w22/challenges/submask-queries-

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

var (
	r                        *Reader
	a                        []int64
	allPrintingElementsSlice IntSlice
)

type IntSlice []int64

func (is IntSlice) Less(i, j int) bool {
	return is[i] < is[j]
}

func (is IntSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func (is IntSlice) Len() int {
	return len(is)
}

type Reader struct {
	data []string
	p    int
}

func (r *Reader) spaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return ' '
		}
		return r
	}, str)
}

func (r *Reader) ReadAll() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	r.data = strings.Split(r.spaceMap(string(bytes)), " ")
}

func (r *Reader) Next() (s string) {
	s = r.data[r.p]
	r.p += 1

	return s
}

func (r *Reader) NextInt() (n int) {
	n, err := strconv.Atoi(r.data[r.p])
	if err != nil {
		panic(err)
	}
	r.p += 1
	return
}

func (r *Reader) NextInt64() (n int64) {
	n, err := strconv.ParseInt(r.data[r.p], 10, 64)
	if err != nil {
		panic(err)
	}
	r.p += 1
	return
}

func init() {
	r = &Reader{}
	r.ReadAll()
}

// Solution starts here
func setValues(n int64, x int64) {
	for _, v := range allPrintingElementsSlice {
		if v > n {
			return
		}
		if n|v == n {
			a[v] = x
		}
	}
}

func printValues(n int64) {
	fmt.Println(a[n])
}

func xor(n int64, x int64) {
	for _, v := range allPrintingElementsSlice {
		if v > n {
			return
		}
		if n|v == n {
			a[v] ^= x
		}
	}
}

func main() {
	r.NextInt()
	m := r.NextInt()

	a = make([]int64, 65536)

	allPrintingElements := make(map[int64]bool, 0)
	operations := make([][3]int64, 0)

	for ; m > 0; m-- {
		op := r.NextInt64()

		if op == 3 {
			s := r.Next()
			subsetInInt, err := strconv.ParseInt(s, 2, 64)
			if err != nil {
				panic(err)
			}

			allPrintingElements[subsetInInt] = true
			operations = append(operations, [3]int64{3, 0, subsetInInt})

		} else {
			x := r.NextInt64()
			s := r.Next()

			subsetInInt, err := strconv.ParseInt(s, 2, 64)
			if err != nil {
				panic(err)
			}

			operations = append(operations, [3]int64{op, x, subsetInInt})
		}
	}

	allPrintingElementsSlice = make(IntSlice, 0)
	for v, _ := range allPrintingElements {
		allPrintingElementsSlice = append(allPrintingElementsSlice, v)
	}

	allPrintingElements = map[int64]bool{}
	sort.Sort(allPrintingElementsSlice)

	for _, o := range operations {
		if o[0] == 1 {
			setValues(o[2], o[1])
			continue
		}

		if o[0] == 2 {
			xor(o[2], o[1])
			continue
		}

		printValues(o[2])
	}
}
