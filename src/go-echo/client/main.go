package main

import (
    "net"
    "fmt"
    "os"
)

func main() {
    //主动连接服务器
    conn, err := net.Dial("tcp", "127.0.0.1:8000")
    if err != nil {
        fmt.Println("net.Dial err = ", err)
        return
    }

    //main 调用完毕,关闭连接
    defer conn.Close()

    go func() {
        buf := make([]byte, 2048)
        for {
            n, err := conn.Read(buf)
            if err != nil {
                fmt.Println("conn.Read err = ", err)
                return
            }
            addr := conn.RemoteAddr().String()
			fmt.Printf("from %s: 服务器应答消息内容: %s", addr, string(buf[:n])) //打印接收到的内容, 转换为字符串再打印
			if string(buf[:n]) == "EXIT" {
				fmt.Printf("Exit")
				break
			}
        }
    }()

    //获取键盘输入的内容
    str := make([]byte, 2048)
    for {
        n, err := os.Stdin.Read(str) //从键盘读取内容， 放在str
        if err != nil {
            fmt.Println("os.Stdin. err = ", err)
            return
        }

        if "exit" == string(str[:n-1]){
            return
        }

        //把输入的内容给服务器发送
        conn.Write(str[:n])
    }
}