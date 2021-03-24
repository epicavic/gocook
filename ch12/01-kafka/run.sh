#!/usr/bin/env bash

# trap ctrl-c and call ctrl_c()
trap ctrl_c INT

function ctrl_c() {
    echo "Trapped CTRL-C"
    docker-compose down &>/dev/null
    exit
}

# run containers
docker-compose up -d &>/dev/null
sleep 10

# run go apps
go run producer/main.go &
go run consumer/main.go

# stop containers
docker-compose down &>/dev/null