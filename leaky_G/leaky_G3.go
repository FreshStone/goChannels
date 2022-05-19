package main

import (
	"fmt"
	"net"
	"time"
)

func sendMsg(msg, addr string) error {
    conn, err := net.Dial("tcp", addr)
    if err != nil {
        return err
    }
    defer conn.Close()
    _, err = fmt.Fprint(conn, msg)
//    fmt.Println(err)
    return err
} 

func broadcastMsg(msg string, addrs []string) error {
    errc := make(chan error)
    for _, addr := range addrs {
        go func(addr string) {
            errc <- sendMsg(msg, addr)
            fmt.Println("done")
        }(addr)
    }
//	var errRec error
    for _ = range addrs {
        if err := <-errc; err != nil {
            return err
        }
    }
    return nil
}

func main() {
//	addr := []string{ "localhost:8080", "http://google.com"}
//	addr := []string{"http://google.com"}
       addr := []string{"localhost:8080"}


    err := broadcastMsg("hi", addr)
    fmt.Println(err)
    time.Sleep(time.Second)

    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("everything went fine")
}
