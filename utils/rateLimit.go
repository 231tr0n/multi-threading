package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var url string
var requests int

func init() {
	log.SetFlags(0)
	log.SetPrefix("INFO: ")

	if len(os.Args) != 4 {
		log.Fatalln("Provide the right arguments in order: website-url, requests")
	}

	url = os.Args[1]
	var err error
	requests, err = strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
}

func worker(requestsChan <-chan int) {
	for id := range requestsChan {
		startTime := time.Now()
		res, err := http.Get(url + "?n=" + strconv.Itoa(id))
		if err != nil {
			log.Println("REQ", id, err)
			continue
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println("REQ", id, err)
			continue
		}
		log.Println("REQ", id, time.Since(startTime).Round(time.Second), strings.TrimSpace(string(body)))
	}
}

func main() {
	startTime := time.Now()
	requestsChan := make(chan int, requests)

	for range runtime.NumCPU() {
		go worker(requestsChan)
	}
	for i := range requests {
		requestsChan <- i
	}
	close(requestsChan)

	duration := int(time.Since(startTime).Round(time.Second).Seconds())
	if duration > 0 {
		requestsPerSecond := int(requests / duration)
		log.Printf("Load test completed in %ds with %d rps\n", duration, requestsPerSecond)
	}
}
