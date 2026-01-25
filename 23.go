package main
import (
	"fmt"
	"strconv"
	"os"
)

func main(){
	var start1 string; // Starting of table
	var start2 string; // Startin of multiplier
	var end1 string; // Ending of table
	var end2 string; // Ending of multiplier
	var fname string;

	var Istart1 int; // Starting of table
	var Istart2 int; // Startin of multiplier
	var Iend1 int; // Ending of tabhe
	var Iend2 int; // Ending of multiplier

	fmt.Println("Enter table start, table end, multiplier start, multiplier end and file name below, respectively:")
	fmt.Scanln(&start1)
	fmt.Scanln(&end1)
	fmt.Scanln(&start2)
	fmt.Scanln(&end2)
	fmt.Scanln(&fname)

	outputFile, _ := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	Istart1, _ = strconv.Atoi(start1)
	Istart2, _ = strconv.Atoi(start2)
	Iend1, _ = strconv.Atoi(end1)
	Iend2, _ = strconv.Atoi(end2)

	for i := Istart1; i <= Iend1; i++{
		for j := Istart2; j <= Iend2; j++{
			fmt.Printf("%vx%v=%v\n", j, i, j*i)
			temp1 := fmt.Sprintf("%vx%v=%v\n", j, i, j*i) 
			outputFile.WriteString(temp1)
		}
		outputFile.WriteString("\n")
		fmt.Println()
	}

	outputFile.Close()
}
