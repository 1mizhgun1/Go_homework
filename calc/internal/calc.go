package calc

import (
	"calc/internal/utils"
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
		return errors.New("пустое выражение")
	}

	runedExpression := []rune(expression)
	for i, char := range runedExpression {
		if isOperation(char) {
			if i == 0 || !isDigit(runedExpression[i-1]) && !isClosedBracket(runedExpression[i-1]) {
				return errors.New("сразу до знака операции должен идти операнд, либо закрывающая скобка")
			} else if i == size-1 || !isDigit(runedExpression[i+1]) && !isOpenedBracket(runedExpression[i+1]) {
				return errors.New("сразу после знака операции должен идти операнд, либо открывающая скобка")
			}
			continue
		}

		if isOpenedBracket(char) {
			balance++
			if i == size-1 || !isDigit(runedExpression[i+1]) && !isOpenedBracket(runedExpression[i+1]) {
				return errors.New("сразу после открывающей скобки должен идти операнд, либо открывающая скобка")
			}
			continue
		}

		if isClosedBracket(char) {
			balance--
			if i == 0 || !isDigit(runedExpression[i-1]) && !isClosedBracket(runedExpression[i-1]) {
				return errors.New("закрывающая скобка должна идти сразу после операнда, либо закрывающей скобки")
			}
			continue
		}

		if !isDigit(char) && !isEndline(char) {
			return errors.New("выражение может состоять только из этих символов: 0123456789()+-*/")
		}
	}

	if balance != 0 {
		return errors.New("количество открывающих и закрывающих скобок должно быть одинаково")
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
	var (
		postfix []string
		stack   = utils.CreateStack()
	)
	precedence := map[string]int{"+": 1, "-": 1, "*": 2, "/": 2}

	for _, token := range infix {
		switch token {
		case "+", "-", "*", "/":
			for !stack.IsEmpty() && precedence[stack.Top().(string)] >= precedence[token] {
				postfix = append(postfix, stack.Top().(string))
				stack.Pop()
			}
			stack.Push(token)
		case "(":
			stack.Push(token)
		case ")":
			for stack.Top() != "(" {
				postfix = append(postfix, stack.Top().(string))
				stack.Pop()
			}
			stack.Pop()
		default:
			postfix = append(postfix, token)
		}
	}

	for !stack.IsEmpty() {
		postfix = append(postfix, stack.Top().(string))
		stack.Pop()
	}

	return postfix
}

func extractOperands(stack *utils.Stack) (float64, float64, bool, bool) {
	op2, ok2 := stack.Top().(float64)
	stack.Pop()
	op1, ok1 := stack.Top().(float64)
	stack.Pop()
	return op1, op2, ok1, ok2
}

func solve(postfix []string) (float64, error) {
	stack := utils.CreateStack()

	for _, token := range postfix {
		switch token {
		case "+":
			op1, op2, ok1, ok2 := extractOperands(stack)
			if ok2 && ok1 {
				stack.Push(op1 + op2)
			} else {
				return answerWhenError, errors.New("internal error")
			}
		case "-":
			op1, op2, ok1, ok2 := extractOperands(stack)
			if ok2 && ok1 {
				stack.Push(op1 - op2)
			} else {
				return answerWhenError, errors.New("internal error")
			}
		case "*":
			op1, op2, ok1, ok2 := extractOperands(stack)
			if ok2 && ok1 {
				stack.Push(op1 * op2)
			} else {
				return answerWhenError, errors.New("internal error")
			}
		case "/":
			op1, op2, ok1, ok2 := extractOperands(stack)
			if ok2 && ok1 {
				stack.Push(op1 / op2)
			} else {
				return answerWhenError, errors.New("internal error")
			}
		default:
			num, _ := strconv.ParseFloat(token, 64)
			stack.Push(num)
		}
	}

	result, ok := stack.Top().(float64)
	if !ok {
		return answerWhenError, errors.New("internal error")
	}

	return result, nil
}

func Calc(expression string) (float64, error) {
	preparedExpression := prepare(expression)

	if err := validate(preparedExpression); err != nil {
		return answerWhenError, err
	}

	slice := convertToSlice(preparedExpression)
	postfix := covertToPostfix(slice)

	return solve(postfix)
}
