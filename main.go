package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 1
	for sum < 1000 {
		fmt.Println("1.Rock 2.Paper Or 3.Scissors?")

		// Read input line by line
		for scanner.Scan() {
			text := scanner.Text() // Get the current line of text
			if text == "" {
				break // Exit loop if an empty line is entered
			}
			fmt.Println("You entered:", text)
			i, err := strconv.Atoi(text)
    		if err != nil {
    		    // ... handle error
    		    panic(err)
    		}
			min := 1
			max := 4
			random := rand.Intn(max - min) + min
			fmt.Println(random)
			if i != random {
			fmt.Println("You lose :(")
			}else {
				fmt.Println("Congrats!!")
			}
			break
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error:", err)
		}
		sum += sum
	}

	
}