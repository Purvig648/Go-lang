package calculattor

func Square(n int) int {
	Res := n * n
	return Res
}
func Sum(a int, b int) int {
	Res := a + b
	return Res
}
func Diff(a int, b int) int {
	Res := a - b
	return Res
}
func Prod(a int, b int) int {
	Res := a * b
	return Res
}
func Div(a int, b int) (int, int) {
	Res1 := a / b
	Res2 := a % b
	return Res1, Res2
}