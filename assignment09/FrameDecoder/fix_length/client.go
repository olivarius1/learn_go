package fix_length

import (
	"net"
)

func ClientTcpFixLength(conn net.Conn) {
	sendByte := make([]byte, 1024)
	sendMessage := "{\"test1\":1,\"test2\",2}"
	for i := 0; i < 1000; i++ {
		temp := []byte(sendMessage)
		// 发一个完整的数据包
		for j := 0; j < len(temp) && j < 1024; j++ {
			sendByte[j] = temp[j]
		}
		_, err := conn.Write(sendByte)
		if err != nil {
			panic(err)
		}
	}
}