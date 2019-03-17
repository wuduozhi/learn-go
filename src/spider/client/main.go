package main

import (
    "fmt"
    "os"
    "context"
    "github.com/apache/thrift/lib/go/thrift"
    "gen-go/spider/bks"
)

const (
    HOST = "localhost"
    PORT = "9090"
)

func main() {
    socket,err := thrift.NewTSocket("localhost:9090")

    transport := thrift.NewTBufferedTransport(socket, 8192)
    // serialize
    protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

    if err != nil {
        fmt.Fprintln(os.Stderr, "Error resolving address:", err)
        os.Exit(1)
    }
    
    
    client := bks.NewBksClientFactory(transport, protocolFactory)
    
    

    if err := transport.Open(); err != nil {
        fmt.Fprintln(os.Stderr, "Error opening socket to "+HOST+":"+PORT, " ", err)
        os.Exit(1)
    }

    defer socket.Close()

	xn := "2018"
	xq := "1"
	stuid := "201626010520"
	password := "WudUozhI"

    data,err := client.GetGrade(context.Background(),xn,xq,stuid,password)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error func client:", err)
    }

    data,err = client.GetClassTable(context.Background(),xn,xq,stuid,password)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error func client:", err)
    }


    fmt.Printf(data)
    fmt.Printf("Hello World")

}

