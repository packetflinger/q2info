# Q2info
A command line tool for getting Quake 2 server information.

## Usage
`q2info ip:port`

## Examples
```
q2info 192.0.2.45:27910
q2info gaming.packetflinger.com:27912
q2info dc.packetflinger.com:27910 | grep player_count | cut -d" " -f2
```