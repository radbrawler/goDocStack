package intro

import "fmt"

type GoMap map[int]string

func main() {
	helloWorld()
	fizzBuzz()
}

func helloWorld() {
	fmt.Print("Hello World")
}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		s := ""       // Empty string
		if i%3 == 0 { // divisible by 3: add "Fizz"
			s += "Fizz"
		}
		if i%5 == 0 { // divisible by 5: append "Buzz"
			s += "Buzz"
		}
		if s != "" {
			fmt.Println(s) // print our FizzBuzz string
		} else {
			fmt.Println(i) // print the number
		}
	}
}
