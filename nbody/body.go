package nbody

import "math"

const (
	solarMass   = 4 * math.Pi * math.Pi
	daysPerYear = 365.24
)

// Body represents a body orbiting annother body
type Body struct {
	X    float64
	Y    float64
	Z    float64
	Vx   float64
	Vy   float64
	Vz   float64
	Mass float64
}

// Sun creates the sun
func Sun() *Body {
	return &Body{
		Mass: solarMass,
	}
}

// Jupiter creates the Jupiter body
func Jupiter() *Body {
	return &Body{
		X:    4.84143144246472090e+00,
		Y:    -1.16032004402742839e+00,
		Z:    -1.03622044471123109e-01,
		Vx:   1.66007664274403694e-03 * daysPerYear,
		Vy:   7.69901118419740425e-03 * daysPerYear,
		Vz:   -6.90460016972063023e-05 * daysPerYear,
		Mass: 9.54791938424326609e-04 * solarMass,
	}
}

// Saturn creates the Saturnian body
func Saturn() *Body {
	return &Body{
		X:    8.34336671824457987e+00,
		Y:    4.12479856412430479e+00,
		Z:    -4.03523417114321381e-01,
		Vx:   -2.76742510726862411e-03 * daysPerYear,
		Vy:   4.99852801234917238e-03 * daysPerYear,
		Vz:   2.30417297573763929e-05 * daysPerYear,
		Mass: 2.85885980666130812e-04 * solarMass,
	}
}

// Uranus creates the Uranian body
func Uranus() *Body {
	return &Body{
		X:    1.28943695621391310e+01,
		Y:    -1.51111514016986312e+01,
		Z:    -2.23307578892655734e-01,
		Vx:   2.96460137564761618e-03 * daysPerYear,
		Vy:   2.37847173959480950e-03 * daysPerYear,
		Vz:   -2.96589568540237556e-05 * daysPerYear,
		Mass: 4.36624404335156298e-05 * solarMass,
	}
}

// Neptune creates the Neptunian body
func Neptune() *Body {
	return &Body{
		X:    1.53796971148509165e+01,
		Y:    -2.59193146099879641e+01,
		Z:    1.79258772950371181e-01,
		Vx:   2.68067772490389322e-03 * daysPerYear,
		Vy:   1.62824170038242295e-03 * daysPerYear,
		Vz:   -9.51592254519715870e-05 * daysPerYear,
		Mass: 5.15138902046611451e-05 * solarMass,
	}
}

// OffsetMomentum offsets the momentum of the body
func (b *Body) OffsetMomentum(px, py, pz float64) {
	b.Vx = -px / solarMass
	b.Vy = -py / solarMass
	b.Vz = -pz / solarMass
}
