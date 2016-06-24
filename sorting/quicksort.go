package sorting

import "math/rand"

func QuickSort(a *[]float64) {
	partition := func(arr *[]float64, low, high int) int {
		a := *arr
		ind := low + rand.Intn(high-low+1)
		a[ind], a[high] = a[high], a[ind]
		pivot := a[high]
		i := low
		for j := low; j < high; j++ {
			if a[j] <= pivot {
				a[i], a[j] = a[j], a[i]
				i++
			}
		}
		a[i], a[high] = a[high], a[i]
		return i
	}
	var sort func(*[]float64, int, int)
	sort = func(a *[]float64, low, high int) {
		if low < high {
			p := partition(a, low, high)
			sort(a, low, p-1)
			sort(a, p+1, high)
		}
	}
	sort(a, 0, len(*a)-1)
}

func QuickSort2(arr *[]float64) (mergedSolution []float64) {
	a := *arr

	if len(a) <= 1 {
		return a
	}

	pivotIndex := rand.Intn(len(a) - 1)

	less := []float64{}
	larger := []float64{}
	for i, _ := range a {
		if i == pivotIndex {
			continue
		}
		if a[i] < a[pivotIndex] {
			less = append(less, a[i])
		} else {
			larger = append(larger, a[i])
		}
	}

	mergedSolution = QuickSort2(&less)
	mergedSolution = append(mergedSolution, a[pivotIndex])
	mergedSolution = append(mergedSolution, QuickSort2(&larger)...)

	return
}
