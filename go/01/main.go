package main

import (
	"fmt"
)

func main() {
	// str1 := "\""
	// str2 := "helllo"
	// str3 := fmt.Sprintf("%v %v", str1, str2)
	// fmt.Println(str3)
	// str1 := "5-5-5"
	// arr := strings.Split(str1, "-")
	// fmt.Printf("%T\n", arr)
	// str2 := strings.Join(arr, "年")
	// Str3 := []string{"2023", "2024", "2025"}
	// fmt.Println(str2 + Str3[0])
	str1 := "你好 golang"
	for v, i := range str1 {
		fmt.Printf("%v%v(%c)", str1[v], i, i)
	}
}
