package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { //跨域检测
			return true
		},
	}
)

func wsHandler(rw http.ResponseWriter, rq *http.Request) {
	//rw.Write([]byte("hello"))

	var (
		websocketConn *websocket.Conn
		err           error
		messageType   int
		message       []byte
	)

	//websocket.TextMessage

	if websocketConn, err = upgrader.Upgrade(rw, rq, nil); err != nil {
		return
	}

	defer websocketConn.Close()

	fmt.Println("websocket 连接成功")

	for {
		if messageType, message, err = websocketConn.ReadMessage(); err != nil {
			fmt.Println("websocket接收失败: " + err.Error())
			return
		}
		if err = websocketConn.WriteMessage(messageType, message); err != nil {
			fmt.Println("websocket发送失败: " + err.Error())
			return
		}
	}

}

func main() {
	fmt.Println("服务器开始启动 ...")
	http.HandleFunc("/ws", wsHandler)
	err := http.ListenAndServe("0.0.0.0:8888", nil)
	if err != nil {
		fmt.Println("服务器启动失败: " + err.Error())
	}
}
