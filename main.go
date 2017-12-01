package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Advent of Code 2017 in Go")
	args := os.Args[1:]
	if len(args) > 1 {
		day, _ := strconv.Atoi(args[0])
		puzzle, _ := strconv.Atoi(args[1])
		star := day*2 + puzzle - 2
		switch star {
		case 1:
			fmt.Printf("Reverse Captcha for %s is %d\n", args[2], ReverseCaptcha(args[2]))
		case 2:
			fmt.Printf("Reverse Captcha for %s is %d\n", args[2], ReverseCaptchaHalfway(args[2]))
		default:
			fmt.Println("Not done yet.")
		}
	} else {
		fmt.Println("Please provide the correct input (day, 1/2, args)")
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

//ReverseCaptcha returns the sum of halfway repeat integers in a circular sequence
func ReverseCaptchaHalfway(input string) int {
	sum := 0
	for p, c := range input {
		i, _ := strconv.Atoi(string(c))
		var ii int
		p2 := p + len(input)/2
		if p >= (len(input) / 2) {
			p2 = p2 - len(input)
		}
		ii, _ = strconv.Atoi(string(input[p2]))
		if i == ii {
			sum += i
		}
	}
	return sum
}
