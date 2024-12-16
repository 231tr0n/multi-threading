package fibonacci_test

import (
	"testing"

	"visualization.io/fibonacci"
)

func failTestCase(t *testing.T, i, o, w any) {
	t.Helper()
	t.Error("Input:", i, "|", "Output:", o, "|", "Want:", w)
}

func TestFibonacci(t *testing.T) {
	tests := [][2]int{
		{
			0,
			0,
		},
		{
			1,
			1,
		},
		{
			2,
			1,
		},
		{
			3,
			2,
		},
		{
			4,
			3,
		},
		{
			5,
			5,
		},
	}

	for _, test := range tests {
		t.Log(test)
		out := fibonacci.Fibonacci(test[0])

		if out != test[1] {
			failTestCase(t, test[0], out, test[1])
		}
	}
}
