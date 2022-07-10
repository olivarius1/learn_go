package length_field

import (
  "encoding/binary"
  "fmt"
  "net"
)


func main() {
  conn, err := net.Dial("tcp", ":9000")
  if err != nil {
      return
  }
  defer conn.Close()
  for {
      s := "Hello, Server!"
	  // 2字节表示真实数据长度
      sbytes := make([]byte, 2+len(s))
      binary.BigEndian.PutUint16(sbytes, uint16(len(s)))
      copy(sbytes[2:], []byte(s))
      n, err := conn.Write(sbytes)
      if err != nil {
          fmt.Println("Error:", err)
          fmt.Println("Error N:", n)
          return
      }
  }
}