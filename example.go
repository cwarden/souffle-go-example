package main

import "github.com/cwarden/souffle-go-example/SwigInterface"

func main() {
	p := SwigInterface.NewInstance("example")
	p.LoadAll(".")
	p.Run()
	p.PrintAll(".")
}
