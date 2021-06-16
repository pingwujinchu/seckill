#!/bin/bash
bazel build //client:client
cd server
docker build -t sec-kill-server:dev .
docker-compose up