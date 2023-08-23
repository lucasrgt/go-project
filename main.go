package main

import (
	"awesomeProject/internal/infra/database"
	"awesomeProject/internal/usecases"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	Model string
	Color string
}

// metodo

func (c *Car) Start() {
	println(c.Model + " has been started")
}

func (c *Car) ChangeColor(color string) {
	c.Color = color
}

// funcao

func soma(x int, y int) int {
	return x + y
}

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")

	if err != nil {
		panic(err)
	}
	orderRepository := database.NewOrderRepository(db)
	usecase := usecases.NewCalculateFinalPrice(orderRepository)

	input := usecases.OrderInput{
		ID:    "123",
		Price: 10.0,
		Tax:   1.0,
	}

	output, err := usecase.Execute(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(*output)
}
