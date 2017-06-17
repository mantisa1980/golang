package network

import (
	"fmt"
	"net"
	"os"
	"reflect"
	"strconv"
	"strings"
)

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
	SrcIP   string
	SrcPort int
	//   DestIP string
	//   DestPort int
}

type Socket struct {
	SocketInfo
}

func (x Socket) String() string {
	//return fmt.Sprintf("%s:%d:%s:%d", x.SrcIP,x.SrcPort,x.DestIP,x.DestPort)
	return fmt.Sprintf("%s:%d", x.SrcIP, x.SrcPort)
}

func (s *Socket) Send(buffer []byte) bool {
	fmt.Println("socket send data imp", buffer)
	return true
}

func (s *Socket) Close() bool {
	return true
}

// ===========
type SocketServer struct {
	ListenIP   string
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
	l, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", s.ListenIP, s.ListenPort))
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

		adrInfo := strings.Split(conn.RemoteAddr().String(), ":")
		clientIP := adrInfo[0]
		clientPort, err := strconv.Atoi(adrInfo[1])

		fmt.Println(reflect.TypeOf(conn), clientIP, clientPort, reflect.TypeOf(clientIP), reflect.TypeOf(clientPort))

		// callback listeners

		//socket_handle := Socket{SocketInfo{SrcIP:clientIP,SrcPort:client_port}}

		go handleRequest(conn)
	}
}

// private functions
func handleRequest(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2)

	for {
		rdLen, err := conn.Read(buf) // will read length based on size (determined on compile time) of buf each time.
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}
		fmt.Println("read length:", rdLen, buf[:rdLen])
		conn.Write(buf[:rdLen])
	}
}
