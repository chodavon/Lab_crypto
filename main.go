package main

import "fmt"
import (
	"encoding/base64"
	"encoding/hex"
)

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
func Task4(c, d, ch int) {
	fmt.Println("========task4========")
	for {
		fmt.Println("\n===== Mini Calculator =====")
		fmt.Println("1) Add 2) Sub 3) Mul 4) Div 5) Mod 6) Exit")
		fmt.Print("Choose: ")
		fmt.Scan(&ch)
		if ch == 6 {
			fmt.Println("Exiting...")
			break
		}

		fmt.Print("Enter a: ")
		fmt.Scan(&c)
		fmt.Print("Enter b: ")
		fmt.Scan(&d)

		switch ch {
		case 1:
			fmt.Println("Result:", c+d)
		case 2:
			fmt.Println("Result:", c-d)
		case 3:
			fmt.Println("Result:", c*d)
		case 4:
			result := "Error: division by zero"
			switch c {
			case 0:
			default:
				result = fmt.Sprint(c / d)
			}
			fmt.Println("Result:", result)
		case 5:
			result := "Error: division by zero"
			switch c {
			case 0:
			default:
				result = fmt.Sprint(c % d)
			}
			fmt.Println("Result:", result)
		default:
			fmt.Println("Invalid choice!")
		}
	}
}
func task5(a, b int) {
	fmt.Println("========task5========")
	var text string
	fmt.Print("Enter text: ")
	fmt.Scanln(&text)
	data := []byte(text)
	hexResult := hex.EncodeToString(data)
	base64Result := base64.StdEncoding.EncodeToString(data)
	binaryResult := ""
	for _, b := range data {
		binaryResult += fmt.Sprintf("%08b ", b) // 8-bit binary
	}

	fmt.Println("\n--- Encodings ---")
	fmt.Println("Binary :", binaryResult)
	fmt.Println("Hexadecimal :", hexResult)
	fmt.Println("Base64 :", base64Result)
}

func xorEncrypt(text string, key byte) string {
	fmt.Println("========task6========")
	result := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		result[i] = text[i] ^ key // XOR each character with the key
	}
	return string(result)
}

func main() {
	fmt.Println("Hello world")
	var a, b int = 10, 5
	var c, d, ch int
	var text string
	var key byte
	task1(a, b)
	task2(a, b)
	task3(a, b)
	Task4(c, d, ch)
	task5(a, b)
	fmt.Print("Enter text: ")
	fmt.Scanln(&text)
	fmt.Print("Enter key (single character): ")
	fmt.Scanf("%c", &key)

	// Encrypt the text
	encrypted := xorEncrypt(text, key)
	fmt.Println("\nEncrypted:", encrypted)

	// Decrypt using the same function
	decrypted := xorEncrypt(encrypted, key)
	fmt.Println("Decrypted:", decrypted)
}
