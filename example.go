package main

import (
	"fmt"
	"strings"

	"github.com/cwarden/souffle-go-example/SwigInterface"
)

func main() {
	p := SwigInterface.NewInstance("example")
	if p.Swigcptr() == 0 {
		panic("Souffle program not found")
	}
	p.LoadAll(".")
	edges := p.GetRelation("edge")

	edges.Purge()

	t := edges.NewTuple()
	t.PutInteger(4)
	t.PutInteger(5)
	t.Insert()

	t = edges.NewTuple()
	t.PutInteger(5)
	t.PutInteger(6)
	t.Insert()

	t = edges.NewTuple()
	t.PutInteger(1)
	t.PutInteger(4)
	t.Insert()

	p.Run()
	paths := p.GetRelation("path")

	arity := paths.GetArity()
	size := paths.Size()
	fmt.Printf("ARITY: %d\n", arity)
	fmt.Printf("TUPLES: %d\n", size)
	fmt.Printf("SIGNATURE: %s\n", paths.GetSignature())
	var i, j uint
	for i = 0; i < size; i++ {
		tuple := paths.Next()

		var elements []string
		for j = 0; j < arity; j++ {
			elements = append(elements, fmt.Sprintf("%d", tuple.GetInteger(j)))
		}

		fmt.Printf("<%s>\n", strings.Join(elements, ","))
	}
}
