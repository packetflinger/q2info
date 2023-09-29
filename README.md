# Q2info
A command line tool for getting Quake 2 server information.

## Usage
`q2info [--config <config_file>] [--property <server_var>] <alias|ip:port>`

The server can either be an alias listed in a config file or an IPv4/DNS address + port. The `lookup` value is optional and gives you an easy way to get a single value. If `lookup` is omitted, all the available server info variables will be displayed.

## Config File
Server aliases are loaded from a text-format protobuf file named `.q2servers.config` in your home directory by default. You can override this name and path with the `--config` flag.

The proto file this config implements is https://github.com/packetflinger/libq2/proto/servers_file.proto

The `.q2servers.config` file format used by this program is a small subset of the `servers_file.proto` file. The bare minimum required for q2info is as follows:

```
# Add a server stanza for each server
server {
  name: "server1"
  address: "192.0.2.33:27988"
}

server {
  name: "server2"
  address: "q2.example.com:27910"
}
``` 

## Examples
```
$ ./q2info server1
fraglimit: 0
gamename: OpenTDM
port: 27910
q2admin: r238~374e0a8
revision: 215
deathmatch: 1
anticheat: 1
game: opentdm
gamedate: Nov 20 2020
gamedir: opentdm
mapname: q2dm1
maxclients: 20
time_remaining: 0:27
score_a: 51
players: "B100D","WallFly[BZZZ]","S!ckMan","Player-345","scr","Sniper.bg","Idaho"
match_type: TDM
hostname: PacketFlinger.com ~ OpenTDM ~ Germany
cheats: 0
dmflags: 1040
protocol: 34
timelimit: 0
version: q2proded r1828~0c53495 Nov 20 2020 Linux x86_64
uptime: 104 days, 12 hours, 33 mins, 58 secs
player_count: 7
score_b: 42
```
```
$ ./q2info --property=player_count server2
7
```

```
$ ./q2info --property=hostname 86.105.53.128:27910
PacketFlinger.com ~ OpenTDM ~ Germany
```

## Dependencies
```
$ go get google.golang.org/protobuf
$ go get github.com/packetflinger/libq2 
```

## Compiling
```
$ go build .
```
