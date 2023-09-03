#!/usr/bin/env bash

if [ -n "$1" ]; then
  IMAGE_NAME="$1";
  echo "Running image: ${IMAGE_NAME}";
else
  echo "Provide docker image name to run.";
  exit 1;
fi

docker run \
  -v ~/.gitconfig:/etc/gitconfig \
  --privileged \
  --network host \
  -v "${PWD}":/code \
  --entrypoint bash\
  --rm \
  -it \
  "${IMAGE_NAME}"
