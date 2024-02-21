package uniq

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readLines(reader io.Reader) ([]string, error) {
	var data []string

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

func ReadData(args Arguments) ([]string, error) {
	var data []string
	var reader io.Reader = os.Stdin

	if len(args.input) > 0 {
		file, err := os.Open(args.input)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		reader = file
	}

	data, err := readLines(reader)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func writeLines(writer io.Writer, data []string) error {
	for _, line := range data {
		_, err := fmt.Fprintln(writer, line)
		if err != nil {
			return err
		}
	}
	return nil
}

func WriteAnswer(data []string, args Arguments) error {
	var writer io.Writer = os.Stdout

	if len(args.output) > 0 {
		file, err := os.Create(args.output)
		if err != nil {
			return err
		}
		defer file.Close()

		writer = file
	}

	err := writeLines(writer, data)
	if err != nil {
		return err
	}

	return nil
}
