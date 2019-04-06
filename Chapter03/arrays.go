package main

import "fmt"

func printArray(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int // 声明数组
	arr2 := [3]int{1, 3, 5}  // 声明数组并赋值
	arr3 := [...]int{2, 4, 6, 8, 10} // 不输入数组长度，让编译器来计算长度
	var grid [4][5]int // 二维数组
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	//for i := 0; i < len(arr3); i++ {
	//	fmt.Println(arr3[i])
	//}

	// 遍历数组，通过下标来获取值
	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	// 遍历数组，获取下标和值
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	// 遍历数组，只打印值
	for _, v := range arr3 {
		fmt.Println(v)
	}


	//numbers := [6]int{1, 3, 2, 5, 8, 4}
	//for i := 0; i < len(numbers); i++ {
	//	fmt.Println(numbers[i])
	//}
	//
	//maxi := -1
	//maxValue := -1
	//for i, v := range numbers {
	//	if v > maxValue {
	//		maxi, maxValue = i, v
	//	}
	//}
	//fmt.Println(maxi, maxValue)
	fmt.Println("printArray(arr1):")
	printArray(&arr1)

	fmt.Println("printArray(arr3):")
	printArray(&arr3)

	fmt.Println("fmt.Println(arr1, arr3):")
	fmt.Println(arr1, arr3)
}
