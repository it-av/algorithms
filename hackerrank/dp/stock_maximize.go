/*
https://www.hackerrank.com/challenges/stockmax
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var (
	r *Reader
)

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

func (r *Reader) NextInt() (n int) {
	n, err := strconv.Atoi(r.data[r.p])
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

func Solve(a []int) (sum int) {
	l := len(a) - 1

	rem := a[l]
	l--

	for ; l >= 0; l-- {
		if rem > a[l] {
			sum += rem - a[l]
		} else {
			rem = a[l]
		}
	}

	return
}

func main() {
	t := r.NextInt()

	for ; t > 0; t-- {
		n := r.NextInt()
		var a []int
		for i := 0; i < n; i++ {
			a = append(a, r.NextInt())
		}
		fmt.Println(Solve(a))
	}
}
