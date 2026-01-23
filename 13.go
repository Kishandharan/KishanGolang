package main 
import "fmt"

func main(){
	var map1 map[int]string = make(map[int]string)

	map1[1] = "Hello"
	map1[2] = "World"
	map1[7] = "This is so cool"
	fmt.Println(map1[1])
	fmt.Println(map1[2])
	fmt.Println(map1[7])
	fmt.Println(map1)
	delete(map1, 7)
	fmt.Println(map1)
	anonexistantvalue, ok := map1[7]
	fmt.Printf("%v %v \n", anonexistantvalue, ok)
}
