# sandpiper
service monitor

## modes 
- minecraft (integrated with infrared proxy if used)
- http sites
- more to come (contribute one if you like)

## building

either run this locally using

```
go build -o bin/sandpiper
bin/sandpiper <path to config, relative or absolute>
```

or use the dockerfile/makefile to build docker containers.

## config

a json file is used for configuration, port field can be omitted if not a minecraft server.

the templates page checks if a monitor of type x exists before rendering the category.

comments in your config will cause the program to crash

```json
{
    "port": "8022",
    "interval": "10s",
    "monitors": [
        {
            "name": "Google",
            "address": "https://google.ca",
            "mode": "http"
        },
        {
            "name": "Minecraft Server",
            "address": "mc.example.com",
            "port": 25565,
            "mode": "minecraft"
        }
    ]
}
```

## infrared minecraft

set something like this for your infrared config if you want it to display dynamic sleeping

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

## hacking / contributing
- main code goes in main.go 
- all web resources should go in assets/
- modules for supporting different gameservers should be in modes/
- test your PRs with docker and native

## coming soon
- interfaces for modules to provide gameserver info
- mode-agnostic card format if you don't want the themed ones