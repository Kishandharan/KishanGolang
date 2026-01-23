package main
import "fmt"

func main(){
	var float1 float64 = 3.35
	float2 := fmt.Sprintf("%T %v %% %T", "hello", "hello", float1)
	fmt.Printf("%T %v %% %T \n", "hello", "hello", float1)
	fmt.Println(float2)
}
