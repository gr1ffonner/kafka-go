#!/bin/bash

CGO_ENABLED=0 GOOS=linux go build $@;
status=$?;
attempt=1;
while [ $status -ne 0 ] && [ $attempt -le 5 ]; do
    sleep 20;
    (( attempt++ ));
    CGO_ENABLED=0 GOOS=linux go build $@;
    status=$?;
done;
