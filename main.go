package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ops []string = []string{"(", ")", "/", "*", "+", "-"}

func main() {
	runType := "cli"
	args := os.Args
	fmt.Println(args)
	if len(args) > 1 {
		runType = args[1]
	}

	switch runType {
	case "cli":
		cli()

	case "server":
		server()
	}
}

func cli() {
	fmt.Println("Enter the numbers")
	reader := bufio.NewReader(os.Stdin)
	store, er := reader.ReadString('\n')
	if er != nil {
		panic(er)
	}
	temp := strings.Split(store, "")
	cleanedParams := removeSpaces(temp)
	cleanedParams = parser(cleanedParams)
	fmt.Println(cleanedParams)
	Result := (BodmasFunc(cleanedParams))
	ResultConversion, err := strconv.ParseFloat(Result[0], 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("THE ANSWER IS:%g", ResultConversion)
}

func removeSpaces(i []string) []string {
	var output []string
	for _, v := range i {
		c := strings.Trim(v, "\r\n")
		if c != "" {
			output = append(output, c)
		}
	}

	return output
}

func parser(stringComponents []string) []string {
	var buffer []string
	var cache []string
	for _, token := range stringComponents {
		if !isOperator(token) {
			cache = append(cache, token)
		} else {
			if len(cache) > 0 {
				cachedItem := strings.Join(cache, "")
				buffer = append(buffer, cachedItem)
				cache = []string{}
			}
			buffer = append(buffer, token)
		}
	}
	if len(cache) > 0 {
		cachedItem := strings.Join(cache, "")
		buffer = append(buffer, cachedItem)
	}

	return buffer
}

func isOperator(v string) bool {
	for _, i := range ops {
		if i == v {
			return true
		}
	}
	return false
}
