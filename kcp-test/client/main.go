package main

import (
	"encoding/binary"
	"fmt"

	kcp "github.com/xtaci/kcp-go"
)

var (
	eventSequence = 0
)

func getSequence() int {
	eventSequence++
	return eventSequence
}

// func udppacket() {
// 	conn, err := net.Dial("udp", "10.42.0.1:19000")

// 	// conn.SetWriteDeadline(time.Now().Add(time.Second * 2))

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	fmt.Println("send udp data ...")

// 	for {
// 		len, err := conn.Write([]byte("hello baby"))
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			break
// 		}
// 		fmt.Println("send len ", len)

// 		var buf [1024]byte
// 		_, err = conn.Read(buf[:])
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			break
// 		}
// 	}

// }

func kcpRecvData(conn *kcp.UDPSession) {

	for {
		var buf [1024]byte
		len, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		//TODO 会粘包
		fmt.Println(len)
		fmt.Println(string(buf[8:len]))
	}
}

func kcpSendData(conn *kcp.UDPSession) {

	fmt.Println("send data ...")

	for {

		var sendBuf []byte = make([]byte, 8)
		binary.BigEndian.PutUint64(sendBuf[:], uint64(getSequence()))
		sendBuf = append(sendBuf, "hello baby"...)

		_, err := conn.Write(sendBuf)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		// fmt.Println("send len ", len)

	}
}

func main() {

	fmt.Println("kcp client demo")

	var serverAddr = "10.42.0.1:10000"

	kcpconn, err := kcp.DialWithOptions(serverAddr, nil, 10, 3)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("连接服务器成功", serverAddr)
	defer kcpconn.Close()

	kcpconn.SetWriteBuffer(1024)

	go kcpSendData(kcpconn)
	go kcpRecvData(kcpconn)

	select {}

}
