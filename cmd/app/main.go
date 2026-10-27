package main

import (
	"fmt"
	"math"

	"MyProject/gul"
	"MyProject/stringutils"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(gul.Fac(6))
	fmt.Println(math.Pow(4, 6))
	fmt.Println(gul.Polindrom(989))
	fmt.Println(stringutils.Length("Hello World"))
}
