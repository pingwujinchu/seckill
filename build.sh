#!/bin/bash
# bazel build //client:client
docker run \
  -e USER="$(id -u)" \
  -u="$(id -u)" \
  -v $PWD/:/src/workspace \
  -v /tmp/build_output:/tmp/build_output \
  -w /src/workspace \
  l.gcr.io/google/bazel:latest \
  --output_user_root=/tmp/build_output \
  build //client:client --linkopt="-pthread" 

cd server
docker build -t sec-kill-server:dev .
docker-compose up