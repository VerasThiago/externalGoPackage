package main

import "externalGoPackage/number"

func main() {
	println("Main call", number.Mult(10, 20, 30))
}
