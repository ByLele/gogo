package main

/*interface 的函数参数：
interface的变量可以持有任意实现该interface 类型的对象。只要实现接口，我们就可以传入
*/
import (
	"fmt"
	"strconv"
)

type Element interface {
}

type List []Element
type person struct {
	name string
	age  int
}

func (p person) String() string {
	return "(name:" + p.name + "-age" + strconv.Itoa(p.age) + "years)"
}
func main() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "hello"
	list[2] = person{"tom", 23}

	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("int %d %d \n", index, element)
		case string:
			fmt.Printf("string %d %d\n", index, value)
		case person:
			fmt.Printf("Person %d %d \n", index, value)
		default:
			fmt.Printf("list different type %d", index)
		}
	}
}
