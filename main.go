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

	//TESTING :  fmt.Println(TextSum(map[string]int{"A": 5, "B": 2, "C": 8}, "A-B-C"))
	var sum int

	TextInput2 := strings.Split(TextInput1, "-")
	for _, char := range TextInput2 {

		currentNum := input[char]
		sum = sum + currentNum

	}

	return sum
}

//diagonalDiference
func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var suma1 int32
	var suma2 int32

	n := len(arr)
	for i := 0; i < n; i++ {

		for j := 0; j < len(arr); j++ {

			if i == j {
				suma1 += arr[i][j]
			}

			if i == n-j-1 {
				suma2 += arr[i][j]
			}
		}
	}

	if suma1 < suma2 {

		return suma2 - suma1
	} else {
		return suma1 - suma2
	}

}

func breakingRecords(scores []int32) []int32 {

	max := scores[0]
	min := scores[0]
	var countMax int32
	var countMin int32

	for i := 0; i < len(scores); i++ {

		if max < scores[i] {

			max = scores[i]
			countMax++
		}

		if min > scores[i] {

			min = scores[i]
			countMin++
		}
	}

	result := []int32{countMax, countMin}

	return result
}

func sockMerchant(n int32, ar []int32) int32 {

	list := map[int32]int32{}
	var totalPairs int32

	for _, ar := range ar {

		list[ar] += 1
	}

	for _, v := range list {

		if v != 0 {
			totalPairs = totalPairs + (v / 2)
			//log.Print32ln(totalPairs)
		}

	}

	return totalPairs
}

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {

	var jumpx int32
	var jumpy int32

	for {

		x1 += v1
		x2 += v2

		jumpx++
		jumpy++

		if x2 > x1 && v2 > v1 {

			break
		} else {

			if x1 == x2 && jumpx == jumpy {

				fmt.Println(x1)
				fmt.Println(x2)
				return "YES"

			} else {
				if x2 > x1 && v2 >= v1 {

					break
				}

				continue
			}
		}
	}

	return "NO"
}

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {

	var applesInside int32
	var orangeInside int32

	distanceApples := []int32{}
	distanceOranges := []int32{}

	for _, apple := range apples {

		distanceApples = append(distanceApples, a+apple)
	}

	for _, orange := range oranges {

		distanceOranges = append(distanceOranges, b+orange)
	}

	//check if the fruits are inside or outside

	//apples

	for _, distance := range distanceApples {

		if distance >= s && distance <= t {

			applesInside++
		}
	}

	//oranges
	for _, distanceOrg := range distanceOranges {

		if distanceOrg >= s && distanceOrg <= t {

			orangeInside++
		}
	}

	//return the result

	fmt.Println(applesInside)
	fmt.Println(orangeInside)
}
