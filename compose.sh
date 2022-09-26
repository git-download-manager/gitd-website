#!/bin/bash

export SERVICE_BUILD=$(date '+%Y%m%d%H%M') && export SERVICE_COMMIT_ID=$(git describe --always) 

# docker image of all gitd services builds at one big images - gitd-builder
docker-compose -f docker-compose-builder.yml build --compress --progress plain

# separately builds per gitd services images
docker-compose -f docker-compose.yml build --compress --progress plain

# run all gitd images
docker-compose up --build --force-recreate