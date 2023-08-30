package HumanMouse

import (
	"math/rand"
	"time"
)

// Trajectory represents the starting position and behavior of a simulated mouse movement.
type Trajectory struct {
	StartX        float64 // Initial horizontal position
	StartY        float64 // Initial vertical position
	IsStartRandom bool    // Flag to randomize starting position
	Bounds        Bounds  // Movement boundary
}

// Movements struct holds the simulated mouse movement data.
type Movements struct {
	X        float64
	Y        float64
	Duration int64
}

// Bounds struct defines the limits for the mouse movement.
type Bounds struct {
	MinX, MaxX, MinY, MaxY float64
}

// GenerateMouseMovements creates a series of mouse movements.
func (m *Trajectory) GenerateMouseMovements() (movements []Movements) {
	const (
		step      = 0.390625
		randiness = 16
	)

	if m.IsStartRandom {
		for !m.isOutOfBounds(m.StartX, m.StartY) || m.StartY == 0 && m.StartX == 0 {
			m.StartX, m.StartY = getRandomX(), getRandomY()
		}
	}
	x, y := m.StartX, m.StartY
	startTime := time.Now().UnixMilli()

	for m.isOutOfBounds(x, y) {
		m.updateCoordinates(&x, &y, step, randiness)
		randomSleep()
		movements = append(movements, Movements{X: x, Y: y, Duration: time.Now().UnixMilli() - startTime})
	}

	return movements
}

// isOutOfBounds checks if the current position is outside the defined bounds.
func (m *Trajectory) isOutOfBounds(x, y float64) bool {
	return x < m.Bounds.MinX+float64(rand.Intn(10)) ||
		x > m.Bounds.MaxX-float64(rand.Intn(10)) ||
		y < m.Bounds.MinY+float64(rand.Intn(10)) ||
		y > m.Bounds.MaxY-float64(rand.Intn(10))
}

// updateCoordinates modifies the coordinates based on predefined logic to simulate random mouse movement.
func (m *Trajectory) updateCoordinates(x, y *float64, step float64, randiness int) {
	n := rand.Intn(5)
	k := rand.Intn(2)

	switch {
	case k == 0 && *y > m.Bounds.MaxY:
		*y -= step * float64(rand.Intn(randiness))
	case k == 0 && *y < m.Bounds.MinY:
		*y += step * float64(rand.Intn(randiness))
	case k == 0 && rand.Intn(2) == 0 && n == 0:
		*y += step * float64(rand.Intn(randiness))
	case k == 0 && n == 0:
		*y -= step * float64(rand.Intn(randiness))
	case *x > m.Bounds.MaxX:
		*x -= step * float64(rand.Intn(randiness))
	case *x < m.Bounds.MinX:
		*x += step * float64(rand.Intn(randiness))
	case rand.Intn(2) == 0 && n == 0:
		*x += step * float64(rand.Intn(randiness))
	case n == 0:
		*x -= step * float64(rand.Intn(randiness))
	}
}

// randomSleep introduces a random delay to simulate realistic pauses between mouse movements.
func randomSleep() {
	t := rand.Intn(1254)
	switch {
	case t < 1:
		time.Sleep(8 * time.Millisecond)
	case t < 2:
		time.Sleep(7 * time.Millisecond)
	case t < 5:
		time.Sleep(6 * time.Millisecond)
	case t < 11:
		time.Sleep(5 * time.Millisecond)
	case t < 23:
		time.Sleep(4 * time.Millisecond)
	case t < 40:
		time.Sleep(3 * time.Millisecond)
	case t < 93:
		time.Sleep(2 * time.Millisecond)
	case t >= 187 && t < 1252:
		time.Sleep(1 * time.Millisecond)
	}
}

// getRandomX returns a random horizontal starting position.
func getRandomX() float64 {
	units := float64(rand.Intn(1700))
	decimals := float64(390625*(rand.Intn(24)-1)) / 100000000
	return units + decimals
}

// getRandomY returns a random vertical starting position.
func getRandomY() float64 {
	units := float64(rand.Intn(800))
	decimals := float64(390625*(rand.Intn(24)-1)) / 100000000
	return units + decimals
}
