package main

func ShuntingYard(tokens []string) Queue {
	stack := Stack{}
	queue := Queue{}

	for len(tokens) > 0 {

		var token string
		token, tokens = tokens[0], tokens[1:]

		switch {

		case isNum.MatchString(token):
			queue.Enqueue(token)

		case token == "+", token == "*":
			for !stack.Empty() && stack.Top() == "+" {
				queue.Enqueue(stack.Pop())
			}
			stack.Push(token)

		case token == "(":
			stack.Push(token)

		case token == ")":
			for stack.Top() != "(" {
				queue.Enqueue(stack.Pop())
			}
			_ = stack.Pop() // discard left parenthesis
		}
	}

	for !stack.Empty() {
		queue.Enqueue(stack.Pop())
	}

	return queue
}
