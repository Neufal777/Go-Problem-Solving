package main

import (
	"fmt"
	"strings"
)

//stoPrint Given a slice with NUMS (int) , stops printing when STOP (int) is found in the slice
func stoPrint(nums []int, stop int) {

	for _, num := range nums {
		if num == stop {
			break
		} else {
			fmt.Println(num)
		}
	}

}

//birthdayCakeCandles
func birthdayCakeCandles(ar []int32) int32 {
	//	TESTING : birthdayCakeCandles([]int32{2, 5, 87, 23, 87, 87, 65, 87})

	var largest int32
	var rep int32

	largest = 0
	rep = 0

	for _, a := range ar {
		if a >= largest {
			largest = a
		}
	}

	for _, arr := range ar {
		if largest == arr {
			rep = rep + 1
		}
	}

	return rep
}

//TextSum summ strig characters
func TextSum(input map[string]int, TextInput1 string) int {

	var sum int

	TextInput2 := strings.Split(TextInput1, "-")
	for _, char := range TextInput2 {

		currentNum := input[char]
		sum = sum + currentNum

	}

	return sum

	//TESTING :  fmt.Println(TextSum(map[string]int{"A": 5, "B": 2, "C": 8}, "A-B-C"))
}

//diagonalDiference
func diagonalDiference(arr [][]int32) int32 {

	/*

		ar := [][]int32{}

		row1 := []int32{1, 2, 3}
		row2 := []int32{4, 5, 6}
		row3 := []int32{9, 8, 9}

		ar = append(ar, row1)
		ar = append(ar, row2)
		ar = append(ar, row3)

		fmt.Println(diagonalDiference(ar))
	*/

	leftRight := arr[0][0] + arr[1][1] + arr[2][2]
	rightLeft := arr[0][2] + arr[1][1] + arr[2][0]

	var diff int32

	if rightLeft >= leftRight {

		diff = rightLeft - leftRight
	} else {
		diff = leftRight - rightLeft
	}

	return diff
}

type car struct {
	gasPedal      uint16
	breakPedal    uint16
	steeringWheel int16
	topSpeedKmh   float64
	fuelType      string
	make          string
	model         string
	year          int32
}
