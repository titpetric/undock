#!/bin/bash
set -e
rm -rf build
GOOSES=${1:-"darwin linux"}
GOARCHS=${2:-"amd64 386"}
for GOOS in $GOOSES; do
	for GOARCH in $GOARCHS; do
		echo "Building $GOOS/$GOARCH"
		if [ "$GOOS" == "windows" ]; then
			docker run --rm -e CGO_ENABLED=0 -e GOOS=$GOOS -e GOARCH=$GOARCH -v `pwd`:/go/src/app -w /go/src/app golang:1.9-alpine go build -o build/undock-$GOOS-$GOARCH.exe .
		else
			docker run --rm -e CGO_ENABLED=0 -e GOOS=$GOOS -e GOARCH=$GOARCH -v `pwd`:/go/src/app -w /go/src/app golang:1.9-alpine go build -o build/undock-$GOOS-$GOARCH .
			if [ "$GOOS" == "linux" ]; then
				strip build/undock-$GOOS-$GOARCH
			fi
		fi
	done
done
gzip -k build/*