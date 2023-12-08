#!bin/sh
for day in [0-9]*/; do
    echo $day
    go run $day/main.go $day/input.txt
done
