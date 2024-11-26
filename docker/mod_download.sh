#!/bin/bash

go mod download $@;
status=$?;
attempt=1;
while [ $status -ne 0 ] && [ $attempt -le 5 ]; do
    sleep 20;
    (( attempt++ ));
    go mod download $@;
    status=$?;
done;
