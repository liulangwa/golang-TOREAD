/*
* 浏览器推送得例子，如果推送成功，会在console窗口打印‘消息被推送’
*https://rakyll.org/http2push/
 */

package main

import (
	"fmt"
	"log"
	"net/http"
)

const mainJS = `console.log("消息被推送");`

const indexHTML = `<html>
<head>
	<title>Hello</title>
	<script src="/main.js"></script>
</head>
<body>
</body>
</html>
`

func main() {
	http.HandleFunc("/main.js", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, mainJS)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		pusher, ok := w.(http.Pusher)
		if ok { // Push is supported. Try pushing rather than waiting for the browser.
			//服务器端推送消息
			if err := pusher.Push("/main.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		fmt.Fprintf(w, indexHTML)
	})
	// Run crypto/tls/generate_cert.go to generate cert.pem and key.pem.
	// See https://golang.org/src/crypto/tls/generate_cert.go
	// log.Fatal(http.ListenAndServeTLS(":7072", "cert.pem", "key.pem", nil))
	log.Fatal(http.ListenAndServe(":7072", nil))
}
