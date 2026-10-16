package main
import (
	"fmt"
	"sync"
	"time"
)

var wait = sync.WaitGroup{}

func main(){
	wait.Add(1);
	go task1();
	wait.Add(1);
	go task2();
	wait.Add(1);
	go task3();

	wait.Wait();
}

func task1(){
	defer wait.Done();
	time.Sleep(3 * time.Second);
	fmt.Println("Task1 done");
}

func task2(){
	defer wait.Done();
	time.Sleep(2 * time.Second);
	fmt.Println("Task2 done");
}

func task3(){
	defer wait.Done();
	time.Sleep(1 * time.Second);
	fmt.Println("Task3 done");
}
