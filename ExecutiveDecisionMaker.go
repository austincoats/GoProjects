package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var (
		number int
	)
	fmt.Println("What are you looking to do?")
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		rand.Seed(time.Now().UnixNano())
		number = rand.Intn(2)
		if number == 1 {
			fmt.Println("Yes, you should.\n")

		} else {
			fmt.Println("No, you should not.\n")
		}
		fmt.Println("What are you looking to do?")
	}
}
