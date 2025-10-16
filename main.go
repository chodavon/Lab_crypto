package main

import "fmt"

func task1(a, b int) {
	fmt.Print("Enter first number (a): ")
	fmt.Scanln(&a)

	fmt.Print("Enter second number (b): ")
	fmt.Scanln(&b)

	fmt.Println("\n--- Checking Conditions ---")

	// Check if both numbers are positive
	bothPositive := (a > 0) && (b > 0)
	fmt.Println("Are both numbers positive?", bothPositive)

	// Check if one number is greater than the other
	oneGreater := (a > b) || (b > a)
	fmt.Println("Is one number greater than the other?", oneGreater)

	// Check if the numbers are not equal
	notEqual := !(a == b)
	fmt.Println("Are the two numbers different?", notEqual)
}

func task2(a, b int) {
	fmt.Println("========Task2=========")
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)

	fmt.Println("Both positive:", a > 0 && b > 0)
	fmt.Println("One greater:", a > b || b > a)
	fmt.Println("Not equal:", a != b)
}

func task3(a, b int) {
	fmt.Println("=========Task3========")
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)

	fmt.Println("\n--- Bitwise Operators ---")
	fmt.Println("a XOR b =", a^b)
	fmt.Println("NOT a =", ^a, ", NOT b =", ^b)
	fmt.Println("a OR b =", a|b)
	fmt.Println("a AND b =", a&b)
	fmt.Println("a << 1 =", a<<1)
	fmt.Println("a >> 1 =", a>>1)

	fmt.Println("\n--- Assignment Operators ---")
	c := a
	fmt.Println("c =", c)
	c += b
	fmt.Println("c += b →", c)
	c -= b
	fmt.Println("c -= b →", c)
	c *= b
	fmt.Println("c *= b →", c)
	if b != 0 {
		c /= b
		fmt.Println("c /= b →", c)
		c %= b
		fmt.Println("c %= b →", c)
	}
}

func main() {
	fmt.Println("Hello world")
	var a, b int = 10, 5
	task1(a, b)
	task2(a, b)
	task3(a, b)
}
