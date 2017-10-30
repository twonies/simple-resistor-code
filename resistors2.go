package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	const black = 0
	const brown = 1
	const red = 2
	const orange = 3
	const yellow = 4
	const green = 5
	const blue = 6
	const violet = 7
	const grey = 8
	const white = 9
	var x int
	var z int
	var y int
	var b float64
	check := 4
	fmt.Println("RESISTOR COLOUR CODE FINDER V1.1")

	fmt.Println("What's the resistor colour?")

	reader1 := bufio.NewReader(os.Stdin)
	first, _ := reader1.ReadString('\n')
	reader2 := bufio.NewReader(os.Stdin)
	second, _ := reader2.ReadString('\n')
	reader3 := bufio.NewReader(os.Stdin)
	third, _ := reader3.ReadString('\n')

	if strings.TrimRight(first, "\n") == "black" {
		x = black
	} else if strings.TrimRight(first, "\n") == "brown" {
		x = brown
	} else if strings.TrimRight(first, "\n") == "red" {
		x = red
	} else if strings.TrimRight(first, "\n") == "orange" {
		x = orange
	} else if strings.TrimRight(first, "\n") == "yellow" {
		x = yellow
	} else if strings.TrimRight(first, "\n") == "green" {
		x = green
	} else if strings.TrimRight(first, "\n") == "blue" {
		x = blue
	} else if strings.TrimRight(first, "\n") == "violet" {
		x = violet
	} else if strings.TrimRight(first, "\n") == "grey" {
		x = grey
	} else if strings.TrimRight(first, "\n") == "gray" {
		x = grey
	} else if strings.TrimRight(first, "\n") == "white" {
		x = white
	} else {
		check = 1
		// TIME TO GO TO THE SECOND COLOUR
	}
	if strings.TrimRight(second, "\n") == "black" {
		z = black
	} else if strings.TrimRight(second, "\n") == "brown" {
		z = brown
	} else if strings.TrimRight(second, "\n") == "red" {
		z = red
	} else if strings.TrimRight(second, "\n") == "orange" {
		z = orange
	} else if strings.TrimRight(second, "\n") == "yellow" {
		z = yellow
	} else if strings.TrimRight(second, "\n") == "green" {
		z = green
	} else if strings.TrimRight(second, "\n") == "blue" {
		z = blue
	} else if strings.TrimRight(second, "\n") == "violet" {
		z = violet
	} else if strings.TrimRight(second, "\n") == "grey" {
		z = grey
	} else if strings.TrimRight(second, "\n") == "gray" {
		z = grey
	} else if strings.TrimRight(second, "\n") == "white" {
		z = white
	} else {
		check = 2
	}
	// TIME TO GO TO THE THIRD COLOUR
	if strings.TrimRight(third, "\n") == "black" {
		y = black
		b = 1
	} else if strings.TrimRight(third, "\n") == "brown" {
		y = brown
		b = 1
	} else if strings.TrimRight(third, "\n") == "red" {
		y = red
		b = 1000
	} else if strings.TrimRight(third, "\n") == "orange" {
		y = orange
		b = 1000
	} else if strings.TrimRight(third, "\n") == "yellow" {
		y = yellow
		b = 1000
	} else if strings.TrimRight(third, "\n") == "green" {
		y = green
		b = 1000000
	} else if strings.TrimRight(third, "\n") == "blue" {
		y = blue
		b = 1000000
	} else if strings.TrimRight(third, "\n") == "violet" {
		y = violet
		b = 100000000
	} else if strings.TrimRight(third, "\n") == "grey" {
		y = grey
		b = 100000000
	} else if strings.TrimRight(third, "\n") == "gray" {
		y = grey
		b = 100000000
	} else if strings.TrimRight(third, "\n") == "white" {
		y = white
		b = 100000000
	} else {
		check = 3
	}
	a := (math.Pow10(y))
	s := float64(x)
	d := float64(z)
	c := (s*10 + d) * a
	if check == 4 {
		fmt.Print("Your resistance is ")
		fmt.Printf("%.0f", c)
		fmt.Print(" ")
		fmt.Println(" Ohms")
		eng := (c / b)
		if y >= 2 {
			fmt.Println(" or ")
			fmt.Printf("%.1f", eng)
			if y >= 2 && y < 5 {
				fmt.Print(" Kilo Ohms")
			} else if y >= 5 && y < 8 {
				fmt.Print(" Mega Ohms")
			} else if y >= 8 && y < 12 {
				fmt.Print(" Giga Ohms")
			}
		}
	} else if check == 1 {
		fmt.Print("Colour 1 is invalid")
	} else if check == 2 {
		fmt.Print("Colour 2 is invalid")
	} else if check == 4 {
		fmt.Print("Colour 3 is invalid")
	}
}
