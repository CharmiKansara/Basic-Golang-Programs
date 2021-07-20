package main

import "fmt"

func Reverse(str string) string {
	if len(str) <= 0 {
		return str
	}
	temp := str[1:]
	firstchar := string([]rune(str)[0])
	defer fmt.Println("Function called ", len(temp)+1, "time")
	return Reverse(temp) + firstchar
}

func main() {
	var str string
	//Take user Input
	fmt.Println("Enter String to Reverse:")
	fmt.Scan(&str)

	//reversed string
	ReversedString := Reverse(str)
	fmt.Println(ReversedString)
}
