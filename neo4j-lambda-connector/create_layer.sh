#!/bin/bash -x
set -e
rm -rf ../layer
mkdir -p ../layer
docker build -t seabolt-base .
CONTAINER=$(docker run -d seabolt-base false)
docker cp $CONTAINER:/opt/lib ../layer
rm -r ../layer/lib/python2.7
rm -r ../layer/lib/python3.6
docker rm $CONTAINER