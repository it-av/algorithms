package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
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
	scanner := bufio.NewScanner(os.Stdin)

	var buffer bytes.Buffer
	for scanner.Scan() {
		buffer.WriteString(scanner.Text())
		buffer.WriteString(" ")
	}

	r.data = strings.Split(r.spaceMap(buffer.String()), " ")
}

func (r *Reader) NextInt() (n int) {
	n, err := strconv.Atoi(r.data[r.p])
	if err != nil {
		panic(err)
	}
	r.p += 1
	return
}

func (r *Reader) NextLong() (n int64) {
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

func main() {
	n := r.NextInt()

	arr := make([]int, 0)

	start := make(map[int]int, n)
	end := make(map[int]int, n)
	iToL := make(map[int]*int, n)

	result := make(map[int]int, n)

	for i := 1; i <= n; i++ {
		l := r.NextInt()
		r := r.NextInt()

		arr = append(arr, l)
		arr = append(arr, r)

		start[l] = i
		end[r] = i

		iToL[i] = &l
	}

	sort.Sort(sort.IntSlice(arr))

	opens := make(map[int]bool, 0)

	for _, v := range arr {
		if id, ok := start[v]; ok {
			opens[id] = true
			continue
		}

		if id, ok := end[v]; ok {
			delete(opens, id)

			for openId, _ := range opens {
				if *iToL[id] < *iToL[openId] {
					continue
				}
				result[openId]++
			}
		}
	}

	for i := 1; i <= n; i++ {
		fmt.Println(result[i])
	}
}
