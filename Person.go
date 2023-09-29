package main

type Person struct {
	leisure Leisure
}

func (p *Person) setLeisure(leisure Leisure) {
	p.leisure = leisure
}

func (p *Person) executeLeisure() {
	p.leisure.spendingTime(p)
}
