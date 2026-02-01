package main
import ("fmt")

func main(){
	print("I can give this function any type input")
	print(1)
}

func print[T any](data T){
	fmt.Println(data);
}
