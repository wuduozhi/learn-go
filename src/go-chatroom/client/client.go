// one sever to more client chat room
//This is chat client
package main

import (
    "fmt"
	"net"
	"bufio"
	"strings"
	"os"
)

var nick string = ""  //声明聊天室的昵称

func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:8000")  //打开监听端口
    if err != nil {
        fmt.Println("conn fail...")
    }
    defer conn.Close()
    fmt.Println("client connect server successed \n")

    //给自己取一个聊天室的昵称
    fmt.Printf("Make a nickname:")
    fmt.Scanf("%s", &nick)  //输入昵称
    fmt.Println("hello : ", nick)  //客户端输出
    conn.Write([]byte("nick|" + nick))  //将信息发送给服务器端
	
	inputReader := bufio.NewReader(os.Stdin)
	_,_ =inputReader.ReadString('\n') 
	
	go Handle(conn)

    for {
		var msg string  //声明一个空的消息
		msg ,_ =inputReader.ReadString('\n') 
		msg = strings.Replace(msg, "\n", "", -1)

		if msg == "quit" {  //如果消息为quit
            conn.Write([]byte("quit|" + nick))  //将quit字节流发送给服务器端
            break  //程序结束运行
		}
		
		// fmt.Scan(&msg)  //输入消息
		to := ""
		fmt.Printf("Send to who:")
		fmt.Scan(&to)
		if to == "all"{
			conn.Write([]byte("say|" + nick + "|" + msg))  //三段字节流 say | 昵称 | 发送的消息
		}else{
			conn.Write([]byte("to|" + nick + "|" + to + "|" + msg))  //三段字节流 say | 昵称 | 发送的消息
		}
		
		msg ,_ =inputReader.ReadString('\n') 
    }
}

func Handle(conn net.Conn) {

    for {

        data := make([]byte, 255)  //创建一个字节流
        msg_read, err := conn.Read(data)  //将读取的字节流赋值给msg_read和err
        if msg_read == 0 || err != nil {  //如果字节流为0或者有错误
            break
        }

        fmt.Println(string(data[0:msg_read]))  //把字节流转换成字符串
    }
}