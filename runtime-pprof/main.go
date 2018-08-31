// .\runtime-pprof.exe -cpuprofile profile

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var (

	//定义外部输入文件名字
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file.")
)

func doCall() {

	fmt.Println(" ")
}

func main() {

	log.Println("begin")

	flag.Parse()

	if *cpuprofile != "" {

		//创建性能分析文件
		f, err := os.Create(*cpuprofile)
		if err != nil {

			log.Fatal(err)

		}
		//开始分析cpu性能
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()

	}

	for i := 0; i < 30000; i++ {
		go doCall()
	}

	time.Sleep(time.Second * 5)

}
