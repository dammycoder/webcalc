package main

import "fmt"

type Operator interface {
	Evaluate(first, second float64) float64
}

type AddOperator struct {
	Symbol string
}

func (a AddOperator) Evaluate(first, second float64) float64 {
	return first + second
}

type SubstractOperator struct {
	Symbol string
}

type DivOperator struct{
	Symbol string
}

func (a SubstractOperator) Evaluate(first, second float64) float64 {
	return first - second
}

type multiplicationOperator struct {
	Symbol string 
}

func (m multiplicationOperator) Evaluate(first, second float64) float64 {
	return first * second
}

func (d DivOperator) Evaluate(first,second float64)float64{
	return first/second
}

func process(first float64, op string, third float64) float64 {

	opsTable := map[string]Operator{
		"+": AddOperator{Symbol: "+"},
		"-": SubstractOperator{Symbol: "-"},
		"*": multiplicationOperator{Symbol: "*"},
		"/":DivOperator{Symbol:"/"},
	}

	opHandler, ok := opsTable[op]
	if !ok {
		panic(fmt.Sprintf("unspported operator %s", op))
	}

	return opHandler.Evaluate(first, third)

}
