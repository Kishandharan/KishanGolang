package main
import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	sum := 0
	for _, arg := range os.Args{
		argInt, _ := strconv.Atoi(arg)
		sum += argInt }
	fmt.Println(sum)
}
