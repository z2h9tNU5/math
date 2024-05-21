// Package math implement something that is often used in mathematical logic.
package math

import "fmt"

var (
	ErrChineseRemainderTheorem = fmt.Errorf("first and second arguments must be disjoint")
)

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

// ChineseRemainderTheorem Chinese remainder theorem.
func ChineseRemainderTheorem(m1, m2, b1, b2 int) (int, error) {
	if Gcd(m1, m2) != 1 {
		return 0, ErrChineseRemainderTheorem
	}

	var ret int
	for i := 0; i < m1*m2; i++ {
		z := m1*i + b1
		if z%m2 == b2 {
			ret = z
			break
		}
	}

	return ret, nil
}
