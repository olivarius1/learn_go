package delimiter_based

import (
	"fmt"
	"net"
)

func Client(conn net.Conn) {
	var sends string
	sendMsg := "{\"test1\":1,\"test2\",2}\n"
	for i := 0; i < 10; i++ {
		sends += sendMsg
		// 发送一个请求的数据 用\n 作为分割
		s, err := conn.Write([]byte(sends))
		if err != nil {
			fmt.Println("error..")
			return
		}
	}
}