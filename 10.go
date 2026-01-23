package main
import "fmt"

func main(){
	chr1 := "c"
	num1 := 100

	switch chr1{
		case "a":
			fmt.Println(1)
		case "b":
			fmt.Println(2)
		default:
			fmt.Println("Invalid")
	}

	switch{
		case num1 > 10:
			fmt.Println("Greater than ten")
		case num1 < 10:
			fmt.Println("Less than ten")
		default:
			fmt.Println("Invalid")
	}
}
