package main
import "fmt"

type Employee struct{
	name string
	age int
	salary float64
}

func (emp Employee) getSalary() float64{
	return emp.salary
}

func (emp *Employee) setSalary(salary float64){
	emp.salary = salary
}

func main(){
	//employee1 := Employee{"Mukesh", 40, 100} 
	//or
	employee1 := Employee{name:"Mukesh", salary:100, age:40} 
	p_Employee1 := &employee1

	fmt.Println(employee1.name)
	fmt.Println(employee1.age)
	fmt.Println(employee1.salary)

	p_Employee1.name = "Chandrashekar"

	fmt.Println(employee1.name)
	fmt.Println(employee1.age)
	fmt.Println(employee1.salary)

	employee1.setSalary(300)

	fmt.Println(employee1.getSalary())

}
