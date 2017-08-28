package main

import (
	"flag"
	"log"
	"github.com/olanod/espectro"
	"os"
	"time"
)

var (
	n, sampleRate *uint
	interval *time.Duration
)

func init() {
	n = flag.Uint("n", 10, "Distribute output spectrum in n")
	sampleRate = flag.Uint("r", 44100, "Sample rate")
	interval = flag.Duration("i", 50 * time.Millisecond, "Window size")
	flag.Parse()
}

func main() {
	err := espectro.ProcessAudio(os.Stdin, os.Stdout, int(*sampleRate), *interval, int(*n))
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}