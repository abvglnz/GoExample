package gul

func Fac(n int) int {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f
}

func zerkal(n int) int {
	z := 0
	for n > 0 {
		z = z*10 + n%10
		n /= 10
	}
	return z
}

func Polindrom(n int) bool {
	return zerkal(n) == n
}
