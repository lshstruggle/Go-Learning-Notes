package main

import "fmt"

func main() {
	var i int = 1
	var b float64 = 12.34
	str2 := fmt.Sprintf("%.1f", b)
	str1 := fmt.Sprintf("%d", i)
	fmt.Printf("%v-%T\n", str1, str1)
	fmt.Printf("%v-%T", str2, str2)

}
