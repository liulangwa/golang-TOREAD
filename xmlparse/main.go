package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cihub/seelog"
)

var VERSION = "version:1.01 "

var (
	gClient *client
)

type FusClientInfo struct {
	hostIP                 string
	Host                   string `xml:"host,attr"` //属性
	Port                   string `xml:"port,attr"`
	APIGetFusServerInfo    string `xml:"apiGetFusServerInfo,attr"`
	APIGetAllFusServerInfo string `xml:"apiGetAllFusServerInfo,attr"`
	APIAuthentication      string `xml:"apiAuthentication,attr"`
}

type client struct {
	XMLName                  xml.Name        `xml:"client"`
	FusClientInfos           []FusClientInfo `xml:"FusServerInfos>FusServerInfo"` //数组
	AuthenticationServerInfo string          `xml:"AuthenticationServerInfo"`     //值
}

/**
解析xml
*/
func unmarshalXML() error {
	file, err := os.Open("./client.xml")
	if err != nil {
		seelog.Warn(VERSION, "unmarshalXML Open error")
		return err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)

	fmt.Printf(string(data))

	if err != nil {
		seelog.Warn(VERSION, "unmarshalXML ReadAll error")
		return err
	}
	gClient = new(client)
	err = xml.Unmarshal(data, gClient)
	if err != nil {
		seelog.Warn(VERSION, "unmarshalXML Unmarshal error")
		return err
	}

	fmt.Println(gClient)

	return nil
}

func main() {

	//读取log配置
	logger, err := seelog.LoggerFromConfigAsFile("seelog.xml")
	if err != nil {
		seelog.Critical("err parsing config log file", err)
		return
	}
	seelog.ReplaceLogger(logger)
	defer seelog.Flush()
	//解析xml
	if nil == unmarshalXML() {
		seelog.Info("xml 解析成功")
	} else {
		seelog.Error("xml 解析失败")
	}
}
