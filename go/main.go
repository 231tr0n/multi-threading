package main

import (
	"net/http"
	"strconv"
)

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func worker(requestChan <-chan int, responseChan chan<- int) {
	for n := range requestChan {
		ans := Fibonacci(n)
		responseChan <- ans
	}
}

func FibonacciWorkerPool(n int) int {
	ans := 0
	requestChan := make(chan int, n)
	responseChan := make(chan int, n)

	for i := 0; i <= n; i++ {
		go worker(requestChan, responseChan)
	}

	for i := 0; i <= n; i++ {
		requestChan <- i
	}
	close(requestChan)

	for i := 0; i <= n; i++ {
		n := <-responseChan
		ans += n
	}
	close(responseChan)

	return ans
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := strconv.Atoi(r.FormValue("n"))
		if err != nil {
			http.Error(w, "Wrong parameter 'n'", http.StatusBadRequest)
			return
		}
		w.Write([]byte(strconv.Itoa(FibonacciWorkerPool(n))))
	})
	http.ListenAndServe(":8080", nil)
}
