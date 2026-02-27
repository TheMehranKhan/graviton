# Graviton

Gravitational wave pattern generator in Go.

## Physics

### Gravitational Wave Strain

The gravitational wave strain amplitude from a binary inspiral:

$$h(t) = \frac{4\pi^2 G}{c^4} \mathcal{M}_c f^{5/3} \cos(2\pi f t - \phi)$$

Where:
- $G$ = Gravitational constant ($6.67430 \times 10^{-11}$ m³ kg⁻¹ s⁻²)
- $c$ = Speed of light ($2.9979 \times 10^8$ m/s)
- $\mathcal{M}_c$ = Chirp mass
- $f$ = Gravitational wave frequency
- $\phi$ = Phase at coalescence

### Chirp Mass

$$\mathcal{M}_c = \frac{(m_1 m_2)^{3/5}}{(m_1 + m_2)^{1/5}}$$

### Polarization Modes

Gravitational waves have two independent polarization states:

$$h_+ = h(t) \cdot \frac{1 + \cos\iota}{2}$$
$$h_\times = h(t) \cdot \sin\iota$$

Where $\iota$ is the inclination angle of the binary plane.

## Usage

```bash
go run main.go
```

## Output

ASCII visualization of gravitational wave strain over time, simulating LIGO detection of a binary black hole merger.

## References

- [LIGO Scientific Collaboration](https://www.ligo.caltech.edu/)
- Einstein Field Equations: $G_{\mu\nu} = \frac{8\pi G}{c^4} T_{\mu\nu}$
- Quadrupole approximation for GW emission

## License

MIT
