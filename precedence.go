package main

import "fmt"

func BodmasFunc(x []string) []string {
	var math_expr []string
	var Eval_math_expr float64
	var temp_value bool
	var temp int
	temp_value = true
	for temp_value {

		if len(x) == 1 {
			temp_value = false
			break
		}
		for _, value := range ops {
			temp = (contains(x, value))
			temp_value = tocheck(temp)
			if temp_value && value == "(" || value == ")" {
				x = BracketFunc(x)
				continue
			}

			if temp_value {
				math_expr = append(math_expr, x[temp-1])
				math_expr = append(math_expr, x[temp])
				math_expr = append(math_expr, x[temp+1])
				Eval_math_expr = longEquation(math_expr)

				nbuf := []string{}
				nbuf = append(nbuf, x[:temp-1]...)
				nbuf = append(nbuf, fmt.Sprintf("%g", Eval_math_expr))
				nbuf = append(nbuf, x[temp+2:]...)
				fmt.Println(nbuf)
				if len(nbuf) > 1 {
					nbuf = BodmasFunc(nbuf)
				}
				x = nbuf
			}
		}
	}
	return x

}

func contains(s []string, str string) int {
	for p, v := range s {
		if p%2 == 1 {
			if v == str {
				return p
			}
		}
	}

	return -1
}

func tocheck(m int) bool {
	if m > -1 {
		return true
	}
	return false
}

func brackets(s []string, str string) int {
	for p, v := range s {
		if v == str {
			return p
		}
	}
	return -1
}

func openingbrace(s []string,str string) int {
	var k int = -1
	for p, v := range s{
		if v== str{
			k = p
		}
	}
	return k
}


func BracketFunc(x []string) []string {
	if len(x) == 1 {
		return x
	}

	OpenBracene_index:= openingbrace(x, "(")
	Bracket_Checker := tocheck(OpenBracene_index)
	if Bracket_Checker {
		var y []string
		y = append(y, x[OpenBracene_index+1:]...)
		check := brackets(y, ")")
		Bracket_Checker = tocheck(check)
		if Bracket_Checker {
			var store_mathexpr []string
			store_mathexpr = append(store_mathexpr, y[:check]...)
			store_mathexpr= BodmasFunc(store_mathexpr)
			y = []string{}
			y = append(y, x[:OpenBracene_index]...)
			y = append(y, store_mathexpr...)
			Length_Check:= brackets(x, ")")
			if len(x)-1 != Length_Check {
				y = append(y, x[Length_Check+1:]...)
				x = y

			} else {
				x = y
			}
		}

		if !Bracket_Checker{
			panic("Invalid Operation, does not have a closing brace")
		}

	}

	if brackets(x, "(") > -1 {
		return BracketFunc(x)
	}

	return x
}
