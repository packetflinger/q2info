package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	pb "github.com/packetflinger/libq2/proto"
	"github.com/packetflinger/libq2/state"
	"google.golang.org/protobuf/encoding/prototext"
)

var (
	serversFile = ".q2servers.config" // default name, should be in home directory
	config      = flag.String("config", "", "Specify a server data file")
	property    = flag.String("property", "", "Output only this specific server var")
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		showUsage()
		return
	}

	serverspb, err := loadConfig()
	if err != nil {
		log.Println(err)
		return
	}

	addr, port, err := resolveTarget(serverspb, flag.Arg(0))
	if err != nil {
		log.Println(err)
		return
	}

	server := state.Server{
		Address: addr,
		Port:    port,
	}
	info, err := server.FetchInfo()
	if err != nil {
		log.Println(err)
		return
	}

	outputInfo(info, *property)
}

// if no args are supplied
func showUsage() {
	fmt.Printf("Usage: %s <flags> <serveralias>\n", os.Args[0])
	fmt.Printf("  flags:\n")
	flag.PrintDefaults()
}

// Read the text-format proto config file and unmarshal it
func loadConfig() (*pb.ServerFile, error) {
	cfg := &pb.ServerFile{}

	if *config == "" {
		homedir, err := os.UserHomeDir()
		sep := os.PathSeparator
		if err != nil {
			return nil, err
		}
		*config = fmt.Sprintf("%s%c%s", homedir, sep, serversFile)
	}

	raw, err := os.ReadFile(*config)
	if err != nil {
		return nil, err
	}

	err = prototext.Unmarshal(raw, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// Attempt to match the target arg to an identifier in the server config.
// If an exact match isn't found, assume it's a literal address instead of
// an alias.
//
// Returns the ip/host, port, and any errors
func resolveTarget(cfg *pb.ServerFile, targ string) (string, int, error) {
	found := ""
	for _, sv := range cfg.GetServer() {
		if sv.GetIdentifier() == strings.ToLower(targ) {
			found = sv.GetAddress()
		}
	}

	// we didn't match anything, assume the alias is just a server address
	if found == "" {
		found = targ
	}

	tokens := strings.Split(found, ":")
	if len(tokens) == 2 {
		ip := tokens[0]
		port, err := strconv.Atoi(tokens[1])
		if err != nil {
			return "", 0, errors.New("malformed server address, bad port - " + found)
		}
		return ip, port, nil
	}
	return "", 0, errors.New("invalid address, no port specified - " + found)
}

// Display the data we found.
func outputInfo(info state.ServerInfo, property string) {
	if property != "" {
		fmt.Println(info.Server[property])
	} else {
		for k, v := range info.Server {
			fmt.Printf("%s: %s\n", k, v)
		}
	}
}
