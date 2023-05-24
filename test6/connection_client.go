package main

import (
	"fmt"
	"net"
)

//发送交易信息到Ip

func conn_sendTx(ip string, data []byte) {
	// 连接到服务端建立的tcp连接
	conn, err := net.Dial("tcp", ip+":9888")
	// 输出当前建Dial函数的返回值类型, 属于*net.TCPConn类型
	fmt.Printf("客户端: %T\n", conn)
	if err != nil {
		// 连接的时候出现错误
		fmt.Println("err :", err)
		return
	}
	// 当函数返回的时候关闭连接
	defer conn.Close()

	_, err = conn.Write(data)
	if err != nil {
		return
	}
	buf := [512]byte{}
	//	// 读取服务端发送的数据
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("recv failed, err:", err)
		return
	}
	fmt.Println("客户端接收服务端发送的数据: ", string(buf[:n]))
}
