package sorting

func SelectionSort(a []int) {
	n := len(a)

	for i := 0; i < n-1; i++ {
		smallest := i

		for j := i + 1; j < n; j++ {
			if a[j] < a[smallest] {
				smallest = j
			}
		}

		a[i], a[smallest] = a[smallest], a[i]
	}
}
