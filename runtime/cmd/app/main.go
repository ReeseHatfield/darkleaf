package main

import (
	"fmt"

	"github.com/ReeseHatfield/runtime/core"
)

func main() {
	fmt.Println("Runtime application running")

	c := core.NewCore()

	fmt.Println(c.GetKey())
}
