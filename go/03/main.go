package main

import (
	"fmt"
	"time"
)

func main() {
	// var slice = make([]int, 3, 3)
	//

	// sliceA := []string{"ni", "wo"}
	// sliceB := []string{"hao", "ma"}
	// sliceC := append(sliceA, sliceB...)
	// fmt.Println(sliceC)
	// var slice []int
	// for i := 0; i < 10; i++ {
	// 	slice = append(slice, i)
	// 	fmt.Printf("%v,长度是%v,容量是%v\n", slice, len(slice), cap(slice))

	// }
	// s := "你好，golang"
	// r := []byte(s)
	// r[0] = "我"
	// fmt.Println(string(r))
	// var slice = []int{8, 5, 2, 9, 10, 4}
	// max := slice[0]
	// for i := 1; i < len(slice); i++ {
	// 	if slice[i] > max {
	// 		max = slice[i]
	// 	}
	// }
	// fmt.Print(max)
	// var slice = []int{8, 5, 2, 9, 10, 4}
	// for i := 0; i < len(slice); i++ {
	// 	for j := i + 1; j < len(slice); j++ {
	// 		if slice[i] > slice[j] {
	// 			slice[i], slice[j] = slice[j], slice[i]
	// 		}
	// 	}

	// }
	// fmt.Print(slice)
	// var userinfo = map[string]string{
	// 	"name": "陈涛",
	// 	"age":  "18",
	// 	"sex":  "女",
	// }
	// v, ok := userinfo["name"]
	// fmt.Printf("value=%v,ok=%v\n", v, ok)
	// for k, v := range userinfo {
	// 	fmt.Printf("key=%v,value=%v\n", k, v)
	// }
	// delete(userinfo, "sex")
	// fmt.Print(userinfo)
	// var userinfo = make(map[string][]string)
	// userinfo["name"] = []string{
	// 	"陈涛",
	// 	"18",
	// 	"女",
	// }
	// fmt.Println(userinfo)
	// map1 := make(map[int]int, 10)
	// map1[4] = 19
	// map1[3] = 67
	// map1[7] = 98
	// map1[1] = 23
	// map1[9] = 100

	// var keyslice []int
	// for k := range map1 {
	// 	keyslice = append(keyslice, k)
	// }
	// sort.Ints(keyslice)
	// for _, v := range keyslice {
	// 	fmt.Printf("key=%v,value=%v\n", v, map1[v])
	// }
	// str := "chentao shi ge da sha bi wo jue de ye shi"
	// strslices := strings.Split(str, " ")
	// var strMap = make(map[string]int)

	// for _, v := range strslices {
	// 	strMap[v]++

	// }
	// fmt.Println(strMap)
	// fmt.Println(fn1(10, 0))
	// fmt.Println(fn1(10, 2))

	// timeobj := time.Now()
	// unixtime := timeobj.Unix()
	// fmt.Println("当前时间的时间戳是：", unixtime)
	// fmt.Printf("%T\n", unixtime)
	// timeobj2 := time.Unix(unixtime, 0)
	// var str = timeobj2.Format("2006-01-02 15:04:05")
	// fmt.Println("时间字符串是：", str)
	ticker := time.NewTicker(2 * time.Second)
	n := 5
	for t := range ticker.C {
		n--
		fmt.Println("当前时间是：", t)
		if n <= 0 {
			ticker.Stop()
			break
		}
	}

}

// func fn1(a, b int) int {
// 	// defer func() {
// 	// 	err := recover()
// 	// 	if err != nil {
// 	// 		fmt.Println("给管理员发送邮件")
// 	// 	}
// 	// }()
// 	return a / b

// }
