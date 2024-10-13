package main

import (
	"net/http"
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

func fibonacciWorkerPool(n int) int {
	var ans = 0
	var requestChan = make(chan int, n)
	var responseChan = make(chan int, n)

	for i := 0; i <= n; i++ {
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := strconv.Atoi(r.FormValue("n"))
		if err != nil {
			http.Error(w, "Wrong parameter 'n'", http.StatusBadRequest)
			return
		}
		w.Write([]byte(strconv.Itoa(fibonacciWorkerPool(n))))
	})
	http.ListenAndServe(":8080", nil)
}