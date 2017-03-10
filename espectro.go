package espectro

import (
	"math"
	"math/cmplx"
	"time"
	"encoding/binary"
	"bytes"
	"io"
	"fmt"
	"github.com/mjibson/go-dsp/fft"
)

type Signal []float64
type Spectrum []float64

const sampleSize = 2 // 16-bit

// Process reads a raw PCM audio data from an io.Reader
// and writes the frequency spectrum to the provided io.Writer
func ProcessAudio(rd io.Reader, wr io.Writer, sampleRate int, timeBuffer time.Duration, parts int) (err error) {
	chunk := bufferFromDuration(timeBuffer, sampleRate)
	for {
		n, _ := rd.Read(chunk)
		if n == 0 {
			return
		}
		spectrum := processSignal(signalFromBytes(chunk))
		// TODO write bytes(1 or 2 per sample?)
		wr.Write([]byte(fmt.Sprintln(average(spectrum, parts))))
	}
}

func bufferFromDuration(dur time.Duration, sampleRate int) []byte {
	size := sampleSize * sampleRate * int(dur) / int(time.Second)
	return make([]byte, size)
}

func signalFromBytes(data []byte) (signal Signal) {
	sample := make([]byte, sampleSize)
	bb := bytes.NewReader(data)
	for {
		n, _ := bb.Read(sample)
		if n == 0 {
			return
		}
		val := float64(int16(binary.LittleEndian.Uint16(sample)))
		signal = append(signal, val)
	}
}

func processSignal(x Signal) Spectrum {
	spectrum := Spectrum{}
	X := fft.FFTReal(x)
	for i, lX := 0, len(X) / 2; i < lX; i++ {
		spectrum = append(spectrum, cmplx.Abs(X[i]))
	}
	return spectrum
}

func average(in Spectrum, parts int) []float64 {
	if parts == 0 {
		return []float64(in)
	}
	out := []float64{}
	inLength := len(in)
	chunkSize := int(math.Floor(float64(inLength / parts)))
	for i, sum := 1, 0.0; i <= inLength; i++ {
		sum += in[i - 1]
		if i % chunkSize == 0 {
			out = append(out, sum / float64(chunkSize))
			sum = 0.0
		}
	}
	return out
}
