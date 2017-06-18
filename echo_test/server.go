package main

import (
  "net"
  "log"
  "time"
  "math/rand"
)

//!! TODO: renmove useless socket when disconnected

func main() {
    l, err := net.Listen("tcp", ":8888")
    if err != nil {
        log.Println("listen error:", err)
        return
    }
    log.Println("listening on port 8888...")
    defer l.Close()

    //var socket_map map[string]net.Conn = make(map[string]net.Conn)
    var socket_list =[]net.Conn{}
    rand.Seed(time.Now().Unix())

    for {
        c, err := l.Accept()
        if err != nil {
            log.Println("accept error:", err)
            break
        }

        log.Println("accepted", c, "remote addr", c.RemoteAddr().String())

        //socket_map[c.RemoteAddr().String()] = c
        socket_list = append(socket_list, c)
        //log.Println("len=", len(socket_list))

        //go handleConn(c)
        go handleConn(c, socket_list)
    }
}

func handleConn(conn net.Conn, socket_list []net.Conn) {
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
    //time.Sleep(time.Second)
    conn.Write(buf[:dataLen])
    go forward(socket_list, buf[:dataLen])
  }
}

func forward(socket_list []net.Conn, buf[] byte) {
  ln := len(socket_list)
  var lst [] int
  if ln >= 3 {
    lst = rand.Perm(ln)[:3]
  } else {
    lst = rand.Perm(ln)
  }

  log.Println("forward data", string(buf), " to ", socket_list, " lst=", lst)
  
  for index, _ := range lst {
    socket_list[index].Write(buf)
    // index is the index where we are
    // element is the element from someSlice for where we are
  }
}