package main

import (
	"fmt"

	app "github.com/1mizhgun1/Uniq_strings/internal"
)

func main() {
	args := app.Arguments{}
	args.Parse()

	var data []string
	err := app.ReadData(&data, args)
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := app.Uniq(data, args)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = app.WriteAnswer(result, args)
	if err != nil {
		fmt.Println(err)
		return
	}
}
