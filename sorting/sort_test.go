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

func quickSortWrapper(vals []float64) *[]float64 {
	return QuickSort(vals...)
}

func sortWrapper(vals []float64) *[]float64 {
	v := sort.Float64Slice(vals)
	v.Sort()

	f64 := []float64(v)

	return &f64
}

var testCases = []struct {
	in  []float64
	out *[]float64
}{
	{[]float64{2, 3, 1}, &[]float64{1, 2, 3}},
}

func TestQuickSort2(t *testing.T) {
	for _, test := range testCases {
		result := QuickSort(test.in...)
		if !reflect.DeepEqual(result, test.out) {
			t.Errorf("Unexpected listing: want=%v got=%v \n", test.out, result)
		}
	}
}

/*
go test github.com/it-av/algorithms/sorting -v
*/
