package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {

	var cos net.dnsConn
	cos.
	var ipAddr = "10.79.227.174"
	localAddr, err := net.ResolveIPAddr("ip", ipAddr) //客户端绑定的IP
	if err != nil {
		panic(err)
	}

	fmt.Println("绑定本地IP，地址为: ", ipAddr)

	fDialer := net.Dialer{
		LocalAddr: &net.UDPAddr{IP: localAddr.IP},
		Timeout:   5 * time.Second,
		KeepAlive: 5 * time.Second,
	}

	localResolver := &net.Resolver{Dial: fDialer.DialContext, PreferGo: true}

	fmt.Println("配置DNS代理")

	d := net.Dialer{
		LocalAddr: &net.TCPAddr{IP: localAddr.IP},
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		Resolver:  localResolver,
	}

	fmt.Println("配置HTTP代理")

	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial:  d.Dial,

		TLSHandshakeTimeout: 10 * time.Second,
	}

	client := &http.Client{Transport: tr}

	// client := http.Client{}
	url := "http://www.baidu.com"

	// http.Get() == http.NewRequest + client.DO()

	fmt.Println("请求URL ：", url)
	request, err := http.NewRequest("POST", url, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//处理返回结果
	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理

	file, err := os.Create("./index.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()

	// stdout := os.Stdout
	// _, err = io.Copy(stdout, response.Body)
	_, err = io.Copy(file, response.Body)

	//返回的状态码
	status := response.StatusCode

	fmt.Println("\nHTTP状态码: ", status)

}
