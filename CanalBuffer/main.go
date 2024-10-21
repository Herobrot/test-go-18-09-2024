package main

import (
	"fmt"

	"main.go/models"
)

func main() {
	ch := make(chan bool, 3)

	go models.Escritor(ch)
	<-ch

	fmt.Println("Termino de programa")
}
