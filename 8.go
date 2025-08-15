package main
import (
	"os"
	"fmt"
)

func main(){
	err1 := os.Mkdir("friends", 0744)
	err2 := os.Chdir("friends")
	friend1, err3 := os.OpenFile("friend1.txt", os.O_CREATE, 0644)
	friend2, err4 := os.OpenFile("friend2.txt", os.O_CREATE, 0644)
	friend3, err5 := os.OpenFile("friend3.txt", os.O_CREATE, 0644)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil{
		fmt.Println("There was some errors")
	}
	friend1.Close()
	friend2.Close()
	friend3.Close()
}
