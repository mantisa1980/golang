package network

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

type SocketClient struct {
}

func (s *SocketClient) Run() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		// handle error
		fmt.Println("socket client error!", err)
		os.Exit(1)
	}
	//go socket_loop(conn)
	//socket_loop(conn)

	func(conn net.Conn) {
		//data_buffer := bytes.NewBuffer(nil)
		var receive_buffer [16]byte

		for {
			//n, err := conn.Write([]byte("1234"))
			time.Sleep(1000 * time.Millisecond)

			//n, err = conn.Read(receive_buffer[:]) // declared, so use =
			//data_buffer.Write(receive_buffer[0:n]) // save only response data length
			//fmt.Println("Data buffer=", data_buffer)

			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter text: ")
			text, _ := reader.ReadString('\n')
			//fmt.Println(text)
			//fmt.Println("Data buffer=", data_buffer)
			//n, err := conn.Write([]byte(text))
			conn.Write([]byte(text))

			time.Sleep(1000 * time.Millisecond)
			_, err := conn.Read(receive_buffer[:]) // declared, so use =
			//data_buffer.Write(receive_buffer[0:n]) // save only response data length
			fmt.Println("Receive echo data buffer=", receive_buffer)

			if err != nil {
				if err == io.EOF {
					break
				}
			}
		}
	}(conn)
}
