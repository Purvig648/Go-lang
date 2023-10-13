package temperaturee

func Convert(temp float64) float64 {
	i := ((temp - 32) * 5 / 9)
	return i
}
