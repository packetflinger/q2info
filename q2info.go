package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

/**
 * Start here
 */
func main() {
	// no args given, show usage
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <server_ip[:port]> [keylookup]\n", os.Args[0])
		fmt.Printf("  ex: %s 192.0.2.45:27911\n", os.Args[0])
		fmt.Printf("  ex: %s frag.gr:27940\n", os.Args[0])
		fmt.Printf("  ex: %s nj.packetflinger.com\n", os.Args[0])
		fmt.Printf("  ex: %s tastyspleen.net:27916 players\n", os.Args[0])
		return
	}

	server := os.Args[1]
	if !strings.Contains(server, ":") {
		server = server + ":27910"
	}

	// user included a specific value to get
	lookup := ""
	if len(os.Args) == 3 {
		lookup = os.Args[2]
	}

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
	info := ParseServerinfo(lines)

	if lookup != "" {
		PrintSpecificVar(info, lookup)
	} else {
		PrintServerVars(info)
	}
}

/**
 * Load the backslash delimited infostring into a key value map
 */
func ParseServerinfo(s []string) map[string]string {
	serverinfo := s[1][1:]
	playerinfo := s[2 : len(s)-1]

	info := map[string]string{}
	vars := strings.Split(serverinfo, "\\")

	for i := 0; i < len(vars); i += 2 {
		info[strings.ToLower(vars[i])] = vars[i+1]
	}

	playercount := len(playerinfo)
	info["player_count"] = fmt.Sprintf("%d", playercount)

	if playercount > 0 {
		players := ""

		for _, p := range playerinfo {
			player := strings.SplitN(p, " ", 3)
			players = fmt.Sprintf("%s,%s", players, player[2])
		}

		info["players"] = players[1:]
	}
	return info
}

/**
 * Just spit everything to stdout
 */
func PrintServerVars(info map[string]string) {
	for k, v := range info {
		fmt.Printf("%s: %s\n", k, v)
	}
}

/**
 * Print out only the value for the given key
 */
func PrintSpecificVar(info map[string]string, lookup string) {
	for k, v := range info {
		if k == strings.ToLower(lookup) {
			fmt.Printf("%s\n", v)
		}
	}
}
