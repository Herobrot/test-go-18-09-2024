package main

import (
	"fmt"

	"main.go/models"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	p := models.NewProd(ch1)
	c := models.NewCon(ch2)

	go p.Run()
	go c.Run()

	for i := 0; i < 2; i++ {
		select {
		case value := <-ch1:
			fmt.Println(value)
		case value2 := <-ch2:
			fmt.Println(value2)
		}
	}
	fmt.Println("Exit")
}
