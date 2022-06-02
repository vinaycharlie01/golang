package main

import "fmt"

type person struct {
	firstname string
	age       int
}

func changestr(p person) {
	p.age = 20
	p.firstname = "vinay"
}
func changestrpointer(p *person) {
	(*p).age = 39
	(*p).firstname = "vinaykumar"
}

func change(a int, b float64, c bool, d string) {
	a = 20
	b = 23.3
	c = true
	d = "vinay"

}
func changeusingpointer(a *int, b *float64, c *bool, d *string) {
	*a = 20
	*b = 23.3
	*c = true
	*d = "vinay"

}

func main() {
	e, f, g, h := 3, 23.3, false, "viny"
	fmt.Println(e, f, g, h)
	change(e, f, g, h)
	fmt.Println(e, f, g, h)
	changeusingpointer(&e, &f, &g, &h)
	fmt.Println(e, f, g, h)
	p1 := person{
		firstname: "karli",
		age:       20,
	}
	fmt.Println(p1)
	changestr(p1)
	fmt.Println(p1)
	changestrpointer(&p1)
	fmt.Println(p1)

}
