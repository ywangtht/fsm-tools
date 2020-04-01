package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"net"
)

const (
	eventIDGetQueryStat uint32 = 119
	eventLen            uint32 = 0
	processHash         uint32 = 0
	seqID               uint32 = 0
)

var (
	host = flag.String("h", "localhost", "remote server")
	port = flag.String("p", "7918", "remote server port")
)

func main() {
	flag.Parse()
	url := *host + ":" + *port
	conn, err := net.Dial("tcp", url)
	if err != nil {
		fmt.Printf("failed to connect to server with url: %s\n", url)
		return
	}
	buff := new(bytes.Buffer)
	binary.Write(buff, binary.LittleEndian, eventIDGetQueryStat)
	binary.Write(buff, binary.LittleEndian, eventLen)
	binary.Write(buff, binary.LittleEndian, processHash)
	binary.Write(buff, binary.LittleEndian, seqID)
	intByteArray := buff.Bytes()
	conn.Write(intByteArray)

	//receive response
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
