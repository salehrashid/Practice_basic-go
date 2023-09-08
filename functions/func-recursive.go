package main

func main() {
	//5 * 4 * 3 * 2 * 1
	println(factorialLoop(5))
	println(factorialLoopRecursive(5))

	println(doFactorial(5))
}

/** factorial loop without recursive **/
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

/** factorial loop with recursive **/
func factorialLoopRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialLoopRecursive(value-1)
	}
}

func doFactorial(value int) int {
	if value == 1 {
		return value
	}
	return value * doFactorial(value-1)
}
