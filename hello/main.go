package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

func main() {
	example1()
	example2()
	example3()
	example4()
}

func example4() {
	err := validateLength("aaaaaaa")
	if nil != err {
		fmt.Println(err.Error())
	}
	err = validateAge(12)
	if nil != err {
		fmt.Println(err.Error())
	}
}

// validateLength return false when string length less than 8
// Please change return type to error with message "too short string"
func validateLength(s string) error {
	if len([]rune(s)) < 8 {
		return errors.New("too short string")
	}
	return nil
}

var ageError = errors.New("your age is negative!")

// in case of too young age please create a new type ErrTooYoung int` with method `Error() string`
// example error message: "17 is too young"
func validateAge(n int) error {
	if n < 0 {
		return ageError
	}
	if n < 18 {
		return ErrTooYoung(n)
	}
	return nil
}

type ErrTooYoung int

func (e ErrTooYoung) Error() string {
	return fmt.Sprintf("%d is too young", e)
}

func example3() {
	fmt.Println(VolumeOf(cube{
		edge: 10,
	}))

	fmt.Println(VolumeOf(triangularPrism{
		base:     5,
		attitude: 7,
		height:   10,
	}))
}

type volumer interface {
	Volume() float64
}

type cube struct {
	edge float64
} // edge x edge x edge

type triangularPrism struct {
	base     float64
	attitude float64
	height   float64
} // 0.5 x base x attitude x height

func (c cube) Volume() float64 {
	return c.edge * c.edge * c.edge
}

func (t triangularPrism) Volume() float64 {
	return 0.5 * t.base * t.attitude * t.height
}

func VolumeOf(v volumer) float64 {
	return v.Volume()
}

func example2() {
	fmt.Println(wordCount("ake ake boss boss"))

	book := Book{
		Name:   "Harry Potter",
		Author: "J. K. Rowling",
	}

	book.SetName("Eknarong")
	fmt.Println(book.String())

}

type Book struct {
	Name   string
	Author string
}

func (book Book) String() string {
	return book.Name + " by " + book.Author
}

func (book *Book) SetName(name string) {
	book.Name = name
}

func wordCount(s string) map[string]int {
	m := map[string]int{}
	temp := strings.Split(s, " ")
	for _, word := range temp {
		m[word] = m[word] + 1
	}
	return m
}

func example1() {
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
