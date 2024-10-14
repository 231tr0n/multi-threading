package fibonacci

import (
	"log/slog"
	"sync"
)

// Fibonacci is used to generate fibonacci number for a given number.
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func worker(wg *sync.WaitGroup, id int, requestChan <-chan int, responseChan chan<- int) {
	slog.Debug("Starting", "worker", id)
	for n := range requestChan {
		var ans = Fibonacci(n)
		slog.Debug("Processed", "worker", id, "n", n, "ans", ans)
		responseChan <- ans
	}
	slog.Debug("Stopping", "worker", id)
	wg.Done()
}

// FibonacciWorkerPool is used to sum up all the fibonacci numbers till the given number starting from 0.
func FibonacciWorkerPool(n int) int {
	var wg sync.WaitGroup
	var ans = 0
	var requestChan = make(chan int, n)
	var responseChan = make(chan int, n)

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(&wg, i, requestChan, responseChan)
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

	wg.Wait()

	return ans
}
