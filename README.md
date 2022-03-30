# Q2info
A command line tool for getting Quake 2 server information.

## Usage
`q2info ip[:port] [lookup]`

You can use either an IPv4 address or DNS hostname. The `port` is optional, if not supplied the default Quake 2 port of `27910` will be used. The `lookup` value is optional and gives you an easy way to get a single value. If `lookup` is omitted, all the available server info variables will be displayed.

## Examples
```
$ ./q2info 86.105.53.128:27910
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
$ ./q2info 86.105.53.128:27910 player_count
7
```

```
$ ./q2info 86.105.53.128:27910 hostname
PacketFlinger.com ~ OpenTDM ~ Germany
```

## Compiling
`$ go build .`
