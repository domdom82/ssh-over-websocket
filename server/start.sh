#!/usr/bin/env bash

GOOS=linux GOARCH=amd64 go build .

docker rm -f ssh-over-websocket

docker build -t ssh-over-websocket .

docker run --name ssh-over-websocket -p 8080:8080 ssh-over-websocket &

sleep 1

docker exec ssh-over-websocket service ssh start