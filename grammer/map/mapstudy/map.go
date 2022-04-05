package main

import "fmt"

func map1() {
	//定义map，字典，类似java的hashmap，py的dict
	// 只声明  map[key_type]value_type
	hashm1 := make(map[string]int)
	fmt.Printf("hashm1: %v\n", hashm1)
	// 声明且初始化
	hashm2 := map[string]string{
		"Sam":   "Male",
		"Alice": "Female",
	}
	hashm1["Tom"] = 18

	// 遍历map
	for i, val := range hashm2 {
		fmt.Printf("%v is %v\n", i, val)
	}

	fmt.Printf("hashm1: %v\n", hashm1)
	fmt.Printf("hashm2: %v\n", hashm2)

	fmt.Println(len(hashm2))
}

func map2() {
	// 对数组遍历创建哈希表

	// 数组长度为9,但是遍历了十次
	arr1 := [10]int{5, 6, 7, 1, 3, 6, 3, 9, 15}
	m1 := make(map[int]int, cap(arr1))
	for _, val := range arr1 {
		if _, ok := m1[val]; ok {
			m1[val] += 1
		} else {
			m1[val] = 1
		}
	}
	for i, val := range m1 {
		fmt.Printf("%v %v\n", i, val)
	}
}

func main() {
	// map1()
	map2()
}
