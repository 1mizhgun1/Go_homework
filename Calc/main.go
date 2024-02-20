package main

import (
	"fmt"
	app "github.com/1mizhgun1/Calc_expressions/internal"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Пустое выражение.")
		return
	}
	expression := os.Args[1]

	result, err := app.Calc(expression)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%f\n", result)
}
