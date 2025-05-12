package main
import "fmt"
import "os"

func main(){
	file, err := os.Create("Example.txt")
	if err != nil{
		fmt.Println("We have got an error!");
	}
	defer file.Close()
}
