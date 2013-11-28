package main

import "fmt"

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fib()
	// Function calls are evaluated left-to-right.
	fmt.Println(f(), f(), f(), f(), f(),f(),f(),f())
}



/*
a=b   ;
b=a+b ;
	a	b	p(a)
0	0	1	0
1	1	1	1
2	1	2	1
3	2	3	2
4	3	5	3
5	5	8	5
6	8	13	8
7	13	21	13
8	21	34	21

*/
