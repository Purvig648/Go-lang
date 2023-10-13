package shape

import (
	"assignments/assignment8d/models"
	"math"
)

func Area(u models.Circle) float64 {
	area := math.Pi * u.Radius * u.Radius
	return area
}
