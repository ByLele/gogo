package main

/*
Panic 是一个内建函数，可以中断原有的控制流程，进入一个令人恐慌的流程中。
当 F 函数调用 panic，F 的执行逻辑被中断，但是 F 中的延迟函数会正常执行，然后 F 返回调用它的地方。
在调用的地方，F 的行为就像调用了 panic，继续向上直到发生 panic 的 goroutine 中所有调用的函数返回，此时程序退出。
panic 可以手动或者运行时错误产生（越界数组）

 recover：内建函数，可以让进入恐慌的流程中的 goroutine 回复过来。recover 仅在延迟函数中有效。
 正常逻辑调用 recover 返回 nil。当前的 goroutine 先入恐慌可以用 recover 捕获 panic 的输入值，并且回复正常的执行。
*/
import (
	. "fmt"
)

import (
	//f "fmt"
	_ "os"
)
import (
	_ "github.com/ziutek/mymysql/godrv" //_引入包，不直接使用包里面的函数，调用包里的init
)

func main() {
	Println(123)
}
