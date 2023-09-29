package main

func main() {
	person := Person{}
	playing := &Playing{}
	reading := &Reading{}
	training := &Training{}

	person.setLeisure(playing)
	person.executeLeisure()

	person.setLeisure(reading)
	person.executeLeisure()

	person.setLeisure(training)
	person.executeLeisure()
}
