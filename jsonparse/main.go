package main

import (
	"encoding/json"
	"fmt"
)

type TestObject struct {
	Delay string
}

//注册session
type JRegistSession struct {
	Cmd          string        `json:"cmd"`
	Seq          string        `json:"seq"`
	DstIP        string        `json:"dstip"`
	DstPort      string        `json:"dstport"`
	ProxySrcPort string        `json:"proxysrcport"`
	GUID         string        `json:"GUID"`
	TagID        string        `json:"tagID"`
	QueueTime    string        `json:"queueTime"` //本地和服务端标准一致,单位为毫秒
	Extends      []interface{} `json:"extends"`
}

func main() {

	var registerSession JRegistSession

	registerSession.Extends = append(registerSession.Extends, "delay=20")
	registerSession.Extends = append(registerSession.Extends, 300)
	registerSession.Extends = append(registerSession.Extends, TestObject{Delay: "3000"})

	dateMarshal, err := json.Marshal(registerSession)
	if err != nil {
		fmt.Println("Marsharl failed,err:", err.Error())
	}

	fmt.Println(string(dateMarshal))

	var registerSessionUnMarshal JRegistSession

	err = json.Unmarshal(dateMarshal, &registerSessionUnMarshal)
	if err != nil {
		fmt.Println("Marsharl failed,err:", err.Error())
	}

	fmt.Println(registerSessionUnMarshal.Extends[2])

}
