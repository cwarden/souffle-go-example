package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/cwarden/souffle-go-example/SwigInterface"
)

func main() {
	var mem runtime.MemStats
	for i := 0; i < 1000000; i++ {
		p := SwigInterface.NewInstance("example")
		if p.Swigcptr() == 0 {
			panic("Souffle program not found")
		}
		p.LoadAll(".")
		edges := p.GetRelation("edge")

		threads := p.GetNumThreads()
		fmt.Print("\033[H\033[2J")
		fmt.Printf("Theads: %d\n", threads)

		p.SetNumThreads(int64(runtime.NumCPU()))

		threads = p.GetNumThreads()
		fmt.Printf("Theads: %d\n", threads)

		edges.Purge()

		t := edges.NewTuple()
		t.PutInteger(4)
		t.PutInteger(5)
		t.Insert()
		SwigInterface.DeleteSWIGSouffleTuple(t)

		t = edges.NewTuple()
		t.PutInteger(5)
		t.PutInteger(6)
		t.Insert()
		SwigInterface.DeleteSWIGSouffleTuple(t)

		t = edges.NewTuple()
		t.PutInteger(1)
		t.PutInteger(4)
		t.Insert()
		SwigInterface.DeleteSWIGSouffleTuple(t)
		SwigInterface.DeleteSWIGSouffleRelation(edges)

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
			SwigInterface.DeleteSWIGSouffleTuple(tuple)

			_ = fmt.Sprintf("<%s>\n", strings.Join(elements, ","))
		}
		SwigInterface.DeleteSWIGSouffleRelation(paths)
		runtime.ReadMemStats(&mem)
		fmt.Printf("Alloc %d; Sys: %d\n", mem.Alloc, mem.Sys)
		SwigInterface.DeleteSWIGSouffleProgram(p)
	}
}
