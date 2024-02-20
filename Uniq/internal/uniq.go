package uniq

import (
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"
)

func cutWords(data string, num int) string {
	words := strings.Fields(data)
	if num >= len(words) {
		return ""
	}
	return strings.Join(words[num:], " ")
}

func cutChars(data string, chars int) string {
	if chars >= utf8.RuneCountInString(data) {
		return ""
	}
	return string([]rune(data)[chars:])
}

func performString(data string, args Arguments) string {
	result := data

	if args.i {
		result = strings.ToLower(result)
	}

	if args.num > 0 {
		result = cutWords(result, args.num)
	} else if args.chars > 0 {
		result = cutChars(result, args.chars)
	}

	return result
}

type uniqCounter struct {
	count int
	data  string
	index int // index of original string in original slice
}

func Uniq(data []string, args Arguments) ([]string, error) {
	if len(data) == 0 {
		return nil, errors.New("Input file is empty\n")
	}

	if !args.isValid() {
		return nil, errors.New("Command should be in this format:\nuniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
	}

	performedData := []string{}
	for _, item := range data {
		performedData = append(performedData, performString(item, args))
	}

	var counter []uniqCounter
	currentState := uniqCounter{count: 1, data: performedData[0], index: 0}

	for i, item := range performedData {
		if i == 0 {
			continue
		}

		if item == currentState.data {
			currentState.count++
		} else {
			counter = append(counter, currentState)
			currentState = uniqCounter{count: 1, data: item, index: i}
		}
	}
	counter = append(counter, currentState)

	result := []string{}
	for _, item := range counter {
		if args.c {
			result = append(result, strconv.Itoa(item.count)+" "+data[item.index])
		} else if args.d && item.count > 1 || args.u && item.count == 1 || !args.d && !args.u {
			result = append(result, data[item.index])
		}
	}

	return result, nil
}
