package main

import (
	"strconv"
	"testing"
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
		out := Fibonacci(test[0])

		if out != test[1] {
			failTestCase(t, strconv.Itoa(test[0]), strconv.Itoa(out), strconv.Itoa(test[1]))
		}
	}
}

func TestFibonacciWorkerPool(t *testing.T) {
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
			2,
		},
		{
			3,
			4,
		},
		{
			4,
			7,
		},
		{
			5,
			12,
		},
	}

	for _, test := range tests {
		t.Log(test)
		out := FibonacciWorkerPool(test[0])

		if out != test[1] {
			failTestCase(t, strconv.Itoa(test[0]), strconv.Itoa(out), strconv.Itoa(test[1]))
		}
	}
}
