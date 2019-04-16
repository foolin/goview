#!/usr/bin/env bash

#rice embed-go
#echo " => rice embed-go"

go build -o bin/gorice
echo " => go build -o bin/gorice"

echo " => ./bin/gorice"
./bin/gorice