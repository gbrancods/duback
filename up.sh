#!/bin/bash

docker build -t duback .
#docker-compose up -d
docker-compose run --rm duback