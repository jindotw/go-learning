package main

import "fmt"

type Person struct {
	Name string
}

type Hero struct {
	*Person
	HeroName string
	HeroRank int64
}

func (hero *Hero) fly() {
	fmt.Printf("Hero %s can fly\n", hero.HeroName)
}

func (person *Person) walk() {
	fmt.Printf("Person %s can walk\n", person.Name)
}

func main() {
	tony := &Hero{&Person{"Tony Stark"}, "IronMan", 1}
	fmt.Println(tony.Name, tony.HeroName, tony.HeroRank)
	tony.fly()
	tony.walk()
}
