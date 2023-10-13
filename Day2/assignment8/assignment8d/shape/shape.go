package shape

import (
	"assignments/models"
	"math"
)

func Area(u models.Circle) float64 {
	area := math.Pi * u.Radius * u.Radius
	return area
}
