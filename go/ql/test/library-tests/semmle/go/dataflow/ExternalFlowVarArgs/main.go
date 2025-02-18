package main

import (
	"github.com/nonexistent/test"
)

func source() string {
	return "untrusted data"
}

func sink(string) {
}

func main() {
	s := source()
	sink(test.FunctionWithParameter(s)) // $ taintflow dataflow

	stringSlice := []string{source()}
	sink(stringSlice[0]) // $ taintflow dataflow

	s0 := ""
	s1 := source()
	sSlice := []string{s0, s1}
	sink(test.FunctionWithParameter(sSlice[1]))        // $ taintflow dataflow
	sink(test.FunctionWithSliceParameter(sSlice))      // $ taintflow dataflow
	sink(test.FunctionWithVarArgsParameter(sSlice...)) // $ taintflow dataflow
	sink(test.FunctionWithVarArgsParameter(s0, s1))    // $ taintflow dataflow

	sliceOfStructs := []test.A{{Field: source()}}
	sink(sliceOfStructs[0].Field) // $ taintflow dataflow

	a0 := test.A{Field: ""}
	a1 := test.A{Field: source()}
	aSlice := []test.A{a0, a1}
	sink(test.FunctionWithSliceOfStructsParameter(aSlice))      // $ taintflow dataflow
	sink(test.FunctionWithVarArgsOfStructsParameter(aSlice...)) // $ taintflow dataflow
	sink(test.FunctionWithVarArgsOfStructsParameter(a0, a1))    // $ taintflow dataflow
}
