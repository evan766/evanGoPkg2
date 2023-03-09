package main

import (
	"fmt"
	"github.com/evan766/evanGoPkg2/utils"
)

func main() {
	fmt.Println("main func of evanGoPkg2")

	a := []int{1, 2, 3, 4, 5}
	b := []int{10, 2, 3, 100, 20}

	c := utils.Intersect(a, b)

	fmt.Println(c)
}
