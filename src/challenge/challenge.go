package main

import "fmt"

const boilingK float64 = 373.2

func main() {

	tempK := boilingK
	tempC := (tempK - 32) / 1.8

	fmt.Printf("The boiling temperature of water in ºC is %.2f", tempC)

}
