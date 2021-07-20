package main

import "fmt"

func TOH(n, A, C, B int) {

	if n > 0 {
		TOH(n-1, A, B, C)
		fmt.Println("Move disk ", n, "from ", A, "to ", C)
		TOH(n-1, B, C, A)
	}
}
func main() {
	//number of disks
	n := 4
	//1,2,3 indicate the towers
	TOH(n, 1, 3, 2)
}
