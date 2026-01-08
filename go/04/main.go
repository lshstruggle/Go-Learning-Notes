package main

import "fmt"

type Person struct {
	age  int
	name string
	sex  string
}

func (p Person) PrintInfo() {
	fmt.Printf("name=%v,age=%v,sex=%v\n", p.name, p.age, p.sex)
}

func (p *Person) SetInfo(age int, name string) {
	p.age = age
	p.name = name
}
func main() {
	var p1 = Person{
		age:  18,
		name: "陈涛",
		sex:  "女",
	}
	p1.PrintInfo()
	p1.SetInfo(30, "chentao")
	p1.PrintInfo()

}
