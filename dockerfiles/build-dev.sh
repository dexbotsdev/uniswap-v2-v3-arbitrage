## Check if the first argument is --no-cache

if [ "$1" == "--no-cache" ]; then
    CACHE_OPTION="--no-cache"
    echo "Building with no cache..."
else
    CACHE_OPTION=""
fi


docker build \
  ${CACHE_OPTION} \
  -f Dockerfile.dev \
  --build-arg USERNAME=$(whoami) \
  --build-arg USER_UID=$(id -u) \
  --build-arg USER_GID=$(id -g) \
  -t uni_arb .