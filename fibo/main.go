package fibo

import "fmt"

func fibonacciMultiple() func() int {

	current, next := 0, 1
	return func() int {

		current, next = next, current+next
		return current
	}
}

func fibonacciSingle() func() int {

	current, next := 0, 1
	return func() int {

		// first phase, evaluation, left-to-right
		t1 := next
		t2 := current + next

		// second phase, assignmemt, left-to-right
		current = t1
		next = t2

		return current
	}
}

func main() {
	// multiple
	m := fibonacciMultiple()

	// single
	s := fibonacciSingle()

	fmt.Println(m(), m(), m(), m(), m(), m())
	fmt.Println(s(), s(), s(), s(), s(), s())
}
