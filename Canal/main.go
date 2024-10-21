package main

import (
	"fmt"
	"main/models"
	"sync"
)

func main() {
	ch := make(chan bool)
	wd := sync.WaitGroup{}
	//wd.Add(1)
	go models.Lector(ch, &wd)
	value := <-ch
	fmt.Println(value)
	//wd.Wait()
	fmt.Println("Fin de programa")

}
