package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Msg struct {
	From     string
	Content  string
	SendTime string
}

func read(conn net.Conn) {
	for {
		res := make([]byte, 1024)
		n, err := conn.Read(res)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := res[:n]
		var msg Msg
		err2 := json.Unmarshal(result, &msg)
		if err2 != nil {
			fmt.Println(err2)
		}
		fmt.Printf("[%s]在%s说：%s\n", msg.From, msg.SendTime, msg.Content)
	}
}

func write(conn net.Conn) {
	for true {
		reader := bufio.NewReader(os.Stdin)
		content, _ := reader.ReadBytes('\n') //遇到换行
		conn.Write(content)
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	go read(conn)
	write(conn)
}
