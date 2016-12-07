package fft

var (
	worker_pool_size = 0
)

func FFTReal(x []float64) []complex128 {
	xl := len(x)
	if xl <= 1 || (xl & (xl - 1) != 0) {
		return []complex128{}
	}
	return radix2FFT(toComplex(x))
}

func toComplex(x []float64) []complex128 {
	y := make([]complex128, len(x))
	for n, v := range x {
		y[n] = complex(v, 0)
	}
	return y
}
