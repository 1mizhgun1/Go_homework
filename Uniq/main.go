package main

import (
	"flag"
	"fmt"

	app "github.com/1mizhgun1/Uniq_strings/internal"
)

func main() {
	var args = app.Arguments{}
	args.Parse()
	if !args.IsValid() {
		flag.Usage()
		return
	}

	data, err := app.ReadData(args)
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
