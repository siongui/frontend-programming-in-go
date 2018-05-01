package main

import (
	"strconv"

	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-vue"
)

type Model struct {
	*js.Object        // this is needed for bidirectional data bindings
	Intn       string `js:"intn"`
	Result     string `js:"result"`
}

func SieveOfEratosthenes(n int) []int {
	// Create a boolean array "prime[0..n]" and initialize
	// all entries it as true. A value in prime[i] will
	// finally be false if i is Not a prime, else true.
	integers := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		integers[i] = true
	}

	for p := 2; p*p <= n; p++ {
		// If integers[p] is not changed, then it is a prime
		if integers[p] == true {
			// Update all multiples of p
			for i := p * 2; i <= n; i += p {
				integers[i] = false
			}
		}
	}

	// return all prime numbers <= n
	var primes []int
	for p := 2; p <= n; p++ {
		if integers[p] == true {
			primes = append(primes, p)
		}
	}

	return primes
}

func main() {
	m := &Model{
		Object: js.Global.Get("Object").New(),
	}
	// field assignment is required in this way to make data passing works
	m.Intn = "4"
	m.Result = ""

	// create the VueJS viewModel using a struct pointer
	app := vue.New("#vueapp", m)

	opt := js.Global.Get("Object").New()
	opt.Set("immediate", true)
	app.Call("$watch", "intn", func(newVal, oldVal string) {
		n, err := strconv.Atoi(newVal)
		if err != nil {
			m.Result = "input must be integer"
			return
		}
		m.Result = ""

		primes := SieveOfEratosthenes(n)
		for _, prime := range primes {
			m.Result += (strconv.Itoa(prime) + " ")
		}
	}, opt)
}
