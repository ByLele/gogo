package main

import (
	"fmt"
	"unicode/utf8"
)

/*
package
int8, int16, int32, int64, uint8, uint16, uint32, uint64 rune <=> int32 , byte <=> uint8
字符串字面量: "Hello，世界" (双引号) 。go 源文件总是用 utf8编码，go 文本字符串被解释成 utf8。
raw string literal: 反引号包含的字符串，转义符不会被处理，经常用来写正则
unicode: go 术语里叫做 rune(int32)
*/
func main() {
	s := "hello 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s)) //
}
