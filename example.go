package main

import "github.com/cwarden/souffle-go-example/SwigInterface"

func main() {
	p := SwigInterface.NewInstance("example")
	if p.Swigcptr() == 0 {
		panic("Souffle program not found")
	}
	p.LoadAll(".")
	p.Run()
	p.PrintAll(".")
}
