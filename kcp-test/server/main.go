package main

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/xtaci/kcp-go"
)

func kcpRecv(conn *kcp.UDPSession) {

	for {

		var buf [1024]byte

		len, err := conn.Read(buf[:])

		if err != nil {
			break
		}

		var sequence = binary.BigEndian.Uint64(buf[0:8])
		fmt.Printf("%d,%s\n", sequence, string(buf[8:len]))

		_, err = conn.Write(buf[0:len])
		if err != nil {
			fmt.Println(err.Error())
			break
		}

	}

}

func main() {

	fmt.Println("kcp server demo")

	lis, err := kcp.ListenWithOptions(":10000", nil, 10, 3)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := lis.SetReadBuffer(1024); err != nil {
		log.Println("SetReadBuffer:", err)
	}
	if err := lis.SetWriteBuffer(1024); err != nil {
		log.Println("SetWriteBuffer:", err)
	}

	for {
		if conn, err := lis.AcceptKCP(); err == nil {
			log.Println("remote address:", conn.RemoteAddr())
			conn.SetStreamMode(true)
			conn.SetWriteDelay(true)

			fmt.Println("accept sucessed ", conn.RemoteAddr())

			go kcpRecv(conn)

		} else {
			log.Printf("%+v\n", err)
		}
	}

}
