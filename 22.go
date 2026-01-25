package main
import (
	"fmt"
	"strconv"
)

func main(){
	var start1 string; // Starting of table
	var start2 string; // Startin of multiplier
	var end1 string; // Ending of table
	var end2 string; // Ending of multiplier

	var Istart1 int; // Starting of table
	var Istart2 int; // Startin of multiplier
	var Iend1 int; // Ending of tabhe
	var Iend2 int; // Ending of multiplier

	fmt.Println("Enter table start, table end, multiplier start and multiplier end, below, respectively:")
	fmt.Scanln(&start1)
	fmt.Scanln(&end1)
	fmt.Scanln(&start2)
	fmt.Scanln(&end2)

	Istart1, _ = strconv.Atoi(start1)
	Istart2, _ = strconv.Atoi(start2)
	Iend1, _ = strconv.Atoi(end1)
	Iend2, _ = strconv.Atoi(end2)

	for i := Istart1; i <= Iend1; i++{
		for j := Istart2; j <= Iend2; j++{
			fmt.Printf("%vx%v=%v\n", j, i, j*i)
		}
		fmt.Println()
	}
}
