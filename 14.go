package main
import "fmt"

func greet(name string){
	fmt.Printf("Hello %v \n", name)
}

func main(){
	var greet2 = func(){
		fmt.Println("Hello!")
	}

	var greet3 = func() string{
		return "Hello again!"
	}()
	
	var greet4 = func()(string, string){
		return "hello", "world"
	}
	greet("Kishan")
	greet2()
	fmt.Println(greet3)
	fmt.Println(greet4())
	// Like this, functions can also be passed as parameters and can be returned
}
