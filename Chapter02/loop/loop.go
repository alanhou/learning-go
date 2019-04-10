package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 死循环
func forever() {
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
		convertToBin(5),  // 101
		convertToBin(13), // 1101
		convertToBin(789798),
		convertToBin(0),
	)

	printFile("Chapter02/branch/abc.txt")

	s := `abc"d"
		kkkk
		123

		p`
	printFileContents(strings.NewReader(s))

	//forever()
}
