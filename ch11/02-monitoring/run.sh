#!/usr/bin/env bash

# trap ctrl-c and call ctrl_c()
trap ctrl_c INT
container="prom/prometheus"
args=(
    "-p" "9090:9090"
    "-v" "$(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml"
)

function sanitize {
    echo "${*}" | tr -d '[:punct:]'
}
container_name=$(sanitize ${container})

function ctrl_c() {
    echo "Trapped CTRL-C"
    docker stop "${container_name}" >/dev/null 
    exit
}

# run container
docker run --rm --name "${container_name}" "${args[@]}" -d "${container}" >/dev/null

# run go main
go run main.go

# stop container
docker stop "${container_name}" >/dev/null