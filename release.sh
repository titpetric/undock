#!/bin/bash
set -e
if [ ! -d "build" ]; then
	./build.sh
fi

date +"%y%m%d-%H%M" > build/.date

## Get latest tag ID
CI_TAG_ID=$(git tag | tail -n 1)
if [ -z "${CI_TAG_ID}" ]; then
	CI_TAG_ID="v0.0.0";
fi
CI_TAG_AUTO="$(echo ${CI_TAG_ID} | awk -F'.' '{print $1 "." $2}').$(<build/.date)"

function github_release {
	TAG="$1"
	NAME="$2"
	latest_tag=$(git describe --tags `git rev-list --tags --max-count=1`)
	comparison="$latest_tag..HEAD"
	if [ -z "$latest_tag" ]; then
		comparison="";
	fi
	changelog=$(git log $comparison --oneline --no-merges)
	echo "Creating release $1: $2"
	github-release release \
		--user titpetric \
		--repo undock \
		--tag "$1" \
		--name "$2" \
		--description "$changelog"
}

function github_upload {
	echo "Uploading $2 to $1"
	github-release upload \
		--user titpetric \
		--repo undock \
		--tag "$1" \
		--name "$(basename $2)" \
		--file "$2"
}

## Release to GitHub
github_release ${CI_TAG_AUTO} "$(date)"
FILES=$(find build -type f | grep gz$)
for FILE in $FILES; do
	github_upload ${CI_TAG_AUTO} "$FILE"
done
