# Project Sprint : Project 1

## Requirements

- Go 1.22.x
- Docker / Docker Compose
- Make

## How to run?

To run this project, you need to run it per services example :

```bash
docker compose up -d # run databases
go run ./services/ms-user-svc
```

## Running Test

You can run test per services by using this command

```bash
go test ./services/ms-user-svc
```

## How to create services/libs?

The idea here is to use Go workspaces, please read [here to understand](https://go.dev/doc/tutorial/workspaces)

```bash
$ mkdir services/ms-file-svc
$ cd services/ms-file-svc
# Create Go Mod
$ go mod init github.com/mandaputtra/projectsprint-projects2/services/ms-file-svc
# Use it on the workspaces to tell this is a go module
# cd ..
$ go work use ./services/ms-user-svc/
```

## Build

## Deploy

## CI
