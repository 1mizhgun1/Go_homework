package uniq

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

func readLines(reader io.Reader, data *[]string) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		*data = append(*data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func ReadData(data *[]string, args Arguments) error {
	if utf8.RuneCountInString(args.input) > 0 {
		file, err := os.Open(args.input)
		if err != nil {
			return err
		}
		defer file.Close()

		err = readLines(file, data)
		if err != nil {
			return err
		}
	} else {
		err := readLines(os.Stdin, data)
		if err != nil {
			return err
		}
	}
	return nil
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
	if utf8.RuneCountInString(args.output) > 0 {
		file, err := os.Create(args.output)
		if err != nil {
			return err
		}
		defer file.Close()

		err = writeLines(file, data)
		if err != nil {
			return err
		}
	} else {
		err := writeLines(os.Stdout, data)
		if err != nil {
			return err
		}
	}
	return nil
}
