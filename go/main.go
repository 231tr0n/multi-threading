package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

var question *int

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func worker(requestChan <-chan int, responseChan chan<- int) {
	for n := range requestChan {
		var ans = fibonacci(n)
		responseChan <- ans
	}
}

func fibonacciWorkerPool(threads int, n int) int {
	var ans = 0
	var requestChan = make(chan int, n)
	var responseChan = make(chan int, n)

	for i := 1; i <= threads; i++ {
		go worker(requestChan, responseChan)
	}

	for i := 0; i <= n; i++ {
		requestChan <- i
	}
	close(requestChan)

	for i := 0; i <= n; i++ {
		var n = <-responseChan
		ans += n
	}
	close(responseChan)

	return ans
}

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Total:", fibonacciWorkerPool(runtime.NumCPU(), n))
}
