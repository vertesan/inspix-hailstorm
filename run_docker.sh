mkdir -p cache && mkdir -p masterdata && \
docker run -it --rm \
  -v `pwd`/cache:/app/cache \
  -v `pwd`/masterdata:/app/masterdata \
  inspix-hailstorm:latest
