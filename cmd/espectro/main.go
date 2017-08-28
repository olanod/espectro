package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/olanod/espectro"
)

var version string

func main() {
	n := flag.Uint("n", 10, "Distribute output spectrum in n values")
	sampleRate := flag.Uint("r", 44100, "Sample rate of the raw audio signal")
	interval := flag.Duration("i", 50*time.Millisecond, "Window size,")
	ver := flag.Bool("v", false, "Show version and exit")
	flag.Parse()

	if *ver {
		log.Println(version)
		os.Exit(0)
	}

	err := espectro.ProcessAudio(os.Stdin, os.Stdout, int(*sampleRate), *interval, int(*n))
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
