/*
This is the main entry of my chatroom project.

moduels:

SocketServer: basic wrapper for data transmission
SocketDataAdapter: binary/ json / ...
MessageDispatcher
ChatRoomManager: dispatch message to correct rooms
 ChatRoom
 PrivateChatRoom
 PublicChatRoom
HistoryKeeper

*/


package main
import (NT "network")

const (
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

// function main
func main() {
    
    //var isoc NT.ISocket = &NT.Socket{NT.SocketInfo{"serverIp",1234}}
    //var buffer [8] byte;
    //buffer := make([]byte,8)
    //isoc.Recv(buffer[:]) // pass slice in
    //fmt.Println(buffer)
    //isoc.Send(buffer[:])
    //fmt.Println("server side: socket object local ip=", psock.LocalIP)

    server := NT.SocketServer{"0.0.0.0", 8080}
    server.Run()
}
