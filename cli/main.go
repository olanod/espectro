package main

import (
	"flag"
	"log"
	"github.com/olanod/espectro"
	"os"
	"time"
)

func main() {
	n := *flag.Uint("n", 5, "Distribute output spectrum in n")
	sampleRate := *flag.Int("r", 44100, "Sample rate")
	interval := *flag.Duration("w", 50 * time.Millisecond, "Window size")
	flag.Parse()

	err := espectro.Process(sampleRate, interval, n)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}