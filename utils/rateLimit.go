package main

import (
	"flag"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var url *string
var requests *int
var debug *bool

func init() {
	log.SetFlags(0)
	log.SetPrefix("")

	url = flag.String("host", "", "Host to request")
	requests = flag.Int("requests", 48, "Number of requests")
	debug = flag.Bool("debug", false, "Set logger level to debug")
	flag.Parse()

	_, err := http.Get(*url + "?n=1")
	if err != nil {
		slog.Error("Host not valid")
		os.Exit(1)
	}
	if *requests > 48 && *requests >= 0 {
		slog.Error("Requests should be less than or equal to 48 and greater than or equal to 0")
		os.Exit(1)
	}
	if *debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
}

func worker(wg *sync.WaitGroup, requestsChan <-chan int) {
	for n := range requestsChan {
		startTime := time.Now()
		res, err := http.Get(*url + "?n=" + strconv.Itoa(n))
		if err != nil {
			slog.Debug("Request", "n", n, "err", err)
			continue
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			slog.Debug("Request", "n", n, "err", err)
			continue
		}
		slog.Debug("Request", "n", n, "time", time.Since(startTime).Round(time.Nanosecond), "ans", strings.TrimSpace(string(body)))
	}
	wg.Done()
}

func main() {
	startTime := time.Now()
	requestsChan := make(chan int, *requests)
	var wg sync.WaitGroup

	for range runtime.NumCPU() {
		wg.Add(1)
		go worker(&wg, requestsChan)
	}
	for i := 0; i <= *requests; i++ {
		requestsChan <- i
	}
	close(requestsChan)

	wg.Wait()

	slog.Info("Completed", "duration", time.Since(startTime).Round(time.Nanosecond))
}
