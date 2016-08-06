package sorting

import (
	"math"
)

func MergeSort(a []int) {
	var merge func([]int, int, int, int)
	merge = func(a []int, p, q, r int) {
		n1 := q - p + 1
		n2 := r - q

		b := make([]int, n1+1)
		c := make([]int, n2+1)

		copy(b, a[p:q+1])
		copy(c, a[q+1:r+1])

		b[n1] = math.MaxInt32
		c[n2] = math.MaxInt32

		i := 0
		j := 0

		for k := p; k <= r; k++ {
			if b[i] <= c[j] {
				a[k] = b[i]
				i++
			} else {
				a[k] = c[j]
				j++
			}
		}
	}

	var sort func([]int, int, int)
	sort = func(a []int, p, r int) {
		if p >= r {
			return
		}

		q := (p + r) / 2
		sort(a, p, q)
		sort(a, q+1, r)
		merge(a, p, q, r)
	}
	sort(a, 0, len(a)-1)
}
