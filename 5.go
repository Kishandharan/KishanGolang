package main
import (
	"fmt"
	"os"
	"strconv"
)

func main(){

	var filename string
	fmt.Print("Enter the file name for the file to output the tables to: ")
	fmt.Scanln(&filename)
	fmt.Println("All the tables will be outputed to a file named "+filename)
	fmt.Println()
	output_file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	var ms int
	var me int
	var ts int
	var te int
	if err != nil{
		panic("Oopsie!")
	}
	fmt.Print("Enter Multiplier Start: ")
	fmt.Scanln(&ms)
	fmt.Print("Enter Multiplier End: ")
	fmt.Scanln(&me)
	fmt.Print("Enter Table Start: ")
	fmt.Scanln(&ts)
	fmt.Print("Enter Table End: ")
	fmt.Scanln(&te)
	for x:=ts; x<=te; x++{
		for y:=ms; y<=me; y++{
			output_file.WriteString( (strconv.Itoa(y)+"x"+strconv.Itoa(x)+"="+strconv.Itoa(y*x))+"\n" );
		}
		output_file.WriteString("\n");
	} 
	output_file.Close()
}
