package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Error helper
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// RegEx expressions
const numbers_expression string = "[0-9]"
const words_expression string = "(one|two|three|four|five|six|seven|eight|nine)"

func check_if_line_contains_number(line string) bool {
	//Use regex to check if string contains digits
	match, _ := regexp.MatchString(numbers_expression+"|"+words_expression, line)
	return match
}
func get_first_char(text string) string {
	return string(text[0])
}

func get_last_char(text string) string {
	return string(text[len(text)-1])
}

func text_to_digit(text string) string {

	switch text {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	}
	return text
}

func get_calibration_value(line string) int {
	//Find all numbers in a line
	digits := regexp.MustCompile(numbers_expression+"|"+words_expression).FindAllString(line, -1)

	for index, elemet := range digits {
		digits[index] = text_to_digit(elemet)
	}

	//Get the first and last character
	first_char := get_first_char(text_to_digit(digits[0]))
	last_char := get_last_char(text_to_digit(digits[len(digits)-1]))
	fmt.Print(first_char + last_char + "\n")
	// var calibration_value string = get_first_char(text_to_digit(digits[0])) + get_last_char(text_to_digit(digits[len(digits)-1]))
	var calibration_value string = get_first_char(digits[0]) + get_last_char(digits[len(digits)-1])

	//Convert string to int
	var calibration_value_int, err = strconv.Atoi(calibration_value)
	check(err)
	return calibration_value_int
}

func main() {
	var calibration_values_sum int = 0

	//Read file and check for errors
	input, err := os.Open("./input.txt")
	check(err)

	//Split file line by line
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	//Iterate over all values in file
	for scanner.Scan() {
		var line string = scanner.Text()

		//If line contains digits process it
		if check_if_line_contains_number(line) {
			var val int = get_calibration_value(line)
			calibration_values_sum += val
		}
	}

	fmt.Print(calibration_values_sum)
	input.Close()
}
