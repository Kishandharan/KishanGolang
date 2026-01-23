package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Dei, yenna venumo podu, yenakku pasikkidhu: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Println(input)
	fmt.Printf("%T", input)
}
