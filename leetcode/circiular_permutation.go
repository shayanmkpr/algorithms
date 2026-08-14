package main

import "fmt"

func circularPermutation(n int, start int) []int {
	// starting from "start", end with 2^n-1, and they
	// should have a single bit difference.
	var result []int
	value := bin(start)
	number := power(2, n)
	result = append(result, dec(start))
	for i := 0; i < number-1; i++ {
		fmt.Println(value)
		value = binAddOne(value, n)
		result = append(result, dec(value))
	}
	return result
}

func dec(n int) int {
	var result int
	multiplier := 1
	for n > 0 {
		result += n % 10 * multiplier
		multiplier *= 2
		n = n / 10
	}
	return result
}

func bin(n int) int {
	if n == 0 {
		return 0
	}
	var result, place int
	place = 1
	for n > 0 {
		if n&1 == 1 {
			result += place
		}
		place *= 10
		n >>= 1
	}
	return result
}

func binAddOne(n, number int) int {
	var result, place, cntr int
	place = 1
	for n&1 == 1 {
		cntr++
		place *= 10
		n /= 10
	}
	if cntr > number-1 {
		return 0
	}
	result = n*place + place
	return result
}

func power(x, y int) int {
	i := 0
	multiplier := x
	for range y - 1 {
		i++
		x *= multiplier
	}
	return x
}
