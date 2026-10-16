package main
import "fmt"

type person struct {
	name string
	age int
	gender bool
}

type rectangle struct {}
func (r rectangle) speak(){
	fmt.Println("I AM THE RECTANGLE!! I AM EVERYWHERE!! HAHAHAHA");
}

type circle struct {}
func (r circle) speak(){
	fmt.Println("I AM THE CIRCLE! I AM THE SHAPE OF YOUR WATER BOTTLE! HAHAHAA");
}

type shape interface{
	speak()
}
func speakCommon(s shape){
	s.speak();
}

func (p person) displayDetails(){
	var gender string; 

	if p.gender == false{
		gender = "Boy";
	}else{
		gender = "Girl";
	}

	fmt.Println("Name:", p.name);
	fmt.Println("Age:", p.age);
	fmt.Println("Gender:", gender);
}


func main(){
	var person1 person = person{"Kishan", 13, false};
	person1.displayDetails();

	var rectangle1 rectangle;
	var circle1 circle;
	speakCommon(rectangle1);
	speakCommon(circle1);
}
