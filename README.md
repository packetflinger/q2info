# Q2info
A command line tool for getting Quake 2 server information.

## Usage
`q2info [-v] [-c <config_file>] <server|ip[:port]> [lookup]`

The server can either be an alias listed in a config file or an IPv4 address or DNS name. If not using an alias, the `:port` is optional, if not supplied the default Quake 2 port of `27910` will be used. The `lookup` value is optional and gives you an easy way to get a single value. If `lookup` is omitted, all the available server info variables will be displayed.

## Config File
Server aliases are loaded from a file named `.q2servers.json` in your home directory by default. You can override this name and path with the `-c` argument. 

The `.q2servers.json` file format:

```
{
  "passwords": [
    {
      "name": "pass1",
      "password": "thisisareallybadpassword"
    },
    {
      "name": "pass2",
      "password": "ca00b050-495b-11ed-93de-cb7ecce5017d"
    }
  ],
  "servers": [
    {
      "name": "server1",
      "groups": "deathmatch usa",
      "addr": "100.64.3.5:27910",
      "password": "pass2"
    },
    {
      "name": "server2",
      "groups": "rocketarena usa",
      "addr": "192.0.2.44:27919",
      "password": "pass1"
    },
  ]
}
```
This file format is shared with other tools (like q2rcon) and has more infomation than is necessary for this program. The only required data for q2info is the `servers` array with each entry having a `name` and `addr`. You can supply the hostname/ip and port at the commandline directly if you like. 

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
$ ./q2info server2 player_count
7
```

```
$ ./q2info 86.105.53.128:27910 hostname
PacketFlinger.com ~ OpenTDM ~ Germany
```

## Compiling
`$ go build .`
