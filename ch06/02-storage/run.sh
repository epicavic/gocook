#!/usr/bin/env bash

# trap ctrl-c and call ctrl_c()
trap ctrl_c INT
container="mongo"
args=(
    "-p" "27017:27017"
    "-v" "$(pwd)/db:/data/db"
)

function ctrl_c() {
    echo "Trapped CTRL-C"
    docker stop "${container}"
    exit
}

# run container
docker run --rm --name "${container}" "${args[@]}" -d "${container}" >/dev/null

# run go main
go run main.go

# stop container
docker stop "${container}" >/dev/null