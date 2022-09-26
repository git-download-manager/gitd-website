#!/bin/bash

# Before run: chmod a+x upload.sh
WHEREIAM=$(pwd)
IP_ADDR=$(cat $WHEREIAM/IP_ADDR)

echo "Build go files"
cd dmcli && go run . build go --config .build.prod.yaml

echo "Rsync file to server"
rsync -av --progress -e "ssh -i ~/.ssh/gitdm_hetzner" --exclude={'*.dev.env','*.dev.yaml'} $WHEREIAM/build/prod/* root@$IP_ADDR:/usr/local/bin