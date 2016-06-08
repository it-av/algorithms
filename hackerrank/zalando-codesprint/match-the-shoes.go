package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Shoe struct {
	Id         int
	Popularity int
}

type Shoes []*Shoe

func (s Shoes) Len() int {
	return len(s)
}

func (s Shoes) Less(i, j int) bool {
	if s[i].Popularity == s[j].Popularity {
		return s[i].Id < s[j].Id
	}
	return s[i].Popularity > s[j].Popularity
}

func (s Shoes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	var K, M, N int
	fmt.Scanf("%d %d %d", &K, &M, &N)

	A := make(map[int]int, M)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < N; i++ {
		scanner.Scan()
		number, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		A[number] += 1
	}

	shoes := Shoes{}
	for k, v := range A {
		shoes = append(shoes, &Shoe{
			Id:         k,
			Popularity: v,
		})
	}

	sort.Sort(shoes)

	for i := 0; i < K; i++ {
		fmt.Println(shoes[i].Id)
	}
}

/*
Sample Input:
3 4 8
2
1
2
0
3
3
1
2
*/

/*
Sample Output:
2
1
3
*/

/* results:
Test Case #0: 0.22 s
Test Case #1: 0.18 s -> 0.17 s
Test Case #2: 0.23 s
Test Case #3: 0.16 s
Test Case #4: 0.05 s
Test Case #5: 0.23 s
Test Case #6: 0.05 s -> 0.04 s
Test Case #7: 0.28 s
Test Case #8: 0.03 s -> 0.02 s
Test Case #9: 0.02 s
Test Case #10: 0.15 s
Test Case #11: 0.25 s -> 0.24 s
Test Case #12: 0.09 s
Test Case #13: 0.03 s
Test Case #14: 0.15 s
Test Case #15: 0.06 s
Test Case #16: 0.09 s
Test Case #17: 0.26 s
Test Case #18: 0.03 s -> 0.02 s
Test Case #19: 0.27 s
Test Case #20: 0 s
Test Case #21: 0 s
Test Case #22: 0.31 s -> 0.3 s
*/
