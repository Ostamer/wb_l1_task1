package main

import "fmt"

// Структура Human
type Human struct {
	Name   string
	Age    int
	Height int
	Weight int
}

// Метод структуры Human
func (h Human) Speak() {
	fmt.Printf("Привет, меня зовут %s, мой возраст %d, мой рост %d см, а вес %d кг.\n", h.Name, h.Age, h.Height, h.Weight)
}

// Стурктура Action от стурктуры Human
type Action struct {
	Human
}

// Метод структуры Action
func (a Action) Sport() {
	if a.Height >= 180 {
		fmt.Println(a.Name, "Бежит")
	} else {
		fmt.Println(a.Name, "Приседает")
	}
}

func main() {
	// Создаем экземпляр Action от экземпляра Human
	action := Action{
		Human: Human{
			Name:   "Артур",
			Age:    21,
			Height: 182,
			Weight: 60,
		},
	}

	// Используем методы от структуры Human
	action.Speak()

	// Используем метод структуры Action
	action.Sport()
}
