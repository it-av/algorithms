package structures

import "testing"

func Add(args ...int) int {
	res := 0
	for _, v := range args {
		res += v
	}
	return res
}

var cases = []struct {
	args []int
	want int
}{
	{[]int{3, 5}, 8},
}

func TestAdd(t *testing.T) {
	for _, test := range cases {
		if got := Add(test.args...); got != test.want {
			t.Errorf("Add(%v) = %s; want %s", test.args, got, test.want)
		}
	}
}
