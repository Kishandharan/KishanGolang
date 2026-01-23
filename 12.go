package main
import "fmt"

func main(){
	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)
	slice1 = append(slice1, 10)
	fmt.Println(slice1)

	slice2 := make([]int, 1)
	fmt.Println(slice2)
	slice2 = append(slice2, 2)
	fmt.Println(slice2)
	fmt.Println(cap(slice2))
	fmt.Println(len(slice2))

	for i, element := range slice2{
		fmt.Println(i)
		fmt.Println(element)
	}
}
