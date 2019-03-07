package main

import (
    "fmt"
    "net"
    "os"
    "context"
    "github.com/apache/thrift/lib/go/thrift"
    "gen-go/spider/bks"
)

const (
    HOST = "127.0.0.1"
    PORT = "9090"
)

func main() {
    transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
    protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

    transport, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
    if err != nil {
        fmt.Fprintln(os.Stderr, "error resolving address:", err)
        os.Exit(1)
    }

    useTransport,_ := transportFactory.GetTransport(transport)
    client := bks.NewBksClientFactory(useTransport, protocolFactory)
    if err := transport.Open(); err != nil {
        fmt.Fprintln(os.Stderr, "Error opening socket to "+HOST+":"+PORT, " ", err)
        os.Exit(1)
    }
    defer transport.Close()

	xn := "2018"
	xq := "1"
	stuid := "201626010520"
	password := "WudUozhI"

	data,_ := client.GetClassTable(context.Background(),xn,xq,stuid,password)
	fmt.Printf(data)

}

