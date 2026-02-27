package main

import (
	"flag"
	"fmt"
	"math"
)

const (
	G      = 6.67430e-11
	C      = 2.99792458e8
	MSolar = 1.989e30
)

type Config struct {
	M1, M2      float64
	Distance    float64
	Inclination float64
	Duration    float64
	FPS         int
}

func h(t, f, mc, D float64) float64 {
	M := mc * MSolar
	prefactor := 4 * math.Pow(math.Pi, 2) * G / math.Pow(C, 4)
	amplitude := prefactor * M * math.Pow(f, 5.0/3.0)
	return amplitude * math.Cos(2*math.Pi*f*t) / D
}

func chirpMass(m1, m2 float64) float64 {
	return math.Pow(m1*m2, 3.0/5.0) / math.Pow(m1+m2, 1.0/5.0)
}

func orbitalFrequency(m1, m2 float64, r float64) float64 {
	_ = m1 * m2 / (m1 + m2)
	return math.Sqrt(G*(m1+m2)/math.Pow(r*1e10, 3.0)) / (2 * math.Pi)
}

func inspiralFrequency(t, f0, tau float64) float64 {
	if tau <= 0 || t >= tau {
		return 0
	}
	return f0 * math.Pow(1-t/tau, -3.0/8.0)
}

func detectorStrain(t float64, cfg Config) float64 {
	mc := chirpMass(cfg.M1, cfg.M2)
	f0 := 150.0
	tau := cfg.Duration * 0.8

	f := inspiralFrequency(t, f0, tau)
	if f < 10 || f > 2000 {
		return 0
	}

	strain := h(t, f, mc, cfg.Distance)
	hp := strain * (1 + math.Cos(cfg.Inclination)) / 2
	return hp * 1e21
}

func polarization(t float64, cfg Config) (float64, float64) {
	mc := chirpMass(cfg.M1, cfg.M2)
	f0 := 150.0
	tau := cfg.Duration * 0.8

	f := inspiralFrequency(t, f0, tau)
	if f < 10 || f > 2000 {
		return 0, 0
	}

	strain := h(t, f, mc, cfg.Distance)
	hp := strain * (1 + math.Cos(cfg.Inclination)) / 2
	hx := strain * math.Sin(cfg.Inclination)
	return hp * 1e21, hx * 1e21
}

func runAnimation(cfg Config) {
	rows, cols := 35, 90
	frameTime := 1.0 / float64(cfg.FPS)

	for frame := 0; frame < int(float64(cfg.FPS)*cfg.Duration); frame++ {
		t := float64(frame) * frameTime

		fmt.Printf("\033[2J\033[H")
		fmt.Printf("=== Gravitational Wave Detector Simulation ===\n")
		fmt.Printf("Binary: %.1f M☉ + %.1f M☉ | Distance: %.0f Mpc | Inclination: %.0f°\n\n",
			cfg.M1, cfg.M2, cfg.Distance, cfg.Inclination*180/math.Pi)

		centerY := rows / 2

		for y := 0; y < rows; y++ {
			line := make([]rune, cols)
			for x := 0; x < cols; x++ {
				tLocal := t + float64(x-cols/2)*0.02
				hp, hx := polarization(tLocal, cfg)

				offset := int(hp * float64(y-centerY) * 2)
				offsetX := int(hx * float64(x-cols/2) * 0.5)

				dist := y - centerY - offset
				distX := x - cols/2 - offsetX
				distFinal := int(math.Sqrt(float64(dist*dist + distX*distX)))

				if distFinal < 2 {
					line[x] = '█'
				} else if distFinal < 4 {
					line[x] = '▓'
				} else if distFinal < 6 {
					line[x] = '░'
				} else {
					line[x] = ' '
				}
			}
			fmt.Println(string(line))
		}

		phase := "inspiral"
		f0 := 150.0
		tau := cfg.Duration * 0.8
		f := inspiralFrequency(t, f0, tau)
		if t >= tau {
			phase = "merger"
			f = 0
		} else if f > 500 {
			phase = "ringdown"
		}

		fmt.Printf("\nPhase: %s | Frequency: %.0f Hz | Time: %.2fs / %.2fs\n",
			phase, f, t, cfg.Duration)

		progress := int(t / cfg.Duration * 40)
		fmt.Printf("[%s%s] %.1f%%\n",
			string(make([]byte, progress)),
			string(make([]byte, 40-progress)),
			t/cfg.Duration*100)
	}
}

func main() {
	m1 := flag.Float64("m1", 30, "Mass of first body (solar masses)")
	m2 := flag.Float64("m2", 30, "Mass of second body (solar masses)")
	distance := flag.Float64("d", 400, "Distance (Mpc)")
	inc := flag.Float64("i", 0, "Inclination angle (degrees)")
	duration := flag.Float64("t", 10, "Duration (seconds)")
	fps := flag.Int("fps", 10, "Frames per second")

	flag.Parse()

	cfg := Config{
		M1:          *m1,
		M2:          *m2,
		Distance:    *distance,
		Inclination: *inc * math.Pi / 180,
		Duration:    *duration,
		FPS:         *fps,
	}

	runAnimation(cfg)
}
