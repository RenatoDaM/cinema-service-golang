package main

import (
	databaseConnector "com.cinema/cinema-app/common"
	"fmt"
    "net"
    "os"
	"strconv"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "8080"
    CONN_TYPE = "tcp"
)

func main() {
	databaseConnector.Init()
	listen()
}

func listen() {
	listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)

	if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }

	defer listener.Close()
    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        go handleRequest(conn)
    }
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)

	reqLen, err := conn.Read(buf)

	if err != nil {
	  fmt.Println("Error reading:", err.Error())
	}
	
	print("request length: " + strconv.Itoa(reqLen))
	conn.Write([]byte("Message received."))
	conn.Close()
  }
