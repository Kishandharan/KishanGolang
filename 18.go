package main
import "fmt"

type Speakable interface{
	speak() 
}

type Dog struct{}
type Human struct{}

func (h Human) speak(){
	fmt.Println("I can speak")
}

func main(){
	speakables := []Speakable{}
	speakables = append(speakables, Human{})
	speakables = append(speakables, Human{})	

	for _, speakable := range speakables{
		speakable.speak()
	}
}
