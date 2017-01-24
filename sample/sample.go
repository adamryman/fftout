package sample

import "math"
import . "github.com/y0ssar1an/q"

func Sin(sampleRate int, frequency, amplitude, phase float64) []float64 {
	samples := make([]float64, sampleRate)
	for i := 0; i < sampleRate; i++ {
		Q(1 / frequency)
		samples[i] = amplitude *
			math.Sin(
				frequency*(float64(i)/float64(sampleRate))*2*math.Pi+
					phase*(math.Pi/180))
		Q(samples[i])
	}
	return samples
}
