package models

import "fmt"

func Escritor(ch chan bool) {
	fmt.Println("Entre de esc")
	ch <- true
	ch <- true
	ch <- true
	ch <- true
	fmt.Println("Sali de esc")
}
