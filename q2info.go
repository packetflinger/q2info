package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sort"
	"strings"
	"time"
)

//
// Structure to hold our config file ~/.q2servers.json
// This is a small subset of the data in that file,
// just what we need to perform our duty
//
type ServerJSON struct {
	Servers []struct {
		Name string `JSON:"name"`
		Addr string `JSON:"addr"`
	} `JSON:"servers"`
}

const (
	DefaultPort = ":27910"
)

var (
	Config  ServerJSON
	SrvFile *string // flag
	Verbose *bool   // flag
)

//
// init() is called before this
//
func main() {
	server := GetServer(flag.Arg(0))

	// user included a specific value to get
	lookup := ""
	if len(flag.Args()) == 2 {
		lookup = flag.Arg(1)
	}

	p := make([]byte, 1500)

	if *Verbose {
		log.Printf("Connecting to %s\n", server)
	}

	// only use IPv4
	conn, err := net.Dial("udp4", server)
	if err != nil {
		fmt.Printf("Connection error %v\n", err)
		return
	}
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))

	cmd := []byte{0xff, 0xff, 0xff, 0xff}
	cmd = append(cmd, "status"...)

	starttime := time.Now()
	fmt.Fprintln(conn, string(cmd))
	_, err = bufio.NewReader(conn).Read(p)
	duration := time.Since(starttime)

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

	if *Verbose {
		log.Printf("Results fetched in %s\n", duration.String())
	}
}

//
// Either locate the server reference in the config file or assume
// it was given manually at runtime and format it properly with
// a port
//
func GetServer(alias string) string {
	for _, s := range Config.Servers {
		if strings.EqualFold(s.Name, alias) {
			return s.Addr
		}
	}

	// not found in config file, assuming arg is the address
	if !strings.Contains(alias, ":") {
		alias = alias + DefaultPort
	}

	return alias
}

//
// Load the backslash delimited infostring into a key value map
//
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

//
// Just spit everything to stdout after sorting
//
func PrintServerVars(info map[string]string) {
	keys := make([]string, 0, len(info))
	for k := range info {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s: %s\n", k, info[k])
	}
}

//
// Print out only the value for the given key
//
func PrintSpecificVar(info map[string]string, lookup string) {
	for k, v := range info {
		if strings.EqualFold(k, lookup) {
			fmt.Printf("%s\n", v)
		}
	}
}

//
// Called before main()
//
func init() {

	// parse args
	SrvFile = flag.String("c", "", "Specify a server data file")
	Verbose = flag.Bool("v", false, "Show some more info")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Printf("Usage: %s <flags> <serveralias> [property]\n", os.Args[0])
		fmt.Printf("  flags:\n")
		flag.PrintDefaults()
		fmt.Println("  [property] is optional and can be any one key value returned")
		fmt.Println("    ex: players, hostname, dmflags")
		os.Exit(0)
	}

	if *SrvFile == "" {
		homedir, err := os.UserHomeDir()
		sep := os.PathSeparator
		if err != nil {
			log.Fatal(err)
		}
		*SrvFile = fmt.Sprintf("%s%c.q2servers.json", homedir, sep)
	}

	if *Verbose {
		log.Printf("Loading passwords/servers from %s\n", *SrvFile)
	}

	raw, err := os.ReadFile(*SrvFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(raw, &Config)
	if err != nil {
		log.Fatal(err)
	}

	if *Verbose {
		log.Printf("  %d servers found\n", len(Config.Servers))
	}
}
