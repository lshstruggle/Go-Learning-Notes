package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Age int
	// Address Address
	Name string
}
type Address struct {
	City   string
	Person []Person
}

func main() {
	// var p1 = Person{
	// 	Age: 18,
	// 	Address: Address{
	// 		City: "北京",
	// 	},
	// 	Name: "陈涛",
	// }
	// // p1.address.city = "北京"
	// // println(p1.address.city)
	// jsonByte, _ := json.Marshal(p1)
	// fmt.Printf("%T\n", string(jsonByte))
	// var p2 Person
	// json.Unmarshal([]byte(jsonByte), &p2)
	// fmt.Println(p2)
	var c = Address{
		City:   "北京",
		Person: make([]Person, 0),
	}
	for i := 0; i <= 10; i++ {
		p := Person{
			Age:  i,
			Name: fmt.Sprintf("name%d", i),
		}
		c.Person = append(c.Person, p)
	}
	// fmt.Println(c)
	strByte, _ := json.Marshal(c)
	fmt.Println(string(strByte))
}
