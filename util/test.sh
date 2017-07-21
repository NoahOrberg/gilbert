#!/usr/bin/env bash

for d in $(go list $(glide novendor)); do
    go test -v $d
done
