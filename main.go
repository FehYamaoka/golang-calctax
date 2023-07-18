package main

import (
	"database/sql"
	"fmt"

	"github.com/FehYamaoka/devfullcycle/golang-calctax/internal/infra/database"
	"github.com/FehYamaoka/devfullcycle/golang-calctax/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	Model string
	Color string
}

// metodo
func (c Car) Start() {
	println(c.Model + " has been started")
}

func (c *Car) ChangeColor(color string) {
	c.Color = color // duplicando o valor de c.color na memoria
	println("New Color: " + c.Color)
}

// funcao
func soma(x, y int) int {
	return x + y
}

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	orderRepository := database.NewOrderRepository(db)

	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    "124",
		Price: 10.0,
		Tax:   1.0,
	}

	output, err := uc.Execute(input)

	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}
