package calculation

import (
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
	var curr strings.Builder
	runes := []rune(expr)

	for i := 0; i < len(runes); i++ {
		c := runes[i]
		if c == ' ' {
			continue
		}

		if c == '-' {
			if i == 0 {
				curr.WriteRune(c)
				continue
			}
			last := ""
			if len(tokens) > 0 {
				last = tokens[len(tokens)-1]
			}
			if isOperator(last) || last == "(" {
				curr.WriteRune(c)
				continue
			}
			if curr.Len() > 0 {
				tokens = append(tokens, curr.String())
				curr.Reset()
			}
			tokens = append(tokens, string(c))
			continue
		}

		switch c {
		case '+', '*', '/', '(', ')':
			if curr.Len() > 0 {
				tokens = append(tokens, curr.String())
				curr.Reset()
			}
			tokens = append(tokens, string(c))
		default:
			curr.WriteRune(c)
		}
	}
	if curr.Len() > 0 {
		tokens = append(tokens, curr.String())
	}
	return tokens
}

func infixToPostfix(tokens []string) ([]string, error) {
	var out, ops []string
	for _, t := range tokens {
		switch {
		case isNumber(t):
			out = append(out, t)
		case t == "(":
			ops = append(ops, t)
		case t == ")":
			found := false
			for len(ops) > 0 {
				top := ops[len(ops)-1]
				ops = ops[:len(ops)-1]
				if top == "(" {
					found = true
					break
				}
				out = append(out, top)
			}
			if !found {
				return nil, ErrInvalidExpression
			}
		case isOperator(t):
			for len(ops) > 0 && precedence(ops[len(ops)-1]) >= precedence(t) {
				out = append(out, ops[len(ops)-1])
				ops = ops[:len(ops)-1]
			}
			ops = append(ops, t)
		default:
			return nil, ErrInvalidExpression
		}
	}
	for len(ops) > 0 {
		top := ops[len(ops)-1]
		ops = ops[:len(ops)-1]
		if top == "(" {
			return nil, ErrInvalidExpression
		}
		out = append(out, top)
	}
	return out, nil
}

func evaluatePostfix(postfix []string) (float64, error) {
	var stack []float64
	for _, token := range postfix {
		switch {
		case isNumber(token):
			val, e := strconv.ParseFloat(token, 64)
			if e != nil {
				return 0, ErrInvalidExpression
			}
			stack = append(stack, val)
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
				return 0, ErrInvalidExpression
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

func isNumber(t string) bool {
	_, err := strconv.ParseFloat(t, 64)
	return err == nil
}

func isOperator(t string) bool {
	return t == "+" || t == "-" || t == "*" || t == "/"
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
