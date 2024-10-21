package main

import (
	"fmt"

	"main.go/models"
)

func main() {
	ch1 := make(chan string, 5)
	ch2 := make(chan string, 5)
	p := models.NewProd(ch1)
	c := models.NewCon(ch2)

	go p.Run()
	go c.Run()

	for i := 1; i <= 5; i++ {
		select {
		case value := <-ch1:
			fmt.Println(value)
		case value2 := <-ch2:
			fmt.Println(value2)
		}
	}
	p.SetFlag(false)
	c.SetFlag(false)
	fmt.Println("Exit")
}
