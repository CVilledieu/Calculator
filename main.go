package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	welcomeMessage()
	var eq []string
	for {
		scanner.Scan()
		//fmt.Printf("Current first input: %f, Current second input: %f \n", numInput[len((numInput))-2], numInput[len(numInput)-1])

		output := scanner.Text()
		eq = append(eq, output)
		// num, err := strconv.ParseFloat(output, 64)
		// if err != nil {
		if len(output) > 1 {
			if output == "quit" {
				return
			}
			// if output == "clear" {
			// 	numInput = [][]float64{0, 0}
			// }
		}

		// }
		// numInput = append(numInput, num)

		fmt.Println(eq[:])
	}
}

func welcomeMessage() {
	fmt.Println("Welcome!")
	help()
	fmt.Println("What do you want calculated?")
}

func help() {
	fmt.Println("supported operators: '+', '-', '*', or '/'")
	fmt.Println("Alt commands: 'clear', 'history', or 'quit'")
}
