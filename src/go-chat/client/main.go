package main

import (
	"bufio"
    "fmt"
    "net"
	"os"
	"strings"
)

//定义通道
var ch chan int = make(chan int)

//定义昵称
var nickname string

func reader(conn *net.TCPConn) {
    buff := make([]byte, 128)
    for {
        j, err := conn.Read(buff)
        if err != nil {
            ch <- 1
            break
        }

        fmt.Println(string(buff[0:j]))
    }
}

func main() {

    tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
    conn, err := net.DialTCP("tcp", nil, tcpAddr)

    if err != nil {
        fmt.Println("Server is not starting")
        os.Exit(0)
    }

    defer conn.Close()

    go reader(conn)

    fmt.Println("请输入昵称")

    fmt.Scanln(&nickname)

    fmt.Println("你的昵称为:", nickname)

    for {
        var msg string
        // fmt.Scanln(&msg)
        // b := []byte("<" + nickname + ">" + "说:" + msg)
        // conn.Write(b)

		inputReader := bufio.NewReader(os.Stdin)
		msg ,_ =inputReader.ReadString('\n') 
		msg = strings.Replace(msg, "\n", "", -1)
		// fmt.Println(msg)
		b := []byte("<" + nickname + ">" + ":" + msg)
        conn.Write(b)
        //select 为非阻塞的
        select {
        case <-ch:
            fmt.Println("Server错误!请重新连接!")
            os.Exit(1)
        default:
            //不加default的话，那么 <-ch 会阻塞for， 下一个输入就没有法进行
        }
        msg ,_ =inputReader.ReadString('\n') 
    }
}
