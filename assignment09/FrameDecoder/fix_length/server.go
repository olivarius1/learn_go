package fix_length

import (
	"net"
	"fmt"
)


func Server(server net.Conn) {
	const LENGTH = 1024

	for {
		buf:= make([]byte, LENGTH)
		_, err := server.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("recv:",string(buf))
	}
}