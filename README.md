# sandpiper
service monitor

## modes 
- minecraft (integrated with infrared proxy if used)
- http sites

## infrared minecraft

set something like this if you want it to display dynamic sleeping

the docker bit is unnesessary but thats how I manage sleeping

```json
{
    "domainName": "origin.example.com",
    "proxyTo": "hostname:port",
    "disconnectMessage": "Sorry {{username}}, the server is still offline or starting up.",
    "offlineStatus": {
        "versionName": "Sleeping",
        "protocolNumber": 754,
        "maxPlayers": 0,
        "playersOnline": 0,
        "motd": "Server is currently sleeping. Join to wake up!"
    },
    "docker": {
        "containerName": "docker-container",
        "timeout": 30000
    }
}
```