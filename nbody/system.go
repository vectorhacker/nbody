package nbody

import (
	"math"
)

// System represents an NBody System
type System struct {
	Bodies []*Body
}

// NewSystem creates an NBody system
func NewSystem() *System {
	bodies := []*Body{
		Sun(),
		Jupiter(),
		Saturn(),
		Uranus(),
		Neptune(),
	}

	var px, py, pz float64

	for _, body := range bodies {
		px += body.Vx * body.Mass
		py += body.Vy * body.Mass
		pz += body.Vz * body.Mass
	}

	bodies[0].OffsetMomentum(px, py, pz)

	return &System{
		Bodies: bodies,
	}
}

// Advance advances the system by dt
func (s *System) Advance(dt float64) {
	for i, iBody := range s.Bodies {
		for _, jBody := range s.Bodies[i+1:] {
			dx := iBody.X - jBody.X
			dy := iBody.Y - jBody.Y
			dz := iBody.Z - jBody.Z

			dSquared := dx*dx + dy*dy + dz*dz
			distance := math.Sqrt(dSquared)
			mag := dt / (dSquared * distance)

			iBody.Vx -= dx * jBody.Mass * mag
			iBody.Vy -= dy * jBody.Mass * mag
			iBody.Vz -= dz * jBody.Mass * mag

			jBody.Vx += dx * iBody.Mass * mag
			jBody.Vy += dy * iBody.Mass * mag
			jBody.Vz += dz * iBody.Mass * mag
		}
	}

	for _, body := range s.Bodies {
		body.X += dt * body.Vx
		body.Y += dt * body.Vy
		body.Z += dt * body.Vz
	}
}

// Energy calculates the energy of the system
func (s *System) Energy() float64 {
	e := 0.0

	for i, iBody := range s.Bodies {

		e += 0.5 * iBody.Mass * (iBody.Vx*iBody.Vx + iBody.Vy*iBody.Vy + iBody.Vz*iBody.Vz)

		for _, jBody := range s.Bodies[i+1:] {
			dx := iBody.X - jBody.X
			dy := iBody.Y - jBody.Y
			dz := iBody.Z - jBody.Z

			distance := math.Sqrt(dx*dx + dy*dy + dz*dz)

			e -= (iBody.Mass * jBody.Mass) / distance
		}
	}

	return e
}
