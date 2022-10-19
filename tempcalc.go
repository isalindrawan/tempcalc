package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func celcius(value string) (fahrenheit float32, reamur float32, kelvin float32) {

	// cast string to integer
	c, _ := strconv.Atoi(value)

	// calculate and cast to float32
	fahrenheit = float32(((9 / 5) * c) + 32)
	reamur = float32((4 / 5) * c)
	kelvin = float32(c + 273)

	return
}

func reamur(value string) (fahrenheit float32, celcius float32, kelvin float32) {

	// cast string to integer
	r, _ := strconv.Atoi(value)

	// calculate and cast to float32
	fahrenheit = float32(((9 / 4) * r) + 32)
	celcius = float32((5 / 4) * r)
	kelvin = float32((5/4)*r + 273)

	return
}

func fahrenheit(value string) (celcius float32, reamur float32) {

	// cast string to integer
	f, _ := strconv.Atoi(value)

	// calculate and cast to float32
	celcius = float32((5 / 9) * (f - 32))
	reamur = float32((4 / 9) * (f - 32))

	return
}

func kelvin(value string) (celcius float32, reamur float32) {

	// cast string to integer
	k, _ := strconv.Atoi(value)

	// calculate and cast to float32
	celcius = float32(k - 273)
	reamur = float32((4 / 5) * (k - 273))

	return
}

func main() {

	var temp string

	// init flags
	flag.StringVar(&temp, "C", "", "Celcius value to be converted to other temperature")
	flag.StringVar(&temp, "F", "", "Fahrenheit value to be converted to other temperature")
	flag.StringVar(&temp, "R", "", "Reamur value to be converted to other temperature")
	flag.StringVar(&temp, "K", "", "Kelvin value to be converted to other temperature")

	flag.Parse()

	// check argument length from terminal, if more than 3 then error
	if len(os.Args) > 3 {

		fmt.Println("Only 1 flag allowed")

	} else {

		// switch mode
		switch os.Args[1] {
		case "-C":
			fahrenheit, reamur, kelvin := celcius(temp)
			fmt.Printf("Fahrenheit : %f\n", fahrenheit)
			fmt.Printf("Reamur : %f\n", reamur)
			fmt.Printf("Kelvin : %f\n", kelvin)
			break
		case "-R":
			fahrenheit, celcius, kelvin := reamur(temp)
			fmt.Printf("Fahrenheit : %f\n", fahrenheit)
			fmt.Printf("celcius : %f\n", celcius)
			fmt.Printf("Kelvin : %f\n", kelvin)
			break
		case "-F":
			celcius, reamur := fahrenheit(temp)
			fmt.Printf("Celcius : %f\n", celcius)
			fmt.Printf("Reamur : %f\n", reamur)
			fmt.Println("Kelvin : -")
			break
		case "-K":
			celcius, reamur := kelvin(temp)
			fmt.Println("Fahrenheit : -")
			fmt.Printf("Celcius : %f\n", celcius)
			fmt.Printf("Reamur : %f\n", reamur)
			break
		default:
			fmt.Println("Flag unknown")
			break
		}
	}
}
