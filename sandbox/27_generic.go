package main

import "fmt"

type Box[T any] struct {
	name string
	items []T
}

func NewBox[T any](name string) *Box[T] {
	return &Box[T]{
		name: name,
		items: []T{},
	}
}

func (b *Box[T]) Add(item T) {
	b.items = append(b.items, item)
}

func main() {
	box := NewBox[string]("Hello")
	box.Add("go")
	box.Add("lang")

	fmt.Println(box.items)
}