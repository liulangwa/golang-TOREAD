/*

#ulimit -c unlimited
#echo core > /proc/sys/kernel/core_pattern
#env GOTRACEBACK=crash ./panic-to-file

*/

package main

import (
	"fmt"
)

func main() {

	fmt.Printf("panic 生成dumpcore文件")
	var slice = make([]byte, 20)
	slice[30] = 3 //越界异常
}
