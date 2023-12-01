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

func check_if_line_contains_int(line string) bool {
	//Use regex to check if string contains digits
	match, _ := regexp.MatchString("[0-9]+", line)
	return match
}
func get_first_char(word string) string {
	return string(word[0])
}

func get_last_char(word string) string {
	return string(word[len(word)-1])
}

func get_calibration_value(line string) int {
	//Find all digits in string
	digits := regexp.MustCompile("[0-9]+").FindAllString(line, -1)

	//Get the first and last character
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
		if check_if_line_contains_int(line) {
			var val int = get_calibration_value(line)
			calibration_values_sum += val
		}
	}

	fmt.Print(calibration_values_sum)
	input.Close()
}
