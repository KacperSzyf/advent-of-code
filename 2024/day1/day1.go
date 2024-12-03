package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	//read input file
	file, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	//convert bytes to string
	var content string = string(file)

	//remove all spaces and create an array where data is stored: [left col, right col, left col, ... right col]
	ids := strings.Fields(content)

	//separate the id's into left and right cols and sort into ascending order
	leftCol, rightCol := splitColumns(ids)

	fmt.Println(leftCol, rightCol)

	//calculate sum
	sum := calculateDistance(leftCol, rightCol)

	fmt.Println("The difference is: ", sum)

}

func splitColumns(input []string) ([]int, []int) {
	//find left and right column size
	var colSize int = len(input) / 2

	//initialise left and right columns
	leftCol := make([]int, colSize+1)
	rightCol := make([]int, colSize+1)

	//grab every other value into left and right columns
	for i := 1; i < len(input); i += 2 {
		leftCol[i/2], _ = strconv.Atoi(input[i-1])
		rightCol[i/2], _ = strconv.Atoi(input[i])

	}

	//sort both the columns into ascending orders
	slices.Sort(leftCol)
	slices.Sort(rightCol)

	return leftCol, rightCol
}

func calculateDistance(left []int, right []int) int {
	var sum int = 0

	//calculate the sum of differences
	if len(left) == len((right)) {
		for i := 0; i < len(left); i++ {
			current_sum := left[i] - right[i]
			//get absolute value
			if current_sum < 0 {
				current_sum = current_sum * -1
			}
			sum += current_sum
		}
	}

	return sum
}
