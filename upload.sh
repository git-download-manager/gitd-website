#!/bin/bash

# Before run: chmod a+x upload.sh
WHEREIAM=$(pwd)
IP_ADDR=$(cat $WHEREIAM/IP_ADDR)

ARGS=("$@")
SSH_PRIVATE_KEY="${ARGS[0]}"

echo "Build go files"
cd dmcli && go run . build go --config .build.prod.yaml

echo "Rsync file to server"
rsync -av --progress -e "ssh -i ~/$SSH_PRIVATE_KEY" --exclude={'*.dev.env','*.dev.yaml'} $WHEREIAM/build/prod/* root@$IP_ADDR:/usr/local/bin