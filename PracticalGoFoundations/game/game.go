package main

import (
	"fmt"
	"slices"
)

func main() {
	var i Item
	fmt.Printf("i: %#v\n", i)

	i = Item{10, 20} // Must specify all fields
	fmt.Printf("i: %#v\n", i)

	// Can be in any order, can be partial
	i = Item{
		Y: 22,
		//X: 11,
	}
	fmt.Printf("i: %#v\n", i)

	fmt.Println(NewItem(10, 20))
	fmt.Println(NewItem(10, 2000))

	/* Aside: Use %#v for debugging/logging
	a, b := 1, "1"
	fmt.Printf("a=%v, b=%v\n", a, b)
	fmt.Printf("a=%#v, b=%#v\n", a, b)
	*/

	i.Move(10, 20)
	fmt.Printf("i (move): %#v\n", i)

	p1 := Player{
		Name: "Parzival",
	}
	fmt.Printf("p1: %+v\n", p1)
	fmt.Println("p1.X:", p1.X)
	p1.Move(100, 200)
	fmt.Printf("p1 (move): %+v\n", p1)

	fmt.Println(p1.Found(Copper)) // <nil>
	fmt.Println(p1.Found(Copper)) // <nil>
	fmt.Println(p1.Found(Key(7))) // unknown key: "gold"
	fmt.Println("keys:", p1.Keys) // keys: [copper]

	ms := []Mover{
		&i,
		&p1,
	}

	moveAll(ms, 50, 70)
	for _, m := range ms {
		fmt.Println(m)
	}
}

// Rule of thumb: Accept interface, return types

// go install golang.org/x/tools/cmd/stringer@latest
// In ~/.zshrc
// export PATH="$(go env GOPATH)/bin:${PATH}"

// String implement the fmt.Stringer interface.
func (k Key) String() string {
	switch k {
	case Copper:
		return "copper"
	case Jade:
		return "jade"
	case Crystal:
		return "crystal"
	}

	return fmt.Sprintf("<Key %d>", k)
}

type Key byte

const (
	Copper Key = iota + 1
	Jade
	Crystal
)

/*
	Thought exercise: Sorting

	func Sort(s Sortable) {
		// ...
	}

	type Sortable interface {
		Less(i, j int) bool
		Swap(i, j int)
		Len() int
	}

/*

	Intefaces

- Set of methods (and types)
- We define interfaces as "what you need", not "what you provide"
  - Interfaces are small (stdlib average ~2)
  - If you have interface with more that 4 methods, think again

- Start with concrete types, discover interface
*/
type Mover interface {
	Move(int, int)
}

func moveAll(ms []Mover, dx, dy int) {
	for _, m := range ms {
		m.Move(dx, dy)
	}
}

/* Exercise:
- Add a "Keys" field to Player with is a slice of strings
- Add a "Found(key string)" method to player
	- It should err if key is not one of "jade", "copper", "crystal"
	- It should add a key only once
*/

func (p *Player) Found(key Key) error {
	switch key {
	case Copper, Jade, Crystal:
		// OK
	default:
		return fmt.Errorf("unknown key: %q", key)
	}

	if !slices.Contains(p.Keys, key) {
		p.Keys = append(p.Keys, key)
	}

	return nil
}

type Player struct {
	Name string
	Item // Player embeds Item
	Keys []Key
}

// Move moves i by delta x & delta y.
// "i" is called "the receiver"
// i is a pointer receiver
/* Value vs pointer receiver
- In general use value semantics
- Try to keep same semantics on all methods
- When you must use pointer receiver
	- If you have lock field
	- If you need to mutate the struct
	- Decoding/unmarshalling
*/
func (i *Item) Move(dx, dy int) {
	i.X += dx
	i.Y += dy
}

/* Types of "new" or factory functions.
func NewItem(x, y int) Item
func NewItem(x, y int) *Item
func NewItem(x, y int) (Item, error)
func NewItem(x, y int) (*Item, error)

Value semantics: everyone has their own copy
Pointer semantics: everyong share the same copy (heap, lock)
*/

func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		// Value semantics
		// return Item{}, fmt.Errorf("%d/%d of of bounds %d/%d", x, y, maxX, maxY)
		return nil, fmt.Errorf("%d/%d of of bounds %d/%d", x, y, maxX, maxY)
	}

	i := Item{
		X: x,
		Y: y,
	}
	return &i, nil
}

const (
	maxX = 600
	maxY = 400
)

type Item struct {
	X int
	Y int
}