package main

import (
	"errors"
)

func IsPrime(n int) bool {
	// Corner cases
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}

	// This is checked so that we can skip
	// middle five numbers in below loop
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i = i + 6 {
		if n%i == 0 || n%(i+2) == 0 {
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
