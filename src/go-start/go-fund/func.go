package main

import "fmt"

// type typeName func(input1 inputType1, intpu2 inputType2 [, ...]) (result1 resultType1 [, ...])
// 可以把这个类型的函数当做值来传递

/*go中函数也是一种变量，tyep定义类型就是多余拥有相同的参数，相同返回值的一种类型*/

type testInt func(int) bool //函数类型

func isOdd(i int) bool {
	if i%2 == 0 {
		return false
	}
	return true
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range [...]int{1, 2, 3} {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3}
	odd := filter(slice, isOdd)
	fmt.Println("Odd: ", odd)
}
