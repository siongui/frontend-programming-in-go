package main

import (
	"errors"
)

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Lemoine(n int) (map[int]int, error) {
	if n <= 5 || n%2 == 0 {
		return nil, errors.New("n must be greater than 5 and must be odd")
	}

	m := make(map[int]int)

	for q := 1; q <= n/2; q++ {
		p := n - 2*q

		if IsPrime(p) && IsPrime(q) {
			m[p] = q
		}
	}
	return m, nil
}
