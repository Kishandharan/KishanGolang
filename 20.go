package main
import (
	"os"
)

func main(){
	_, _ = os.OpenFile("MyFile.txt", os.O_CREATE | os.O_TRUNC, 0644)
}
