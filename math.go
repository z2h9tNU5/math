// Package math implement something that is often used in mathematical logic.
package math

// Gcd returns the greatest common divisor.
func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

// MultipleGcd returns the greatest common divisor of two or more numbers
func MultipleGcd(val []int) int {
	z := val[0]
	for _, v := range val {
		z = Gcd(v, z)
	}

	return z
}
