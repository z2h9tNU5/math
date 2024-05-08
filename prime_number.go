package math

import (
	"fmt"
	"math/big"
)

var (
	ErrGoldbachConjecture = fmt.Errorf("must be an even number greater than 3")
)

// SieveEratosthenes returns an array indicating whether the number is prime or not,
// less than or equal to the value specified by the argument.
// If true, it is not a prime number; if false, it is a prime number.
// Sieve of Eratosthenes.
func SieveEratosthenes(n int) []bool {
	if n < 0 {
		return nil
	}
	if n == 0 {
		return []bool{true}
	}
	if n == 1 {
		return []bool{true, true}
	}

	res := make([]bool, n+1)
	res[0] = true
	res[1] = true
	for i := 2; i <= n; i++ {
		if res[i] {
			continue
		}

		if i*i > n {
			break
		}

		for j := 2; j <= n; j++ {
			x := i * j
			if x > n {
				break
			}

			res[x] = true
		}
	}

	return res
}

// IsPrime returns whether the number is prime.
func IsPrime(n int64) bool {
	var res big.Int
	res.SetInt64(n)
	return res.ProbablyPrime(0)
}

// GoldbachConjecture returns two prime numbers.
// The sum of its prime numbers is the same as the argument.
// If there are multiple patterns, the one with the largest product is returned.
// This does not prove the Goldbach conjecture.
func GoldbachConjecture(n int) (int, int, error) {
	if n <= 3 || n%2 != 0 {
		return 0, 0, ErrGoldbachConjecture
	}

	se := SieveEratosthenes(n)

	m := map[int]int{}
	for i, v := range se {
		if !v {
			if i == n {
				return 0, i, nil
			}

			num := n - i
			if !se[num] {
				m[i] = num
			}
		}
	}

	var res, maxVal int

	for key, val := range m {
		if key*val > maxVal {
			maxVal = key * val
			res = key
		}
	}

	if res > m[res] {
		return m[res], res, nil
	} else {
		return res, m[res], nil
	}
}
