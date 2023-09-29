package main

import "fmt"

type Playing struct {
}

func (playing *Playing) spendingTime(p *Person) {
	fmt.Println("Playing something")
}
