package uniq

import (
	"errors"
	"strconv"
	"strings"
)

const UsageError = "incorrect usage of flags"

func cutWords(data string, num int) string {
	words := strings.Fields(data)
	if num >= len(words) {
		return ""
	}
	return strings.Join(words[num:], " ")
}

func cutChars(data string, chars int) string {
	var runeData []rune = []rune(data)
	if chars >= len(runeData) {
		return ""
	}
	return string(runeData[chars:])
}

// processing i, f, s flags for one string
func performString(data string, args Arguments) string {
	result := data

	if args.i {
		result = strings.ToLower(result)
	}

	if args.num > 0 {
		result = cutWords(result, args.num)
	}

	if args.chars > 0 {
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
		return []string{}, nil
	}

	if !args.IsValid() {
		return nil, errors.New(UsageError)
	}

	// preparing data by processing i, f, s flags
	var performedData []string
	for _, item := range data {
		performedData = append(performedData, performString(item, args))
	}

	var counter []uniqCounter
	currentState := uniqCounter{count: 1, data: performedData[0], index: 0}

	for i, item := range performedData {
		if i == 0 {
			continue
		}

		// count number of equal strings
		if item == currentState.data {
			currentState.count++
		} else {
			counter = append(counter, currentState)
			currentState = uniqCounter{count: 1, data: item, index: i}
		}
	}
	counter = append(counter, currentState)

	// make result strings, depending from c | d | u flags
	var result []string
	for _, item := range counter {
		if args.c {
			result = append(result, strconv.Itoa(item.count)+" "+data[item.index])
		} else if args.d && item.count > 1 || args.u && item.count == 1 || !args.d && !args.u {
			result = append(result, data[item.index])
		}
	}

	return result, nil
}
