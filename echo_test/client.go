package main

import "net"
//import "fmt"
//import "bufio"
//import "os"
import "log"
import "time"
import "sync"
import "bytes"
import "strconv"
// read or write on conn

func ddos_worker(wg *sync.WaitGroup, conn net.Conn, no int) {
  //defer log.Println("defer ddos")
  //defer conn.Close()  // cannot close on exiting this function, since goroutine is not yet finished
  //defer wg.Done()
  
  var sendBuf bytes.Buffer
  
  //pattern := strconv.Itoa(no) + "--------------------------------------------------------------------------------------" 
  pattern := strconv.Itoa(no)
  //pattern := no 
  log.Println("pattern=", pattern)

  go func(pattern string, conn net.Conn) {
    counter :=0
    for counter < 10 {
      sendBuf.WriteString(pattern)
      conn.Write([] byte(pattern))
      counter+=1
    }
    c1 := make(chan int)

    //log.Println("send len=", sendBuf.Len())
    //time.Sleep(time.Millisecond * 1000)

    go func(conn net.Conn, channel chan int, sb bytes.Buffer) {
      defer conn.Close()
      defer wg.Done()
      var readBuf bytes.Buffer
      var buf = make([]byte, 10)
      readLen := 0
      //for readLen < sb.Len()  {
      for {
        dataLen, _ := conn.Read(buf)

        readLen += dataLen
        log.Println("bot", no, " read len:", dataLen, "now len:", readLen, "data",  string(buf))  // or just print bytes buf
        readBuf.Write(buf[:dataLen] )
      }
      //log.Println("fin reading:", readBuf.String())
      channel <- 9999
    }(conn, c1, sendBuf)

    signal := <-c1
    log.Println("receive signal", signal)
  }(pattern, conn)
  //log.Println("end of ddos")

  //time.Sleep(time.Millisecond * 1000)  // sleep to see  forwarded data
}

  

func main() {
  var wg sync.WaitGroup

  log.Println("dial ok")
  for i := 0; i < 4; i++ {
    conn, err := net.DialTimeout("tcp", ":8888", 2 * time.Second)
    if err != nil {
      log.Println("dial error:", i, err)
      return
    }
    wg.Add(1)
    go ddos_worker(&wg, conn, i)
  }
  
  wg.Wait()
  time.Sleep(time.Second * 1)
  log.Println("ddos Done")

  
  //data := os.Args[1]
  //conn.Write([]byte(data))
  //time.Sleep(time.Second * 10000)
}