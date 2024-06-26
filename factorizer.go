package main

import(
	"fmt"
	"math/big"
	"os"
)

type point struct {
	a big.Int
	b big.Int	
}
/**
	We have 4 important variables:
		1. N: The number we want to factorize
		2. p: The first factor of N
		3. q: The second factor of N
		4. step_width: How many steps we make from the current position
		5. (a,b): Current position
**/

var steps int
func main()  {

	steps = 0;
	if(len(os.Args) < 2){
		fmt.Println("You need to proide a number N to factorize as a command line argument!")
		return
	}
	string_arg := os.Args[1]

	num := big.NewInt(0)
	num.SetString( string_arg , 10)
	fmt.Println("Finding prime Factors for", num)
	find_factors(*num)

	fmt.Println("Took", steps, "steps")
}

func find_factors(N big.Int) (int){
	done := make(chan big.Int)
	
	//var upper_bound int64 = 1
	//primes := sieveOfEratosthenes(upper_bound)
	//primes = append(primes, 1)
	//primes := []big.Int{ *big.NewInt(upper_bound) }

	numbers := getRangeUpTo(*big.NewInt(1048575))
	
	for _,s:= range numbers {
		go execute_algorithm( s, done, N)
	}
	fmt.Println("Started", len(numbers), "Workers")
	
	lucky_worker := <- done
	fmt.Println("Worker with step size:", lucky_worker.String(), "found it")

	return 0;
}

func execute_algorithm(step_len big.Int, done chan big.Int, N big.Int) {

	guess_init := initial_guess(N)
	current_point := guess_init
	last_point := guess_init
	next_point := guess_init
	dist := big.NewInt(1)
	
	// means that one of the factors is greater than N i.e. we can stop searching with this step size
	overshoot := false

	for ; ( dist.Cmp( big.NewInt(0)) != 0 ) && !overshoot;  {
	
		next_point = make_step(current_point, last_point, step_len, N)
		last_point = current_point
		current_point = next_point
		steps = steps + 1
		res := calculate_distance(current_point, N)
		dist = &res
		overshoot = ( N.Cmp( &current_point.a ) <  0 ||  N.Cmp( &current_point.b ) <  0 )
	}
	
	if( overshoot ){
		fmt.Println("Worker", step_len.String(), "is out of bounds")
		return
	}

	fmt.Println("Found the factors!", current_point.a.String(), current_point.b.String())
	done <- step_len
}

func initial_guess(N big.Int) (point) {
	coord := big.NewInt(0)
	coord = coord.Sqrt( &N )  
	return point{ a: *coord, b: *coord}
}

func make_step(current_position point, last_position point, step_len big.Int, N big.Int) (next_position point)  {
	
	// Imagine standing in a 2D-plane where (a,b) is a point and abs( N - a*b ) is its value
	// The plane goes from top left to bottom right (like in a excel sheet) where (p, ) are the columns and (, q) the rows

	b_up, _   := big.NewInt(0).SetString(current_position.b.String(), 10)
	b_up 	   = b_up.Add(b_up, &step_len)
	a_down, _ := big.NewInt(0).SetString(current_position.a.String(), 10)
	a_down     = a_down.Sub(a_down, &step_len)
	// 1. Step to the bottom -> (a, b+1)
	step_bottom := point{ a: current_position.a, b: *b_up}

	// 2. Step to the left -> (a-1, b)
	step_left := point{ a: *a_down, b: current_position.b}

	
	dist_bottom := calculate_distance(step_bottom, N)
	dist_left := calculate_distance(step_left, N)

	distances := []big.Int{dist_bottom, dist_left}
	
	min_dist := Min(distances)

	if dist_bottom.Cmp(&min_dist) == 0 {
		return step_bottom
	} else  {
		return step_left
	}
}

func calculate_distance(point point, N big.Int) (big.Int){
	res := big.NewInt(0).Abs( big.NewInt(0).Sub( &N, big.NewInt(0).Mul( &point.a, &point.b )))
	return *res
}

func Min(array []big.Int) (big.Int) {
    var min big.Int = array[0]
    for _, value := range array {
        if min.Cmp(&value) > 0 {
            min = value
        }
    }
    return min
}


func getRangeUpTo(X big.Int) ([]big.Int){

	var i = X
	var numbers []big.Int
	for ; i.Cmp( big.NewInt(0)) > 0;  {
		

		
		numbers = append(numbers, i)
		var j = *big.NewInt(0).Sub(&i, big.NewInt(1))

		i = j
    }

	return numbers
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


