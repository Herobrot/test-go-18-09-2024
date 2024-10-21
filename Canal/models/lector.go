package models

import (
	"fmt"
	"sync"
)

func Lector(ch chan bool, wd *sync.WaitGroup) {
	//defer wd.Done()
	fmt.Println("Empece en Lector")
	ch <- true
	fmt.Println("Termine Lector")
}
