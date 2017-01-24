package network

import ("fmt"
        "net"
        "strconv"
        "reflect"
        "strings"
        "os")


type INet interface {
    Send(data []byte) bool
}

type IUDP interface {
    INet
}

type ISocket interface {
    INet
    Close() bool
}

type SocketInfo struct {
    SrcIP string
    SrcPort int
 //   DestIP string
 //   DestPort int
}

type Socket struct {
    SocketInfo
}

func (x Socket) String() string {
    //return fmt.Sprintf("%s:%d:%s:%d", x.SrcIP,x.SrcPort,x.DestIP,x.DestPort)  
    return fmt.Sprintf("%s:%d", x.SrcIP,x.SrcPort)  
}

func (s* Socket) Send(buffer []byte) bool {
    fmt.Println("socket send data imp", buffer)
    return true
}

func (s* Socket) Close() bool {
    return true
}

// ===========
type SocketServer struct {
    ListenIP string
    ListenPort int
}

type ISocketEventListener interface {
    OnConnect(hnd ISocket)
    OnClose(hnd ISocket)
    OnReceiveData(buffer []byte)
}

type ISocketServer interface {
    RegisterSocketEventListener(handler *ISocketEventListener)
    Run()
}

// implementation of ISocketServer by SocketServer
func (s *SocketServer) RegisterSocketEventListener(handler *ISocketEventListener) {
    // todo: put handler in s.listen_ist
}

func (s *SocketServer) Run() {
    //l, err := net.Listen("tcp", fmt.Sprintf("%s:%d",s.ListenIP,s.ListenPort))  // ip will be resolved as ipv6...
    l, err := net.Listen("tcp4", fmt.Sprintf("%s:%d",s.ListenIP,s.ListenPort))
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    // Close the listener when the application closes.
    defer l.Close()
    fmt.Println("Listening on " + s.ListenIP + ":" + strconv.Itoa(s.ListenPort))
    for {
        // Listen for an incoming connection.
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        // Handle connections in a new goroutine.
        // !!todo: save socket objects

        adr_info := strings.Split(conn.RemoteAddr().String(), ":")
        client_ip := adr_info[0]
        client_port,err := strconv.Atoi(adr_info[1])

        fmt.Println(reflect.TypeOf(conn), client_ip, client_port, reflect.TypeOf(client_ip), reflect.TypeOf(client_port) )

        // callback listeners

        //socket_handle := Socket{SocketInfo{SrcIP:client_ip,SrcPort:client_port}}

        go handleRequest(conn)
    }
}

// private functions
func handleRequest(conn net.Conn) {
    defer conn.Close()
    buf := make([]byte, 2)

    for {
        rd_len, err := conn.Read(buf)  // will read length based on size (determined on compile time) of buf each time.
        if err != nil {
          fmt.Println("Error reading:", err.Error())
          break
        }
        fmt.Println("read length:",rd_len, buf[:rd_len])
        conn.Write(buf[:rd_len])
    }
}
