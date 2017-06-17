package main
import ("fmt";"network")

// function main
func main() {
    
    //var isoc NT.ISocket = &NT.Socket{NT.SocketInfo{"serverIp",1234}}
    //var buffer [8] byte;
    //buffer := make([]byte,8)
    //isoc.Recv(buffer[:]) // pass slice in
    //fmt.Println(buffer)
    //isoc.Send(buffer[:])
    //fmt.Println("server side: socket object local ip=", psock.LocalIP)
    fmt.Println("12345")
    client := network.SocketClient{}
    client.Run()
}
