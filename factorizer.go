package main

import(
	"fmt"
	"math"
	"os"
	"strconv"
)

type point struct {
	a int
	b int	
}
/**
	We have 4 important variables:
		1. N: The number we want to factorize
		2. p: The first factor of N
		3. q: The second factor of N
		4. step_width: How many steps we make from the current position
		5. (a,b): Current position
**/


func main()  {
	if(len(os.Args) < 2){
		fmt.Println("You need to proide a number N to factorize as a command line argument!")
		return
	}
	arg := os.Args[1]

	num, _ := strconv.Atoi(arg)
	
	fmt.Println("Finding prime Factors for", num)
	find_factors(num)
}

func find_factors(N int) (int){
	done := make(chan int)
	
	upper_bound := 20
	primes := sieveOfEratosthenes(upper_bound)
	primes = append(primes, 1)


	for _,s:= range primes {
		go execute_algorithm(s, done, N)
	}
	fmt.Println("Started", len(primes), "Workers")
	
	lucky_worker := <- done
	fmt.Println("Worker with step size:", lucky_worker, "found it")

	return 0;
}

func execute_algorithm(step_len int, done chan int, N int) {

	guess_init := initial_guess(N)
	current_point := guess_init
	last_point := guess_init
	next_point := guess_init

	for ; calculate_distance(current_point, N) != 0;  {
		
		next_point = make_step(current_point, last_point, step_len, N)
		last_point = current_point
		current_point = next_point

	}
	
	fmt.Println("Found the factors!", current_point.a, current_point.b)
	done <- step_len
}

func initial_guess(N int) (point) {
	coord := int( math.Round( math.Sqrt( float64(N) ) ) )
	return point{ a:coord, b:coord}
}

func make_step(current_position point, last_position point, step_len int, N int) (next_position point)  {
	
	// Imagine standing in a 2D-plane where (a,b) is a point and abs( N - a*b ) is its value
	// The plane goes from top left to bottom right (like in a excel sheet) where (p, ) are the columns and (, q) the rows

	// 1. Step to the bottom -> (a, b+1)
	step_bottom := point{ a: current_position.a, b: current_position.b + step_len}

	// 2. Step to the left -> (a-1, b)
	step_left := point{ a: current_position.a - step_len, b: current_position.b}


	dist_bottom := calculate_distance(step_bottom, N)
	dist_left := calculate_distance(step_left, N)

	distances := []int{dist_bottom, dist_left}
	
	min_dist := Min(distances)


	if dist_bottom == min_dist {
		return step_bottom
	} else  {
		return step_left
	}
}

func calculate_distance(point point, N int) (int){
	return int(math.Abs(float64(N - point.a * point.b )))
}

func Min(array []int) (int) {
    var min int = array[0]
    for _, value := range array {
        if min > value {
            min = value
        }
    }
    return min
}


// return list of primes less than N
func sieveOfEratosthenes(N int) (primes []int) {
    b := make([]bool, N)
    for i := 2; i < N; i++ {
        if b[i] == true { continue }
        primes = append(primes, i)
        for k := i * i; k < N; k += i {
            b[k] = true
        }
    }
    return
}


