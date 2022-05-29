package server

import (
    "io"
    "log"
    "net"
    "fmt"
    "os"
)

type Listener struct {
    Port string
}

func (l *Listener) Start() {

    ln, err := net.Listen("tcp", ":" + l.Port) // ln can be struct data
    if err != nil {
        log.Fatal(err)
    }
    defer ln.Close()

    for {
        fmt.Println("accepting connections")
        conn, err := ln.Accept()
        if err != nil {
            log.Fatal(err)
        }
        go func(c net.Conn) {
            io.Copy(os.Stdout, c)
            c.Close()
        }(conn)
    }
}

type Dialer struct {
    Conn net.Conn
}

func (d *Dialer) Send(buffer []byte) {
    if _, err := d.Conn.Write(buffer); err != nil {
        log.Fatal(err)
    }
}
