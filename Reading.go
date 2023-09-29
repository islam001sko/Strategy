package main

import "fmt"

type Reading struct {
}

func (r *Reading) spendingTime(p *Person) {
	fmt.Println("Reading some book")
}
