package delimiter_based

import (
	"bufio"
	"fmt"
	"net"
)

func Server(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		slice, err := reader.ReadSlice('\n')
		if err != nil {
			fmt.Println("error", err)
			continue
		}
		fmt.Printf("slice %s", slice)
	}
}