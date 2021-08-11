package mathOps

func Fact(f int) int {

	if f == 1 {
		return f
	}
	return f * Fact(f-1)

}

func SumN(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}
