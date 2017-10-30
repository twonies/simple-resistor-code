package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	// Set some Constants
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
	// set some of the variables we'll be using
	var x int
	var z int
	var y int
	var b float64
	var bands int
	var foband int
	var forth string
	check := 4
	//Start actually on the code itself.
	fmt.Println("RESISTOR COLOUR CODE FINDER V1.2")
	// Take in three inputs and set them to three variable strings first,second,third
	fmt.Println("How many bands are there on the resistor?")
	fmt.Scan(&bands)
	if bands == 3 || bands == 4 {
		fmt.Println("What's the resistor colour?")
		fmt.Println("First Colour")
		reader1 := bufio.NewReader(os.Stdin)
		first, _ := reader1.ReadString('\n')
		fmt.Println("Second Colour")
		reader2 := bufio.NewReader(os.Stdin)
		second, _ := reader2.ReadString('\n')
		fmt.Println("Third Colour")
		reader3 := bufio.NewReader(os.Stdin)
		third, _ := reader3.ReadString('\n')
		fmt.Println("Fourth Colour")
		if bands == 4 {
			reader4 := bufio.NewReader(os.Stdin)
			forth, _ := reader4.ReadString('\n')
		}

		// After inputs have been set look to see if they match any constants
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
			// if they do not meet any of the names then i sent a "troubleshooting" variable to 1
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
		if strings.TrimRight(forth, "\n") == "black" {
			foband = black
		} else if strings.TrimRight(forth, "\n") == "brown" {
			foband = brown
		} else if strings.TrimRight(forth, "\n") == "red" {
			foband = red
		} else if strings.TrimRight(forth, "\n") == "orange" {
			foband = orange
		} else if strings.TrimRight(forth, "\n") == "yellow" {
			foband = yellow
		} else if strings.TrimRight(forth, "\n") == "green" {
			foband = green
		} else if strings.TrimRight(forth, "\n") == "blue" {
			foband = blue
		} else if strings.TrimRight(forth, "\n") == "violet" {
			foband = violet
		} else if strings.TrimRight(forth, "\n") == "grey" {
			foband = grey
		} else if strings.TrimRight(forth, "\n") == "gray" {
			foband = grey
		} else if strings.TrimRight(forth, "\n") == "white" {
			foband = white
		} else {
			check = 5
		}
		// TIME TO GO TO THE THIRD COLOUR
		// in this i set b as a division variable to get the right eng notation ohms
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
		// all the const settings are done time to do math to find the resistance
		multipler := (math.Pow10(y))
		firstband := float64(x)
		secondband := float64(z)
		forthband := float64(foband)
		c := (firstband*10 + secondband + forthband) * multipler
		if bands == 4 {
			c * 10
		}
		// this sets so it will only do the math if you input all 3 colours correctly
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
			fmt.Print(" Multipler Colour is invalid")
		} else if check == 5 {
			fmt.Print("Colour 3 is invalid")
		}
	}
}
