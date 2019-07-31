#!/usr/bin/env bash

case "$1" in
enable)
    sed -i 's://fmt.Print:fmt.Print:' game.go
    ;;
disable)
    sed -i 's:fmt.Print://fmt.Print:' game.go
    ;;
*)
    echo "write enable or disable"
    ;;
esac

goimports -w game.go