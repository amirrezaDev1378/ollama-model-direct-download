#!/bin/bash

#set current working directory to the directory of the script
cd "$(dirname "$0")" || exit

OUTPUT_DIR="../bin"
mkdir -p $OUTPUT_DIR
APP_PREFIX="omdd"
PLATFORMS=("linux/amd64" "linux/386" "linux/arm64" "linux/arm" "darwin/amd64" "darwin/arm64" "windows/amd64" "windows/386" "freebsd/amd64" "freebsd/386" "freebsd/arm" "openbsd/amd64" "openbsd/386")
for PLATFORM in "${PLATFORMS[@]}"
do
    OS=$(echo "$PLATFORM" | cut -d'/' -f1)
    ARCH=$(echo "$PLATFORM" | cut -d'/' -f2)
    OUTPUT_NAME=$OUTPUT_DIR/$APP_PREFIX-$OS-$ARCH

    if [ "$OS" = "windows" ]; then
        OUTPUT_NAME+='.exe'
    fi

    echo "Building for $OS/$ARCH..."
    if ! GOOS=$OS GOARCH=$ARCH go build  -o "$OUTPUT_NAME" ../cmd/main.go
         then
            echo "An error occurred during the build for $PLATFORM"
            exit 1
        fi
    done

echo "Build completed!"
