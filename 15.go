package main
import "fmt"

func main(){
	/*
	x := 11 
	y := x
	y = 100
	fmt.Println(x, y)

	This changes the y variable only, not the actual x variable. 
	Because int values copy by value, not reference.
	*/

	x := []int{1, 2, 3}
	y := x
	y[0] = 218418083
	fmt.Println(x, y)
	
	// They both have the same value even if I changed only y 
	// This is because slices copy by reference, in contrast to ints
}
