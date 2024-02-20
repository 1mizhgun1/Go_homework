package calc

import (
	"errors"
	"strconv"
	"strings"
)

const answerWhenError = 0.0

func prepare(expression string) string {
	expression = strings.ReplaceAll(expression, " ", "")

	expression = strings.ReplaceAll(expression, "(-", "(0-")
	if getSize(expression) >= 2 && expression[0] == '-' && isDigit(rune(expression[1])) {
		expression = "0" + expression
	}

	return expression
}

func validate(expression string) error {
	balance := 0
	size := getSize(expression)

	if size == 0 {
		return errors.New("Пустое выражение.\n")
	}

	for i, char := range expression {
		if isOperation(char) {
			if i == 0 || !isDigit(rune(expression[i-1])) && !isClosedBracket(rune(expression[i-1])) {
				return errors.New("Сразу до знака операции должен идти операнд, либо закрывающая скобка.\n")
			} else if i == size-1 || !isDigit(rune(expression[i+1])) && !isOpenedBracket(rune(expression[i+1])) {
				return errors.New("Сразу после знака операции должен идти операнд, либо открывающая скобка.\n")
			}
		} else if isOpenedBracket(char) {
			balance++
			if i == size-1 || !isDigit(rune(expression[i+1])) && !isOpenedBracket(rune(expression[i+1])) {
				return errors.New("Сразу после открывающей скобки должен идти операнд, либо открывающая скобка.\n")
			}
		} else if isClosedBracket(char) {
			balance--
			if i == 0 || !isDigit(rune(expression[i-1])) && !isClosedBracket(rune(expression[i-1])) {
				return errors.New("Закрывающая скобка должна идти сразу после операнда, либо закрывающей скобки.\n")
			}
		} else if !isDigit(char) && !isEndline(char) {
			return errors.New("Выражение может состоять только из этих символов: 0123456789()+-*/\n")
		}

		if balance < 0 {
			return errors.New("Закрывающая скобка должна идти после соответствующей открывающей.\n")
		}
	}

	if balance != 0 {
		return errors.New("Количество открывающих и закрывающих скобок должно быть одинаково.\n")
	}

	return nil
}

func convertToSlice(expression string) []string {
	var tokens []string
	currentToken := ""

	for _, char := range expression {
		if isOperation(char) || isOpenedBracket(char) || isClosedBracket(char) {
			if currentToken != "" {
				tokens = append(tokens, currentToken)
				currentToken = ""
			}
			tokens = append(tokens, string(char))
		} else if isDigit(char) {
			currentToken += string(char)
		}
	}

	if currentToken != "" {
		tokens = append(tokens, currentToken)
	}

	return tokens
}

func covertToPostfix(infix []string) []string {
	var postfix []string
	var stack []string
	precedence := map[string]int{"+": 1, "-": 1, "*": 2, "/": 2}

	for _, token := range infix {
		switch token {
		case "+", "-", "*", "/":
			for len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[token] {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		case "(":
			stack = append(stack, token)
		case ")":
			for stack[len(stack)-1] != "(" {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		default:
			postfix = append(postfix, token)
		}
	}

	for len(stack) > 0 {
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return postfix
}

func solve(postfix []string) float64 {
	var stack []float64

	for _, token := range postfix {
		switch token {
		case "+":
			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, op1+op2)
		case "-":
			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, op1-op2)
		case "*":
			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, op1*op2)
		case "/":
			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, op1/op2)
		default:
			num, _ := strconv.ParseFloat(token, 64)
			stack = append(stack, num)
		}
	}

	return stack[0]
}

func Calc(expression string) (float64, error) {
	preparedExpression := prepare(expression)

	if err := validate(preparedExpression); err != nil {
		return answerWhenError, err
	}

	slice := convertToSlice(preparedExpression)
	postfix := covertToPostfix(slice)

	return solve(postfix), nil
}
