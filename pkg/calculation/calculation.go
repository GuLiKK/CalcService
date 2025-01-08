package calculation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	tokens := tokenize(expression)
	postfix, err := infixToPostfix(tokens)
	if err != nil {
		return 0, err
	}
	return evaluatePostfix(postfix)
}

func tokenize(expr string) []string {
	var tokens []string
	var currentToken strings.Builder

	for i, char := range expr {
		if char == ' ' {
			continue
		}

		if char == '-' {
			// Проверяем, является ли минус унарным (например, "-5" или "2*-3")
			// Условие: либо первый символ, либо после оператора/скобки
			if i == 0 || len(tokens) == 0 ||
				isOperator(tokens[len(tokens)-1]) ||
				tokens[len(tokens)-1] == "(" {
				currentToken.WriteRune(char)
				continue
			}
		}

		switch char {
		case '+', '-', '*', '/', '(', ')':
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}
			tokens = append(tokens, string(char))
		default:
			currentToken.WriteRune(char)
		}
	}
	if currentToken.Len() > 0 {
		tokens = append(tokens, currentToken.String())
	}
	return tokens
}

func infixToPostfix(tokens []string) ([]string, error) {
	var output []string
	var operators []string

	for _, token := range tokens {
		switch {
		case isNumber(token):
			output = append(output, token)
		case token == "(":
			operators = append(operators, token)
		case token == ")":
			foundParen := false
			for len(operators) > 0 {
				top := operators[len(operators)-1]
				operators = operators[:len(operators)-1]
				if top == "(" {
					foundParen = true
					break
				}
				output = append(output, top)
			}
			if !foundParen {
				return nil, ErrInvalidExpression
			}
		case isOperator(token):
			for len(operators) > 0 && precedence(operators[len(operators)-1]) >= precedence(token) {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
		default:
			return nil, ErrInvalidExpression
		}
	}

	for len(operators) > 0 {
		top := operators[len(operators)-1]
		operators = operators[:len(operators)-1]
		if top == "(" {
			return nil, ErrInvalidExpression
		}
		output = append(output, top)
	}
	return output, nil
}

func evaluatePostfix(postfix []string) (float64, error) {
	var stack []float64

	for _, token := range postfix {
		switch {
		case isNumber(token):
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, ErrInvalidExpression
			}
			stack = append(stack, num)
		case isOperator(token):
			if len(stack) < 2 {
				return 0, ErrInvalidExpression
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					return 0, ErrDivisionByZero
				}
				stack = append(stack, a/b)
			default:
				return 0, fmt.Errorf("Неизвестный оператор: %s", token)
			}
		default:
			return 0, ErrInvalidExpression
		}
	}

	if len(stack) != 1 {
		return 0, ErrInvalidExpression
	}
	return stack[0], nil
}

func isNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return 0
}
