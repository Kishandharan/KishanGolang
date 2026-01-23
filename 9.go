package main
import "fmt"

func main(){
	x := 0
	for x <= 5{
		fmt.Println("Hello!")
		if(x == 3){
			break
		}
		x++
	}

	fmt.Println("")

	for i := 1; i <= 10; i++{
		if(i==2){
			continue
		}
		fmt.Println("Hello!")
	}
}
