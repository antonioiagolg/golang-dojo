package main

import "fmt"
import "math"

// Here I will use the memory approach:
// I defined all possible magic squares in memory, as a lookup table,
// and I will perform operations with these squares with the input,
// based on the result, i will check what is the cost.
// I will subtract each of these matrixes with the input and check
// what is the minimu cost
// TODO: Generate the magic squares dynamic
var squares [][][]int32 = [][][]int32{
	{
		{4,9,2},
		{3,5,7},
		{8,1,6},
	},
	{
		{8,3,4},
		{1,5,9},
		{6,7,2},
	},
	{
		{6,1,8},
		{7,5,3},
		{2,9,4},
	},
	{
		{2,7,6},
		{9,5,1},
		{4,3,8},
	},
	{
		{8,1,6},
		{3,5,7},
		{4,9,2},
	},
	{
		{6,7,2},
		{1,5,9},
		{8,3,4},
	},
	{
		{2,9,4},
		{7,5,3},
		{6,1,8},
	},
	{
		{4,3,8},
		{9,5,1},
		{2,7,6},
	},
}

func formingMagicSquare(s [][]int32) int32 {
	var cost int32 = 1000 // Defining a high cost at the beginning because we want the minimun.
	var matrix [][] int32
	for _, matrix = range squares {
		actualCost := getCost(s, matrix)
		if actualCost < cost {
			cost = actualCost
		}
	}
	return cost
}

func getCost(s, m [][]int32) int32 {
	var sum int32
	for i := range s {
		for j := range s[i] {
			sum += int32(math.Abs(float64(s[i][j] - m[i][j])))
		}
	}
	return sum
}

func main() {
	var test [][][]int32 = [][][]int32{
		{
			{4,5,8},
			{2,4,1},
			{1,9,7},
		},
		{
			{4,8,2},
			{4,5,7},
			{6,1,6},
		},
		
	}
	var matrix [][]int32
	for _, matrix = range test {
		fmt.Println(formingMagicSquare(matrix))
	}
}
