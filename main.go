package main

import (
    "log"
    "net"
    "sync"
    "tcpc/server"
)

func startTcpListener(wg *sync.WaitGroup) {
    defer wg.Done()
    var listener = &server.Listener{
        Port: "8080",
    }
    listener.Start()
}


func main() {

    var wg sync.WaitGroup
    wg.Add(1)
    go startTcpListener(&wg)

    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        log.Fatalf("Failed to dial: %v", err)
    }

    var tcpclient = &server.Dialer {
        Conn: conn,
    }
    defer conn.Close()

    var messages = []byte("abcdef\n")
    for i := 0; i < 5; i++ {
        tcpclient.Send(messages)
    }

    wg.Wait()
}
