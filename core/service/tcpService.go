package service

import (
	"bufio"
	"fmt"
	"net"
)

var TCPClients []net.Conn // 用于存储客户端连接的列表

func TCPService() {
	// 创建TCP服务器
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("TCP server is running on port 8080")

	// 接受客户端连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			return
		}

		// 将新连接加入客户端列表
		TCPClients = append(TCPClients, conn)

		// 处理客户端连接
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// 读取客户端消息
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}
		fmt.Println("Received message:", message)

		// 广播消息给所有连接的客户端
		broadcastMessage(message)
	}
}

func broadcastMessage(message string) {
	// 遍历所有连接的客户端，并发送消息
	for _, client := range TCPClients {
		_, err := client.Write([]byte(message))
		if err != nil {
			fmt.Println("Error broadcasting:", err.Error())
		}
	}
}
