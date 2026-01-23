package main
import "fmt"

func main(){
	var str1 = "hello";
	var pStr1 = &str1;
	*pStr1 = "bye";
	fmt.Println(str1)
}
