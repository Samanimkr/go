package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) []int {
	return append(slice, slice...)
}

func mapSlice(f func(a int) int, slice []int) {
	i := 0

	for _, s := range slice {
		slice[i] = f(s)
		i++
	}
}


func mapArray(f func(a int) int, array [5]int) [5]int {
	var newArray [5]int

	for i := 0; i < len(array); i++ {
		newArray[i] = f(array[i])
	}
	
	return newArray
}

func main() {
	intsSlice := make([]int, 5)
	intsSlice[0] = 1
	intsSlice[1] = 2
	intsSlice[2] = 3
	intsSlice[3] = 4
	intsSlice[4] = 5

	mapSlice(addOne, intsSlice)
	fmt.Println(intsSlice)

	intsArray := [5]int{1, 2, 3, 4, 5}
	intsArray = mapArray(addOne, intsArray)
	fmt.Println(intsArray)

	newSlice := intsSlice[1:3]
	mapSlice(square, newSlice)
	fmt.Println(intsSlice, newSlice)

	intsSlice = double(intsSlice)
	fmt.Println(intsSlice)
}
