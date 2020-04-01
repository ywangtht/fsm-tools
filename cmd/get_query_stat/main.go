package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"net"
)

func main() {
	host := flag.String("h", "localhost", "remote server")
	port := flag.String("p", "7918", "remote server port")
	flag.Parse()
	url := *host + ":" + *port
	conn, err := net.Dial("tcp", url)
	if err != nil {
		fmt.Printf("failed to connect to server with url: %s\n", url)
		return
	}
	buff := new(bytes.Buffer)
	binary.Write(buff, binary.LittleEndian, uint32(119))
	binary.Write(buff, binary.LittleEndian, uint32(0))
	binary.Write(buff, binary.LittleEndian, uint32(0))
	binary.Write(buff, binary.LittleEndian, uint32(0))
	intByteArray := buff.Bytes()
	conn.Write(intByteArray)
	p := make([]byte, 4)
	io.ReadFull(conn, p)
	statLen := binary.LittleEndian.Uint32(p)
	fmt.Printf("stat json length: %d\n", statLen)
	message := make([]byte, statLen)
	data := make([]byte, 0)
	len, err := conn.Read(message)
	if err != nil || len == 0 {
		fmt.Println(err)
	}
	if len > 0 {
		data = append(data, message[:len]...)
	}
	fmt.Println("received: " + string(data))
}
