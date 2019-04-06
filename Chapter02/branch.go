package Chapter02

import (
	"io/ioutil"
	"fmt"
)

//func bounded(v int) int {
//	if v > 100 {
//		return 100
//	} else if v < 0 {
//		return 0
//	} else {
//		return v
//	}
//}

//func eval(a, b int, op string) int {
//	var result int
//	switch op {
//	case "+":
//		result = a + b
//	case "-":
//		result = a - b
//	case "*":
//		result = a * b
//	case "/":
//		result = a / b
//	default:
//		panic("unsupported operator:" + op)
//	}
//	return result
//}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	const filename = "abc.txt"
	//contents, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Print(err)
	//} else {
	//	fmt.Printf("%s\n", contents)
	//}

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	}else{
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
		grade(101),
	)
}
