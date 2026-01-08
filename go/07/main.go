package main

type Animaler interface {
	GetName() string
	Setname(name string)
}

type Cat struct {
	name string
}

func (c Cat) GetName() string {
	return c.name
}

func (c *Cat) Setname(name string) {
	c.name = name
}

func main() {
	c := Cat{
		name: "tom",
	}
	var a Animaler = &c
	println(a.GetName())
	a.Setname("jerry")
	println(a.GetName())
}
