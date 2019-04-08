package main

import (
	"fmt"
	"net"
	"strings"
)

var ConnMap map[string]net.Conn = make(map[string]net.Conn)

func main(){
	listen_socket,err := net.Listen("tcp","127.0.0.1:8000")

	if err != nil {
		fmt.Println("server start error")
	}

	defer listen_socket.Close()

	fmt.Println("server is waiting ...")

	for{
		conn,err := listen_socket.Accept()
		if err != nil {
			fmt.Println("conn fail ...")
		}

		fmt.Println(conn.RemoteAddr()," connect successed")

		go handle(conn)
	}

}

func handle(conn net.Conn){
	for {
		data := make([]byte,255)
		msg_read,err := conn.Read(data)

		if msg_read == 0 || err != nil {
			continue
		}

		msg_str := strings.Split(string(data[0:msg_read]),"|")

		switch msg_str[0] {
		case "nick": //加入聊天室
			fmt.Println(conn.RemoteAddr,"-->",msg_str[1])

			for k,v := range ConnMap{
				if k != msg_str[1]{
					v.Write([]byte("[" + msg_str[1] + "]: join..."))
				}
			}
			ConnMap[msg_str[1]] = conn
			case "say":   //转发消息
			for k, v := range ConnMap {  //k指客户端昵称   v指客户端连接服务器端后的地址
				if k != msg_str[1] {  //判断是不是给自己发，如果不是
					fmt.Println("Send "+msg_str[2]+" to ", k)  //服务器端将消息转发给集合中的每一个客户端
					v.Write([]byte("[" + msg_str[1] + "]: " + msg_str[2]))  //给除了自己的每一个客户端发送自己之前要发送的消息
				}
			}
		case "quit":  //退出
			for k, v := range ConnMap {  //遍历集合中的客户端昵称
				if k != msg_str[1] {  //如果昵称不是自己
					v.Write([]byte("[" + msg_str[1] + "]: quit"))  //给除了自己的其他客户端昵称发送退出的消息，并使Write方法阻塞
				}
			}
			delete(ConnMap, msg_str[1])  //退出聊天室

		}

	}
}