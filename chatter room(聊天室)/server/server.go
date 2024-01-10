package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type Msg struct {
	From     string
	Content  string
	SendTime string
}

var onlineClientMap = make(map[string]net.Conn)

func broadcast(content []byte) {
	// 向所有的客户端广播消息
	fmt.Println("客户端:::", onlineClientMap)
	for _, client := range onlineClientMap {
		client.Write(content)
	}
}

func handleConn(conn net.Conn) {
	// (1)在线注册用户
	addr := conn.RemoteAddr().String()
	onlineClientMap[addr] = conn

	//提示用户上线
	msg := Msg{
		From:     "系统消息",
		Content:  addr + "上线了!",
		SendTime: time.Now().String()[:19],
	}
	resByte, _ := json.Marshal(msg)
	broadcast(resByte)

	//循环读取消息，并写入
	for true {
		data := make([]byte, 1024)
		n, _ := conn.Read(data)
		msg := Msg{
			From:     addr,
			Content:  string(data[:n]),
			SendTime: time.Now().String()[:19],
		}
		if n == 0 {
			break
		}
		jsonByte, _ := json.Marshal(msg)
		broadcast(jsonByte)
	}
}

func main() {
	//起服务
	listent, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println("连接错误:", err)
	}
	defer listent.Close()
	// 开启循环，一直监听
	for true {
		fmt.Println("服务器监听客户端连接....")
		conn, _ := listent.Accept()
		fmt.Println("conn", conn)
		go handleConn(conn)
	}

}
