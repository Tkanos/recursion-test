package main

func main() {
	/*fmt.Println(fiboR(10))
	fmt.Println(fiboI(10))
	fmt.Println(fiboTail(10))*/
}

// Fibonacci in a recursive version
func fiboR(n int) int {
	if n < 2 {
		return n
	}
	return fiboR(n-2) + fiboR(n-1)
}

// Fibonacci in an iterative version
func fiboI(n int) int {
	var result int

	for i, first, second := 0, 0, 1; i <= n; i, first, second = i+1, first+second, first {
		if i == n {
			result = first
		}
	}

	return result
}

func fiboTail(n int) int {
	return fiboT(n, 0, 1)
}

//Fibonacci in a tail-recursive version
func fiboT(n, first, second int) int {
	if n == 0 {
		return first
	}

	return fiboT(n-1, second, first+second)
}

//http://codereview.stackexchange.com/questions/48020/four-algorithms-to-find-the-nth-fibonacci-number
//http://codereview.stackexchange.com/questions/28386/fibonacci-generator-with-golang/28445#28445
//http://www.golangpatterns.info/object-oriented/iterators
