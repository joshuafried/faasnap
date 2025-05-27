#!/bin/bash

set -x

REDIS=$(docker ps --format "{{.ID}} {{.Image}}" | grep '^.* redis' | awk '{print $1}')
if [ ! -z "$REDIS" ]; then
    docker container stop $REDIS
    docker container remove $REDIS
fi

ZIPKIN=$(docker ps --format "{{.ID}} {{.Image}}" | grep '^.* openzipkin' | awk '{print $1}')
if [ ! -z "$ZIPKIN" ]; then
    docker container stop $ZIPKIN
    docker container remove $ZIPKIN
fi

for i in {1..100}; do sudo ip netns delete fc$i ;done
