package main

import "fmt"

type Training struct {
}

func (t *Training) spendingTime(p *Person) {
	fmt.Println("Training at gym")
}
