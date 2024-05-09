package main

import (
	"fmt"
	"math/rand"
	"time"
)

func partition(array []int, left int, right int, random_pivot bool) int {
	var pivot int

	if random_pivot {
		random_i := left + (rand.Int() % (right - left + 1))
		pivot = array[random_i]
	} else {
		pivot = array[left]
	}

	i, j := left-1, right+1

	for i < j {
		i++
		for array[i] < pivot {
			i++
		}

		j--
		for array[j] > pivot {
			j--
		}

		if i < j {
			array[i], array[j] = array[j], array[i]
		}
	}

	return j
}

func k_largest_integers(array []int, k int, random_pivot bool) []int {
	l := len(array) - k

	left, right := 0, len(array)-1

	p := partition(array, left, right, random_pivot)
	// Elements in array[0:p] will be smaller than array[p+1 : N].

	for p+1 != l {
		if p+1 < l {
			left = p + 1
		} else { // p+1 > l
			right = p
		}
		p = partition(array, left, right, random_pivot)
	}

	return array[p+1:]
}

func main() {
	// array := []int{45, 76, 98, 100, 22, 31, 17, 86, 59}

	size := 100_000_000

	// Generate numbers.
	array := make([]int, size)
	start := time.Now()
	for i := range array {
		array[i] = rand.Int()
	}
	duration := time.Since(start)
	fmt.Printf("Generating %v numbers took %v seconds\n", size, duration.Seconds())

	// Duplicate the array.
	array2 := make([]int, size)
	copy(array2, array)

	// Time quick select with pivot strat 1: first element is pivot.
	start1 := time.Now()
	k_largest := k_largest_integers(array, 3, false)
	duration1 := time.Since(start1)

	// Time quick select with pivot strat 2: random pivot.
	start2 := time.Now()
	k_largest2 := k_largest_integers(array2, 3, true)
	duration2 := time.Since(start2)

	fmt.Printf("First element is pivot:\n\tAnswer: %v, time taken: %v seconds\n", k_largest, duration1.Seconds())
	fmt.Printf("Random pivot:\n\tAnswer: %v, time taken: %v seconds\n", k_largest2, duration2.Seconds())
}
