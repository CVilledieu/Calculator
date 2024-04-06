package main

type Calculator struct {
	CurrentValue float64
	SecondValue  float64
	Operator     string
}

func clearScreen() Calculator {
	return Calculator{
		CurrentValue: 0,
		SecondValue:  0,
		Operator:     "",
	}
}

type history = []Calculator

func doMath(c Calculator) Calculator {
	switch c.Operator {
	case "+":
		c.CurrentValue += c.SecondValue
	case "-":
		c.CurrentValue -= c.SecondValue

	case "*":
		c.CurrentValue = c.CurrentValue * c.SecondValue
	case "/":
		c.CurrentValue = c.CurrentValue / c.SecondValue
	}
	c.SecondValue = 0
	return c
}
