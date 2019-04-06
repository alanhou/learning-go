package Chapter02

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func eval(a, b int, op string) (int ,error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		//return a / b
		q, _ := div(a, b) // 下划线表示接收的变量无需使用
		return q, nil
	default:
		//panic("unsupported operation: " + op)
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}
}

// 函数返回多个值 带余除法：13 / 3 = 4 ... 1
//func div(a, b int) (int, int) {
//	return a / b, a % b
//}
func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

//func apply(op func(int, int) int, a, b int) int {
//	p := reflect.ValueOf(op).Pointer()
//	opName := runtime.FuncForPC(p).Name()
//	fmt.Printf("Calling function %s with args " +
//		"(%d, %d)\n", opName, a, b)
//	return op(a, b)
//}

func pow(a, b int) int{
	return int(math.Pow(float64(a), float64(b)))
}

func apply(op func(int, int) int, a, b int) int{
	fmt.Printf("Calling %s with %d, %d\n",
		runtime.FuncForPC(reflect.ValueOf(op).Pointer()).Name(),
		a, b)
	return op(a, b)
}

func sumArgs(values ...int) int {
	sum := 0
	for i := range values{
		sum += values[i]
	}
	return sum
}

//func swap(a, b *int){
//	*b, *a = *a, *b
//}
func swap(a, b int) (int, int){
	return b, a
}

func main() {
	//fmt.Println(eval(3, 4, "x"))
	if result, err := eval(3, 4, "x"); err != nil{
		fmt.Println("Error:", err)
	}else{
		fmt.Println(result)
	}
	//fmt.Println(div(13, 3))
	q, r := div(13, 3)
	fmt.Println(q, r)

	//fmt.Println(apply(pow, 3, 4))
	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Println(sumArgs(1, 2, 3, 4, 5))
	a, b := 3, 4
	//swap(&a, &b)
	a, b = swap(a, b)
	fmt.Println(a, b)
}

