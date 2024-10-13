package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var url string
var requests int

func init() {
	log.SetFlags(0)
	log.SetPrefix("INFO: ")

	if len(os.Args) != 3 {
		log.Fatalln("Provide the right arguments in order: website-url, requests")
	}

	url = os.Args[1]
	var err error
	requests, err = strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
}

func worker(wg *sync.WaitGroup, requestsChan <-chan int) {
	for n := range requestsChan {
		startTime := time.Now()
		res, err := http.Get(url + "?n=" + strconv.Itoa(n))
		if err != nil {
			log.Println("REQ", n, err)
			continue
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println("REQ", n, err)
			continue
		}
		log.Println("REQ", n, time.Since(startTime).Round(time.Second), strings.TrimSpace(string(body)))
	}
	wg.Done()
}

func main() {
	startTime := time.Now()
	requestsChan := make(chan int, requests)
	var wg sync.WaitGroup

	for range runtime.NumCPU() {
		wg.Add(1)
		go worker(&wg, requestsChan)
	}
	for i := 0; i <= requests; i++ {
		requestsChan <- i
	}
	close(requestsChan)

	wg.Wait()

	duration := int(time.Since(startTime).Round(time.Second).Seconds())
	if duration > 0 {
		requestsPerSecond := int(requests / duration)
		log.Printf("Load test completed in %ds with %d rps\n", duration, requestsPerSecond)
	}
}
