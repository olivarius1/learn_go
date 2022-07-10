package length_field

import (
  "encoding/binary"
  "fmt"
  "net"
)

func main() {
  ln, err := net.Listen("tcp", ":9000")
  if err != nil {
      return
  }
  for {
      conn, err := ln.Accept()
      if err != nil {
          continue
      }
      go handleConnection(conn)
  }
}
func handleConnection(conn net.Conn) {
  defer conn.Close()
  tmp := []byte{}
  for {
      buf := make([]byte, 1024)
      n, err := conn.Read(buf)
      if err != nil {
          if e, ok := err.(*net.OpError); ok {
              fmt.Println(e.Source, e.Addr, e.Net, e.Op, e.Err)
              if e.Timeout() {
                  fmt.Println("Timeout Error")
              }
          }
          fmt.Println("Read Error:", err)
          fmt.Println("Read N:", n)
          return
      }
      if n == 0 {
          fmt.Println("Read N:", n)
          return
      }
      tmp = append(tmp, buf[:n]...)
      length := len(tmp)
	  // 封包总长度不足2字节（这种情况不能完整获取包头），缓存起来与下次获取的数据拼接
      if length < 2 {
          continue
      }
      if length >= 2 {
          head := make([]byte, 2)
          copy(head, tmp[:2])
          dataLength := binary.BigEndian.Uint16(head)
          data := make([]byte, dataLength)
          copy(data, tmp[2:dataLength+2])
          fmt.Println(string(data)) // 得到数据
		  // 刚好读一个包 清空 tmp
          if uint16(length) == 2+dataLength {
              tmp = []byte{}
          } else if uint16(length) > 2+dataLength {
			  // 粘包 剩余数据给下一个请求用
              tmp = tmp[dataLength+2:]
          }
      }
      // fmt.Println(string(buf))
  }
}