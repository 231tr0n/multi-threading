package main

import (
	"flag"
	"log"
	"log/slog"

	"visualization.io/fibonacci"
)

var question *int

func init() {
	log.SetFlags(0)
	log.SetPrefix("")
}

func main() {
	question = flag.Int("n", 47, "Value to find the answer")
	debug := flag.Bool("debug", false, "Log level debug")
	flag.Parse()
	if *debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	slog.Info("Parameters", "FACTORIALSUM", *question)
	slog.Info("Solution", "ans", fibonacci.FibonacciWorkerPool(*question))
}
