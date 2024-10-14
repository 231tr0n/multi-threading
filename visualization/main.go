package main

import (
	"flag"
	"log"
	"log/slog"
	"sync"
)

var question *int

func init() {
	log.SetFlags(0)
	log.SetPrefix("")
}

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

func main() {
	question = flag.Int("n", 47, "Value to find the answer")
	debug := flag.Bool("debug", false, "Log level debug")
	flag.Parse()
	if *debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	slog.Info("Parameters", "FACTORIALSUM", *question)
	slog.Info("Solution", "ans", FibonacciWorkerPool(*question))
}
