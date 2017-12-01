package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Advent of Code 2017 in Go")
	args := os.Args[1:]
	if len(args) > 0 {
		fmt.Printf("Reverse Captcha for %s is %d\n", args[0], ReverseCaptcha(args[0]))
	} else {
		fmt.Println("Please provide an input for the Captcha")
	}
}

//ReverseCaptcha returns the sum of repeat integers in a circular sequence
func ReverseCaptcha(input string) int {
	sum := 0
	for p, c := range input {
		i, _ := strconv.Atoi(string(c))
		var ii int
		if p == len(input)-1 {
			ii, _ = strconv.Atoi(string(input[0]))
		} else {
			ii, _ = strconv.Atoi(string(input[p+1]))
		}
		if i == ii {
			sum += i
		}
	}
	return sum
}
