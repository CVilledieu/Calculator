package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	welcomeMessage()
	var calc Calculator
	for scanner.Scan() {
		var curQuestion []string
		input := scanner.Text()
		scanInput := bufio.NewScanner(strings.NewReader(input))
		scanInput.Split(bufio.ScanWords)
		for scanInput.Scan() {
			eachPiece := scanInput.Text()
			curQuestion = append(curQuestion, eachPiece)
		}

		if len(curQuestion) == 1 {
			if curQuestion[0] == "quit" {
				return
			}
			if curQuestion[0] == "history" {
				fmt.Println(calc.History)
				continue
			}
		}

		calc.UpdateHistory(curQuestion)
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

type Calculator struct {
	History   [][]string
	CurrentEQ []string
	FirstVal  float64
	SecondVal float64
	Operator  string
}

func (c *Calculator) UpdateHistory(currentQuestion []string) {
	c.History = append(c.History, currentQuestion)
}

func (c *Calculator) Addition() float64 {
	return c.FirstVal + c.SecondVal
}

func (c *Calculator) Subtraction() float64 {
	return c.FirstVal - c.SecondVal
}

func (c *Calculator) Multiply() float64 {
	return c.FirstVal * c.SecondVal
}

func (c *Calculator) Divid() float64 {
	return c.FirstVal / c.SecondVal
}
