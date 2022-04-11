package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, ake!")
	fmt.Println(greeting("Pallat"))
	fmt.Println(greetingWithAge("Pallat", 30))
	fmt.Println(greetingWithAgeAndDrink("Pallat", 30, "Cola"))
	fmt.Println(isOdd(2))
	fmt.Println(isOdd(5))
	fmt.Println(sumOfFirst(3))
	fmt.Println(isPrime(1))
	fmt.Println(isPrime(10))
	fmt.Println(isPrime(13))
}

// greeting("Pallat") should return "Hello, Pallat"
func greeting(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

// greetingWithAge("Pallat", 30) should return "Hello, Pallat. You are 30 years old."
func greetingWithAge(name string, age uint) string {
	return fmt.Sprintf("%s. You are %d years old.", greeting(name), age)
}

// greetingWithAgeAndDrink("Pallat", 30, "Cola") should return "Hello, Pallat. You are 30 years old and your favorite drink is Cola."
func greetingWithAgeAndDrink(name string, age uint, drink string) string {
	return fmt.Sprintf("%s. You are %d years old and your favorite drink is %s", greeting(name), age, drink)
}

func isOdd(n int) bool {
	if value := math.Mod(float64(n), 2); value == 0 {
		return false
	} else {
		return true
	}
}

func sumOfFirst(n int) int {
	var sum int = 0
	for i := n; i > 0; i-- {
		sum += i
	}
	return sum
}

// A prime number is a number greater than 1 with only two factors – themselves and 1
func isPrime(n int) bool {
	sqrtn := math.Sqrt(float64(n))
	if n < 2 {
		return false // o and 1 are not primes
	}
	if n < 4 {
		return true // 2 and 3 are the first primes
	}
	if n%2 == 0 {
		return false // 2 is the only even prime
	}
	for d := 3; float64(d) <= sqrtn; d += 2 {
		if n%d == 0 {
			return false // n has a nontrival divisor
		}
	}
	return true
}
