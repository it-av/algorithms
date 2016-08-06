package sorting

import (
	"reflect"
	"sort"
	"testing"
	"testing/quick"
)

func TestQuickSort(t *testing.T) {
	err := quick.CheckEqual(quickSortWrapper, sortWrapper, nil)

	if err != nil {
		t.Error(err)
	}
}

func quickSortWrapper(vals *[]float64) *[]float64 {
	if vals == nil {
		return nil
	}
	QuickSort(vals)
	return vals
}

func sortWrapper(vals *[]float64) *[]float64 {
	if vals == nil {
		return nil
	}
	v := sort.Float64Slice(*vals)
	v.Sort()

	f64 := []float64(v)

	return &f64
}

var testCases = []struct {
	in  *[]float64
	out *[]float64
}{
	{&[]float64{2, 3, 1}, &[]float64{1, 2, 3}},
}

func TestQuickSort2(t *testing.T) {
	for _, test := range testCases {
		QuickSort(test.in)
		if !reflect.DeepEqual(test.in, test.out) {
			t.Errorf("Unexpected listing: want=%v got=%v \n", test.out, test.in)
		}
	}
}

var testIntCases = []struct {
	in  []int
	out []int
}{
	{[]int{12, 9, 3, 7, 14, 11, 6, 2, 10, 5}, []int{2, 3, 5, 6, 7, 9, 10, 11, 12, 14}},
}

func TestSelectionSort(t *testing.T) {
	for _, test := range testIntCases {
		arr := make([]int, len(test.in))
		copy(arr, test.in)

		SelectionSort(arr)
		if !reflect.DeepEqual(arr, test.out) {
			t.Errorf("Unexpected listing: want=%v got=%v \n", test.out, arr)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	for _, test := range testIntCases {
		arr := make([]int, len(test.in))
		copy(arr, test.in)
		InsertionSort(arr)
		if !reflect.DeepEqual(arr, test.out) {
			t.Errorf("Unexpected listing: want=%v got=%v \n", test.out, arr)
		}
	}
}

func TestMergeSort(t *testing.T) {
	for _, test := range testIntCases {
		arr := make([]int, len(test.in))
		copy(arr, test.in)
		MergeSort(arr)
		if !reflect.DeepEqual(arr, test.out) {
			t.Errorf("Unexpected listing: want=%v got=%v \n", test.out, arr)
		}
	}
}

func BenchmarkQuickSort(b *testing.B) {
	n := 10000
	a := make([]float64, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < n; i++ {
			a[i] = float64(n - i)
		}
		QuickSort(&a)
	}
}

func BenchmarkStandartQuickSort(b *testing.B) {
	n := 10000
	a := make([]float64, n)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < n; i++ {
			a[i] = float64(n - i)
		}
		v := sort.Float64Slice(a)
		v.Sort()
	}
}

func BenchmarkStandartQuickSort2(b *testing.B) {
	n := 10000
	a := make([]float64, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < n; i++ {
			a[i] = float64(n - i)
		}
		QuickSort2(&a)
	}
}

/*
go test github.com/it-av/algorithms/sorting -v
go test github.com/it-av/algorithms/sorting -v --bench=.
*/
