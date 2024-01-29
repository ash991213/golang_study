package main

import "fmt"

func main() {
	// for_loop()
	// while_loop()
	// infinite_loop()
	// for_range()
	// loop_continue_break()
} 

func for_loop() {
	sum := 0

	for i := 0; i <= 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}

func while_loop() {
	sum := 1
	for sum < 100 {
		sum += sum
	}

	fmt.Println((sum))
}

func infinite_loop() {
	for {

	}
}

func for_range() {
	runes := []rune{'A', 'B', 'C'}

	for _, r := range runes {
		fmt.Println(r)
	}
}

func loop_continue_break() {
	numbers := []int{2, 4, 6, 8, 10, 11, 14}

	for _, num := range numbers {
		if num%2 == 0 {
			fmt.Println(num)
			continue
		}
		break
	}
}
