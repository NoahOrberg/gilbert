#!/usr/bin/env bash

for d in $(go list); do
    go test -v $d
done
