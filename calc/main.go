package main

import (
	"fmt"
	"os"

	app "calc/internal"
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
