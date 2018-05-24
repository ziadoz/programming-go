package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/ex2.2/conv"
)

func main() {
	if len(os.Args[1:]) > 0 {
		parseArgs(os.Args[1:])
	} else {
		fmt.Println("Input numbers. Press Ctrl + C to exit.")
		parseStdIn()
	}
}

func parseArgs([]string) {
	for _, arg := range os.Args[1:] {
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Println("Invalid argument: " + arg)
			os.Exit(1)
		}

		convert(num)
	}
}

func parseStdIn() {
	for true {
		var arg string
		if _, err := fmt.Scan(&arg); err != nil {
			continue
		}

		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Println("Invalid argument: " + arg)
			continue
		}

		convert(num)
	}
}

func convert(num float64) {
	{
		f := conv.Farenheit(num)
		c := conv.Celcius(num)

		fmt.Printf("%s = %s\n%s = %s\n\n", f, conv.FToC(f), c, conv.CToF(c))
	}
	{
		f := conv.Feet(num)
		m := conv.Metre(num)

		fmt.Printf("%s = %s\n%s = %s\n\n", f, conv.FToM(f), m, conv.MToF(m))
	}
	{
		p := conv.Pound(num)
		k := conv.Kilogram(num)

		fmt.Printf("%s = %s\n%s = %s\n\n", p, conv.PToK(p), k, conv.KToP(k))
	}
}
