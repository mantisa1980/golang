package main

import (
  "net"
  "log"
  "time"
)

//import "bufio"
//import "strings" // only needed below for sample processing


func main() {
    l, err := net.Listen("tcp", ":8888")
    if err != nil {
        log.Println("listen error:", err)
        return
    }
    defer l.Close()

    for {
        c, err := l.Accept()
        if err != nil {
            log.Println("accept error:", err)
            break
        }
        // start a new goroutine to handle
        // the new connection.
        log.Println("accepted", c)
        go handleConn(c)
    }
}

func handleConn(conn net.Conn) {
  defer conn.Close()

  log.Println("start to read from conn")
  var buf = make([]byte, 10)
  for {
    // read from the connection
    dataLen, err := conn.Read(buf)
    if err != nil {
      log.Println("conn read error:", err)
      return
    }
    log.Printf("read %d bytes, content is %s\n", dataLen, string(buf[:dataLen]))
    time.Sleep(time.Second)
    conn.Write(buf[:dataLen])
  }
}