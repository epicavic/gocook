#!/usr/bin/env bash

# trap ctrl-c and call ctrl_c()
trap ctrl_c INT
container="mysql"
args=(
    "-e" "MYSQL_ROOT_PASSWORD=root"
    "-e" "MYSQL_DATABASE=cook"
    "-e" "MYSQL_USER=user"
    "-e" "MYSQL_PASSWORD=pass"
    "-p" "3306:3306"
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
MYSQLUSERNAME='user' \
MYSQLPASSWORD='pass' \
MYSQLDATABASE='cook' \
go run main.go

# stop container
docker stop "${container}" >/dev/null