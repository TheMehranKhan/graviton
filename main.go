package main

import (
	"fmt"
	"math"
)

// Gravitational wave strain amplitude
func h(t, f, D float64) float64 {
	// h ≈ (4π^2 G / c^4) * (M_chirp * f)^(5/3) * cos(2π f t)
	G := 6.67430e-11  // Gravitational constant (m³ kg⁻¹ s⁻²)
	c := 2.99792458e8 // Speed of light (m/s)
	
	const M_chirp = 30.0      // Chirp mass in solar masses
	M_solar := 1.989e30      // Solar mass in kg
	M := M_chirp * M_solar    // Mass in kg
	
	prefactor := 4 * math.Pow(math.Pi, 2) * G / math.Pow(c, 4)
	amplitude := prefactor * M * math.Pow(f, 5.0/3.0)
	
	return amplitude * math.Cos(2*math.Pi*f*t) / D
}

// Polarization modes (plus and cross)
func polarizations(t, f, D, phi float64) (hp, hx float64) {
	hp = h(t, f, D) * (1 + math.Cos(phi)) / 2
	hx = h(t, f, D) * math.Sin(phi)
	return
}

// Strain at detector from binary inspiral
func detectorStrain(t float64) float64 {
	// LIGO parameters
	f0 := 150.0  // Initial frequency (Hz)
	_ = f0
	D := 400e6  // Distance in parsecs
	
	// Frequency evolution during inspiral
	f := f0 * math.Pow(1-t/10.0, -3.0/8.0)
	if f < 10 || f > 500 {
		return 0
	}
	
	return h(t, f, D) * 1e21 // Scale for visualization
}

func main() {
	fmt.Println("Gravitational Wave Pattern Generator")
	fmt.Println("=====================================")
	
	// ASCII visualization
	rows, cols := 30, 80
	centerY := rows / 2
	
	for y := 0; y < rows; y++ {
		line := make([]byte, cols)
		for x := 0; x < cols; x++ {
			t := float64(x) * 0.1
			strain := detectorStrain(t)
			
			offset := int(strain * float64(y-centerY))
			dist := y - centerY - offset
			
			if abs(dist) < 2 {
				line[x] = '*'
			} else if abs(dist) < 4 {
				line[x] = '.'
			} else {
				line[x] = ' '
			}
		}
		fmt.Println(string(line))
	}
	
	fmt.Println("\nWave Parameters:")
	fmt.Printf("Frequency range: 10-500 Hz\n")
	fmt.Printf("Strain amplitude (scaled): 10⁻²¹\n")
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
