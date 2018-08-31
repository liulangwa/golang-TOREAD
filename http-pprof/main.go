/*
http://127.0.0.1:9876/debug/pprof/

*/

package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {

}

func main() {

	http.HandleFunc("/test", handler)
	log.Fatal(http.ListenAndServe(":9876", nil))

	for {
		fmt.Println("pprof test ")
		time.Sleep(time.Second)
	}
}
