#!/bin/bash

# Usage: ./build service-file-name 
# Before run: chmod a+x build.sh

# Local Variables
ARGS=("$@")
SERVICE_NAME="${ARGS[0]}"
#SERVICE_VERSION="$(cat VERSION)"
OS_TYPE="${ARGS[1]}"
ARCH_TYPE="${ARGS[2]}"
BUILD_PATH="${ARGS[3]}"
BUILD_DEST="${ARGS[4]}"
BUILD_NAME="${ARGS[5]}"
COPY_TO="${ARGS[6]}"
SENV="${ARGS[7]}"
SERVICE_BUILD=$(date '+%Y%m%d%H%M')
SERVICE_SHORT_COMMIT_ID=$(git describe --always)
#SERVICE_COMMIT_ID=$(git rev-parse HEAD)

# Templates Folder Copy
TEMPLATESFOLDER=$BUILD_DEST../templates
if [ ! -d "$TEMPLATESFOLDER" ]; then
    echo "$TEMPLATESFOLDER not exists."

    mkdir -p $TEMPLATESFOLDER
fi

cp -vr $BUILD_PATH../data/templates $BUILD_DEST../

# Cron Folder Copy
CRONLISTFOLDER=$BUILD_DEST../cronlist
if [ ! -d "$CRONLISTFOLDER" ]; then
    echo "$CRONLISTFOLDER not exists."

    mkdir -p $CRONLISTFOLDER
fi

cp -vr $BUILD_PATH../data/cronlist $BUILD_DEST../

if [[ $# -eq 0 ]] ; then
    echo "Service file name not found"
    exit 1
fi

# Nginx Conf Folder Copy
NGINXCONFFOLDER=$BUILD_DEST../nginx-conf
if [ ! -d "$NGINXCONFFOLDER" ]; then
    echo "$NGINXCONFFOLDER not exists."

    mkdir -p $NGINXCONFFOLDER
fi

cp -vr $BUILD_PATH../data/nginx-conf/$SENV/ $BUILD_DEST../nginx-conf/

if [[ $# -eq 0 ]] ; then
    echo "Service file name not found"
    exit 1
fi

cd $BUILD_PATH
GOOS=$OS_TYPE GOARCH=$ARCH_TYPE go build -ldflags="-w -s -X main.ServiceName=$SERVICE_NAME -X main.ServiceVersion=$SERVICE_SHORT_COMMIT_ID -X main.ServiceBuild=$SERVICE_BUILD -X main.ServiceCommitId=$SERVICE_SHORT_COMMIT_ID" -o "$BUILD_NAME"

if [ -f "$BUILD_NAME" ]; then
    echo "$SERVICE_NAME with $SERVICE_SHORT_COMMIT_ID version build successfully"

    chmod +x $BUILD_NAME

    mv -f $BUILD_NAME $BUILD_DEST
    echo "$BUILD_NAME file moved into build folder"

    # Via: https://superuser.com/posts/1329136/revisions
    shopt -s dotglob
    cp -vr $COPY_TO $BUILD_DEST
    echo "Additional files/folders moved into build folder"

else
    cp -vr $COPY_TO $BUILD_DEST
    echo "Additional files/folders moved into build folder"
    
    echo "Something wrong for building service"
fi