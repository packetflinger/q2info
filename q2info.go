package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	server := os.Args[1]
	p := make([]byte, 1500)

	// only use IPv4
	conn, err := net.Dial("udp4", server)
	if err != nil {
		fmt.Printf("Connection error %v", err)
		return
	}
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))

	cmd := []byte{0xff, 0xff, 0xff, 0xff}
	cmd = append(cmd, "status"...)
	fmt.Fprintln(conn, string(cmd))

	_, err = bufio.NewReader(conn).Read(p)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	lines := strings.Split(strings.Trim(string(p), " \n\t"), "\n")
	PrintServerVars(lines[1][1:])
	PrintPlayerInfo(lines[2:len(lines)-1])
}

func PrintServerVars(s string) {
	vars := strings.Split(s, "\\")
	for i := 0; i < len(vars); i++ {
		fmt.Printf("%s: ", vars[i])
		i++
		fmt.Printf("%s\n", vars[i])
	}
}

func PrintPlayerInfo(s []string) {
	playerscount := len(s)

	fmt.Println("player_count:", playerscount)

	if playerscount > 0 {
		players := ""
		for _, p := range s {
			player := strings.SplitN(p, " ", 3)
			players = fmt.Sprintf("%s,%s", players, player[2])
		}

		fmt.Println("players:", players[1:])
	}
}
