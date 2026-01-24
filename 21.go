package main
import (
	"os"
)

func main(){
	f1, _ := os.OpenFile("MyFile2.txt", os.O_CREATE | os.O_RDWR, 0644)
	defer f1.Close()
	f1.WriteString("Hello! Do you hear me??")
}

