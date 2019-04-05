package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

// 整数转二进制
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string){
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

// 死循环
func forever(){
	for {
		fmt.Println("abc")
	}
}

func main() {
	//sum := 0
	//for i := 1; i <= 100; i++ {
	//	sum += i
	//}
	fmt.Println(
		convertToBin(5), // 101
		convertToBin(13), // 1101
		convertToBin(789798),
		convertToBin(0),
	)

	printFile("abc.txt")

	forever()
}
