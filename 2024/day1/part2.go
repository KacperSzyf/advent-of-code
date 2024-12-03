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

	//count occurances in both lists
	leftOccurance := countOccurances(leftCol)
	rightOccurance := countOccurances(rightCol)

	sum := calculateSum(leftOccurance, rightOccurance)

	//calculate the sum
	fmt.Println(sum)

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

func countOccurances(col []int) map[int]int {
	occurances := make(map[int]int)

	for _, v := range col {
		occurances[v] = occurances[v] + 1
	}

	return occurances
}

func calculateSum(left map[int]int, right map[int]int) int {

	var sum int = 0

	//calculate te sum
	for i, v := range left {
		sum += v * (i * right[i])
	}

	return sum
}
