package geom

import (
	"fmt"
	"math"
)

// Point defines a location.
type Point struct {
	X, Y float64
}

// Align modifies this Point to align with integer coordinates.
func (p *Point) Align() {
	p.X = math.Floor(p.X)
	p.Y = math.Floor(p.Y)
}

// Add modifies this Point by adding the supplied coordinates.
func (p *Point) Add(pt Point) {
	p.X += pt.X
	p.Y += pt.Y
}

// Subtract modifies this Point by subtracting the supplied coordinates.
func (p *Point) Subtract(pt Point) {
	p.X -= pt.X
	p.Y -= pt.Y
}

// String implements the fmt.Stringer interface.
func (p Point) String() string {
	return fmt.Sprintf("%v, %v", p.X, p.Y)
}
