package main
import (
	"fmt"
)

type Point struct{
	x int
	y int
}

func main(){
	point1 := Point{10, 49}
	fmt.Println(point1.x)
	fmt.Println(point1.y)
}
