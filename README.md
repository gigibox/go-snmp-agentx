# go-snmp-agent

## Overview
This project is an SNMP AgentX subagent implemented in Golang. It extends an SNMP master agent by providing real-time system metrics, including CPU usage, temperature, memory status, and wireless network information.

## Build
```azure
# MT7621
GOOS=linux GOARCH=mipsle GOMIPS=softfloat CGO_ENABLED=0 go build -ldflags="-s -w"

# 7981
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w"
```
## Project
The implementation was provided by [posteo/go-agentx)](https://github.com/posteo/go-agentx)

## License
The project is licensed under LGPL 3.0 (see LICENSE file).
