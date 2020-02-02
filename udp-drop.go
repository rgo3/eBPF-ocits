package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	s, err := net.ResolveUDPAddr("udp4", "0.0.0.0:5001")
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listening on 0.0.0.0:5001")

	defer connection.Close()
	buffer := make([]byte, 1470)

	var dropped int
	var bytes int
	var seconds int

	n, _, _ := connection.ReadFromUDP(buffer)
	dropped++
	bytes += n
	timer := time.NewTicker(time.Second)
	for {
		select {
		case <-timer.C:
			seconds++
			fmt.Printf("dropping %d pps\n", dropped/seconds)
		default:
			n, _, _ = connection.ReadFromUDP(buffer)
			dropped++
			bytes += n
		}

	}
}
