package main

import(
	"fmt"
	"math"
)

/**
	We have 4 important variables:
		1. N: The number we want to factorize
		2. p: The first factor of N
		3. q: The second factor of N
		4. step_width: How many steps we make from the current position
		5. (a,b): Current position
**/

var N int
var p int
var q int

func main()  {
	p := 31
	q := 7
	N := p * q

	a, b := initial_guess(N)
	fmt.Println("Hello World!", a, b)
}

func initial_guess(N int) (a_init int, b_init int) {
	guess := int( math.Round( math.Sqrt( float64(N) ) ) )
	return guess, guess
}