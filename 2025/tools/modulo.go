package tools

func Mod(a, n int) int {
	return (a%n + n) % n
}
