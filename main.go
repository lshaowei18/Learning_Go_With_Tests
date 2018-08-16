package main

import "fmt"

func main() {
	// Create buffer and slice
	var buffer [256]byte
	slice := buffer[10:20]
	//Print their types out
	fmt.Printf("buffer type : %T\n", buffer)
	fmt.Printf("slice type : %T\n", slice)
	//Loop through and add values to the slice
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}

	//Show that slice can pass a pointer in
	fmt.Println("Before:", slice)
	AddOneToEachElement(slice)
	fmt.Println("After:", slice)

	fmt.Println("Before: len(slice) = ", len(slice))
	newSlice := SubtractOneFromLength(slice)
	fmt.Println("After length of slice =", len(slice))
	fmt.Println("After length of newSlice =", len(newSlice))
}

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}

func SubtractOneFromLength(slice []byte) []byte {
	slice = slice[0 : len(slice)-1]
	return slice
}
