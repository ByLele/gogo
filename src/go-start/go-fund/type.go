package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZero Celsius = -273.15
	FreezingC    Celsius = 0
	Biiling      Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*95 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

/*go程序里每个包都通过一个唯一的 import path 标识，比如 "gopl.io/ch2/tempconv"
包初始化的顺序：以包级别的变量开始按照声明顺序初始化，但是被依赖的值先初始化*/
/*go tools 通过把文件名排序后给交给编译器 可以在文件里定义 init() 函数，
他们不能被引用，程序执行的时候会按照被声明顺序自动运行(比如初始化一个查找表)*/
func main() {
	fmt.Println(CToF(334))
	fmt.Println(FToC(31762))
}
