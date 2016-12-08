package espectro

import (
	"math"
	"math/cmplx"
	"github.com/olanod/espectro/fft"
	"time"
	"errors"
	"encoding/binary"
	"bytes"
)

type Signal []float64
type Spectrum []float64

// Process reads a raw PCM audio data from the stdIn
// and outputs averaged frequency data to the stdOut by chunks
// and
func Process(rate int, chunkTime time.Duration, n uint) (err error) {
	// TODO
	return nil
}

func signalFromBytes(buf *bytes.Buffer) (signal Signal) {
	sample := make([]byte, 2)
	for {
		n, _ := buf.Read(sample)
		if n == 0 {
			return
		}
		// TODO convert to float in -1,1 range
		val := float64(int16(binary.LittleEndian.Uint16(sample)))
		signal = append(signal, val)
	}
}

func processSignal(x Signal) (Spectrum, error) {
	if !powerOf2(len(x)) {
		return nil, errors.New("The number of samples most be a power of 2")
	}
	spectrum := Spectrum{}
	X := fft.FFTReal(x)
	for i, lX := 0, len(X) / 2; i < lX; i++ {
		spectrum = append(spectrum, cmplx.Abs(X[i]))
	}
	return spectrum, nil
}

func average(in Spectrum, n int) []float64 {
	out := []float64{}
	inLength := len(in)
	chunkSize := int(math.Floor(float64(inLength / n)))
	for i, sum := 1, 0.0; i <= inLength; i++ {
		sum += in[i - 1]
		if i % chunkSize == 0 {
			out = append(out, sum / float64(chunkSize))
			sum = 0.0
		}
	}
	return out
}

func powerOf2(n int) bool {
	return n & (n - 1) == 0
}