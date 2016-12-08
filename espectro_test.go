package espectro

import (
	"testing"
	"math"
	"bytes"
)

func TestProcessSignal(t *testing.T) {
	for i, tt := range []struct {
		signal     Signal
		wantedOut  Spectrum
		shouldFail bool
	}{
		// sin(x); 0 < x < 2*PI; 16 samples
		{
			Signal{0, 0.3826834323650898, 0.7071067811865475, 0.9238795325112867, 1, 0.9238795325112867, 0.7071067811865476, 0.3826834323650899, 1.2246467991473532e-16, -0.38268343236508967, -0.7071067811865475, -0.9238795325112865, -1, -0.9238795325112866, -0.707106781186547, -0.3826834323650887},
			Spectrum{0, 8, 0, 0, 0, 0, 0, 0},
			false,
		},
		// sin(x); 0 < x < 4*PI; 32 samples
		{
			Signal{0, 0.3826834323650898, 0.7071067811865475, 0.9238795325112867, 1, 0.9238795325112867, 0.7071067811865476, 0.3826834323650899, 1.2246467991473532e-16, -0.38268343236508967, -0.7071067811865475, -0.9238795325112865, -1, -0.9238795325112866, -0.707106781186547, -0.3826834323650887, 1.5314274795707798e-15, 0.38268343236509156, 0.7071067811865492, 0.9238795325112878, 1, 0.9238795325112853, 0.7071067811865446, 0.38268343236508556, -4.961676478456545e-15, -0.3826834323650948, -0.7071067811865517, -0.9238795325112892, -1, -0.9238795325112841, -0.7071067811865422, -0.3826834323650824},
			Spectrum{0, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			false,
		},
		// 2*sin(x/2-PI/2); 0 < x < 4*PI; 32 samples
		{
			Signal{-2, -1.9615705608064609, -1.8477590650225735, -1.6629392246050905, -1.414213562373095, -1.1111404660392044, -0.7653668647301796, -0.3901806440322565, 0, 0.3901806440322565, 0.7653668647301796, 1.1111404660392041, 1.414213562373095, 1.6629392246050907, 1.847759065022574, 1.961570560806461, 2, 1.9615705608064604, 1.8477590650225726, 1.662939224605089, 1.4142135623730927, 1.1111404660392012, 0.7653668647301757, 0.39018064403225194, -5.084141158371281e-15, -0.39018064403226194, -0.765366864730185, -1.1111404660392097, -1.4142135623730998, -1.6629392246050945, -1.8477590650225764, -1.9615705608064624},
			Spectrum{0, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			false,
		},
		// sin(x) + 2*sin(x/2-PI/2); 0 < x < 4*PI; 32 samples
		{
			Signal{-2, -1.578887128441371, -1.140652283836026, -0.7390596920938037, -0.4142135623730949, -0.18726093352791762, -0.05826008354363199, -0.007497211667166603, 1.2246467991473532e-16, 0.007497211667166825, 0.0582600835436321, 0.18726093352791762, 0.4142135623730949, 0.7390596920938041, 1.140652283836027, 1.5788871284413724, 2.0000000000000013, 2.344253993171552, 2.554865846209122, 2.5868187571163768, 2.4142135623730927, 2.0350199985504864, 1.4724736459167203, 0.7728640763973376, -1.0045817636827827e-14, -0.7728640763973567, -1.4724736459167367, -2.035019998550499, -2.4142135623731, -2.5868187571163785, -2.5548658462091187, -2.3442539931715447},
			Spectrum{0, 32, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			false,
		},
		// odd samples are not allowed
		{
			Signal{0, 0, 0, 0, 0},
			Spectrum{},
			true,
		},
	} {
		out, err := processSignal(tt.signal)
		if err != nil {
			if !tt.shouldFail {
				t.Errorf(err.Error())
			}
			continue
		}
		outLen := len(out)
		if oL := len(tt.signal) / 2; outLen != oL {
			t.Errorf("Signal %d: length of output should equal %d, got %d", i, oL, outLen)
			continue
		}
		if outLen != len(tt.wantedOut) {
			t.Errorf("Signal %d: Wanted output and processed output should have the same length", i)
			continue
		}
		for i, n := range tt.wantedOut {
			// are numbers approximately the same?
			const error = 1e-12
			diff := math.Abs(n - out[i])
			if diff > error {
				t.Errorf("Spectrum %v should be approximately %v", out, tt.wantedOut)
				break
			}
		}
	}
}

func TestAverage(t *testing.T) {
	const channels = 5
	spectrum := Spectrum{
		1.1, 2.2, 1.5, 2.1,
		0.5, 2.0, 0.7, 1.4,
		4.1, 3.2, 5.3, 3.4,
		1.4, 0.2, 1.0, 0.2,
		2.3, 2.7, 1.8, 2.1,
		3.0, 3.0, 2.0,
	}
	out := average(spectrum, channels)

	if len(out) != channels {
		t.Errorf("Output lenght must be equal to %d, got %d", channels, len(out))
	}

	for i, want := range []float64{
		1.725,
		1.15,
		4,
		0.7,
		2.225,
	} {
		if want != out[i] {
			t.Errorf("Wanted: %f at %d, got: %f", want, i, out[i])
		}
	}
}

func TestSignalFromBytes(t *testing.T) {
	for _, tt := range []struct {
		rawData []byte
		signal  Signal
	}{
		{
			[]byte{0x00, 0x00, 0xFF, 0x7F, 0x00, 0x80}, // little-endian 0,32767,-32768
			Signal{0, math.MaxInt16, math.MinInt16},
		},
	} {
		signal := signalFromBytes(bytes.NewBuffer(tt.rawData))
		for i, samp := range signal {
			if samp != tt.signal[i] {
				t.Errorf("wanted %f, got %f", tt.signal[i], samp)
				break
			}
		}
	}
}