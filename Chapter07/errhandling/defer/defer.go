package main

import (
	"bufio"
	"fmt"
	"learninggo/Chapter07/errhandling/fib"
	"os"
)

func tryDefer() {
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(3)
	//panic("error occured")
	//fmt.Println(4)

	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	//file, err := os.Create(filename)
	file, err := os.OpenFile(
		filename, os.O_EXCL|os.O_CREATE, 0666)

	//err = errors.New("this is a custom error")

	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		//fmt.Println("Error:", err.Error())
		return
		//panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()
	writeFile("fib.txt")
}
