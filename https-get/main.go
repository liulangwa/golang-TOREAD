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

	localAddr, err := net.ResolveIPAddr("ip", "10.160.71.111") //客户端绑定的IP
	if err != nil {
		panic(err)
	}

	localTCPAddr := net.TCPAddr{
		IP: localAddr.IP,
	}

	d := net.Dialer{
		LocalAddr: &localTCPAddr,
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	tr := &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		Dial:                d.Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	client := &http.Client{Transport: tr}

	// client := http.Client{}
	url := "https://wx.qq.com/"

	// http.Get() == http.NewRequest + client.DO()
	request, err := http.NewRequest("GET", url, nil) //构造请求

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//设置header
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.87 Safari/537.36")
	request.Header.Add("Connection", "keep-alive")

	response, err := client.Do(request) //发送请求，返回处理结果
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer response.Body.Close()

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

	fmt.Println(status)

}
