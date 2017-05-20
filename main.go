package main

import (
	//"fmt"
	"image"
	"math"
	"math/cmplx"

	"github.com/adamryman/fftout/dft"
	"github.com/adamryman/fftout/sample"
	ui "github.com/gizak/termui"
	"github.com/mjibson/go-dsp/fft"
	flag "github.com/spf13/pflag"
	. "github.com/y0ssar1an/q"
)

var (
	sampleRate = flag.IntP("sampleRate", "s", 12, "sample rate")
	frequency  = flag.Float64P("frequency", "f", 4.0, "frequency")
	amplitude  = flag.Float64P("amplitude", "a", 1.0, "amplitude of signal from 0 to max 1")
	phase      = flag.Float64P("phase", "p", 0.0, "phase in degrees")
	dtf        = flag.BoolP("dtf true, fft default", "d", false, "switch true to use dft rather than fft")
)

func init() {

}

type canvas struct {
	full  *image.Gray
	graph *image.Gray
}

func newCanvas(sampleRate int, vertGanularity int) canvas {
	padding := 1
	hozLine := 1
	hozTick := 2
	vertLine := 1
	vertTick := 1

	f := image.Rect(0, 0, padding*2+vertLine+vertTick+vertGanularity, padding*2+hozLine+hozTick+sampleRate)
	full := image.NewGray(f)

	return canvas{
		full:  full,
		graph: full.SubImage(image.Rect(padding+vertLine+vertTick, padding, f.Max.X-padding, f.Max.Y-padding)).(*image.Gray),
	}

}

func main() {
	flag.Parse()
	Q(*sampleRate, *frequency, *amplitude, *phase)
	r := newCanvas(*sampleRate, 100)
	_ = r

	time := sample.Sin(*sampleRate, *frequency, *amplitude, *phase)
	//for i, v := range time {
	//r.graph.SetGray(kkkkkkk)
	//}

	amp := make([]float64, *sampleRate)
	pha := make([]float64, *sampleRate)

	switch *dtf {
	case true:
		for i, v := range dft.DFTReal(time) {
			amp[i], pha[i] = cmplx.Polar(v)
			pha[i] = pha[i] * 180 / math.Pi
		}

	case false:
		for i, v := range fft.FFTReal(time) {
			amp[i], pha[i] = cmplx.Polar(v)
			pha[i] = pha[i] * 180 / math.Pi
		}
	}

	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	//g := ui.NewGauge()
	//g.Percent = 50
	//g.Width = 50
	//g.BorderLabel = "Gauge"
	timeLine := ui.NewLineChart()
	timeLine.DotStyle = 'ξ'
	timeLine.Width = 17 * *sampleRate / 12
	timeLine.Height = 8
	timeLine.Data = time

	freqLine := ui.NewLineChart()
	freqLine.DotStyle = 'ξ'
	freqLine.Width = 17 * *sampleRate / 12

	freqLine.Height = 8
	freqLine.Data = amp

	phaLine := ui.NewLineChart()
	phaLine.DotStyle = 'ξ'
	phaLine.Width = 17 * *sampleRate / 12

	phaLine.Height = 8
	phaLine.Data = pha
	ui.Body.AddRows(ui.NewRow(ui.NewCol(12, 0, timeLine, freqLine, phaLine)))
	ui.Body.Align()
	ui.Render(ui.Body)
	// handle key q pressing
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		// press q to quit
		ui.StopLoop()
	})
	ui.Handle("/sys/kbd/^C", func(ui.Event) {
		// press q to quit
		ui.StopLoop()
	})
	ui.Handle("/sys/kbd/^D", func(ui.Event) {
		// press q to quit
		ui.StopLoop()
	})

	ui.Loop()
}
