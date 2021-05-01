package main

import(
	"fmt"
	"math"
	"reflect"
	
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

var N int


func main()  {
	
	p := 31
	q := 7
	N = p*q	

	guess_init := initial_guess(N)
	current_point := guess_init
	last_point := guess_init
	next_point := guess_init

	for ; calculate_distance(current_point) != 0;  {
		
		next_point = make_step(current_point, last_point)
		last_point = current_point
		current_point = next_point

	}
	
	fmt.Println("Found the factors!", current_point.a, current_point.b)
}

func initial_guess(N int) (point) {
	coord := int( math.Round( math.Sqrt( float64(N) ) ) )
	return point{ a:coord, b:coord}
}

func make_step(current_position point, last_position point) (next_position point)  {
	
	// Imagine standing in a 2D-plane where (a,b) is a point and abs( N - a*b ) is its value
	// The plane goes from top left to bottom right (like in a excel sheet) where (p, ) are the columns and (, q) the rows

	// 1. Step to the top -> (a, b-1)
	step_top := point{ a: current_position.a, b: current_position.b -1}

	// 2. Step to the bottom -> (a, b+1)
	step_bottom := point{ a: current_position.a, b: current_position.b +1}

	// 3. Step to the left -> (a-1, b)
	step_left := point{ a: current_position.a -1, b: current_position.b}

	// 4. Step to the right -> (a+1, b)
	step_right := point{ a: current_position.a +1, b: current_position.b}

	dist_top := calculate_distance(step_top)
	dist_bottom := calculate_distance(step_bottom)
	dist_left := calculate_distance(step_left)
	dist_right := calculate_distance(step_right)

	// Now make a step into the direction IF 
	// 1. The step does not make us land on the previous position
	// 2. From the non-previous positions, go to the one with the smallest distance

	distances := []int{dist_top, dist_bottom, dist_left, dist_right}
	
	if reflect.DeepEqual(last_position, step_top) {
		distances = remove(distances, 0)
		// fmt.Println("Removing dist_top..")
	} else if reflect.DeepEqual(last_position, step_bottom) {
		distances = remove(distances, 1)
		// fmt.Println("Removing dist_bottom..")
	} else if reflect.DeepEqual(last_position, step_left) {
		distances = remove(distances, 2)
		// fmt.Println("Removing dist_left..")
	} else if reflect.DeepEqual(last_position, step_right) {
		distances = remove(distances, 3)
		//fmt.Println("Removing dist_right..")
	} else {
		// fmt.Println("ERR: Removing nothing")
	}
	
	min_dist := Min(distances)

	//fmt.Println("dist_top: ", dist_top)
	//fmt.Println("dist_bottom: ", dist_bottom)
	//fmt.Println("dist_left: ", dist_left)
	//fmt.Println("dist_right: ", dist_right)
	//fmt.Println("min_dist: ", min_dist)
	

	if dist_top == min_dist && !reflect.DeepEqual(last_position, step_top){
		return step_top
	} else if dist_bottom == min_dist && !reflect.DeepEqual(last_position, step_bottom) {
		return step_bottom
	} else if dist_left == min_dist && !reflect.DeepEqual(last_position, step_left) {
		return step_left
	} else  {
		return step_right
	}
}

func calculate_distance(point point) (dist int){
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

func remove(s []int, i int) []int {
    s[i] = s[len(s)-1]
    // We do not need to put s[i] at the end, as it will be discarded anyway
    return s[:len(s)-1]
}