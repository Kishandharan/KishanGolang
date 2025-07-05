package main
import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main(){
	cmd_args := os.Args[1:]
	var maintimer int 
	for _, x := range cmd_args{
		temp, _ := strconv.Atoi(x)
		maintimer += temp 
	}
	
	maintimer += 1

	for _, x := range cmd_args{
		go func(){
			temp, _ := strconv.Atoi(x)
			time.Sleep(time.Duration(temp) * time.Second)
			fmt.Println("Beep Beep!!!")
		}()
	}

	time.Sleep(time.Duration(maintimer) * time.Second)
}
